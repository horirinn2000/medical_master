package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportDentalPractices(db *gorm.DB) {
	importCsv(db, 66, "csv/dental_practice/h_20251201.csv", func(record []string) model.DentalPractice {
		return model.DentalPractice{
			ChangeCategory:                  record[0],
			MasterType:                      record[1],
			DentalPracticeCode:              record[2],
			SectionChar:                     record[3],
			SectionNumber:                   record[4],
			SectionBranchNumber:             record[5],
			SectionItemNumber:               record[6],
			AdditionCode:                    record[7],
			BasicKanjiName:                  record[8],
			AbbreviatedKanjiName:            record[9],
			ScoreType:                       record[10],
			Score:                           parseFloat(record[11]),
			OldScoreType:                    record[12],
			OldScore:                        parseFloat(record[13]),
			InpatientOutpatientCategory:     record[14],
			ElderlyMedicalCategory:          record[15],
			TimeAdditionCategory:            record[16],
			HospitalClinicCategory:          record[17],
			NursingAddition:                 record[18],
			Reserved1:                       record[19],
			Reserved2:                       record[20],
			RegionAddition:                  record[21],
			DiseaseRelatedCategory:          record[22],
			MedicineRelatedCategory:         record[23],
			BedCountCategory:                record[24],
			NotificationCategory:            record[25],
			FutureVisitCategory:             record[26],
			ShortStaySurgery:                record[27],
			SpecialNote:                     record[28],
			TestJudgmentCategory:            record[29],
			TestJudgmentGroupCategory:       record[30],
			ReductionTargetCategory:         record[31],
			ComprehensiveReductionCategory:  record[32],
			ConformityCategory:              record[33],
			TargetFacilityStandard:          record[34],
			FacilityStandardCode1:           record[35],
			FacilityStandardCode2:           record[36],
			FacilityStandardCode3:           record[37],
			FacilityStandardCode4:           record[38],
			FacilityStandardCode5:           record[39],
			FacilityStandardCode6:           record[40],
			FacilityStandardCode7:           record[41],
			FacilityStandardCode8:           record[42],
			FacilityStandardCode9:           record[43],
			FacilityStandardCode10:          record[44],
			GeneralAdditionGroup:            record[45],
			BasicAdditionGroup:              record[46],
			NoteAdditionGroup:               record[47],
			TechniqueMaterialAdditionGroup:  record[48],
			CountLimitTableId:               record[49],
			StepTableId:                     record[50],
			AgeLimitTableId:                 record[51],
			ConflictTableId:                 record[52],
			ActualDaysTableId:               record[53],
			Reserved3:                       record[54],
			Reserved4:                       record[55],
			UpdateDate:                      record[56],
			ExpiryDate:                      record[57],
			LongTermAnesthesiaAddition:      record[58],
			MalignantTumorPathologyAddition: record[59],
			NursingStaffTreatmentEvaluation: record[60],
			DentalOutpatientBaseUp1:         record[61],
			DentalOutpatientBaseUp2:         record[62],
			Reserved5:                       record[63],
			Reserved6:                       record[64],
			PublicationOrder:                record[65],
		}
	})
}

func ImportDentalPracticeSupports(db *gorm.DB) {
	importCsv(db, 25, "csv/dental_practice/01補助マスターテーブル（歯科）.csv", func(record []string) model.DentalPracticeSupport {
		return model.DentalPracticeSupport{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			SupportCode:         record[2],
			Name:                record[3],
			Flag1:               record[4],
			Flag2:               record[5],
			Flag3:               record[6],
			Flag4:               record[7],
			Flag5:               record[8],
			Flag6:               record[9],
			Flag7:               record[10],
			Flag8:               record[11],
			Flag9:               record[12],
			Flag10:              record[13],
			Flag11:              record[14],
			Flag12:              record[15],
			Flag13:              record[16],
			Flag14:              record[17],
			Flag15:              record[18],
			Flag16:              record[19],
			Flag17:              record[20],
			Flag18:              record[21],
			Flag19:              record[22],
			UpdateDate:          record[23],
			ExpiryDate:          record[24],
		}
	})
}

func ImportDentalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 8, "csv/dental_practice/02包括テーブル（歯科）.csv", func(record []string) model.DentalPracticeInclusion {
		return model.DentalPracticeInclusion{
			ChangeCategory:       record[0],
			GroupNumber:          record[1],
			IncludedPracticeCode: record[2],
			AdditionCode:         record[3],
			Name:                 record[4],
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
		importCsv(db, 12, filePath, func(record []string) model.DentalPracticeConflictDetail {
			return model.DentalPracticeConflictDetail{
				ChangeCategory:       record[0],
				MedicalPracticeCode1: record[1],
				AdditionCode1:        record[2],
				Name1:                record[3],
				MedicalPracticeCode2: record[4],
				AdditionCode2:        record[5],
				Name2:                record[6],
				ConflictCategory:     record[7],
				SpecialCondition:     record[8],
				Reserved:             record[9],
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
			AdditionCode:        record[2],
			Name:                record[3],
			UnitCode:            record[4],
			UnitName:            record[5],
			CountLimit:          parseInt(record[6]),
			SpecialCondition:    record[7],
			Reserved1:           record[8],
			Reserved2:           record[9],
			UpdateDate:          record[10],
			ExpiryDate:          record[11],
		}
	})
}

func ImportDentalPracticeAdditionRelations(db *gorm.DB) {
	filePaths := []string{
		"csv/dental_practice/h-2_20240628.csv",
		"csv/dental_practice/h-3_20250901.csv",
		"csv/dental_practice/h-4_20250307.csv",
		"csv/dental_practice/h-5_20251201.csv",
	}

	for _, filePath := range filePaths {
		importCsv(db, 9, filePath, func(record []string) model.DentalPracticeAdditionRelation {
			return model.DentalPracticeAdditionRelation{
				ChangeCategory:     record[0],
				GroupNumber:        record[1],
				AdditionCode:       record[2],
				DentalPracticeCode: record[3],
				BasicKanjiName:     record[4],
				AdditionIdentifier: record[5],
				UpdateDate:         record[6],
				ExpiryDate:         record[7],
				Reserved:           record[8],
			}
		})
	}
}

func ImportDentalPracticeCalculationCountsMaster(db *gorm.DB) {
	importCsv(db, 15, "csv/dental_practice/h-6_20250307.csv", func(record []string) model.DentalPracticeCalculationCount {
		return model.DentalPracticeCalculationCount{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			SectionChar:          record[2],
			SectionNumber:        record[3],
			SectionBranchNumber:  record[4],
			SectionItemNumber:    record[5],
			AdditionCode:         record[6],
			BasicKanjiName:       record[7],
			AbbreviatedKanjiName: record[8],
			UnitCode:             record[9],
			CountLimit:           parseInt(record[10]),
			CountErrorTreatment:  record[11],
			UpdateDate:           record[12],
			ExpiryDate:           record[13],
			Reserved:             record[14],
		}
	})
}

func ImportDentalPracticeSteps(db *gorm.DB) {
	importCsv(db, 20, "csv/dental_practice/h-7_20251201.csv", func(record []string) model.DentalPracticeStep {
		return model.DentalPracticeStep{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			SectionChar:          record[2],
			SectionNumber:        record[3],
			SectionBranchNumber:  record[4],
			SectionItemNumber:    record[5],
			AdditionCode:         record[6],
			BasicKanjiName:       record[7],
			AbbreviatedKanjiName: record[8],
			ScoreType:            record[9],
			Score:                parseFloat(record[10]),
			StepUnit:             record[11],
			StepLowerLimit:       parseFloat(record[12]),
			StepUpperLimit:       parseFloat(record[13]),
			StepValue:            parseFloat(record[14]),
			StepScore:            parseFloat(record[15]),
			StepErrorTreatment:   record[16],
			UpdateDate:           record[17],
			ExpiryDate:           record[18],
			Reserved:             record[19],
		}
	})
}

func ImportDentalPracticeAgeConstraints(db *gorm.DB) {
	importCsv(db, 14, "csv/dental_practice/h-8_20240531.csv", func(record []string) model.DentalPracticeAgeConstraint {
		return model.DentalPracticeAgeConstraint{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			SectionChar:          record[2],
			SectionNumber:        record[3],
			SectionBranchNumber:  record[4],
			SectionItemNumber:    record[5],
			AdditionCode:         record[6],
			BasicKanjiName:       record[7],
			AbbreviatedKanjiName: record[8],
			LowerAge:             record[9],
			UpperAge:             record[10],
			UpdateDate:           record[11],
			ExpiryDate:           record[12],
			Reserved:             record[13],
		}
	})
}

func ImportDentalPracticeConflictsMaster(db *gorm.DB) {
	importCsv(db, 104, "csv/dental_practice/h-9_20240628.csv", func(record []string) model.DentalPracticeConflict {
		return model.DentalPracticeConflict{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			SectionChar:          record[2],
			SectionNumber:        record[3],
			SectionBranchNumber:  record[4],
			SectionItemNumber:    record[5],
			AdditionCode:         record[6],
			BasicKanjiName:       record[7],
			AbbreviatedKanjiName: record[8],
			Conflict1_Flag:       record[9],
			Conflict1_Code:       record[10],
			Conflict1_Section:    record[11],
			Conflict1_SecNum:     record[12],
			Conflict1_Branch:     record[13],
			Conflict1_Item:       record[14],
			Conflict1_AddCode:    record[15],
			Conflict1_Name:       record[16],
			Conflict1_Abbr:       record[17],
			Conflict2_Flag:       record[18],
			Conflict2_Code:       record[19],
			Conflict2_Section:    record[20],
			Conflict2_SecNum:     record[21],
			Conflict2_Branch:     record[22],
			Conflict2_Item:       record[23],
			Conflict2_AddCode:    record[24],
			Conflict2_Name:       record[25],
			Conflict2_Abbr:       record[26],
			Conflict3_Flag:       record[27],
			Conflict3_Code:       record[28],
			Conflict3_Section:    record[29],
			Conflict3_SecNum:     record[30],
			Conflict3_Branch:     record[31],
			Conflict3_Item:       record[32],
			Conflict3_AddCode:    record[33],
			Conflict3_Name:       record[34],
			Conflict3_Abbr:       record[35],
			Conflict4_Flag:       record[36],
			Conflict4_Code:       record[37],
			Conflict4_Section:    record[38],
			Conflict4_SecNum:     record[39],
			Conflict4_Branch:     record[40],
			Conflict4_Item:       record[41],
			Conflict4_AddCode:    record[42],
			Conflict4_Name:       record[43],
			Conflict4_Abbr:       record[44],
			Conflict5_Flag:       record[45],
			Conflict5_Code:       record[46],
			Conflict5_Section:    record[47],
			Conflict5_SecNum:     record[48],
			Conflict5_Branch:     record[49],
			Conflict5_Item:       record[50],
			Conflict5_AddCode:    record[51],
			Conflict5_Name:       record[52],
			Conflict5_Abbr:       record[53],
			Conflict6_Flag:       record[54],
			Conflict6_Code:       record[55],
			Conflict6_Section:    record[56],
			Conflict6_SecNum:     record[57],
			Conflict6_Branch:     record[58],
			Conflict6_Item:       record[59],
			Conflict6_AddCode:    record[60],
			Conflict6_Name:       record[61],
			Conflict6_Abbr:       record[62],
			Conflict7_Flag:       record[63],
			Conflict7_Code:       record[64],
			Conflict7_Section:    record[65],
			Conflict7_SecNum:     record[66],
			Conflict7_Branch:     record[67],
			Conflict7_Item:       record[68],
			Conflict7_AddCode:    record[69],
			Conflict7_Name:       record[70],
			Conflict7_Abbr:       record[71],
			Conflict8_Flag:       record[72],
			Conflict8_Code:       record[73],
			Conflict8_Section:    record[74],
			Conflict8_SecNum:     record[75],
			Conflict8_Branch:     record[76],
			Conflict8_Item:       record[77],
			Conflict8_AddCode:    record[78],
			Conflict8_Name:       record[79],
			Conflict8_Abbr:       record[80],
			Conflict9_Flag:       record[81],
			Conflict9_Code:       record[82],
			Conflict9_Section:    record[83],
			Conflict9_SecNum:     record[84],
			Conflict9_Branch:     record[85],
			Conflict9_Item:       record[86],
			Conflict9_AddCode:    record[87],
			Conflict9_Name:       record[88],
			Conflict9_Abbr:       record[89],
			Conflict10_Flag:      record[90],
			Conflict10_Code:      record[91],
			Conflict10_Section:   record[92],
			Conflict10_SecNum:    record[93],
			Conflict10_Branch:    record[94],
			Conflict10_Item:      record[95],
			Conflict10_AddCode:   record[96],
			Conflict10_Name:      record[97],
			Conflict10_Abbr:      record[98],
			UpdateDate:           record[99],
			ExpiryDate:           record[100],
			Reserved1:            record[101],
			Reserved2:            record[102],
			Reserved3:            record[103],
		}
	})
}

func ImportDentalPracticeActualDays(db *gorm.DB) {
	importCsv(db, 14, "csv/dental_practice/h-10_20240726.csv", func(record []string) model.DentalPracticeActualDays {
		return model.DentalPracticeActualDays{
			ChangeCategory:       record[0],
			DentalPracticeCode:   record[1],
			SectionChar:          record[2],
			SectionNumber:        record[3],
			SectionBranchNumber:  record[4],
			SectionItemNumber:    record[5],
			AdditionCode:         record[6],
			BasicKanjiName:       record[7],
			AbbreviatedKanjiName: record[8],
			ActualDays:           record[9],
			DaysTimes:            record[10],
			UpdateDate:           record[11],
			ExpiryDate:           record[12],
			Reserved:             record[13],
		}
	})
}
