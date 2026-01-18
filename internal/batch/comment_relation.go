package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportCommentRelations(db *gorm.DB) {
	// 全30項目チェック (R06rec3.pdf 5-7ページ)
	importCsv(db, 30, "csv/comment_relation/ck_ALL_20250901.csv", func(record []string) model.CommentRelation {
		return model.CommentRelation{
			ChangeCategory:           record[0],
			CommentNotificationType:  record[1],
			SectionNumber:            record[2],
			Category:                 record[3],
			BranchNumber:             record[4],
			ActCode:                  record[5],
			AdditionCode:             record[6],
			AbbreviatedKanjiName:     record[7],
			CommentCode:              record[8],
			PatientStatusCode:        record[9],
			CommentText:              record[10],
			UpdateDate:               record[11],
			ExpiryDate:               record[12],
			ConditionCategory:        record[13],
			NonCalculationReasonFlag: record[14],
			InpatientOutpatientType:  record[15],
			CalculationCount:         parseInt(record[16]),
			PublicationOrder:         parseInt(record[17]),
			Reserved19:               record[18],
			Reserved20:               record[19],
			Reserved21:               record[20],
			Reserved22:               record[21],
			Reserved23:               record[22],
			Reserved24:               record[23],
			Reserved25:               record[24],
			Reserved26:               record[25],
			Reserved27:               record[26],
			Reserved28:               record[27],
			Reserved29:               record[28],
			Reserved30:               record[29],
		}
	})
}
