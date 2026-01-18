package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportWards(db *gorm.DB) {
	importCsv(db, 113, "csv/ward/k.csv", func(record []string) model.Ward {
		return model.Ward{
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
	})
}
