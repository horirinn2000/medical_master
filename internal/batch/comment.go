package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportComments(db *gorm.DB) {
	importCsv(db, 30, "csv/comment/c_20250715.csv", func(record []string) model.Comment {
		return model.Comment{
			UpdateCategory:      record[0],
			MasterType:          record[1],
			Category:            record[2],
			Pattern:             record[3],
			SerialNo:            record[4],
			NameKanjiLen:        parseInt(record[5]),
			NameKanji:           record[6],
			NameKanaLen:         parseInt(record[7]),
			NameKana:            record[8],
			ReceiptEditInfo1Pos: parseInt(record[9]),
			ReceiptEditInfo1Len: parseInt(record[10]),
			ReceiptEditInfo2Pos: parseInt(record[11]),
			ReceiptEditInfo2Len: parseInt(record[12]),
			ReceiptEditInfo3Pos: parseInt(record[13]),
			ReceiptEditInfo3Len: parseInt(record[14]),
			ReceiptEditInfo4Pos: parseInt(record[15]),
			ReceiptEditInfo4Len: parseInt(record[16]),
			Reserved1:           record[17],
			Reserved2:           record[18],
			SelectionCategory:   record[19],
			UpdateDate:          record[20],
			DiscontinuedDate:    record[21],
			Code:                record[22],
			PublishOrder:        parseInt(record[23]),
			Reserved3:           record[24],
			Reserved4:           record[25],
			Reserved5:           record[26],
			Reserved6:           record[27],
			Reserved7:           record[28],
			Reserved8:           record[29],
		}
	})
}
