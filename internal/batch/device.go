package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDevices(db *gorm.DB) {
	importCsv(db, 38, "csv/device/t_20251128.csv", func(record []string) model.SpecialMedicalDevice {
		return model.SpecialMedicalDevice{
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
	})
}
