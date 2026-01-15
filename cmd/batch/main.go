package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"medical_master/model"
	"os"
	"strconv"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB接続設定
	dsn := "host=localhost user=postgres password=postgres dbname=medical_master port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&model.SpecialMedicalDevice{}, &model.Comment{}, &model.Disease{}, &model.Medication{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 特定器材マスターのインポート
	importDevices(db)

	// コメントマスターのインポート
	importComments(db)

	// 傷病名マスターのインポート
	importDiseases(db, "csv/b_20260101.txt", false)

	// 旧傷病名管理ファイルのインポート
	importDiseases(db, "csv/hb_20260101.txt", true)

	// 調剤行為マスターのインポート
	importMedications(db)
}

func importDiseases(db *gorm.DB, filePath string, isOld bool) {
	fmt.Printf("Starting import from %s (isOld=%v)...\n", filePath, isOld)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var diseases []model.Disease
	batchSize := 500
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read record: %v", err)
			continue
		}

		if len(record) < 46 {
			log.Printf("invalid record length: expected at least 46, got %d", len(record))
			continue
		}

		disease := model.Disease{
			IsOld:                         isOld,
			UpdateCategory:                record[0],
			MasterType:                    record[1],
			Code:                          record[2],
			RelatedCode:                   record[3],
			NameKanjiLen:                  parseInt(record[4]),
			NameKanji:                     record[5],
			NameAbbrLen:                   parseInt(record[6]),
			NameAbbr:                      record[7],
			NameKanaLen:                   parseInt(record[8]),
			NameKana:                      record[9],
			DiseaseManagementNumber:       record[10],
			AdoptionCategory:              record[11],
			DiseaseExchangeCode:           record[12],
			Reserved1:                     record[13],
			Reserved2:                     record[14],
			Icd10_1:                       record[15],
			Icd10_2:                       record[16],
			Reserved3:                     record[17],
			SingleUseForbiddenCategory:    record[18],
			InsuranceOutCategory:          record[19],
			SpecificDiseaseCategory:       record[20],
			ImportDate:                    record[21],
			UpdateDate:                    record[22],
			DiscontinuedDate:              record[23],
			NameKanjiUpdateFlag:           record[24],
			NameAbbrUpdateFlag:            record[25],
			NameKanaUpdateFlag:            record[26],
			AdoptionUpdateFlag:            record[27],
			DiseaseExchangeUpdateFlag:     record[28],
			Reserved4:                     record[29],
			Reserved5:                     record[30],
			DentalNameAbbrUpdateFlag:      record[31],
			IncurableDiseaseUpdateFlag:    record[32],
			DentalSpecificUpdateFlag:      record[33],
			SingleUseForbiddenUpdateFlag:  record[34],
			InsuranceOutUpdateFlag:        record[35],
			SpecificDiseaseUpdateFlag:     record[36],
			MigrationManagementNumber:     record[37],
			DentalNameAbbr:                record[38],
			Reserved6:                     record[39],
			Reserved7:                     record[40],
			DentalNameAbbrLen:             parseInt(record[41]),
			IncurableDiseaseCategory:      record[42],
			DentalSpecificDiseaseCategory: record[43],
			Icd10_1UpdateFlag:             record[44],
			Icd10_2UpdateFlag:             record[45],
		}

		diseases = append(diseases, disease)

		if len(diseases) >= batchSize {
			if err := db.Create(&diseases).Error; err != nil {
				log.Fatalf("failed to bulk insert diseases: %v", err)
			}
			count += len(diseases)
			fmt.Printf("Inserted %d diseases...\n", count)
			diseases = nil
		}
	}

	if len(diseases) > 0 {
		if err := db.Create(&diseases).Error; err != nil {
			log.Fatalf("failed to bulk insert diseases (last batch): %v", err)
		}
		count += len(diseases)
	}

	fmt.Printf("Disease import finished. Total %d records inserted from %s.\n", count, filePath)
}

