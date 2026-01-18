package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportMedicalPractices(db *gorm.DB) {
	// 全150項目チェック
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
			Reserved17:                     record[16],
			DPCApplicabilityCategory:       record[17],
			HospitalClinicCategory:         record[18],
			SurgerySupportCategory:         record[19],
			MedicalObservationCategory:     record[20],
			NursingAddition:                record[21],
			AnesthesiaCategory:             record[22],
			Reserved24:                     record[23],
			DiseaseRelatedCategory:         record[24],
			Reserved26:                     record[25],
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
			KanjiNameUpdateFlag:            record[58],
			KanaNameUpdateFlag:             record[59],
			TestCommentCategory:            record[60],
			GeneralAdditionApplicability:   record[61],
			ComprehensiveReductionCategory: record[62],
			EndoscopicSurgeryAddition:      record[63],
			Reserved65:                     record[64],
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
			RelatedSectionChapter:          record[94],
			RelatedSectionPart:             record[95],
			RelatedSectionNumber:           record[96],
			RelatedSectionSubNumber:        record[97],
			RelatedSectionItemNumber:       record[98],
			AgeAdditionLower1:              record[99],
			AgeAdditionUpper1:              record[100],
			AgeAdditionCode1:               record[101],
			AgeAdditionLower2:              record[102],
			AgeAdditionUpper2:              record[103],
			AgeAdditionCode2:               record[104],
			AgeAdditionLower3:              record[105],
			AgeAdditionUpper3:              record[106],
			AgeAdditionCode3:               record[107],
			AgeAdditionLower4:              record[108],
			AgeAdditionUpper4:              record[109],
			AgeAdditionCode4:               record[110],
			MigrationRelatedCode:           record[111],
			BasicKanjiName:                 record[112],
			SinusSurgeryEndoscope:          record[113],
			SinusSurgeryBoneTissueDevice:   record[114],
			LongTimeAnesthesia:             record[115],
			PointTableNumber:               record[116],
			MonitoringAddition:             record[117],
			FrozenTissueAddition:           record[118],
			MalignantTumorPathology:        record[119],
			ExternalFixationAddition:       record[120],
			UltrasoundCuttingAddition:      record[121],
			LeftAtrialAppendageClosure:     record[122],
			OutpatientInfectionControl:     record[123],
			OtolaryngologyInfantTreatment:  record[124],
			OtolaryngologyAntimicrobial:    record[125],
			NegativePressureDressing:       record[126],
			NursingStaffTreatmentImprove:   record[127],
			OutpatientBaseUpEvaluation1:    record[128],
			OutpatientBaseUpEvaluation2:    record[129],
			RemanufacturedSingleUseDevice:  record[130],
			Reserved132:                    record[131],
			Reserved133:                    record[132],
			Reserved134:                    record[133],
			Reserved135:                    record[134],
			Reserved136:                    record[135],
			Reserved137:                    record[136],
			Reserved138:                    record[137],
			Reserved139:                    record[138],
			Reserved140:                    record[139],
			Reserved141:                    record[140],
			Reserved142:                    record[141],
			Reserved143:                    record[142],
			Reserved144:                    record[143],
			Reserved145:                    record[144],
			Reserved146:                    record[145],
			Reserved147:                    record[146],
			Reserved148:                    record[147],
			Reserved149:                    record[148],
			Reserved150:                    record[149],
		}
	})
}

func ImportMedicalPracticeSupports(db *gorm.DB) {
	importCsv(db, 27, "csv/medical_practice/01補助マスターテーブル.csv", func(record []string) model.MedicalPracticeSupport {
		return model.MedicalPracticeSupport{
			ChangeCategory:           record[0],
			MedicalPracticeCode:      record[1],
			AbbreviatedKanjiName:     record[2],
			ComprehensiveUnit1:       record[3],
			ComprehensiveGroup1:      record[4],
			ComprehensiveUnit2:       record[5],
			ComprehensiveGroup2:      record[6],
			ComprehensiveUnit3:       record[7],
			ComprehensiveGroup3:      record[8],
			ConflictDaily:            record[9],
			ConflictMonthly:          record[10],
			ConflictSimultaneous:     record[11],
			ConflictWeekly:           record[12],
			Reserved14:               record[13],
			Reserved15:               record[14],
			Reserved16:               record[15],
			Reserved17:               record[16],
			Reserved18:               record[17],
			Reserved19:               record[18],
			InpatientBasicFeeGroup:   record[19],
			CalculationCountAddition: record[20],
			Reserved22:               record[21],
			Reserved23:               record[22],
			Reserved24:               record[23],
			Reserved25:               record[24],
			UpdateDate:               record[25],
			ExpiryDate:               record[26],
		}
	})
}

func ImportMedicalPracticeInclusions(db *gorm.DB) {
	importCsv(db, 7, "csv/medical_practice/02包括テーブル.csv", func(record []string) model.MedicalPracticeInclusion {
		return model.MedicalPracticeInclusion{
			ChangeCategory:            record[0],
			ComprehensiveGroupNumber:  record[1],
			ComprehensivePracticeCode: record[2],
			AbbreviatedKanjiName:      record[3],
			SpecialCondition:          record[4],
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
		importCsv(db, 10, filePath, func(record []string) model.MedicalPracticeConflict {
			return model.MedicalPracticeConflict{
				ChangeCategory:        record[0],
				MedicalPracticeCode1:  record[1],
				AbbreviatedKanjiName1: record[2],
				MedicalPracticeCode2:  record[3],
				AbbreviatedKanjiName2: record[4],
				ConflictCategory:      record[5],
				SpecialCondition:      record[6],
				Reserved8:             record[7],
				UpdateDate:            record[8],
				ExpiryDate:            record[9],
			}
		})
	}
}

func ImportInpatientBasicFees(db *gorm.DB) {
	importCsv(db, 8, "csv/medical_practice/04入院基本料テーブル.csv", func(record []string) model.InpatientBasicFee {
		return model.InpatientBasicFee{
			ChangeCategory:       record[0],
			BasicFeeGroup:        record[1],
			MedicalPracticeCode:  record[2],
			AbbreviatedKanjiName: record[3],
			AdditionIdentifier:   record[4],
			Reserved6:            record[5],
			UpdateDate:           record[6],
			ExpiryDate:           record[7],
		}
	})
}

func ImportCalculationCounts(db *gorm.DB) {
	importCsv(db, 14, "csv/medical_practice/05算定回数テーブル.csv", func(record []string) model.CalculationCount {
		return model.CalculationCount{
			ChangeCategory:       record[0],
			MedicalPracticeCode:  record[1],
			AbbreviatedKanjiName: record[2],
			UnitCode:             record[3],
			UnitName:             record[4],
			MaxTimes:             parseInt(record[5]),
			SpecialCondition:     record[6],
			Reserved8:            record[7],
			Reserved9:            record[8],
			Reserved10:           record[9],
			Reserved11:           record[10],
			Reserved12:           record[11],
			UpdateDate:           record[12],
			ExpiryDate:           record[13],
		}
	})
}
