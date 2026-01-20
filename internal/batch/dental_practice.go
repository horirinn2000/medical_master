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
				ChangeCategory:       record[0],
				SectionNumber:        record[1],
				AdditionCode:         record[2],
				MedicalPracticeCode:  record[3],
				AbbreviatedKanjiName: record[4],
				AdditionCategory:     record[5],
				UpdateDate:           record[6],
				ExpiryDate:           record[7],
				SearchPriority:       record[8],
			}
		})
	}
}

func ImportDentalPracticeCalculationCountsMaster(db *gorm.DB) {
	importCsv(db, 15, "csv/dental_practice/h-6_20250307.csv", func(record []string) model.DentalPracticeCalculationCount {
		return model.DentalPracticeCalculationCount{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			Column3:             record[2],
			Column4:             record[3],
			Column5:             record[4],
			Column6:             record[5],
			Column7:             record[6],
			Name1:               record[7],
			Name2:               record[8],
			CountLimit:          parseInt(record[9]),
			Column11:            record[10],
			Column12:            record[11],
			UpdateDate:          record[12],
			ExpiryDate:          record[13],
			Column15:            record[14],
		}
	})
}

func ImportDentalPracticeSteps(db *gorm.DB) {
	importCsv(db, 20, "csv/dental_practice/h-7_20251201.csv", func(record []string) model.DentalPracticeStep {
		return model.DentalPracticeStep{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			Column3:             record[2],
			Column4:             record[3],
			Column5:             record[4],
			Column6:             record[5],
			Column7:             record[6],
			Name1:               record[7],
			Name2:               record[8],
			Column10:            record[9],
			StepLimit:           parseFloat(record[10]),
			Column12:            record[11],
			Column13:            record[12],
			Column14:            record[13],
			Column15:            record[14],
			StepScore:           parseFloat(record[15]),
			Column17:            record[16],
			UpdateDate:          record[17],
			ExpiryDate:          record[18],
			Column20:            record[19],
		}
	})
}

func ImportDentalPracticeAgeConstraints(db *gorm.DB) {
	importCsv(db, 14, "csv/dental_practice/h-8_20240531.csv", func(record []string) model.DentalPracticeAgeConstraint {
		return model.DentalPracticeAgeConstraint{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			Column3:             record[2],
			Column4:             record[3],
			Column5:             record[4],
			Column6:             record[5],
			Column7:             record[6],
			Name1:               record[7],
			Name2:               record[8],
			LowerAge:            record[9],
			UpperAge:            record[10],
			UpdateDate:          record[11],
			ExpiryDate:          record[12],
			Column14:            record[13],
		}
	})
}

func ImportDentalPracticeConflictsMaster(db *gorm.DB) {
	importCsv(db, 104, "csv/dental_practice/h-9_20240628.csv", func(record []string) model.DentalPracticeConflict {
		return model.DentalPracticeConflict{
			ChangeCategory: record[0],
			SourceCode:     record[1],
			SourceType:     record[2],
			SourceSection:  record[3],
			Column5:        record[4],
			Column6:        record[5],
			Column7:        record[6],
			SourceName1:    record[7],
			SourceName2:    record[8],
			ConflictFlag:   record[9],
			TargetCode:     record[10],
			TargetType:     record[11],
			TargetSection:  record[12],
			Column14:       record[13],
			Column15:       record[14],
			Column16:       record[15],
			TargetName1:    record[16],
			TargetName2:    record[17],
			Column19:       record[18],
			Column20:       record[19],
			Column21:       record[20],
			Column22:       record[21],
			Column23:       record[22],
			Column24:       record[23],
			Column25:       record[24],
			Column26:       record[25],
			Column27:       record[26],
			Column28:       record[27],
			Column29:       record[28],
			Column30:       record[29],
			Column31:       record[30],
			Column32:       record[31],
			Column33:       record[32],
			Column34:       record[33],
			Column35:       record[34],
			Column36:       record[35],
			Column37:       record[36],
			Column38:       record[37],
			Column39:       record[38],
			Column40:       record[39],
			Column41:       record[40],
			Column42:       record[41],
			Column43:       record[42],
			Column44:       record[43],
			Column45:       record[44],
			Column46:       record[45],
			Column47:       record[46],
			Column48:       record[47],
			Column49:       record[48],
			Column50:       record[49],
			Column51:       record[50],
			Column52:       record[51],
			Column53:       record[52],
			Column54:       record[53],
			Column55:       record[54],
			Column56:       record[55],
			Column57:       record[56],
			Column58:       record[57],
			Column59:       record[58],
			Column60:       record[59],
			Column61:       record[60],
			Column62:       record[61],
			Column63:       record[62],
			Column64:       record[63],
			Column65:       record[64],
			Column66:       record[65],
			Column67:       record[66],
			Column68:       record[67],
			Column69:       record[68],
			Column70:       record[69],
			Column71:       record[70],
			Column72:       record[71],
			Column73:       record[72],
			Column74:       record[73],
			Column75:       record[74],
			Column76:       record[75],
			Column77:       record[76],
			Column78:       record[77],
			Column79:       record[78],
			Column80:       record[79],
			Column81:       record[80],
			Column82:       record[81],
			Column83:       record[82],
			Column84:       record[83],
			Column85:       record[84],
			Column86:       record[85],
			Column87:       record[86],
			Column88:       record[87],
			Column89:       record[88],
			Column90:       record[89],
			Column91:       record[90],
			Column92:       record[91],
			Column93:       record[92],
			Column94:       record[93],
			Column95:       record[94],
			Column96:       record[95],
			Column97:       record[96],
			Column98:       record[97],
			Column99:       record[98],
			UpdateDate:     record[99],
			ExpiryDate:     record[100],
			Column102:      record[101],
			Column103:      record[102],
			Column104:      record[103],
		}
	})
}

func ImportDentalPracticeActualDays(db *gorm.DB) {
	importCsv(db, 14, "csv/dental_practice/h-10_20240726.csv", func(record []string) model.DentalPracticeActualDays {
		return model.DentalPracticeActualDays{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			Column3:             record[2],
			Column4:             record[3],
			Column5:             record[4],
			Column6:             record[5],
			Column7:             record[6],
			Name1:               record[7],
			Name2:               record[8],
			DaysLimit:           record[9],
			Column11:            record[10],
			UpdateDate:          record[11],
			ExpiryDate:          record[12],
			Column14:            record[13],
		}
	})
}
