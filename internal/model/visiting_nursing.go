package model

// VisitingNursingFee 訪問看護療養費マスター・基本テーブル (r_*.csv)
// 全72項目を網羅 (R06rec3.pdf 218-219ページに準拠)
type VisitingNursingFee struct {
	ChangeCategory               string   `csv:"1" json:"change_category"`               // 1: 変更区分
	MasterType                   string   `csv:"2" json:"master_type"`                   // 2: マスター種別 ('R'固定)
	VisitingNursingFeeCode       string   `csv:"3" json:"visiting_nursing_fee_code"`     // 3: 訪問看護療養費コード
	NotificationSection          string   `csv:"4" json:"notification_section"`          // 4: 区分番号 (告示番号)
	NotificationBranch           string   `csv:"5" json:"notification_branch"`           // 5: 枝番 (告示番号)
	NotificationItem             string   `csv:"6" json:"notification_item"`             // 6: 項番 (告示番号)
	BasicName                    string   `csv:"7" json:"basic_name"`                    // 7: 基本名称
	AbbreviatedNameLength        int      `csv:"8" json:"abbreviated_name_length"`       // 8: 省略名称有効桁数
	AbbreviatedName              string   `csv:"9" json:"abbreviated_name"`              // 9: 省略名称
	AbbreviatedKanaNameLength    int      `csv:"10" json:"abbreviated_kana_name_length"` // 10: 省略カナ名称有効桁数
	AbbreviatedKanaName          string   `csv:"11" json:"abbreviated_kana_name"`        // 11: 省略カナ名称
	DataStandardCode             string   `csv:"12" json:"data_standard_code"`           // 12: データ規格コード
	KanjiUnitNameLength          int      `csv:"13" json:"kanji_unit_name_length"`       // 13: 漢字有効桁数 (データ規格名)
	KanjiUnitName                string   `csv:"14" json:"kanji_unit_name"`              // 14: 漢字名称 (データ規格名)
	PriceCategory                string   `csv:"15" json:"price_category"`               // 15: 金額識別 (新又は現金額)
	Price                        float64  `csv:"16" json:"price"`                        // 16: 新又は現金額
	OldPriceCategory             string   `csv:"17" json:"old_price_category"`           // 17: 金額識別 (旧金額)
	OldPrice                     float64  `csv:"18" json:"old_price"`                    // 18: 旧金額
	StepCalculationCategory      string   `csv:"19" json:"step_calculation_category"`    // 19: きざみ値計算識別
	StepLowerLimit               float64  `csv:"20" json:"step_lower_limit"`             // 20: 下限値
	StepUpperLimit               float64  `csv:"21" json:"step_upper_limit"`             // 21: 上限値
	StepValue                    float64  `csv:"22" json:"step_value"`                   // 22: きざみ値
	StepPrice                    float64  `csv:"23" json:"step_price"`                   // 23: きざみ金額
	StepErrorCategory            string   `csv:"24" json:"step_error_category"`          // 24: 上下限エラー処理
	LowerAge                     string   `csv:"25" json:"lower_age"`                    // 25: 下限年齢 (上下限年齢)
	UpperAge                     string   `csv:"26" json:"upper_age"`                    // 26: 上限年齢 (上下限年齢)
	ElderlyMedicalCategory       string   `csv:"27" json:"elderly_medical_category"`     // 27: 後期高齢者医療適用区分
	MedicalObservationCategory   string   `csv:"28" json:"medical_observation_category"` // 28: 医療観察法対象区分
	JobCategoryCodes             []string `gorm:"-" json:"job_category_codes"`           // 29-43: 職種等コード(15項目)
	JobCategory1                 string   `csv:"29" json:"job_category_1"`
	JobCategory2                 string   `csv:"30" json:"job_category_2"`
	JobCategory3                 string   `csv:"31" json:"job_category_3"`
	JobCategory4                 string   `csv:"32" json:"job_category_4"`
	JobCategory5                 string   `csv:"33" json:"job_category_5"`
	JobCategory6                 string   `csv:"34" json:"job_category_6"`
	JobCategory7                 string   `csv:"35" json:"job_category_7"`
	JobCategory8                 string   `csv:"36" json:"job_category_8"`
	JobCategory9                 string   `csv:"37" json:"job_category_9"`
	JobCategory10                string   `csv:"38" json:"job_category_10"`
	JobCategory11                string   `csv:"39" json:"job_category_11"`
	JobCategory12                string   `csv:"40" json:"job_category_12"`
	JobCategory13                string   `csv:"41" json:"job_category_13"`
	JobCategory14                string   `csv:"42" json:"job_category_14"`
	JobCategory15                string   `csv:"43" json:"job_category_15"`
	VisitTimesCategory           string   `csv:"44" json:"visit_times_category"`            // 44: 実施回数区分
	NursingInstructionCategory   string   `csv:"45" json:"nursing_instruction_category"`    // 45: 訪問看護指示区分
	SpecialInstructionCategory   string   `csv:"46" json:"special_instruction_category"`    // 46: 特別訪問看護指示区分
	SoloAdditionCategory         string   `csv:"47" json:"solo_addition_category"`          // 47: 加算単独算定区分
	AdditionGroup                string   `csv:"48" json:"addition_group"`                  // 48: 加算グループ
	FacilityStandardGroup        string   `csv:"49" json:"facility_standard_group"`         // 49: 施設基準グループ
	AdditionTableRelated         string   `csv:"50" json:"addition_table_related"`          // 50: 基本・加算対応テーブル関連識別
	CalculationCountTableRelated string   `csv:"51" json:"calculation_count_table_related"` // 51: 算定回数限度テーブル関連識別
	ConflictTableRelated         string   `csv:"52" json:"conflict_table_related"`          // 52: 併算定背反テーブル関連識別
	ReceiptDisplaySection        string   `csv:"53" json:"receipt_display_section"`         // 53: レセプト表示欄
	ReceiptDisplayItem           string   `csv:"54" json:"receipt_display_item"`            // 54: レセプト表示項
	ReceiptDisplaySerial         string   `csv:"55" json:"receipt_display_serial"`          // 55: レセプト表示連番
	ReceiptSymbols               []string `gorm:"-" json:"receipt_symbols"`                 // 56-64: レセプト表示用記号(9項目)
	ReceiptSymbol1               string   `csv:"56" json:"receipt_symbol_1"`
	ReceiptSymbol2               string   `csv:"57" json:"receipt_symbol_2"`
	ReceiptSymbol3               string   `csv:"58" json:"receipt_symbol_3"`
	ReceiptSymbol4               string   `csv:"59" json:"receipt_symbol_4"`
	ReceiptSymbol5               string   `csv:"60" json:"receipt_symbol_5"`
	ReceiptSymbol6               string   `csv:"61" json:"receipt_symbol_6"`
	ReceiptSymbol7               string   `csv:"62" json:"receipt_symbol_7"`
	ReceiptSymbol8               string   `csv:"63" json:"receipt_symbol_8"`
	ReceiptSymbol9               string   `csv:"64" json:"receipt_symbol_9"`
	Reserved65                   string   `csv:"65" json:"reserved_65"`           // 65: 予備
	PublicationOrder             int      `csv:"66" json:"publication_order"`     // 66: 公表順序番号
	VisitingNursingType          string   `csv:"67" json:"visiting_nursing_type"` // 67: 訪問看護療養費種類
	Reserved68                   string   `csv:"68" json:"reserved_68"`           // 68: 予備
	Reserved69                   string   `csv:"69" json:"reserved_69"`           // 69: 予備
	Reserved70                   string   `csv:"70" json:"reserved_70"`           // 70: 予備
	UpdateDate                   string   `csv:"71" json:"update_date"`           // 71: 変更年月日
	ExpiryDate                   string   `csv:"72" json:"expiry_date"`           // 72: 廃止年月日
}

