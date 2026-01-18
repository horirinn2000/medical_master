package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportMedicalPractices(db *gorm.DB) {
	importCsv(db, 117, "csv/medical_practice/s_20251201.csv", func(record []string) model.MedicalPractice {
		return model.MedicalPractice{
			ChangeCategory:                 record[0],
			MasterType:                     record[1],
			MedicalPracticeCode:            record[2],
			AbbreviatedKanjiNameLength:     parseInt(record[3]),
			AbbreviatedKanjiName:           record[4],
			AbbreviatedKanaNameLength:      parseInt(record[5]),
			AbbreviatedKanaName:            record[6],
			DataStandardCode:               record[7],
			KanjiUnitNameLength:            parseInt(record[8]),
			KanjiUnitName:                  record[9],
			ScoreType:                      record[10],
			Score:                          parseFloat(record[11]),
			InpatientOutpatientCategory:    record[12],
			ElderlyMedicalCategory:         record[13],
			AggregationCategoryOutpatient:  record[14],
			ComprehensiveTargetTest:        record[15],
			DPCApplicabilityCategory:       record[17],
			HospitalClinicCategory:         record[18],
			SurgerySupportCategory:         record[19],
			MedicalObservationCategory:     record[20],
			NursingAddition:                record[21],
			AnesthesiaCategory:             record[22],
			DiseaseRelatedCategory:         record[24],
			ActualDaysCategory:             record[26],
			DaysTimesCategory:              record[27],
			MedicineRelatedCategory:        record[28],
			StepCalculationCategory:        record[29],
			StepLowerLimit:                 parseFloat(record[30]),
			StepUpperLimit:                 parseFloat(record[31]),
			StepValue:                      parseFloat(record[32]),
			StepScore:                      parseFloat(record[33]),
			StepErrorCategory:              record[34],
			MaxTimes:                       parseInt(record[35]),
			MaxTimesErrorCategory:          record[36],
			NoteAdditionCode:               record[37],
			NoteAdditionSequence:           record[38],
			GeneralAgeCategory:             record[39],
			LowerAge:                       record[40],
			UpperAge:                       record[41],
			TimeAdditionCategory:           record[42],
			FacilityStandardCategory:       record[43],
			FacilityStandardCode:           record[44],
			TreatmentInfantAddition:        record[45],
			VeryLowBirthWeightAddition:     record[46],
			InpatientReductionCategory:     record[47],
			DonorCollectionCategory:        record[48],
			TestJudgmentCategory:           record[49],
			TestJudgmentGroupCategory:      record[50],
			ReductionApplicability:         record[51],
			SpinalEvokedPotentialCategory:  record[52],
			NeckDissectionCategory:         record[53],
			AutoSutureCategory:             record[54],
			OutpatientManagementAddition:   record[55],
			OldScoreType:                   record[56],
			OldScore:                       parseFloat(record[57]),
			TestCommentCategory:            record[60],
			GeneralAdditionApplicability:   record[61],
			ComprehensiveReductionCategory: record[62],
			EndoscopicSurgeryAddition:      record[63],
			AggregationCategoryInpatient:   record[65],
			AutoAnastomosisCategory:        record[66],
			NotificationCategory1:          record[67],
			NotificationCategory2:          record[68],
			RegionAddition:                 record[69],
			BedCountCategory:               record[70],
			FacilityStandard1:              record[71],
			FacilityStandard2:              record[72],
			FacilityStandard3:              record[73],
			FacilityStandard4:              record[74],
			FacilityStandard5:              record[75],
			FacilityStandard6:              record[76],
			FacilityStandard7:              record[77],
			FacilityStandard8:              record[78],
			FacilityStandard9:              record[79],
			FacilityStandard10:             record[80],
			UltrasoundCoagulationCategory:  record[81],
			ShortStaySurgery:               record[82],
			DentalApplicability:            record[83],
			SectionCodeAlpha:               record[84],
			NotificationRelatedAlpha:       record[85],
			UpdateDate:                     record[86],
			ExpiryDate:                     record[87],
			PublicationOrder:               record[88],
			SectionChapter:                 record[89],
			SectionPart:                    record[90],
			SectionNumber:                  record[91],
			SectionSubNumber:               record[92],
			SectionItemNumber:              record[93],
			BasicKanjiName:                 record[112],
			PointTableNumber:               record[116],
		}
	})
}

func ImportMedicalPracticeSupports(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/01補助マスターテーブル.csv", func(record []string) model.MedicalPracticeSupport {
		return model.MedicalPracticeSupport{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			SupportInfo:         record[3],
			UpdateDate:          record[25],
			ExpiryDate:          record[26],
		}
	})
}

func ImportMedicalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/02包括テーブル.csv", func(record []string) model.MedicalPracticeInclusion {
		return model.MedicalPracticeInclusion{
			ChangeCategory:            record[0],
			ComprehensivePracticeCode: record[1],
			IncludedPracticeCode:      record[2],
			UpdateDate:                record[5],
			ExpiryDate:                record[6],
		}
	})
}

func ImportMedicalPracticeConflicts(db *gorm.DB) {
	filePaths := []string{
		"csv/medical_practice/03-1背反テーブル1.csv",
		"csv/medical_practice/03-2背反テーブル2.csv",
		"csv/medical_practice/03-3背反テーブル3.csv",
		"csv/medical_practice/03-4背反テーブル4.csv",
	}

	for _, filePath := range filePaths {
		importCsv(db, 7, filePath, func(record []string) model.MedicalPracticeConflict {
			return model.MedicalPracticeConflict{
				ChangeCategory:       record[0],
				MedicalPracticeCode:  record[1],
				ConflictPracticeCode: record[3],
				UpdateDate:           record[8],
				ExpiryDate:           record[9],
			}
		})
	}
}

func ImportInpatientBasicFees(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/04入院基本料テーブル.csv", func(record []string) model.InpatientBasicFee {
		return model.InpatientBasicFee{
			ChangeCategory:      record[0],
			BasicFeeCode:        record[1],
			MedicalPracticeCode: record[2],
			UpdateDate:          record[6],
			ExpiryDate:          record[7],
		}
	})
}

func ImportCalculationCounts(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/05算定回数テーブル.csv", func(record []string) model.CalculationCount {
		return model.CalculationCount{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			CountLimit:          parseInt(record[5]),
			UpdateDate:          record[12],
			ExpiryDate:          record[13],
		}
	})
}
