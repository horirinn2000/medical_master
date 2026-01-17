package model

import "time"

type Tooth struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 歯式マスター仕様 (7カラム)
	UpdateCategory   string `json:"update_category"`   // 1: 変更区分
	MasterType       string `json:"master_type"`       // 2: マスター種別 ('F'固定)
	Code             string `gorm:"index" json:"code"` // 3: 歯式コード (6桁)
	Reserved         string `json:"reserved"`          // 4: 予備
	Name             string `gorm:"index" json:"name"` // 5: 歯式名称
	UpdateDate       string `json:"update_date"`       // 6: 変更年月日
	DiscontinuedDate string `json:"discontinued_date"` // 7: 廃止年月日
}