// VisitingNursingAddition 訪問看護・基本加算対応テーブル (イ)
// 全8項目 (R06rec3.pdf 220ページ)
type VisitingNursingAddition struct {
	ChangeCategory         string `csv:"1" json:"change_category"`
	GroupNumber            string `csv:"2" json:"group_number"`
	VisitingNursingFeeCode string `csv:"3" json:"visiting_nursing_fee_code"`
	AbbreviatedName        string `csv:"4" json:"abbreviated_name"`
	AdditionIdentifier     string `csv:"5" json:"addition_identifier"`
	UpdateDate             string `csv:"6" json:"update_date"`
	ExpiryDate             string `csv:"7" json:"expiry_date"`
	Reserved8              string `csv:"8" json:"reserved_8"`
}

// VisitingNursingCalculationCount 訪問看護・算定回数限度テーブル (ウ)
// 全10項目 (R06rec3.pdf 221ページ)
type VisitingNursingCalculationCount struct {
	ChangeCategory         string `csv:"1" json:"change_category"`
	VisitingNursingFeeCode string `csv:"2" json:"visiting_nursing_fee_code"`
	AbbreviatedName        string `csv:"3" json:"abbreviated_name"`
	UnitCode               string `csv:"4" json:"unit_code"`
	UnitName               string `csv:"5" json:"unit_name"`
	MaxTimes               int    `csv:"6" json:"max_times"`
	MaxTimesErrorCategory  string `csv:"7" json:"max_times_error_category"`
	UpdateDate             string `csv:"8" json:"update_date"`
	ExpiryDate             string `csv:"9" json:"expiry_date"`
	Reserved10             string `csv:"10" json:"reserved_10"`
}

