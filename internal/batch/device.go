package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDevices(db *gorm.DB) {
	// 全38項目チェック (R06rec3.pdf 198-199ページ)
	importCsv(db, 38, "csv/device/t_20251128.csv", func(record []string) model.SpecialMedicalDevice {
		return model.SpecialMedicalDevice{
			ChangeCategory:               record[0],
			MasterType:                   record[1],
			Code:                         record[2],
			KanjiNameLength:              parseInt(record[3]),
			KanjiName:                    record[4],
			KanaNameLength:               parseInt(record[5]),
			KanaName:                     record[6],
			UnitCode:                     record[7],
			UnitKanjiNameLength:          parseInt(record[8]),
			UnitKanjiName:                record[9],
			PriceType:                    record[10],
			Price:                        parseFloat(record[11]),
			Reserved13:                   record[12],
			AgeAdditionCategory:          record[13],
			MinAge:                       record[14],
			MaxAge:                       record[15],
			OldPriceType:                 record[16],
			OldPrice:                     parseFloat(record[17]),
			KanjiNameUpdateFlag:          record[18],
			KanaNameUpdateFlag:           record[19],
			OxygenCategory:               record[20],
			DeviceCategory:               record[21],
			MaxPriceFlag:                 record[22],
			MaxPoints:                    parseInt(record[23]),
			Reserved25:                   record[24],
			PublicationOrder:             parseInt(record[25]),
			RelatedCode:                  record[26],
			UpdateDate:                   record[27],
			TransitionalDate:             record[28],
			DiscontinuedDate:             record[29],
			AppendixNumber:               record[30],
			SectionNumber:                record[31],
			DPCApplicabilityCategory:     record[32],
			Reserved34:                   record[33],
			Reserved35:                   record[34],
			Reserved36:                   record[35],
			BasicKanjiName:               record[36],
			RemanufacturedDeviceCategory: record[37],
		}
	})
}
