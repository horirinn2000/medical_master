package batch

import (
	"medical_master/internal/model"

	"gorm.io/gorm"
)

func ImportCommentRelations(db *gorm.DB) {
	importCsv(db, 18, "csv/comment_relation/ck_ALL_20250901.csv", func(record []string) model.CommentRelation {
		return model.CommentRelation{
			UpdateCategory:    record[0],
			NotificationType:  record[1],
			ItemNumber:        record[2],
			Section:           record[3],
			BranchNumber:      record[4],
			ActCode:           record[5],
			AdditionCode:      record[6],
			ActNameAbbr:       record[7],
			CommentCode:       record[8],
			PatientStatusCode: record[9],
			CommentText:       record[10],
			UpdateDate:        record[11],
			DiscontinuedDate:  record[12],
			ConditionCategory: record[13],
			NonPaymentReason:  record[14],
			AdmissionCategory: record[15],
			CalculateCount:    parseInt(record[16]),
			PublishOrder:      record[17],
		}
	})
}
