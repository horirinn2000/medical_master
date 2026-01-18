package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDiseases(db *gorm.DB) {
	// 全46項目チェック (R06rec3.pdf 192-193ページ)
	importCsv(db, 46, "csv/disease/b_20260101.txt", func(record []string) model.Disease {
		d := toModel(record)
		d.IsOld = false
		return d
	})

	importCsv(db, 46, "csv/disease/hb_20260101.txt", func(record []string) model.Disease {
		d := toModel(record)
		d.IsOld = true
		return d
	})
}

func toModel(record []string) model.Disease {
	return model.Disease{
		ChangeCategory:                record[0],
		MasterType:                    record[1],
		Code:                          record[2],
		MigrationCode:                 record[3],
		BasicNameLength:               parseInt(record[4]),
		BasicName:                     record[5],
		AbbreviatedNameLength:         parseInt(record[6]),
		AbbreviatedName:               record[7],
		KanaNameLength:                parseInt(record[8]),
		KanaName:                      record[9],
		ManagementNumber:              record[10],
		SelectionCategory:             record[11],
		ExchangeCode:                  record[12],
		Reserved14:                    record[13],
		Reserved15:                    record[14],
		ICD10_1:                       record[15],
		ICD10_2:                       record[16],
		Reserved18:                    record[17],
		SingleUseForbiddenCategory:    record[18],
		InsuranceOutCategory:          record[19],
		SpecificDiseaseCategory:       record[20],
		ListingDate:                   record[21],
		UpdateDate:                    record[22],
		ExpiryDate:                    record[23],
		BasicNameUpdateFlag:           record[24],
		AbbreviatedNameUpdateFlag:     record[25],
		KanaNameUpdateFlag:            record[26],
		SelectionCategoryUpdateFlag:   record[27],
		ExchangeCodeUpdateFlag:        record[28],
		Reserved30:                    record[29],
		Reserved31:                    record[30],
		DentalAbbreviatedUpdateFlag:   record[31],
		IncurableDiseaseUpdateFlag:    record[32],
		DentalSpecificUpdateFlag:      record[33],
		SingleUseForbiddenUpdateFlag:  record[34],
		InsuranceOutUpdateFlag:        record[35],
		SpecificDiseaseUpdateFlag:     record[36],
		MigrationManagementNumber:     record[37],
		DentalAbbreviatedName:         record[38],
		Reserved40:                    record[39],
		Reserved41:                    record[40],
		DentalAbbreviatedNameLength:   parseInt(record[41]),
		IncurableDiseaseCategory:      record[42],
		DentalSpecificDiseaseCategory: record[43],
		ICD10_1UpdateFlag:             record[44],
		ICD10_2UpdateFlag:             record[45],
	}
}
