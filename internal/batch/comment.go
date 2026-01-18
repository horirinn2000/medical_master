package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportComments(db *gorm.DB) {
	// 全30項目チェック (R06rec3.pdf 200ページ)
	importCsv(db, 30, "csv/comment/c_20250715.csv", func(record []string) model.Comment {
		return model.Comment{
			ChangeCategory:    record[0],
			MasterType:        record[1],
			Category:          record[2],
			Pattern:           record[3],
			SerialNo:          record[4],
			KanjiNameLength:   parseInt(record[5]),
			KanjiName:         record[6],
			KanaNameLength:    parseInt(record[7]),
			KanaName:          record[8],
			EditInfo1Pos:      parseInt(record[9]),
			EditInfo1Len:      parseInt(record[10]),
			EditInfo2Pos:      parseInt(record[11]),
			EditInfo2Len:      parseInt(record[12]),
			EditInfo3Pos:      parseInt(record[13]),
			EditInfo3Len:      parseInt(record[14]),
			EditInfo4Pos:      parseInt(record[15]),
			EditInfo4Len:      parseInt(record[16]),
			Reserved18:        record[17],
			Reserved19:        record[18],
			SelectionCategory: record[19],
			UpdateDate:        record[20],
			ExpiryDate:        record[21],
			Code:              record[22],
			PublicationOrder:  parseInt(record[23]),
			Reserved25:        record[24],
			Reserved26:        record[25],
			Reserved27:        record[26],
			Reserved28:        record[27],
			Reserved29:        record[28],
			Reserved30:        record[29],
		}
	})
}
