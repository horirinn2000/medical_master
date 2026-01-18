package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportMedications(db *gorm.DB) {
	// 全65項目チェック (R06rec3.pdf 216-217ページ)
	importCsv(db, 65, "csv/medication/m20250307.csv", func(record []string) model.Medication {
		return model.Medication{
			ChangeCategory:             record[0],
			MasterType:                 record[1],
			Code:                       record[2],
			KanjiNameLength:            parseInt(record[3]),
			KanjiName:                  record[4],
			KanaNameLength:             parseInt(record[5]),
			KanaName:                   record[6],
			ReceiptSymbolCode:          record[7],
			ReceiptOrderNumber:         parseInt(record[8]),
			ScoreType:                  record[9],
			QuantityCalculationFlag:    record[10],
			BasicScore:                 parseInt(record[11]),
			StepCalculationCategory:    record[13], // 13: きざみ値計算識別
			StepLowerLimit:             parseInt(record[14]),
			StepUpperLimit:             parseInt(record[15]),
			StepValue:                  parseInt(record[16]),
			StepScore:                  parseInt(record[17]),
			StepErrorCategory:          record[18],
			ReductionCategory:          record[19],
			ReductionTargetCategory:    record[20],
			DispensingCategory:         record[21],
			ComprehensiveCategory:      record[22],
			Reserved23:                 record[23],
			Reserved24:                 record[24],
			Reserved25:                 record[25],
			Reserved26:                 record[26],
			Reserved27:                 record[27],
			DispensingType1:            record[28],
			DispensingType2:            record[29],
			ElderlyMedicalCategory:     record[30],
			FacilityStandard1:          record[31],
			FacilityStandard2:          record[32],
			FacilityStandard3:          record[33],
			FacilityStandard4:          record[34],
			FacilityStandard5:          record[35],
			FacilityStandard6:          record[36],
			FacilityStandard7:          record[37],
			FacilityStandard8:          record[38],
			FacilityStandard9:          record[39],
			FacilityStandard10:         record[40],
			ReceiptConflictCode:        record[41],
			PrescriptionConflictCode:   record[42],
			DispensingConflictCode:     record[43],
			NarcoticToxicDrugCategory:  record[44],
			TimeAdditionCategory:       record[45],
			DosageForm:                 record[46],
			ReceiptMaxTimes:            parseInt(record[47]),
			ReceiptMaxTimesError:       record[48],
			PrescriptionMaxTimes:       parseInt(record[49]),
			PrescriptionMaxTimesError:  record[50],
			NoteAdditionCode:           record[51],
			NoteAdditionSequence:       record[52],
			LowerAge:                   record[53],
			UpperAge:                   record[54],
			MedicineManagementCategory: record[55],
			NotificationCategory1:      record[56],
			NotificationCategory2:      record[57],
			OldScoreType:               record[58],
			OldScore:                   parseInt(record[59]),
			UpdateDate:                 record[60],
			ExpiryDate:                 record[61],
			PublicationOrder:           parseInt(record[62]),
			SectionCode:                record[63],
			RelatedSectionCode:         record[64],
			MigrationRelatedCode:       record[65],
		}
	})
}