func importDevices(db *gorm.DB) {
	// CSVファイルオープン
	f, err := os.Open("csv/t_20251128.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	// Shift-JISからUTF-8への変換
	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var devices []model.SpecialMedicalDevice
	batchSize := 100
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read record: %v", err)
			continue
		}

		// カラム数がPDFの仕様と一致するか確認（38カラムを想定）
		// 提供されたサンプルの末尾の空要素などにより38以上になる可能性あり
		if len(record) < 38 {
			log.Printf("invalid record length: expected at least 38, got %d", len(record))
			continue
		}

		device := model.SpecialMedicalDevice{
			UpdateCategory:               record[0],
			MasterType:                   record[1],
			Code:                         record[2],
			NameKanjiLen:                 parseInt(record[3]),
			NameKanji:                    record[4],
			NameKanaLen:                  parseInt(record[5]),
			NameKana:                     record[6],
			UnitCode:                     record[7],
			UnitNameKanjiLen:             parseInt(record[8]),
			UnitNameKanji:                record[9],
			PriceType:                    record[10],
			Price:                        parseFloat(record[11]),
			Reserved1:                    record[12],
			AgeAdditionCategory:          record[13],
			MinAge:                       record[14],
			MaxAge:                       record[15],
			OldPriceType:                 record[16],
			OldPrice:                     parseFloat(record[17]),
			NameKanjiUpdateFlag:          record[18],
			NameKanaUpdateFlag:           record[19],
			OxygenCategory:               record[20],
			DeviceCategory:               record[21],
			MaxPriceFlag:                 record[22],
			MaxPoints:                    parseInt(record[23]),
			Reserved2:                    record[24],
			PublishOrder:                 parseInt(record[25]),
			RelatedCode:                  record[26],
			UpdateDate:                   record[27],
			TransitionalDate:             record[28],
			DiscontinuedDate:             record[29],
			AppendixNumber:               record[30],
			SectionNumber:                record[31],
			DpcCategory:                  record[32],
			Reserved3:                    record[33],
			Reserved4:                    record[34],
			Reserved5:                    record[35],
			Fullname:                     record[36],
			RemanufacturedDeviceCategory: record[37],
		}

		devices = append(devices, device)

		if len(devices) >= batchSize {
			if err := db.Create(&devices).Error; err != nil {
				log.Fatalf("failed to bulk insert: %v", err)
			}
			count += len(devices)
			fmt.Printf("Inserted %d records...\n", count)
			devices = nil
		}
	}

	// 残りのデータをインサート
	if len(devices) > 0 {
		if err := db.Create(&devices).Error; err != nil {
			log.Fatalf("failed to bulk insert (last batch): %v", err)
		}
		count += len(devices)
	}

	fmt.Printf("Device Master Batch process finished. Total %d records inserted.\n", count)
}

