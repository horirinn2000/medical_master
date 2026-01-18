package model

import (
	"time"
)

// CommentRelation コメント関連テーブルモデル
// doc/ck_1_20240823.pdf に基づく
type CommentRelation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UpdateCategory    string `gorm:"size:1" json:"update_category"`     // 1: 変更区分 (数字 1)
	NotificationType  string `gorm:"size:1" json:"notification_type"`   // 2: コメント記載通知等 (数字 1)
	ItemNumber        string `gorm:"size:4" json:"item_number"`         // 3: 項番 (数字 4)
	Section           string `gorm:"size:64" json:"section"`            // 4: 区分 (可変 64)
	BranchNumber      string `gorm:"size:2" json:"branch_number"`       // 5: 枝番 (数字 2)
	ActCode           string `gorm:"size:9;index" json:"act_code"`      // 6: 診療(調剤)行為コード (数字 9)
	AdditionCode      string `gorm:"size:5" json:"addition_code"`       // 7: 加算コード (可変 5)
	ActNameAbbr       string `gorm:"size:64" json:"act_name_abbr"`      // 8: 省略漢字名称 (可変 64)
	CommentCode       string `gorm:"size:9;index" json:"comment_code"`  // 9: コメントコード (数字 9)
	PatientStatusCode string `gorm:"size:3" json:"patient_status_code"` // 10: 患者の状態コード (可変 3)
	CommentText       string `gorm:"size:300" json:"comment_text"`      // 11: コメント文 (可変 300)
	UpdateDate        string `gorm:"size:8" json:"update_date"`         // 12: 変更年月日 (数字 8)
	DiscontinuedDate  string `gorm:"size:8" json:"discontinued_date"`   // 13: 廃止年月日 (数字 8)
	ConditionCategory string `gorm:"size:2" json:"condition_category"`  // 14: 条件区分 (数字 2)
	NonPaymentReason  string `gorm:"size:1" json:"non_payment_reason"`  // 15: 非算定理由コメント (数字 1)
	AdmissionCategory string `gorm:"size:1" json:"admission_category"`  // 16: 入外区分 (数字 1)
	CalculateCount    int    `json:"calculate_count"`                   // 17: 算定回数 (数字 3)
	PublishOrder      string `gorm:"size:9" json:"publish_order"`       // 18: 公表順序番号 (数字 9)

	// リレーション
	Comment Comment `gorm:"foreignKey:CommentCode;references:Code" json:"comment"`
}
