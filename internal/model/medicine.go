package model

import "time"

// Medicine 医薬品マスター
// 仕様書: doc/R06rec1.pdf (P.16-19)
type Medicine struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 1: 変更区分
	UpdateCategory string `json:"update_category"`
	// 2: マスター種別 (Y)
	MasterType string `json:"master_type"`
	// 3: 医薬品コード (9桁)
	Code string `gorm:"index" json:"code"`

	// 医薬品名・規格名
	NameKanjiLen int    `json:"name_kanji_len"` // 4: 漢字有効桁数
	NameKanji    string `json:"name_kanji"`     // 5: 漢字名称
	NameKanaLen  int    `json:"name_kana_len"`  // 6: カナ有効桁数
	NameKana     string `json:"name_kana"`      // 7: カナ名称

	// 単位
	UnitCode         string `json:"unit_code"`           // 8: コード
	UnitNameKanjiLen int    `json:"unit_name_kanji_len"` // 9: 漢字有効桁数
	UnitNameKanji    string `json:"unit_name_kanji"`     // 10: 漢字名称

	// 新又は現金額
	PriceType string  `json:"price_type"` // 11: 金額種別
	Price     float64 `json:"price"`      // 12: 新又は現金額

	Reserved1 string `json:"reserved_1"` // 13: 予備

	// 14: 麻薬・毒薬・覚醒剤原料・向精神薬
	NarcoticsCategory string `json:"narcotics_category"`
	// 15: 神経破壊剤
	NeurolyticFlag string `json:"neurolytic_flag"`
	// 16: 生物学的製剤
	BiologicFlag string `json:"biologic_flag"`
	// 17: 後発品
	GenericFlag string `json:"generic_flag"`

	Reserved2 string `json:"reserved_2"` // 18: 予備

	// 19: 歯科特定薬剤
	DentalSpecificFlag string `json:"dental_specific_flag"`
	// 20: 造影（補助）剤
	ContrastAgentCategory string `json:"contrast_agent_category"`
	// 21: 注射容量
	InjectionVolume int `json:"injection_volume"`
	// 22: 収載方式等識別
	ListingMethodType string `json:"listing_method_type"`
	// 23: 商品名等関連
	BrandNameCode string `json:"brand_name_code"`

	// 旧金額
	OldPriceType string  `json:"old_price_type"` // 24: 金額種別
	OldPrice     float64 `json:"old_price"`      // 25: 旧金額

	NameKanjiUpdateFlag string `json:"name_kanji_update_flag"` // 26: 漢字名称変更区分
	NameKanaUpdateFlag  string `json:"name_kana_update_flag"`  // 27: カナ名称変更区分

	// 28: 剤形
	DosageForm string `json:"dosage_form"`

	Reserved3 string `json:"reserved_3"` // 29: 予備

	UpdateDate       string `json:"update_date"`       // 30: 変更年月日
	DiscontinuedDate string `json:"discontinued_date"` // 31: 廃止年月日

	// 32: 薬価基準収載医薬品コード (12桁)
	NationalDrugCode string `gorm:"index" json:"national_drug_code"`

	PublishOrder     int    `json:"publish_order"`     // 33: 公表順序番号
	TransitionalDate string `json:"transitional_date"` // 34: 経過措置年月日又は商品名医薬品コード使用期限
	BasicName        string `json:"basic_name"`        // 35: 基本漢字名称 (品名)
	ListingDate      string `json:"listing_date"`      // 36: 薬価基準収載年月日

	// 一般名関連
	GenericNameCode     string `json:"generic_name_code"`     // 37: 一般名コード
	GenericNameStandard string `json:"generic_name_standard"` // 38: 一般名処方の標準的な記載
	GenericAdditionType string `json:"generic_addition_type"` // 39: 一般名処方加算対象区分

	// 40: 抗ＨＩＶ薬区分
	AntiHivFlag string `json:"anti_hiv_flag"`
	// 41: 長期収載品関連
	LongTermListingCode string `json:"long_term_listing_code"`
	// 42: 選定療養区分
	SelectionMedicalCategory string `json:"selection_medical_category"`
}

// MedicinePrice 薬価履歴
type MedicinePrice struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	MedicineCode string    `gorm:"index" json:"medicine_code"`
	Price        float64   `json:"price"`
	PriceType    string    `json:"price_type"` // 1:現金額, 2:新金額
	StartDate    string    `json:"start_date"` // 適用開始日
}

