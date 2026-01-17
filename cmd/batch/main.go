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
	if err := db.AutoMigrate(
		&model.SpecialMedicalDevice{},
		&model.Comment{},
		&model.Disease{},
		&model.Medication{},
		&model.CommentRelation{},
		&model.Ward{},
		&model.Tooth{},
		&model.MedicalPractice{},
		&model.MedicalPracticeInclusion{},
		&model.MedicalPracticeConflict{},
		&model.MedicalPracticeSupport{},
		&model.InpatientBasicFee{},
		&model.CalculationCount{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 病棟マスターのインポート
	importWards(db)

	// 歯式マスターのインポート
	importTeeth(db)

	// 特定器材マスターのインポート
	importDevices(db)

	// コメントマスターのインポート
	importComments(db)

	// コメント関連テーブルのインポート
	importCommentRelations(db)

	// 傷病名マスターのインポート
	importDiseases(db, "csv/b_20260101.txt", false)

	// 旧傷病名管理ファイルのインポート
	importDiseases(db, "csv/hb_20260101.txt", true)

	// 調剤行為マスターのインポート
	importMedications(db)

	// 医科診療行為マスターのインポート
	importMedicalPractices(db)

	// 包括テーブルのインポート
	importMedicalPracticeInclusions(db)

	// 背反テーブルのインポート
	importMedicalPracticeConflicts(db)

	// 補助マスターテーブルのインポート
	importMedicalPracticeSupports(db)

	// 入院基本料テーブルのインポート
	importInpatientBasicFees(db)

	// 算定回数テーブルのインポート
	importCalculationCounts(db)
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

func importWards(db *gorm.DB) {
	fmt.Println("Starting import from csv/k.csv...")
	f, err := os.Open("csv/k.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var wards []model.Ward
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

		if len(record) < 113 {
			log.Printf("invalid record length: expected at least 113, got %d", len(record))
			continue
		}

		ward := model.Ward{
			UpdateCategory:      record[0],
			MasterType:          record[1],
			Code:                record[2],
			NameKanjiLen:        parseInt(record[3]),
			NameKanji:           record[4],
			NameKanaLen:         parseInt(record[5]),
			NameKana:            record[6],
			PointCategory:       record[10],
			Point:               parseFloat(record[11]),
			InOuterCategory:     record[12],
			ElderlyCategory:     record[13],
			DpcCategory:         record[17],
			InstitutionCategory: record[18],
			MedicalObservation:  record[20],
			LowerLimit:          record[30],
			UpperLimit:          record[31],
			NotificationType:    record[67],
			NotificationType2:   record[68],
			DentalCategory:      record[83],
			UpdateDate:          record[86],
			DiscontinuedDate:    record[87],
			PublishOrder:        record[88],
			Chapter:             record[89],
			Section:             record[90],
			SubsectionCode:      record[91],
			BranchNumber:        record[92],
			ItemNumber:          record[93],
			Fullname:            record[112],
		}

		wards = append(wards, ward)

		if len(wards) >= batchSize {
			if err := db.Create(&wards).Error; err != nil {
				log.Fatalf("failed to bulk insert wards: %v", err)
			}
			count += len(wards)
			fmt.Printf("Inserted %d wards...\n", count)
			wards = nil
		}
	}

	if len(wards) > 0 {
		if err := db.Create(&wards).Error; err != nil {
			log.Fatalf("failed to bulk insert wards (last batch): %v", err)
		}
		count += len(wards)
	}

	fmt.Printf("Ward Master Batch process finished. Total %d records inserted.\n", count)
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

func importCommentRelations(db *gorm.DB) {
	fmt.Println("Starting import from csv/ck_ALL_20250901.csv...")
	f, err := os.Open("csv/ck_ALL_20250901.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var relations []model.CommentRelation
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

		if len(record) < 18 {
			log.Printf("invalid record length: expected at least 18, got %d", len(record))
			continue
		}

		relation := model.CommentRelation{
			UpdateCategory:    record[0],
			NotificationType:  record[1],
			ItemNumber:        record[2],
			Section:           record[3],
			BranchNumber:      record[4],
			ActCode:           record[5],
			AdditionCode:      record[6],
			ActNameAbbr:       record[7],
			CommentCode:       record[8],
			PatientStatusCode: record[9],
			CommentText:       record[10],
			UpdateDate:        record[11],
			DiscontinuedDate:  record[12],
			ConditionCategory: record[13],
			NonPaymentReason:  record[14],
			AdmissionCategory: record[15],
			CalculateCount:    parseInt(record[16]),
			PublishOrder:      record[17],
		}

		relations = append(relations, relation)

		if len(relations) >= batchSize {
			if err := db.Create(&relations).Error; err != nil {
				log.Fatalf("failed to bulk insert comment relations: %v", err)
			}
			count += len(relations)
			fmt.Printf("Inserted %d comment relations...\n", count)
			relations = nil
		}
	}

	if len(relations) > 0 {
		if err := db.Create(&relations).Error; err != nil {
			log.Fatalf("failed to bulk insert comment relations (last batch): %v", err)
		}
		count += len(relations)
	}

	fmt.Printf("Comment Relation Master Batch process finished. Total %d records inserted.\n", count)
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

func importTeeth(db *gorm.DB) {
	fmt.Println("Starting import from csv/f_20111003.csv...")
	f, err := os.Open("csv/f_20111003.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var teeth []model.Tooth
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

		if len(record) < 7 {
			log.Printf("invalid record length: expected at least 7, got %d", len(record))
			continue
		}

		tooth := model.Tooth{
			UpdateCategory:   record[0],
			MasterType:       record[1],
			Code:             record[2],
			Reserved:         record[3],
			Name:             record[4],
			UpdateDate:       record[5],
			DiscontinuedDate: record[6],
		}

		teeth = append(teeth, tooth)

		if len(teeth) >= batchSize {
			if err := db.Create(&teeth).Error; err != nil {
				log.Fatalf("failed to bulk insert teeth: %v", err)
			}
			count += len(teeth)
			fmt.Printf("Inserted %d teeth...\n", count)
			teeth = nil
		}
	}

	if len(teeth) > 0 {
		if err := db.Create(&teeth).Error; err != nil {
			log.Fatalf("failed to bulk insert teeth (last batch): %v", err)
		}
		count += len(teeth)
	}

	fmt.Printf("Tooth Master Batch process finished. Total %d records inserted.\n", count)
}

func importMedicalPractices(db *gorm.DB) {
	fmt.Println("Starting import from csv/s_20251201.csv...")
	f, err := os.Open("csv/s_20251201.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var practices []model.MedicalPractice
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

		if len(record) < 117 {
			log.Printf("invalid record length: expected at least 117, got %d", len(record))
			continue
		}

		practice := model.MedicalPractice{
			ChangeCategory:                 record[0],
			MasterType:                     record[1],
			MedicalPracticeCode:            record[2],
			AbbreviatedKanjiNameLength:     parseInt(record[3]),
			AbbreviatedKanjiName:           record[4],
			AbbreviatedKanaNameLength:      parseInt(record[5]),
			AbbreviatedKanaName:            record[6],
			DataStandardCode:               record[7],
			KanjiUnitNameLength:            parseInt(record[8]),
			KanjiUnitName:                  record[9],
			ScoreType:                      record[10],
			Score:                          parseFloat(record[11]),
			InpatientOutpatientCategory:    record[12],
			ElderlyMedicalCategory:         record[13],
			AggregationCategoryOutpatient:  record[14],
			ComprehensiveTargetTest:        record[15],
			DPCApplicabilityCategory:       record[17],
			HospitalClinicCategory:         record[18],
			SurgerySupportCategory:         record[19],
			MedicalObservationCategory:     record[20],
			NursingAddition:                record[21],
			AnesthesiaCategory:             record[22],
			DiseaseRelatedCategory:         record[24],
			ActualDaysCategory:             record[26],
			DaysTimesCategory:              record[27],
			MedicineRelatedCategory:        record[28],
			StepCalculationCategory:        record[29],
			StepLowerLimit:                 parseFloat(record[30]),
			StepUpperLimit:                 parseFloat(record[31]),
			StepValue:                      parseFloat(record[32]),
			StepScore:                      parseFloat(record[33]),
			StepErrorCategory:              record[34],
			MaxTimes:                       parseInt(record[35]),
			MaxTimesErrorCategory:          record[36],
			NoteAdditionCode:               record[37],
			NoteAdditionSequence:           record[38],
			GeneralAgeCategory:             record[39],
			LowerAge:                       record[40],
			UpperAge:                       record[41],
			TimeAdditionCategory:           record[42],
			FacilityStandardCategory:       record[43],
			FacilityStandardCode:           record[44],
			TreatmentInfantAddition:        record[45],
			VeryLowBirthWeightAddition:     record[46],
			InpatientReductionCategory:     record[47],
			DonorCollectionCategory:        record[48],
			TestJudgmentCategory:           record[49],
			TestJudgmentGroupCategory:      record[50],
			ReductionApplicability:         record[51],
			SpinalEvokedPotentialCategory:  record[52],
			NeckDissectionCategory:         record[53],
			AutoSutureCategory:             record[54],
			OutpatientManagementAddition:   record[55],
			OldScoreType:                   record[56],
			OldScore:                       parseFloat(record[57]),
			TestCommentCategory:            record[60],
			GeneralAdditionApplicability:   record[61],
			ComprehensiveReductionCategory: record[62],
			EndoscopicSurgeryAddition:      record[63],
			AggregationCategoryInpatient:   record[65],
			AutoAnastomosisCategory:        record[66],
			NotificationCategory1:          record[67],
			NotificationCategory2:          record[68],
			RegionAddition:                 record[69],
			BedCountCategory:               record[70],
			FacilityStandard1:              record[71],
			FacilityStandard2:              record[72],
			FacilityStandard3:              record[73],
			FacilityStandard4:              record[74],
			FacilityStandard5:              record[75],
			FacilityStandard6:              record[76],
			FacilityStandard7:              record[77],
			FacilityStandard8:              record[78],
			FacilityStandard9:              record[79],
			FacilityStandard10:             record[80],
			UltrasoundCoagulationCategory:  record[81],
			ShortStaySurgery:               record[82],
			DentalApplicability:            record[83],
			SectionCodeAlpha:               record[84],
			NotificationRelatedAlpha:       record[85],
			UpdateDate:                     record[86],
			ExpiryDate:                     record[87],
			PublicationOrder:               record[88],
			SectionChapter:                 record[89],
			SectionPart:                    record[90],
			SectionNumber:                  record[91],
			SectionSubNumber:               record[92],
			SectionItemNumber:              record[93],
			BasicKanjiName:                 record[112],
			PointTableNumber:               record[116],
		}

		practices = append(practices, practice)

		if len(practices) >= batchSize {
			if err := db.Create(&practices).Error; err != nil {
				log.Fatalf("failed to bulk insert medical practices: %v", err)
			}
			count += len(practices)
			fmt.Printf("Inserted %d medical practices...\n", count)
			practices = nil
		}
	}

	if len(practices) > 0 {
		if err := db.Create(&practices).Error; err != nil {
			log.Fatalf("failed to bulk insert medical practices (last batch): %v", err)
		}
		count += len(practices)
	}

	fmt.Printf("Medical Practice Master Batch process finished. Total %d records inserted.\n", count)
}

func importMedicalPracticeInclusions(db *gorm.DB) {
	fmt.Println("Starting import from csv/tensuhyo_02/02包括テーブル.csv...")
	f, err := os.Open("csv/tensuhyo_02/02包括テーブル.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var inclusions []model.MedicalPracticeInclusion
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

		if len(record) < 7 {
			log.Printf("invalid record length: expected at least 7, got %d", len(record))
			continue
		}

		inclusion := model.MedicalPracticeInclusion{
			ChangeCategory:            record[0],
			ComprehensivePracticeCode: record[1],
			IncludedPracticeCode:      record[2],
			UpdateDate:                record[5],
			ExpiryDate:                record[6],
		}

		inclusions = append(inclusions, inclusion)

		if len(inclusions) >= batchSize {
			if err := db.Create(&inclusions).Error; err != nil {
				log.Fatalf("failed to bulk insert inclusions: %v", err)
			}
			count += len(inclusions)
			fmt.Printf("Inserted %d inclusions...\n", count)
			inclusions = nil
		}
	}

	if len(inclusions) > 0 {
		if err := db.Create(&inclusions).Error; err != nil {
			log.Fatalf("failed to bulk insert inclusions (last batch): %v", err)
		}
		count += len(inclusions)
	}

	fmt.Printf("Inclusion Master Batch process finished. Total %d records inserted.\n", count)
}

func importMedicalPracticeConflicts(db *gorm.DB) {
	filePaths := []string{
		"csv/tensuhyo_02/03-1背反テーブル1.csv",
		"csv/tensuhyo_02/03-2背反テーブル2.csv",
		"csv/tensuhyo_02/03-3背反テーブル3.csv",
		"csv/tensuhyo_02/03-4背反テーブル4.csv",
	}

	for _, filePath := range filePaths {
		fmt.Printf("Starting import from %s...\n", filePath)
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("failed to open csv file %s: %v", filePath, err)
		}

		reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
		reader.LazyQuotes = true

		var conflicts []model.MedicalPracticeConflict
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

			var conflict model.MedicalPracticeConflict
			if len(record) >= 102 {
				conflict = model.MedicalPracticeConflict{
					ChangeCategory:       record[0],
					MedicalPracticeCode:  record[1],
					ConflictPracticeCode: record[11],
					UpdateDate:           record[99],
					ExpiryDate:           record[100],
				}
			} else if len(record) >= 10 {
				conflict = model.MedicalPracticeConflict{
					ChangeCategory:       record[0],
					MedicalPracticeCode:  record[1],
					ConflictPracticeCode: record[3],
					UpdateDate:           record[8],
					ExpiryDate:           record[9],
				}
			} else {
				log.Printf("invalid record length in %s: expected at least 10, got %d", filePath, len(record))
				continue
			}

			conflicts = append(conflicts, conflict)

			if len(conflicts) >= batchSize {
				if err := db.Create(&conflicts).Error; err != nil {
					log.Fatalf("failed to bulk insert conflicts: %v", err)
				}
				count += len(conflicts)
				fmt.Printf("Inserted %d conflicts from %s...\n", count, filePath)
				conflicts = nil
			}
		}

		if len(conflicts) > 0 {
			if err := db.Create(&conflicts).Error; err != nil {
				log.Fatalf("failed to bulk insert conflicts (last batch): %v", err)
			}
			count += len(conflicts)
		}
		f.Close()
		fmt.Printf("Conflict Master Batch finished for %s. Total %d records inserted.\n", filePath, count)
	}
}

func importMedicalPracticeSupports(db *gorm.DB) {
	fmt.Println("Starting import from csv/tensuhyo_02/01補助マスターテーブル.csv...")
	f, err := os.Open("csv/tensuhyo_02/01補助マスターテーブル.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var supports []model.MedicalPracticeSupport
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

		if len(record) < 27 {
			log.Printf("invalid record length: expected at least 27, got %d", len(record))
			continue
		}

		support := model.MedicalPracticeSupport{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			SupportInfo:         record[3],
			UpdateDate:          record[25],
			ExpiryDate:          record[26],
		}

		supports = append(supports, support)

		if len(supports) >= batchSize {
			if err := db.Create(&supports).Error; err != nil {
				log.Fatalf("failed to bulk insert supports: %v", err)
			}
			count += len(supports)
			fmt.Printf("Inserted %d supports...\n", count)
			supports = nil
		}
	}

	if len(supports) > 0 {
		if err := db.Create(&supports).Error; err != nil {
			log.Fatalf("failed to bulk insert supports (last batch): %v", err)
		}
		count += len(supports)
	}

	fmt.Printf("Support Master Batch process finished. Total %d records inserted.\n", count)
}

func importInpatientBasicFees(db *gorm.DB) {
	fmt.Println("Starting import from csv/tensuhyo_02/04入院基本料テーブル.csv...")
	f, err := os.Open("csv/tensuhyo_02/04入院基本料テーブル.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var fees []model.InpatientBasicFee
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

		if len(record) < 8 {
			log.Printf("invalid record length: expected at least 8, got %d", len(record))
			continue
		}

		fee := model.InpatientBasicFee{
			ChangeCategory:      record[0],
			BasicFeeCode:        record[1],
			MedicalPracticeCode: record[2],
			UpdateDate:          record[6],
			ExpiryDate:          record[7],
		}

		fees = append(fees, fee)

		if len(fees) >= batchSize {
			if err := db.Create(&fees).Error; err != nil {
				log.Fatalf("failed to bulk insert inpatient fees: %v", err)
			}
			count += len(fees)
			fmt.Printf("Inserted %d inpatient fees...\n", count)
			fees = nil
		}
	}

	if len(fees) > 0 {
		if err := db.Create(&fees).Error; err != nil {
			log.Fatalf("failed to bulk insert inpatient fees (last batch): %v", err)
		}
		count += len(fees)
	}

	fmt.Printf("Inpatient Fee Master Batch process finished. Total %d records inserted.\n", count)
}

func importCalculationCounts(db *gorm.DB) {
	fmt.Println("Starting import from csv/tensuhyo_02/05算定回数テーブル.csv...")
	f, err := os.Open("csv/tensuhyo_02/05算定回数テーブル.csv")
	if err != nil {
		log.Fatalf("failed to open csv file: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var counts []model.CalculationCount
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

		if len(record) < 14 {
			log.Printf("invalid record length: expected at least 14, got %d", len(record))
			continue
		}

		calcCount := model.CalculationCount{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			CountLimit:          parseInt(record[5]),
			UpdateDate:          record[12],
			ExpiryDate:          record[13],
		}

		counts = append(counts, calcCount)

		if len(counts) >= batchSize {
			if err := db.Create(&counts).Error; err != nil {
				log.Fatalf("failed to bulk insert calculation counts: %v", err)
			}
			count += len(counts)
			fmt.Printf("Inserted %d calculation counts...\n", count)
			counts = nil
		}
	}

	if len(counts) > 0 {
		if err := db.Create(&counts).Error; err != nil {
			log.Fatalf("failed to bulk insert calculation counts (last batch): %v", err)
		}
		count += len(counts)
	}

	fmt.Printf("Calculation Count Master Batch process finished. Total %d records inserted.\n", count)
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
