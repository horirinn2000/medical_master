package model

// Medicine 医薬品マスター (y_*.csv)
// 全42項目を網羅 (R06rec3.pdf 196-197ページに準拠)
type Medicine struct {
	ChangeCategory                 string  `csv:"1" json:"change_category"`                    // 1: 変更区分
	MasterType                     string  `csv:"2" json:"master_type"`                        // 2: マスター種別
	MedicineCode                   string  `csv:"3" json:"medicine_code"`                      // 3: 医薬品コード
	KanjiNameLength                int     `csv:"4" json:"kanji_name_length"`                  // 4: 漢字有効桁数 (医薬品名・規格名)
	KanjiName                      string  `csv:"5" json:"kanji_name"`                         // 5: 漢字名称 (医薬品名・規格名)
	KanaNameLength                 int     `csv:"6" json:"kana_name_length"`                   // 6: カナ有効桁数 (医薬品名・規格名)
	KanaName                       string  `csv:"7" json:"kana_name"`                          // 7: カナ名称 (医薬品名・規格名)
	UnitCode                       string  `csv:"8" json:"unit_code"`                          // 8: コード (単位)
	UnitKanjiNameLength            int     `csv:"9" json:"unit_kanji_name_length"`             // 9: 漢字有効桁数 (単位)
	UnitKanjiName                  string  `csv:"10" json:"unit_kanji_name"`                   // 10: 漢字名称 (単位)
	PriceType                      string  `csv:"11" json:"price_type"`                        // 11: 金額種別 (新又は現金額)
	Price                          float64 `csv:"12" json:"price"`                             // 12: 新又は現金額
	Reserved13                     string  `csv:"13" json:"reserved_13"`                       // 13: 予備
	NarcoticToxicDrugCategory      string  `csv:"14" json:"narcotic_toxic_drug_category"`      // 14: 麻薬・毒薬・覚醒剤原料・向精神薬
	NerveDestructionAgent          string  `csv:"15" json:"nerve_destruction_agent"`           // 15: 神経破壊剤
	BiologicProduct                string  `csv:"16" json:"biologic_product"`                  // 16: 生物学的製剤
	GenericDrug                    string  `csv:"17" json:"generic_drug"`                      // 17: 後発品
	Reserved18                     string  `csv:"18" json:"reserved_18"`                       // 18: 予備
	DentalSpecificDrug             string  `csv:"19" json:"dental_specific_drug"`              // 19: 歯科特定薬剤
	ContrastAgent                  string  `csv:"20" json:"contrast_agent"`                    // 20: 造影（補助）剤
	InjectionVolume                int     `csv:"21" json:"injection_volume"`                  // 21: 注射容量
	ListingMethod                  string  `csv:"22" json:"listing_method"`                    // 22: 収載方式等識別
	BrandNameRelated               string  `csv:"23" json:"brand_name_related"`                // 23: 商品名等関連
	OldPriceType                   string  `csv:"24" json:"old_price_type"`                    // 24: 金額種別 (旧金額)
	OldPrice                       float64 `csv:"25" json:"old_price"`                         // 25: 旧金額
	KanjiNameUpdateFlag            string  `csv:"26" json:"kanji_name_update_flag"`            // 26: 漢字名称変更区分
	KanaNameUpdateFlag             string  `csv:"27" json:"kana_name_update_flag"`             // 27: カナ名称変更区分
	DosageForm                     string  `csv:"28" json:"dosage_form"`                       // 28: 剤形
	Reserved29                     string  `csv:"29" json:"reserved_29"`                       // 29: 予備
	UpdateDate                     string  `csv:"30" json:"update_date"`                       // 30: 変更年月日
	ExpiryDate                     string  `csv:"31" json:"expiry_date"`                       // 31: 廃止年月日
	NationalDrugCode               string  `csv:"32" json:"national_drug_code"`                // 32: 薬価基準収載医薬品コード
	PublicationOrder               int     `csv:"33" json:"publication_order"`                 // 33: 公表順序番号
	TransitionalDate               string  `csv:"34" json:"transitional_date"`                 // 34: 経過措置年月日又は商品名医薬品コード使用期限
	BasicKanjiName                 string  `csv:"35" json:"basic_kanji_name"`                  // 35: 基本漢字名称
	PriceListingDate               string  `csv:"36" json:"price_listing_date"`                // 36: 薬価基準収載年月日
	GenericNameCode                string  `csv:"37" json:"generic_name_code"`                 // 37: 一般名コード
	GenericNameStandardDescription string  `csv:"38" json:"generic_name_standard_description"` // 38: 一般名処方の標準的な記載
	GenericAdditionCategory        string  `csv:"39" json:"generic_addition_category"`         // 39: 一般名処方加算対象区分
	AntiHIVDrugCategory            string  `csv:"40" json:"anti_hiv_drug_category"`            // 40: 抗ＨＩＶ薬区分
	LongTermListingRelated         string  `csv:"41" json:"long_term_listing_related"`         // 41: 長期収載品関連
	SelectedMedicalCategory        string  `csv:"42" json:"selected_medical_category"`         // 42: 選定療養区分
}

