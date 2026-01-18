package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDentalPractices(db *gorm.DB) {
	importCsv(db, 66, "csv/dental_practice/h_20251201.csv", func(record []string) model.DentalPractice {
		return model.DentalPractice{
			ChangeCategory:                record[0],
			MasterType:                    record[1],
			MedicalPracticeCode:           record[2],
			AbbreviatedKanjiNameLength:    parseInt(record[3]),
			AbbreviatedKanjiName:          record[4],
			AbbreviatedKanaNameLength:     parseInt(record[5]),
			AbbreviatedKanaName:           record[6],
			DataStandardCode:              record[7],
			KanjiUnitNameLength:           parseInt(record[8]),
			KanjiUnitName:                 record[9],
			ScoreType:                     record[10],
			Score:                         parseFloat(record[11]),
			InpatientOutpatientCategory:   record[12],
			ElderlyMedicalCategory:        record[13],
			AggregationCategoryOutpatient: record[14],
			ComprehensiveTargetTest:       record[15],
			DPCApplicabilityCategory:      record[17],
			HospitalClinicCategory:        record[18],
			SurgerySupportCategory:        record[19],
			MedicalObservationCategory:    record[20],
			NursingAddition:               record[21],
			AnesthesiaCategory:            record[22],
			DiseaseRelatedCategory:        record[24],
			ActualDaysCategory:            record[26],
			DaysTimesCategory:             record[27],
			MedicineRelatedCategory:       record[28],
			StepCalculationCategory:       record[29],
			StepLowerLimit:                parseFloat(record[30]),
			StepUpperLimit:                parseFloat(record[31]),
			StepValue:                     parseFloat(record[32]),
			StepScore:                     parseFloat(record[33]),
			StepErrorCategory:             record[34],
			MaxTimes:                      parseInt(record[35]),
			MaxTimesErrorCategory:         record[36],
			NoteAdditionCode:              record[37],
			NoteAdditionSequence:          record[38],
			GeneralAgeCategory:            record[39],
			LowerAge:                      record[40],
			UpperAge:                      record[41],
			TimeAdditionCategory:          record[42],
			FacilityStandardCategory:      record[43],
			FacilityStandardCode:          record[44],
			TreatmentInfantAddition:       record[45],
			VeryLowBirthWeightAddition:    record[46],
			InpatientReductionCategory:    record[47],
			DonorCollectionCategory:       record[48],
			TestJudgmentCategory:          record[49],
			TestJudgmentGroupCategory:     record[50],
			ReductionApplicability:        record[51],
			SpinalEvokedPotentialCategory: record[52],
			NeckDissectionCategory:        record[53],
			AutoSutureCategory:            record[54],
			OutpatientManagementAddition:  record[55],
			UpdateDate:                    record[56],
			ExpiryDate:                    record[57],
			PublicationOrder:              record[58],
			SectionChapter:                record[59],
			SectionPart:                   record[60],
			SectionNumber:                 record[61],
			SectionSubNumber:              record[62],
			SectionItemNumber:             record[63],
			AggregationCategoryInpatient:  record[65],
		}
	})
}

func ImportDentalPracticeSupports(db *gorm.DB) {
	importCsv(db, 25, "csv/dental_practice/01補助マスターテーブル（歯科）.csv", func(record []string) model.DentalPracticeSupport {
		return model.DentalPracticeSupport{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			SupportCode:         record[2],
			SupportInfo:         record[3],
			UpdateDate:          record[23],
			ExpiryDate:          record[24],
		}
	})
}

func ImportDentalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 25, "csv/dental_practice/02包括テーブル（歯科）.csv", func(record []string) model.DentalPracticeInclusion {
		return model.DentalPracticeInclusion{
			ChangeCategory:            record[0],
			ComprehensivePracticeCode: record[1],
			IncludedPracticeCode:      record[2],
			UpdateDate:                record[6],
			ExpiryDate:                record[7],
		}
	})
}

func ImportDentalPracticeConflicts(db *gorm.DB) {
	filePaths := []string{
		"csv/dental_practice/03-1背反テーブル1（歯科）.csv",
		"csv/dental_practice/03-2背反テーブル2（歯科）.csv",
		"csv/dental_practice/03-3背反テーブル3（歯科）.csv",
		"csv/dental_practice/03-4背反テーブル4（歯科）.csv",
		"csv/dental_practice/03-5背反テーブル5（歯科）.csv",
	}

	for _, filePath := range filePaths {
		importCsv(db, 12, filePath, func(record []string) model.DentalPracticeConflictDetail {
			return model.DentalPracticeConflictDetail{
				ChangeCategory:       record[0],
				MedicalPracticeCode:  record[1],
				ConflictPracticeCode: record[3],
				UpdateDate:           record[10],
				ExpiryDate:           record[11],
			}
		})
	}
}

func ImportDentalPracticeCalculationCounts(db *gorm.DB) {
	importCsv(db, 12, "csv/dental_practice/04算定回数テーブル（歯科）.csv", func(record []string) model.DentalPracticeCalculationCountLimit {
		return model.DentalPracticeCalculationCountLimit{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			CountLimit:          parseInt(record[7]),
			UpdateDate:          record[10],
			ExpiryDate:          record[11],
		}
	})
}
