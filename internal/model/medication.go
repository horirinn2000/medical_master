package model

import "time"

type Medication struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 調剤行為マスター仕様 (65カラム)
	UpdateCategory             string  `json:"update_category"`              // 1: 変更区分
	MasterType                 string  `json:"master_type"`                  // 2: マスター種別
	Code                       string  `gorm:"index" json:"code"`            // 3: 調剤行為コード
	NameKanjiLen               int     `json:"name_kanji_len"`               // 4: 漢字有効桁数
	NameKanji                  string  `gorm:"index" json:"name_kanji"`      // 5: 漢字名称
	NameKanaLen                int     `json:"name_kana_len"`                // 6: カナ有効桁数
	NameKana                   string  `gorm:"index" json:"name_kana"`       // 7: カナ名称
	UnitCode                   string  `json:"unit_code"`                    // 8: 単位コード
	UnitNameKanjiLen           int     `json:"unit_name_kanji_len"`          // 9: 単位漢字有効桁数
	UnitNameKanji              string  `json:"unit_name_kanji"`              // 10: 単位漢字名称
	PriceType                  string  `json:"price_type"`                   // 11: 金額種別
	Price                      float64 `json:"price"`                        // 12: 新又は現金額
	Reserved1                  string  `json:"reserved_1"`                   // 13: 予備
	AgeAdditionCategory        string  `json:"age_addition_category"`        // 14: 年齢加算区分
	MinAge                     string  `json:"min_age"`                      // 15: 下限年齢
	MaxAge                     string  `json:"max_age"`                      // 16: 上限年齢
	OldPriceType               string  `json:"old_price_type"`               // 17: 金額種別
	OldPrice                   float64 `json:"old_price"`                    // 18: 旧金額
	NameKanjiUpdateFlag        string  `json:"name_kanji_update_flag"`       // 19: 漢字名称変更区分
	NameKanaUpdateFlag         string  `json:"name_kana_update_flag"`        // 20: カナ名称変更区分
	Reserved2                  string  `json:"reserved_2"`                   // 21: 予備
	Reserved3                  string  `json:"reserved_3"`                   // 22: 予備
	MaxPriceFlag               string  `json:"max_price_flag"`               // 23: 上限価格
	MaxPoints                  int     `json:"max_points"`                   // 24: 上限点数
	Reserved4                  string  `json:"reserved_4"`                   // 25: 予備
	PublishOrder               int     `json:"publish_order"`                // 26: 公表順序番号
	RelatedCode                string  `json:"related_code"`                 // 27: 廃止・新設関連
	UpdateDate                 string  `json:"update_date"`                  // 28: 変更年月日
	TransitionalDate           string  `json:"transitional_date"`            // 29: 経過措置年月日
	DiscontinuedDate           string  `json:"discontinued_date"`            // 30: 廃止年月日
	MedicationSectionNumber    string  `json:"medication_section_number"`    // 31: 区分番号
	Reserved5                  string  `json:"reserved_5"`                   // 32: 予備
	Reserved6                  string  `json:"reserved_6"`                   // 33: 予備
	Reserved7                  string  `json:"reserved_7"`                   // 34: 予備
	Reserved8                  string  `json:"reserved_8"`                   // 35: 予備
	Reserved9                  string  `json:"reserved_9"`                   // 36: 予備
	Reserved10                 string  `json:"reserved_10"`                  // 37: 予備
	Reserved11                 string  `json:"reserved_11"`                  // 38: 予備
	Reserved12                 string  `json:"reserved_12"`                  // 39: 予備
	Reserved13                 string  `json:"reserved_13"`                  // 40: 予備
	Reserved14                 string  `json:"reserved_14"`                  // 41: 予備
	Reserved15                 string  `json:"reserved_15"`                  // 42: 予備
	Reserved16                 string  `json:"reserved_16"`                  // 43: 予備
	Reserved17                 string  `json:"reserved_17"`                  // 44: 予備
	Reserved18                 string  `json:"reserved_18"`                  // 45: 予備
	Reserved19                 string  `json:"reserved_19"`                  // 46: 予備
	Reserved20                 string  `json:"reserved_20"`                  // 47: 予備
	Reserved21                 string  `json:"reserved_21"`                  // 48: 予備
	Reserved22                 string  `json:"reserved_22"`                  // 49: 予備
	Reserved23                 string  `json:"reserved_23"`                  // 50: 予備
	Reserved24                 string  `json:"reserved_24"`                  // 51: 予備
	Reserved25                 string  `json:"reserved_25"`                  // 52: 予備
	Reserved26                 string  `json:"reserved_26"`                  // 53: 予備
	Reserved27                 string  `json:"reserved_27"`                  // 54: 予備
	Reserved28                 string  `json:"reserved_28"`                  // 55: 予備
	Reserved29                 string  `json:"reserved_29"`                  // 56: 予備
	Reserved30                 string  `json:"reserved_30"`                  // 57: 予備
	Reserved31                 string  `json:"reserved_31"`                  // 58: 予備
	PriceTablePrice            float64 `json:"price_table_price"`            // 59: 料金表価格
	ImportDate                 string  `json:"import_date"`                  // 60: 収載年月日
	Reserved32                 string  `json:"reserved_32"`                  // 61: 予備
	MedicationManagementNumber string  `json:"medication_management_number"` // 62: 調剤行為管理番号
	MedicationMigrationNumber  string  `json:"medication_migration_number"`  // 63: 移行先調剤行為管理番号
	MedicationMigrationFlag    string  `json:"medication_migration_flag"`    // 64: 調剤行為管理番号（変更情報）
	MedicationMigrationUpdate  string  `json:"medication_migration_update"`  // 65: 移行先調剤行為管理番号（変更情報）
}
