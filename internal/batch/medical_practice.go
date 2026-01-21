package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportMedicalPractices(db *gorm.DB) {
	importCsv(db, 150, "csv/medical_practice/s_20251201.csv", func(record []string) model.MedicalPractice {
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
			Reserved1:                      record[16],
			DPCApplicabilityCategory:       record[17],
			HospitalClinicCategory:         record[18],
			SurgerySupportCategory:         record[19],
			MedicalObservationCategory:     record[20],
			NursingAddition:                record[21],
			AnesthesiaCategory:             record[22],
			Reserved2:                      record[23],
			DiseaseRelatedCategory:         record[24],
			Reserved3:                      record[25],
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
			KanjiNameChangeCategory:        record[58],
			KanaNameChangeCategory:         record[59],
			TestCommentCategory:            record[60],
			GeneralAdditionApplicability:   record[61],
			ComprehensiveReductionCategory: record[62],
			EndoscopicSurgeryAddition:      record[63],
			Reserved4:                      record[64],
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
			NotificationSectionChapter:     record[94],
			NotificationSectionPart:        record[95],
			NotificationSectionNumber:      record[96],
			NotificationSectionSubNumber:   record[97],
			NotificationSectionItemNumber:  record[98],
			AgeAddition1Lower:              record[99],
			AgeAddition1Upper:              record[100],
			AgeAddition1Code:               record[101],
			AgeAddition2Lower:              record[102],
			AgeAddition2Upper:              record[103],
			AgeAddition2Code:               record[104],
			AgeAddition3Lower:              record[105],
			AgeAddition3Upper:              record[106],
			AgeAddition3Code:               record[107],
			AgeAddition4Lower:              record[108],
			AgeAddition4Upper:              record[109],
			AgeAddition4Code:               record[110],
			MigrationRelated:               record[111],
			BasicKanjiName:                 record[112],
			SinusSurgeryEndoscopy:          record[113],
			SinusSurgeryBoneCutting:        record[114],
			LongTermAnesthesia:             record[115],
			PointTableNumber:               record[116],
			MonitoringAddition:             record[117],
			FrozenTissueAddition:           record[118],
			MalignantTumorAddition:         record[119],
			ExternalFixatorAddition:        record[120],
			UltrasoundCuttingAddition:      record[121],
			LeftAtrialAppendageClosure:     record[122],
			OutpatientInfectionControl:     record[123],
			EntInfantTreatment:             record[124],
			EntAntimicrobialSupport:        record[125],
			NegativePressureClosure:        record[126],
			NursingStaffTreatment:          record[127],
			OutpatientBaseUp1:              record[128],
			OutpatientBaseUp2:              record[129],
			RemanufacturedDevice:           record[130],
			Reserved5:                      record[131],
			Reserved6:                      record[132],
			Reserved7:                      record[133],
			Reserved8:                      record[134],
			Reserved9:                      record[135],
			Reserved10:                     record[136],
			Reserved11:                     record[137],
			Reserved12:                     record[138],
			Reserved13:                     record[139],
			Reserved14:                     record[140],
			Reserved15:                     record[141],
			Reserved16:                     record[142],
			Reserved17:                     record[143],
			Reserved18:                     record[144],
			Reserved19:                     record[145],
			Reserved20:                     record[146],
			Reserved21:                     record[147],
			Reserved22:                     record[148],
			Reserved23:                     record[149],
		}
	})
}

func ImportMedicalPracticeSupports(db *gorm.DB) {
	importCsv(db, 27, "csv/medical_practice/01補助マスターテーブル.csv", func(record []string) model.MedicalPracticeSupport {
		return model.MedicalPracticeSupport{
			ChangeCategory:            record[0],
			MedicalPracticeCode:       record[1],
			AbbreviatedKanjiName:      record[2],
			InclusionUnit1:            record[3],
			InclusionGroup1:           record[4],
			InclusionUnit2:            record[5],
			InclusionGroup2:           record[6],
			InclusionUnit3:            record[7],
			InclusionGroup3:           record[8],
			ConflictPerDay:            record[9],
			ConflictPerMonth:          record[10],
			ConflictSimultaneous:      record[11],
			ConflictPerWeek:           record[12],
			Reserved1:                 record[13],
			Reserved2:                 record[14],
			Reserved3:                 record[15],
			Reserved4:                 record[16],
			Reserved5:                 record[17],
			Reserved6:                 record[18],
			InpatientBasicFeeCategory: record[19],
			CalculationCountCategory:  record[20],
			Reserved7:                 record[21],
			Reserved8:                 record[22],
			Reserved9:                 record[23],
			Reserved10:                record[24],
			NewDate:                   record[25],
			ExpiryDate:                record[26],
		}
	})
}

func ImportMedicalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/02包括テーブル.csv", func(record []string) model.MedicalPracticeInclusion {
		return model.MedicalPracticeInclusion{
			ChangeCategory:       record[0],
			GroupNumber:          record[1],
			IncludedPracticeCode: record[2],
			IncludedPracticeName: record[3],
			SpecialCondition:     record[4],
			NewDate:              record[5],
			ExpiryDate:           record[6],
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
		importCsv(db, 10, filePath, func(record []string) model.MedicalPracticeConflict {
			return model.MedicalPracticeConflict{
				ChangeCategory:       record[0],
				MedicalPracticeCode1: record[1],
				PracticeName1:        record[2],
				MedicalPracticeCode2: record[3],
				PracticeName2:        record[4],
				ConflictCategory:     record[5],
				SpecialCondition:     record[6],
				Reserved1:            record[7],
				NewDate:              record[8],
				ExpiryDate:           record[9],
			}
		})
	}
}

func ImportInpatientBasicFees(db *gorm.DB) {
	importCsv(db, 8, "csv/medical_practice/04入院基本料テーブル.csv", func(record []string) model.InpatientBasicFee {
		return model.InpatientBasicFee{
			ChangeCategory:      record[0],
			GroupNumber:         record[1],
			MedicalPracticeCode: record[2],
			PracticeName:        record[3],
			AdditionIdentifier:  record[4],
			Reserved1:           record[5],
			NewDate:             record[6],
			ExpiryDate:          record[7],
		}
	})
}

func ImportCalculationCounts(db *gorm.DB) {
	importCsv(db, 14, "csv/medical_practice/05算定回数テーブル.csv", func(record []string) model.CalculationCount {
		return model.CalculationCount{
			ChangeCategory:      record[0],
			MedicalPracticeCode: record[1],
			PracticeName:        record[2],
			UnitCode:            record[3],
			UnitName:            record[4],
			CountLimit:          parseInt(record[5]),
			SpecialCondition:    record[6],
			Reserved1:           record[7],
			Reserved2:           record[8],
			Reserved3:           record[9],
			Reserved4:           record[10],
			Reserved5:           record[11],
			NewDate:             record[12],
			ExpiryDate:          record[13],
		}
	})
}