// HotCode 医薬品HOTコードマスター
// 仕様書: doc/標準レイアウト.pdf, doc/オプションレイアウト.pdf
type HotCode struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// --- 標準セット (24項目) ---
	HotCode        string `gorm:"index" json:"hot_code"`       // 1: 基準番号(ＨＯＴコード) 13桁
	Hot7           string `gorm:"index" json:"hot7"`           // 2: 処方用番号(ＨＯＴ７) 7桁
	CompanyCode    string `json:"company_code"`                // 3: 会社識別番号
	DispensingNo   string `json:"dispensing_no"`               // 4: 調剤用番号
	LogisticsNo    string `json:"logistics_no"`                // 5: 物流用番号
	JanCode        string `gorm:"index" json:"jan_code"`       // 6: ＪＡＮコード
	NationalCode   string `gorm:"index" json:"national_code"`  // 7: 薬価基準収載医薬品コード (12桁)
	IndividualCode string `json:"individual_code"`             // 8: 個別医薬品コード
	ReceiptCode1   string `gorm:"index" json:"receipt_code_1"` // 9: レセプト電算処理システムコード(1)
	ReceiptCode2   string `gorm:"index" json:"receipt_code_2"` // 10: レセプト電算処理システムコード(2)

	NotificationName string  `json:"notification_name"` // 11: 告示名称
	SalesName        string  `json:"sales_name"`        // 12: 販売名
	ReceiptName      string  `json:"receipt_name"`      // 13: レセプト電算処理システム医薬品名
	SpecUnit         string  `json:"spec_unit"`         // 14: 規格単位
	PackageForm      string  `json:"package_form"`      // 15: 包装形態
	PackageCount     float64 `json:"package_count"`     // 16: 包装単位数
	PackageUnit      string  `json:"package_unit"`      // 17: 包装単位単位
	TotalVolume      float64 `json:"total_volume"`      // 18: 包装総量数
	TotalVolumeUnit  string  `json:"total_volume_unit"` // 19: 包装総量単位
	Category         string  `json:"category"`          // 20: 区分 (内、外、注、歯)
	Manufacturer     string  `json:"manufacturer"`      // 21: 製造会社
	Distributor      string  `json:"distributor"`       // 22: 販売会社
	RecordType       string  `json:"record_type"`       // 23: レコード区分
	UpdateDate       string  `json:"update_date"`       // 24: 更新年月日

	// --- オプションレイアウト (7項目) ---
	// 1: 基準番号(ＨＯＴコード) は共通
	OptionPackageQuantity       float64 `json:"option_package_quantity"`         // 2: 包装数量(数量)
	OptionPackageQuantityUnit   string  `json:"option_package_quantity_unit"`    // 3: 包装数量(単位)
	OptionPackageInQuantity     float64 `json:"option_package_in_quantity"`      // 4: 包装入数(数量)
	OptionPackageInQuantityUnit string  `json:"option_package_in_quantity_unit"` // 5: 包装入数(単位)
	OptionRecordType            string  `json:"option_record_type"`              // 6: レコード区分
	OptionUpdateDate            string  `json:"option_update_date"`              // 7: 更新年月日

	// --- HOT9マスター (24項目) ---
	Hot9                 string  `gorm:"index" json:"hot9"`      // 0: 基準番号（ＨＯＴコード）
	Hot9Hot7             string  `json:"hot9_hot7"`              // 1: 処方用番号（ＨＯＴ７）
	Hot9CompanyCode      string  `json:"hot9_company_code"`      // 2: 会社識別用番号
	Hot9DispensingNo     string  `json:"hot9_dispensing_no"`     // 3: 調剤用番号
	Hot9LogisticsNo      string  `json:"hot9_logistics_no"`      // 4: 物流用番号
	Hot9JanCode          string  `json:"hot9_jan_code"`          // 5: ＪＡＮコード
	Hot9NationalCode     string  `json:"hot9_national_code"`     // 6: 薬価基準収載医薬品コード
	Hot9IndividualCode   string  `json:"hot9_individual_code"`   // 7: 個別医薬品コード
	Hot9ReceiptCode1     string  `json:"hot9_receipt_code_1"`    // 8: レセプト電算処理システムコード（１）
	Hot9ReceiptCode2     string  `json:"hot9_receipt_code_2"`    // 9: レセプト電算処理システムコード（２）
	Hot9NotificationName string  `json:"hot9_notification_name"` // 10: 告示名称
	Hot9SalesName        string  `json:"hot9_sales_name"`        // 11: 販売名
	Hot9ReceiptName      string  `json:"hot9_receipt_name"`      // 12: レセプト電算処理システム医薬品名
	Hot9SpecUnit         string  `json:"hot9_spec_unit"`         // 13: 規格単位
	Hot9PackageForm      string  `json:"hot9_package_form"`      // 14: 包装形態
	Hot9PackageCount     float64 `json:"hot9_package_count"`     // 15: 包装単位数
	Hot9PackageUnit      string  `json:"hot9_package_unit"`      // 16: 包装単位単位
	Hot9TotalVolume      float64 `json:"hot9_total_volume"`      // 17: 包装総量数
	Hot9TotalVolumeUnit  string  `json:"hot9_total_volume_unit"` // 18: 包装総量単位
	Hot9Category         string  `json:"hot9_category"`          // 19: 区分
	Hot9Manufacturer     string  `json:"hot9_manufacturer"`      // 20: 製造会社
	Hot9Distributor      string  `json:"hot9_distributor"`       // 21: 販売会社
	Hot9UpdateCategory   string  `json:"hot9_update_category"`   // 22: 更新区分
	Hot9UpdateDate       string  `json:"hot9_update_date"`       // 23: 更新年月日
}
