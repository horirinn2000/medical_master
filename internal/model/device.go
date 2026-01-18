package model

import (
	"time"
)

// SpecialMedicalDevice 特定器材マスターモデル
// doc/R06rec1.pdf 198-199ページ（ファイルレイアウト：特定器材マスター）に基づく
type SpecialMedicalDevice struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UpdateCategory               string  `gorm:"size:1" json:"update_category"`                // 1: 変更区分 (数字 1)
	MasterType                   string  `gorm:"size:1" json:"master_type"`                    // 2: マスター種別 (英数 1) - "T" 固定
	Code                         string  `gorm:"size:9;index" json:"code"`                     // 3: 特定器材コード (数字 9)
	NameKanjiLen                 int     `json:"name_kanji_len"`                               // 4: 漢字有効桁数 (数字 2)
	NameKanji                    string  `gorm:"size:64;index" json:"name_kanji"`              // 5: 漢字名称 (漢字 64)
	NameKanaLen                  int     `json:"name_kana_len"`                                // 6: カナ有効桁数 (数字 2)
	NameKana                     string  `gorm:"size:20;index" json:"name_kana"`               // 7: カナ名称 (英数カナ 20)
	UnitCode                     string  `gorm:"size:3" json:"unit_code"`                      // 8: 単位コード (数字 3)
	UnitNameKanjiLen             int     `json:"unit_name_kanji_len"`                          // 9: 単位漢字有効桁数 (数字 1)
	UnitNameKanji                string  `gorm:"size:12" json:"unit_name_kanji"`               // 10: 単位漢字名称 (漢字 12)
	PriceType                    string  `gorm:"size:1" json:"price_type"`                     // 11: 金額種別（新又は現金額） (数字 1)
	Price                        float64 `json:"price"`                                        // 12: 新又は現金額 (数字 13) - 整数10桁、小数点1桁、小数2桁
	Reserved1                    string  `gorm:"size:1" json:"reserved_1"`                     // 13: 予備 (数字 1)
	AgeAdditionCategory          string  `gorm:"size:1" json:"age_addition_category"`          // 14: 年齢加算区分 (数字 1)
	MinAge                       string  `gorm:"size:2" json:"min_age"`                        // 15: 下限年齢 (英数 2)
	MaxAge                       string  `gorm:"size:2" json:"max_age"`                        // 16: 上限年齢 (英数 2)
	OldPriceType                 string  `gorm:"size:1" json:"old_price_type"`                 // 17: 金額種別（旧金額） (数字 1)
	OldPrice                     float64 `json:"old_price"`                                    // 18: 旧金額 (数字 13)
	NameKanjiUpdateFlag          string  `gorm:"size:1" json:"name_kanji_update_flag"`         // 19: 漢字名称変更区分 (数字 1)
	NameKanaUpdateFlag           string  `gorm:"size:1" json:"name_kana_update_flag"`          // 20: カナ名称変更区分 (数字 1)
	OxygenCategory               string  `gorm:"size:1" json:"oxygen_category"`                // 21: 酸素等区分 (数字 1)
	DeviceCategory               string  `gorm:"size:1" json:"device_category"`                // 22: 特定器材種別 (数字 1)
	MaxPriceFlag                 string  `gorm:"size:1" json:"max_price_flag"`                 // 23: 上限価格 (数字 1)
	MaxPoints                    int     `json:"max_points"`                                   // 24: 上限点数 (数字 7)
	Reserved2                    string  `gorm:"size:85" json:"reserved_2"`                    // 25: 予備 (英数 85)
	PublishOrder                 int     `json:"publish_order"`                                // 26: 公表順序番号 (数字 9)
	RelatedCode                  string  `gorm:"size:9" json:"related_code"`                   // 27: 廃止・新設関連 (数字 9)
	UpdateDate                   string  `gorm:"size:8" json:"update_date"`                    // 28: 変更年月日 (数字 8)
	TransitionalDate             string  `gorm:"size:8" json:"transitional_date"`              // 29: 経過措置年月日 (数字 8)
	DiscontinuedDate             string  `gorm:"size:8" json:"discontinued_date"`              // 30: 廃止年月日 (数字 8)
	AppendixNumber               string  `gorm:"size:2" json:"appendix_number"`                // 31: 別表番号 (数字 2)
	SectionNumber                string  `gorm:"size:3" json:"section_number"`                 // 32: 区分番号 (数字 3)
	DpcCategory                  string  `gorm:"size:1" json:"dpc_category"`                   // 33: DPC適用区分 (数字 1)
	Reserved3                    string  `gorm:"size:10" json:"reserved_3"`                    // 34: 予備 (英数 10)
	Reserved4                    string  `gorm:"size:10" json:"reserved_4"`                    // 35: 予備 (英数 10)
	Reserved5                    string  `gorm:"size:10" json:"reserved_5"`                    // 36: 予備 (英数 10)
	Fullname                     string  `gorm:"size:300;index" json:"fullname"`               // 37: 基本漢字名称 (漢字 300)
	RemanufacturedDeviceCategory string  `gorm:"size:3" json:"remanufactured_device_category"` // 38: 再製造単回使用医療機器 (数字 3)
}
