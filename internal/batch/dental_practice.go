package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDentalPractices(db *gorm.DB) {
	// 全66項目チェック (R06rec3.pdf 205-206ページ)
	importCsv(db, 66, "csv/dental_practice/h_20251201.csv", func(record []string) model.DentalPractice {
		return model.DentalPractice{
			ChangeCategory:                 record[0],
			MasterType:                     record[1],
			DentalPracticeCode:             record[2],
			SectionCode:                    record[3],
			SectionNumber:                  record[4],
			SectionSubNumber:               record[5],
			SectionItemNumber:              record[6],
			AdditionCode:                   record[7],
			BasicKanjiName:                 record[8],
			AbbreviatedKanjiName:           record[9],
			ScoreType:                      record[10],
			Score:                          parseFloat(record[11]),
			OldScoreType:                   record[12],
			OldScore:                       parseFloat(record[13]),
			InpatientOutpatientCategory:    record[14],
			ElderlyMedicalCategory:         record[15],
			TimeAdditionCategory:           record[16],
			HospitalClinicCategory:         record[17],
			NursingAddition:                record[18],
			Reserved20:                     record[19],
			Reserved21:                     record[20],
			RegionAddition:                 record[21],
			DiseaseRelatedCategory:         record[22],
			MedicineRelatedCategory:        record[23],
			BedCountCategory:               record[24],
			NotificationCategory:           record[25],
			FutureVisitCategory:            record[26],
			ShortStaySurgery:               record[27],
			SpecialNoteCategory:            record[28],
			TestJudgmentCategory:           record[29],
			TestJudgmentGroupCategory:      record[30],
			ReductionApplicability:         record[31],
			ComprehensiveReductionCategory: record[32],
			FacilityStandardCategory:       record[33],
			FacilityStandardCode:           record[34],
			FacilityStandard1:              record[35],
			FacilityStandard2:              record[36],
			FacilityStandard3:              record[37],
			FacilityStandard4:              record[38],
			FacilityStandard5:              record[39],
			FacilityStandard6:              record[40],
			FacilityStandard7:              record[41],
			FacilityStandard8:              record[42],
			FacilityStandard9:              record[43],
			FacilityStandard10:             record[44],
			GeneralAdditionGroup:           record[45],
			BasicAdditionGroup:             record[46],
			NoteAdditionGroup:              record[47],
			ProcedureMaterialGroup:         record[48],
			CalculationCountTableRelated:   record[49],
			StepTableRelated:               record[50],
			AgeLimitTableRelated:           record[51],
			ConflictTableRelated:           record[52],
			ActualDaysTableRelated:         record[53],
			Reserved55:                     record[54],
			Reserved56:                     record[55],
			UpdateDate:                     record[56],
			ExpiryDate:                     record[57],
			LongTimeAnesthesia:             record[58],
			MalignantTumorPathology:        record[59],
			NursingStaffTreatmentImprove:   record[60],
			DentalBaseUpEvaluation1:        record[61],
			DentalBaseUpEvaluation2:        record[62],
			Reserved64:                     record[63],
			Reserved65:                     record[64],
			PublicationOrder:               record[65],
		}
	})
}

func ImportDentalPracticeSupports(db *gorm.DB) {
	importCsv(db, 25, "csv/dental_practice/01補助マスターテーブル（歯科）.csv", func(record []string) model.DentalPracticeSupport {
		return model.DentalPracticeSupport{
			ChangeCategory:           record[0],
			DentalPracticeCode:       record[1],
			AdditionCode:             record[2],
			AbbreviatedKanjiName:     record[3],
			ComprehensiveUnit1:       record[4],
			ComprehensiveGroup1:      record[5],
			ComprehensiveUnit2:       record[6],
			ComprehensiveGroup2:      record[7],
			ComprehensiveUnit3:       record[8],
			ComprehensiveGroup3:      record[9],
			ConflictDaily:            record[10],
			ConflictMonthly:          record[11],
			ConflictSimultaneous:     record[12],
			ConflictPartSimultaneous: record[13],
			ConflictWeekly:           record[14],
			Reserved16:               record[15],
			CalculationCountRelated:  record[16],
			Reserved18:               record[17],
			Reserved19:               record[18],
			Reserved20:               record[19],
			Reserved21:               record[20],
			Reserved22:               record[21],
			Reserved23:               record[22],
			UpdateDate:               record[23],
			ExpiryDate:               record[24],
		}
	})
}

func ImportDentalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 8, "csv/dental_practice/02包括テーブル（歯科）.csv", func(record []string) model.DentalPracticeInclusion {
		return model.DentalPracticeInclusion{
			ChangeCategory:       record[0],
			GroupNumber:          record[1],
			DentalPracticeCode:   record[2],
			AdditionCode:         record[3],
			AbbreviatedKanjiName: record[4],
			SpecialCondition:     record[5],
			UpdateDate:           record[6],
			ExpiryDate:           record[7],
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
		importCsv(db, 12, filePath, func(record []string) model.DentalPracticeConflict {
			return model.DentalPracticeConflict{
				ChangeCategory:        record[0],
				DentalPracticeCode1:   record[1],
				AdditionCode1:         record[2],
				AbbreviatedKanjiName1: record[3],
				DentalPracticeCode2:   record[4],
				AdditionCode2:         record[5],
				AbbreviatedKanjiName2: record[6],
				ConflictCategory:      record[7],
				SpecialCondition:      record[8],
				Reserved10:            record[9],
				UpdateDate:            record[10],
				ExpiryDate:            record[11],
			}
		})
	}
}

func ImportDentalPracticeCalculationCounts(db *gorm.DB) {
	importCsv(db, 12, "csv/dental_practice/04算定回数テーブル（歯科）.csv", func(record []string) model.DentalPracticeCalculationCount {
		return model.DentalPracticeCalculationCount{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			AdditionCode:         record[2],
			AbbreviatedKanjiName: record[3],
			RequirementCode:      record[4],
			UnitCode:             record[5],
			UnitName:             record[6],
			MaxTimes:             parseInt(record[7]),
			SpecialCondition:     record[8],
			Reserved10:           record[9],
			UpdateDate:           record[10],
			ExpiryDate:           record[11],
		}
	})
}
