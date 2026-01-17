package model

import (
	"time"
)

// Ward 病棟マスター (k.csv)
type Ward struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	UpdateCategory      string    `json:"update_category"`         // 1: 変更区分
	MasterType          string    `json:"master_type"`             // 2: マスター種別
	Code                string    `gorm:"index" json:"code"`       // 3: 診療行為コード (9桁)
	NameKanjiLen        int       `json:"name_kanji_len"`          // 4: 省略漢字有効桁数
	NameKanji           string    `gorm:"index" json:"name_kanji"` // 5: 省略漢字名称
	NameKanaLen         int       `json:"name_kana_len"`           // 6: 省略カナ有効桁数
	NameKana            string    `json:"name_kana"`               // 7: 省略カナ名称
	PointCategory       string    `json:"point_category"`          // 11: 点数識別
	Point               float64   `json:"point"`                   // 12: 新又は現点数
	InOuterCategory     string    `json:"in_outer_category"`       // 13: 入外適用区分
	ElderlyCategory     string    `json:"elderly_category"`        // 14: 後期高齢者医療適用区分
	DpcCategory         string    `json:"dpc_category"`            // 18: DPC適用区分
	InstitutionCategory string    `json:"institution_category"`    // 19: 病院・診療所区分
	MedicalObservation  string    `json:"medical_observation"`     // 21: 医療観察法対象区分
	LowerLimit          string    `json:"lower_limit"`             // 31: 下限値
	UpperLimit          string    `json:"upper_limit"`             // 32: 上限値
	NotificationType    string    `json:"notification_type_1"`     // 68: 告示等識別区分（１）
	NotificationType2   string    `json:"notification_type_2"`     // 69: 告示等識別区分（２）
	DentalCategory      string    `json:"dental_category"`         // 84: 歯科適用区分
	UpdateDate          string    `json:"update_date"`             // 87: 変更年月日
	DiscontinuedDate    string    `json:"discontinued_date"`       // 88: 廃止年月日
	PublishOrder        string    `json:"publish_order"`           // 89: 公表順序番号
	Chapter             string    `json:"chapter"`                 // 90: 章
	Section             string    `json:"section"`                 // 91: 部
	SubsectionCode      string    `json:"subsection_code"`         // 92: 区分番号
	BranchNumber        string    `json:"branch_number"`           // 93: 枝番
	ItemNumber          string    `json:"item_number"`             // 94: 項番
	Fullname            string    `json:"fullname"`                // 113: 基本漢字名称
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
