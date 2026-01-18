package model

// CommentRelation コメント関連テーブル (ck_*.csv)
// 全30項目を網羅 (R06rec3.pdf 5-7ページ [仕様説明書は PDF 5-7, 218] に準拠)
type CommentRelation struct {
	ChangeCategory           string `csv:"1" json:"change_category"`              // 1: 変更区分
	CommentNotificationType  string `csv:"2" json:"comment_notification_type"`    // 2: コメント記載通知等
	SectionNumber            string `csv:"3" json:"section_number"`               // 3: 項番
	Category                 string `csv:"4" json:"category"`                     // 4: 区分
	BranchNumber             string `csv:"5" json:"branch_number"`                // 5: 枝番
	ActCode                  string `csv:"6" json:"act_code"`                     // 6: 診療（調剤）行為コード
	AdditionCode             string `csv:"7" json:"addition_code"`                // 7: 加算コード
	AbbreviatedKanjiName     string `csv:"8" json:"abbreviated_kanji_name"`       // 8: 省略漢字名称
	CommentCode              string `csv:"9" json:"comment_code"`                 // 9: コメントコード
	PatientStatusCode        string `csv:"10" json:"patient_status_code"`         // 10: 患者の状態コード
	CommentText              string `csv:"11" json:"comment_text"`                // 11: コメント文
	UpdateDate               string `csv:"12" json:"update_date"`                 // 12: 変更年月日
	ExpiryDate               string `csv:"13" json:"expiry_date"`                 // 13: 廃止年月日
	ConditionCategory        string `csv:"14" json:"condition_category"`          // 14: 条件区分
	NonCalculationReasonFlag string `csv:"15" json:"non_calculation_reason_flag"` // 15: 非算定理由コメント
	InpatientOutpatientType  string `csv:"16" json:"inpatient_outpatient_type"`   // 16: 入外区分
	CalculationCount         int    `csv:"17" json:"calculation_count"`           // 17: 算定回数
	PublicationOrder         int    `csv:"18" json:"publication_order"`           // 18: 公表順序番号
	Reserved19               string `csv:"19" json:"reserved_19"`                 // 19: 予備
	Reserved20               string `csv:"20" json:"reserved_20"`                 // 20: 予備
	Reserved21               string `csv:"21" json:"reserved_21"`                 // 21: 予備
	Reserved22               string `csv:"22" json:"reserved_22"`                 // 22: 予備
	Reserved23               string `csv:"23" json:"reserved_23"`                 // 23: 予備
	Reserved24               string `csv:"24" json:"reserved_24"`                 // 24: 予備
	Reserved25               string `csv:"25" json:"reserved_25"`                 // 25: 予備
	Reserved26               string `csv:"26" json:"reserved_26"`                 // 26: 予備
	Reserved27               string `csv:"27" json:"reserved_27"`                 // 27: 予備
	Reserved28               string `csv:"28" json:"reserved_28"`                 // 28: 予備
	Reserved29               string `csv:"29" json:"reserved_29"`                 // 29: 予備
	Reserved30               string `csv:"30" json:"reserved_30"`                 // 30: 予備
}
