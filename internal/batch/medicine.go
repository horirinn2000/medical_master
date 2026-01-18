package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportMedicines(db *gorm.DB) {
	// 全42項目チェック (R06rec3.pdf 196-197ページ)
	importCsv(db, 42, "csv/medicine/y_20260116.csv", func(record []string) model.Medicine {
		return model.Medicine{
			ChangeCategory:                 record[0],
			MasterType:                     record[1],
			MedicineCode:                   record[2],
			KanjiNameLength:                parseInt(record[3]),
			KanjiName:                      record[4],
			KanaNameLength:                 parseInt(record[5]),
			KanaName:                       record[6],
			UnitCode:                       record[7],
			UnitKanjiNameLength:            parseInt(record[8]),
			UnitKanjiName:                  record[9],
			PriceType:                      record[10],
			Price:                          parseFloat(record[11]),
			Reserved13:                     record[12],
			NarcoticToxicDrugCategory:      record[13],
			NerveDestructionAgent:          record[14],
			BiologicProduct:                record[15],
			GenericDrug:                    record[16],
			Reserved18:                     record[17],
			DentalSpecificDrug:             record[18],
			ContrastAgent:                  record[19],
			InjectionVolume:                parseInt(record[20]),
			ListingMethod:                  record[21],
			BrandNameRelated:               record[22],
			OldPriceType:                   record[23],
			OldPrice:                       parseFloat(record[24]),
			KanjiNameUpdateFlag:            record[25],
			KanaNameUpdateFlag:             record[26],
			DosageForm:                     record[27],
			Reserved29:                     record[28],
			UpdateDate:                     record[29],
			ExpiryDate:                     record[30],
			NationalDrugCode:               record[31],
			PublicationOrder:               parseInt(record[32]),
			TransitionalDate:               record[33],
			BasicKanjiName:                 record[34],
			PriceListingDate:               record[35],
			GenericNameCode:                record[36],
			GenericNameStandardDescription: record[37],
			GenericAdditionCategory:        record[38],
			AntiHIVDrugCategory:            record[39],
			LongTermListingRelated:         record[40],
			SelectedMedicalCategory:        record[41],
		}
	})
}

func ImportHotCodes(db *gorm.DB) {
	// 標準レイアウト (24項目)
	importCsv(db, 24, "csv/medicine/MEDIS20251130.TXT", func(record []string) model.HotCode {
		return model.HotCode{
			HotCode:             record[0],
			Hot7:                record[1],
			CompanyCode:         record[2],
			DispensingCode:      record[3],
			LogisticsCode:       record[4],
			JanCode:             record[5],
			NationalDrugCode:    record[6],
			IndividualDrugCode:  record[7],
			ReceiptCode1:        record[8],
			ReceiptCode2:        record[9],
			OfficialName:        record[10],
			SalesName:           record[11],
			ReceiptMedicineName: record[12],
			StandardUnit:        record[13],
			PackageType:         record[14],
			PackageQuantityNum:  parseFloat(record[15]),
			PackageQuantityUnit: record[16],
			TotalQuantityNum:    parseFloat(record[17]),
			TotalQuantityUnit:   record[18],
			Category:            record[19],
			Manufacturer:        record[20],
			SalesCompany:        record[21],
			RecordCategory:      record[22],
			UpdateDate:          record[23],
		}
	})

	// オプションレイアウトから補完する処理は必要に応じて別途実装
}
