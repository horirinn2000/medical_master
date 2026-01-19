package batch

import (
	"encoding/csv"
	"fmt"
	"io"
	"medical_master/internal/model"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

func ImportMedicines(db *gorm.DB) {
	importCsv(db, 42, "csv/medicine/y_20260116.csv", func(record []string) model.Medicine {
		m := model.Medicine{
			UpdateCategory:           record[0],
			MasterType:               record[1],
			Code:                     record[2],
			NameKanjiLen:             parseInt(record[3]),
			NameKanji:                record[4],
			NameKanaLen:              parseInt(record[5]),
			NameKana:                 record[6],
			UnitCode:                 record[7],
			UnitNameKanjiLen:         parseInt(record[8]),
			UnitNameKanji:            record[9],
			PriceType:                record[10],
			Price:                    parseFloat(record[11]),
			Reserved1:                record[12],
			NarcoticsCategory:        record[13],
			NeurolyticFlag:           record[14],
			BiologicFlag:             record[15],
			GenericFlag:              record[16],
			Reserved2:                record[17],
			DentalSpecificFlag:       record[18],
			ContrastAgentCategory:    record[19],
			InjectionVolume:          parseInt(record[20]),
			ListingMethodType:        record[21],
			BrandNameCode:            record[22],
			OldPriceType:             record[23],
			OldPrice:                 parseFloat(record[24]),
			NameKanjiUpdateFlag:      record[25],
			NameKanaUpdateFlag:       record[26],
			DosageForm:               record[27],
			Reserved3:                record[28],
			UpdateDate:               record[29],
			DiscontinuedDate:         record[30],
			NationalDrugCode:         record[31],
			PublishOrder:             parseInt(record[32]),
			TransitionalDate:         record[33],
			BasicName:                record[34],
			ListingDate:              record[35],
			GenericNameCode:          record[36],
			GenericNameStandard:      record[37],
			GenericAdditionType:      record[38],
			AntiHivFlag:              record[39],
			LongTermListingCode:      record[40],
			SelectionMedicalCategory: record[41],
		}

		// 価格履歴の保存 (現金額)
		if m.Price > 0 {
			db.Where(model.MedicinePrice{MedicineCode: m.Code, Price: m.Price, PriceType: "1"}).FirstOrCreate(&model.MedicinePrice{
				MedicineCode: m.Code,
				Price:        m.Price,
				PriceType:    "1",
				StartDate:    m.ListingDate,
			})
		}
		// 価格履歴の保存 (旧金額)
		if m.OldPrice > 0 {
			db.Where(model.MedicinePrice{MedicineCode: m.Code, Price: m.OldPrice, PriceType: "2"}).FirstOrCreate(&model.MedicinePrice{
				MedicineCode: m.Code,
				Price:        m.OldPrice,
				PriceType:    "2",
				StartDate:    "", // 旧価格の開始日は不明
			})
		}

		return m
	})
}

func ImportHotCodes(db *gorm.DB) {
	fmt.Println("Starting import HOT codes from MEDIS files...")

	// 1. オプション情報をマップに保持 (全14項目)
	optionsMap := make(map[string]model.HotCode)
	optFile, err := os.Open("csv/medicine/MEDIS20251130_OP.TXT")
	if err == nil {
		defer optFile.Close()
		optReader := csv.NewReader(transform.NewReader(optFile, japanese.ShiftJIS.NewDecoder()))
		optReader.LazyQuotes = true
		for {
			record, err := optReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil || len(record) < 7 {
				continue
			}
			optionsMap[record[0]] = model.HotCode{
				OptionPackageQuantity:       parseFloat(record[1]),
				OptionPackageQuantityUnit:   record[2],
				OptionPackageInQuantity:     parseFloat(record[3]),
				OptionPackageInQuantityUnit: record[4],
				OptionInnerPackageQuantity:  parseFloat(record[5]),
				OptionJanCode:               record[6],
				OptionItfCode:               record[7],
				OptionRssiCode:              record[8],
				OptionHot7:                  record[9],
				OptionReceiptCode:           record[10],
				OptionHot9:                  record[11],
				OptionUpdateCategory:        record[12],
				OptionUpdateDate:            record[13],
			}
		}
	}

	// 2. HOT9情報をマップに保持 (24項目)
	hot9Map := make(map[string]model.HotCode)
	hot9File, err := os.Open("csv/medicine/MEDIS20251130_HOT9.TXT")
	if err == nil {
		defer hot9File.Close()
		hot9Reader := csv.NewReader(transform.NewReader(hot9File, japanese.ShiftJIS.NewDecoder()))
		hot9Reader.LazyQuotes = true
		for {
			record, err := hot9Reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil || len(record) < 24 {
				continue
			}
			// record[0] は 基準番号(HOT9)
			hot9Map[record[0]] = model.HotCode{
				Hot9:           record[0],
				Hot9Hot7:       record[1],
				Hot9Receipt1:   record[8],
				Hot9SalesName:  record[10],
				Hot9Usage:      record[18],
				Hot9Owner:      record[19],
				Hot9Vendor:     record[20],
				Hot9UpdateDate: record[23],
			}
		}
	}

	// 3. 標準セットを読み込んで統合してDB保存
	importCsv(db, 24, "csv/medicine/MEDIS20251130.TXT", func(record []string) model.HotCode {
		hotCode := model.HotCode{
			HotCode:          record[0],
			Hot7:             record[1],
			CompanyCode:      record[2],
			DispensingNo:     record[3],
			LogisticsNo:      record[4],
			JanCode:          record[5],
			NationalCode:     record[6],
			IndividualCode:   record[7],
			ReceiptCode1:     record[8],
			ReceiptCode2:     record[9],
			NotificationName: record[10],
			SalesName:        record[11],
			ReceiptName:      record[12],
			SpecUnit:         record[13],
			PackageForm:      record[14],
			PackageCount:     parseFloat(record[15]),
			PackageUnit:      record[16],
			TotalVolume:      parseFloat(record[17]),
			TotalVolumeUnit:  record[18],
			Category:         record[19],
			Manufacturer:     record[20],
			Distributor:      record[21],
			RecordType:       record[22],
			UpdateDate:       record[23],
		}

		// オプション情報のマージ
		if opt, ok := optionsMap[hotCode.HotCode]; ok {
			hotCode.OptionPackageQuantity = opt.OptionPackageQuantity
			hotCode.OptionPackageQuantityUnit = opt.OptionPackageQuantityUnit
			hotCode.OptionPackageInQuantity = opt.OptionPackageInQuantity
			hotCode.OptionPackageInQuantityUnit = opt.OptionPackageInQuantityUnit
			hotCode.OptionInnerPackageQuantity = opt.OptionInnerPackageQuantity
			hotCode.OptionJanCode = opt.OptionJanCode
			hotCode.OptionItfCode = opt.OptionItfCode
			hotCode.OptionRssiCode = opt.OptionRssiCode
			hotCode.OptionHot7 = opt.OptionHot7
			hotCode.OptionReceiptCode = opt.OptionReceiptCode
			hotCode.OptionHot9 = opt.OptionHot9
			hotCode.OptionUpdateCategory = opt.OptionUpdateCategory
			hotCode.OptionUpdateDate = opt.OptionUpdateDate
		}

		// HOT9情報のマージ
		// 基準番号の頭9桁がHOT9に相当する場合が多いが、MEDIS仕様に合わせる
		hot9Key := hotCode.HotCode
		if len(hot9Key) > 9 {
			hot9Key = hot9Key[:9]
		}
		if h9, ok := hot9Map[hot9Key]; ok {
			hotCode.Hot9 = h9.Hot9
			hotCode.Hot9Hot7 = h9.Hot9Hot7
			hotCode.Hot9Receipt1 = h9.Hot9Receipt1
			hotCode.Hot9SalesName = h9.Hot9SalesName
			hotCode.Hot9Usage = h9.Hot9Usage
			hotCode.Hot9Owner = h9.Hot9Owner
			hotCode.Hot9Vendor = h9.Hot9Vendor
			hotCode.Hot9UpdateDate = h9.Hot9UpdateDate
		}

		return hotCode
	})

	fmt.Println("Finished import HOT codes.")
}
