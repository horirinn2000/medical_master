package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportTeeth(db *gorm.DB) {
	importCsv(db, 7, "csv/tooth/f_20111003.csv", func(record []string) model.Tooth {
		return model.Tooth{
			UpdateCategory:   record[0],
			MasterType:       record[1],
			Code:             record[2],
			Reserved:         record[3],
			Name:             record[4],
			UpdateDate:       record[5],
			DiscontinuedDate: record[6],
		}
	})
}
