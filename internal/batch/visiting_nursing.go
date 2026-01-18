package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportVisitingNursingFees(db *gorm.DB) {
	// 全72項目チェック (R06rec3.pdf 218-219ページ)
	importCsv(db, 72, "csv/visiting_nursing/r_20240823.csv", func(record []string) model.VisitingNursingFee {
		return model.VisitingNursingFee{
			ChangeCategory:               record[0],
			MasterType:                   record[1],
			VisitingNursingFeeCode:       record[2],
			NotificationSection:          record[3],
			NotificationBranch:           record[4],
			NotificationItem:             record[5],
			BasicName:                    record[6],
			AbbreviatedNameLength:        parseInt(record[7]),
			AbbreviatedName:              record[8],
			AbbreviatedKanaNameLength:    parseInt(record[9]),
			AbbreviatedKanaName:          record[10],
			DataStandardCode:             record[11],
			KanjiUnitNameLength:          parseInt(record[12]),
			KanjiUnitName:                record[13],
			PriceCategory:                record[14],
			Price:                        parseFloat(record[15]),
			OldPriceCategory:             record[16],
			OldPrice:                     parseFloat(record[17]),
			StepCalculationCategory:      record[18],
			StepLowerLimit:               parseFloat(record[19]),
			StepUpperLimit:               parseFloat(record[20]),
			StepValue:                    parseFloat(record[21]),
			StepPrice:                    parseFloat(record[22]),
			StepErrorCategory:            record[23],
			LowerAge:                     record[24],
			UpperAge:                     record[25],
			ElderlyMedicalCategory:       record[26],
			MedicalObservationCategory:   record[27],
			JobCategory1:                 record[28],
			JobCategory2:                 record[29],
			JobCategory3:                 record[30],
			JobCategory4:                 record[31],
			JobCategory5:                 record[32],
			JobCategory6:                 record[33],
			JobCategory7:                 record[34],
			JobCategory8:                 record[35],
			JobCategory9:                 record[36],
			JobCategory10:                record[37],
			JobCategory11:                record[38],
			JobCategory12:                record[39],
			JobCategory13:                record[40],
			JobCategory14:                record[41],
			JobCategory15:                record[42],
			VisitTimesCategory:           record[43],
			NursingInstructionCategory:   record[44],
			SpecialInstructionCategory:   record[45],
			SoloAdditionCategory:         record[46],
			AdditionGroup:                record[47],
			FacilityStandardGroup:        record[48],
			AdditionTableRelated:         record[49],
			CalculationCountTableRelated: record[50],
			ConflictTableRelated:         record[51],
			ReceiptDisplaySection:        record[52],
			ReceiptDisplayItem:           record[53],
			ReceiptDisplaySerial:         record[54],
			ReceiptSymbol1:               record[55],
			ReceiptSymbol2:               record[56],
			ReceiptSymbol3:               record[57],
			ReceiptSymbol4:               record[58],
			ReceiptSymbol5:               record[59],
			ReceiptSymbol6:               record[60],
			ReceiptSymbol7:               record[61],
			ReceiptSymbol8:               record[62],
			ReceiptSymbol9:               record[63],
			Reserved65:                   record[64],
			PublicationOrder:             parseInt(record[65]),
			VisitingNursingType:          record[66],
			Reserved68:                   record[67],
			Reserved69:                   record[68],
			Reserved70:                   record[69],
			UpdateDate:                   record[70],
			ExpiryDate:                   record[71],
		}
	})
}

func ImportVisitingNursingAdditions(db *gorm.DB) {
	importCsv(db, 8, "csv/visiting_nursing/r2_20240823.csv", func(record []string) model.VisitingNursingAddition {
		return model.VisitingNursingAddition{
			ChangeCategory:         record[0],
			GroupNumber:            record[1],
			VisitingNursingFeeCode: record[2],
			AbbreviatedName:        record[3],
			AdditionIdentifier:     record[4],
			UpdateDate:             record[5],
			ExpiryDate:             record[6],
			Reserved8:              record[7],
		}
	})
}

func ImportVisitingNursingCalculationCounts(db *gorm.DB) {
	importCsv(db, 10, "csv/visiting_nursing/r3_20240823.csv", func(record []string) model.VisitingNursingCalculationCount {
		return model.VisitingNursingCalculationCount{
			ChangeCategory:         record[0],
			VisitingNursingFeeCode: record[1],
			AbbreviatedName:        record[2],
			UnitCode:               record[3],
			UnitName:               record[4],
			MaxTimes:               parseInt(record[5]),
			MaxTimesErrorCategory:  record[6],
			UpdateDate:             record[7],
			ExpiryDate:             record[8],
			Reserved10:             record[9],
		}
	})
}

func ImportVisitingNursingConflicts(db *gorm.DB) {
	importCsv(db, 13, "csv/visiting_nursing/r4_20240823.csv", func(record []string) model.VisitingNursingConflict {
		return model.VisitingNursingConflict{
			ChangeCategory:          record[0],
			VisitingNursingFeeCode1: record[1],
			AbbreviatedName1:        record[2],
			ConflictCategory:        record[3],
			VisitingNursingFeeCode2: record[4],
			AbbreviatedName2:        record[5],
			ConflictUnit:            record[6],
			SpecialCondition:        record[7],
			UpdateDate:              record[8],
			ExpiryDate:              record[9],
			Reserved11:              record[10],
			Reserved12:              record[11],
			Reserved13:              record[12],
		}
	})
}

func ImportVisitingNursingFacilityStandards(db *gorm.DB) {
	importCsv(db, 9, "csv/visiting_nursing/r5_20240823.csv", func(record []string) model.VisitingNursingFacilityStandard {
		return model.VisitingNursingFacilityStandard{
			ChangeCategory:         record[0],
			GroupNumber:            record[1],
			VisitingNursingFeeCode: record[2],
			AbbreviatedName:        record[3],
			FacilityStandardCode:   record[4],
			FacilityIdentifier:     record[5],
			UpdateDate:             record[6],
			ExpiryDate:             record[7],
			Reserved9:              record[8],
		}
	})
}
