package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDiseases(db *gorm.DB) {
	importCsv(db, 46, "csv/disease/b_20260101.txt", func(record []string) model.Disease {
		model := toModel(record)
		model.IsOld = false
		return model
	})

	importCsv(db, 46, "csv/disease/hb_20260101.txt", func(record []string) model.Disease {
		model := toModel(record)
		model.IsOld = true
		return model
	})
}

func toModel(record []string) model.Disease {
	return model.Disease{
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
}