// HotCode 医薬品ＨＯＴコードマスター (標準セット)
// 全24項目を網羅 (標準レイアウト.pdfに準拠)
type HotCode struct {
	HotCode               string  `csv:"1" json:"hot_code"`                 // 1: 基準番号(ＨＯＴコード)
	Hot7                  string  `csv:"2" json:"hot_7"`                    // 2: 処方用番号(ＨＯＴ７)
	CompanyCode           string  `csv:"3" json:"company_code"`             // 3: 会社識別番号
	DispensingCode        string  `csv:"4" json:"dispensing_code"`          // 4: 調剤用番号
	LogisticsCode         string  `csv:"5" json:"logistics_code"`           // 5: 物流用番号
	JanCode               string  `csv:"6" json:"jan_code"`                 // 6: ＪＡＮコード
	NationalDrugCode      string  `csv:"7" json:"national_drug_code"`       // 7: 薬価基準収載医薬品コード
	IndividualDrugCode    string  `csv:"8" json:"individual_drug_code"`     // 8: 個別医薬品コード
	ReceiptCode1          string  `csv:"9" json:"receipt_code_1"`           // 9: レセプト電算処理システムコード(1)
	ReceiptCode2          string  `csv:"10" json:"receipt_code_2"`          // 10: レセプト電算処理システムコード(2)
	OfficialName          string  `csv:"11" json:"official_name"`           // 11: 告示名称
	SalesName             string  `csv:"12" json:"sales_name"`              // 12: 販売名
	ReceiptMedicineName   string  `csv:"13" json:"receipt_medicine_name"`   // 13: レセプト電算処理システム医薬品名
	StandardUnit          string  `csv:"14" json:"standard_unit"`           // 14: 規格単位
	PackageType           string  `csv:"15" json:"package_type"`            // 15: 包装形態
	PackageQuantityNum    float64 `csv:"16" json:"package_quantity_num"`    // 16: 包装単位数
	PackageQuantityUnit   string  `csv:"17" json:"package_quantity_unit"`   // 17: 包装単位単位
	TotalQuantityNum      float64 `csv:"18" json:"total_quantity_num"`      // 18: 包装総量数
	TotalQuantityUnit     string  `csv:"19" json:"total_quantity_unit"`     // 19: 包装総量単位
	Category              string  `csv:"20" json:"category"`                // 20: 区分
	Manufacturer          string  `csv:"21" json:"manufacturer"`            // 21: 製造会社
	SalesCompany          string  `csv:"22" json:"sales_company"`           // 22: 販売会社
	RecordCategory        string  `csv:"23" json:"record_category"`         // 23: レコード区分
	UpdateDate            string  `csv:"24" json:"update_date"`             // 24: 更新年月日
	OptionPackageQuantity float64 `gorm:"-" json:"option_package_quantity"` // オプションレイアウトから補完
}

// HotCodeOption 医薬品ＨＯＴコードマスター (オプションレイアウト)
// 全7項目を網羅 (オプションレイアウト.pdfに準拠)
type HotCodeOption struct {
	HotCode             string  `csv:"1" json:"hot_code"`              // 1: 基準番号(ＨＯＴコード)
	PackageQuantityNum  float64 `csv:"2" json:"package_quantity_num"`  // 2: 包装数量(数量)
	PackageQuantityUnit string  `csv:"3" json:"package_quantity_unit"` // 3: 包装数量(単位)
	InnerQuantityNum    float64 `csv:"4" json:"inner_quantity_num"`    // 4: 包装入数(数量)
	InnerQuantityUnit   string  `csv:"5" json:"inner_quantity_unit"`   // 5: 包装入数(単位)
	RecordCategory      string  `csv:"6" json:"record_category"`       // 6: レコード区分
	UpdateDate          string  `csv:"7" json:"update_date"`           // 7: 更新年月日
}