func importMedications(db *gorm.DB) {
	fmt.Println("Starting import from csv/m20250307.csv...")
	f, err := os.Open("csv/m20250307.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var medications []model.Medication
	batchSize := 100
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read record: %v", err)
			continue
		}

		if len(record) < 65 {
			log.Printf("invalid record length: expected at least 65, got %d", len(record))
			continue
		}

		medication := model.Medication{
			UpdateCategory:             record[0],
			MasterType:                 record[1],
			Code:                       record[2],
			NameKanjiLen:               parseInt(record[3]),
			NameKanji:                  record[4],
			NameKanaLen:                parseInt(record[5]),
			NameKana:                   record[6],
			UnitCode:                   record[7],
			UnitNameKanjiLen:           parseInt(record[8]),
			UnitNameKanji:              record[9],
			PriceType:                  record[10],
			Price:                      parseFloat(record[11]),
			Reserved1:                  record[12],
			AgeAdditionCategory:        record[13],
			MinAge:                     record[14],
			MaxAge:                     record[15],
			OldPriceType:               record[16],
			OldPrice:                   parseFloat(record[17]),
			NameKanjiUpdateFlag:        record[18],
			NameKanaUpdateFlag:         record[19],
			Reserved2:                  record[20],
			Reserved3:                  record[21],
			MaxPriceFlag:               record[22],
			MaxPoints:                  parseInt(record[23]),
			Reserved4:                  record[24],
			PublishOrder:               parseInt(record[25]),
			RelatedCode:                record[26],
			UpdateDate:                 record[27],
			TransitionalDate:           record[28],
			DiscontinuedDate:           record[29],
			MedicationSectionNumber:    record[30],
			Reserved5:                  record[31],
			Reserved6:                  record[32],
			Reserved7:                  record[33],
			Reserved8:                  record[34],
			Reserved9:                  record[35],
			Reserved10:                 record[36],
			Reserved11:                 record[37],
			Reserved12:                 record[38],
			Reserved13:                 record[39],
			Reserved14:                 record[40],
			Reserved15:                 record[41],
			Reserved16:                 record[42],
			Reserved17:                 record[43],
			Reserved18:                 record[44],
			Reserved19:                 record[45],
			Reserved20:                 record[46],
			Reserved21:                 record[47],
			Reserved22:                 record[48],
			Reserved23:                 record[49],
			Reserved24:                 record[50],
			Reserved25:                 record[51],
			Reserved26:                 record[52],
			Reserved27:                 record[53],
			Reserved28:                 record[54],
			Reserved29:                 record[55],
			Reserved30:                 record[56],
			Reserved31:                 record[57],
			PriceTablePrice:            parseFloat(record[58]),
			ImportDate:                 record[59],
			Reserved32:                 record[60],
			MedicationManagementNumber: record[61],
			MedicationMigrationNumber:  record[62],
			MedicationMigrationFlag:    record[63],
			MedicationMigrationUpdate:  record[64],
		}

		medications = append(medications, medication)

		if len(medications) >= batchSize {
			if err := db.Create(&medications).Error; err != nil {
				log.Fatalf("failed to bulk insert medications: %v", err)
			}
			count += len(medications)
			fmt.Printf("Inserted %d medications...\n", count)
			medications = nil
		}
	}

	if len(medications) > 0 {
		if err := db.Create(&medications).Error; err != nil {
			log.Fatalf("failed to bulk insert medications (last batch): %v", err)
		}
		count += len(medications)
	}

	fmt.Printf("Medication Master Batch process finished. Total %d records inserted.\n", count)
}

func importComments(db *gorm.DB) {
	// CSVファイルオープン
	f, err := os.Open("csv/c_20250715.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	// Shift-JISからUTF-8への変換
	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var comments []model.Comment
	batchSize := 100
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read record: %v", err)
			continue
		}

		// カラム数がPDFの仕様と一致するか確認（30カラムを想定）
		if len(record) < 30 {
			log.Printf("invalid record length: expected at least 30, got %d", len(record))
			continue
		}

		comment := model.Comment{
			UpdateCategory:      record[0],
			MasterType:          record[1],
			Category:            record[2],
			Pattern:             record[3],
			SerialNo:            record[4],
			NameKanjiLen:        parseInt(record[5]),
			NameKanji:           record[6],
			NameKanaLen:         parseInt(record[7]),
			NameKana:            record[8],
			ReceiptEditInfo1Pos: parseInt(record[9]),
			ReceiptEditInfo1Len: parseInt(record[10]),
			ReceiptEditInfo2Pos: parseInt(record[11]),
			ReceiptEditInfo2Len: parseInt(record[12]),
			ReceiptEditInfo3Pos: parseInt(record[13]),
			ReceiptEditInfo3Len: parseInt(record[14]),
			ReceiptEditInfo4Pos: parseInt(record[15]),
			ReceiptEditInfo4Len: parseInt(record[16]),
			Reserved1:           record[17],
			Reserved2:           record[18],
			SelectionCategory:   record[19],
			UpdateDate:          record[20],
			DiscontinuedDate:    record[21],
			Code:                record[22],
			PublishOrder:        parseInt(record[23]),
			Reserved3:           record[24],
			Reserved4:           record[25],
			Reserved5:           record[26],
			Reserved6:           record[27],
			Reserved7:           record[28],
			Reserved8:           record[29],
		}

		comments = append(comments, comment)

		if len(comments) >= batchSize {
			if err := db.Create(&comments).Error; err != nil {
				log.Fatalf("failed to bulk insert comments: %v", err)
			}
			count += len(comments)
			fmt.Printf("Inserted %d comments...\n", count)
			comments = nil
		}
	}

	// 残りのデータをインサート
	if len(comments) > 0 {
		if err := db.Create(&comments).Error; err != nil {
			log.Fatalf("failed to bulk insert comments (last batch): %v", err)
		}
		count += len(comments)
	}

	fmt.Printf("Comment Master Batch process finished. Total %d records inserted.\n", count)
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