// VisitingNursingConflict 訪問看護・併算定背反テーブル (エ)
// 全13項目 (R06rec3.pdf 222ページ)
type VisitingNursingConflict struct {
	ChangeCategory          string `csv:"1" json:"change_category"`
	VisitingNursingFeeCode1 string `csv:"2" json:"visiting_nursing_fee_code_1"`
	AbbreviatedName1        string `csv:"3" json:"abbreviated_name_1"`
	ConflictCategory        string `csv:"4" json:"conflict_category"`
	VisitingNursingFeeCode2 string `csv:"5" json:"visiting_nursing_fee_code_2"`
	AbbreviatedName2        string `csv:"6" json:"abbreviated_name_2"`
	ConflictUnit            string `csv:"7" json:"conflict_unit"`
	SpecialCondition        string `csv:"8" json:"special_condition"`
	UpdateDate              string `csv:"9" json:"update_date"`
	ExpiryDate              string `csv:"10" json:"expiry_date"`
	Reserved11              string `csv:"11" json:"reserved_11"`
	Reserved12              string `csv:"12" json:"reserved_12"`
	Reserved13              string `csv:"13" json:"reserved_13"`
}

// VisitingNursingFacilityStandard 訪問看護・施設基準テーブル (オ)
// 全9項目 (R06rec3.pdf 223ページ)
type VisitingNursingFacilityStandard struct {
	ChangeCategory         string `csv:"1" json:"change_category"`
	GroupNumber            string `csv:"2" json:"group_number"`
	VisitingNursingFeeCode string `csv:"3" json:"visiting_nursing_fee_code"`
	AbbreviatedName        string `csv:"4" json:"abbreviated_name"`
	FacilityStandardCode   string `csv:"5" json:"facility_standard_code"`
	FacilityIdentifier     string `csv:"6" json:"facility_identifier"`
	UpdateDate             string `csv:"7" json:"update_date"`
	ExpiryDate             string `csv:"8" json:"expiry_date"`
	Reserved9              string `csv:"9" json:"reserved_9"`
}
