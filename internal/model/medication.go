package model

// Medication 調剤行為マスター (m*.csv)
// 全65項目を網羅 (R06rec3.pdf 216-217ページに準拠)
type Medication struct {
	ChangeCategory             string `csv:"1" json:"change_category"`               // 1: 変更区分
	MasterType                 string `csv:"2" json:"master_type"`                   // 2: マスター識別 ('M'固定)
	Code                       string `csv:"3" json:"code"`                          // 3: 調剤行為コード
	KanjiNameLength            int    `csv:"4" json:"kanji_name_length"`             // 4: 漢字有効桁数
	KanjiName                  string `csv:"5" json:"kanji_name"`                    // 5: 漢字名称
	KanaNameLength             int    `csv:"6" json:"kana_name_length"`              // 6: カナ有効桁数
	KanaName                   string `csv:"7" json:"kana_name"`                     // 7: カナ名称
	ReceiptSymbolCode          string `csv:"8" json:"receipt_symbol_code"`           // 8: レセプト表示用記号コード
	ReceiptOrderNumber         int    `csv:"9" json:"receipt_order_number"`          // 9: レセプト表示順番号
	ScoreType                  string `csv:"10" json:"score_type"`                   // 10: 新又は現点数点数識別
	QuantityCalculationFlag    string `csv:"11" json:"quantity_calculation_flag"`    // 11: 調剤数量計算フラグ
	BasicScore                 int    `csv:"12" json:"basic_score"`                  // 12: 新又は現点数（基本点数）
	StepCalculationCategory    string `csv:"13" json:"step_calculation_category"`    // 13: きざみ値計算識別
	StepLowerLimit             int    `csv:"14" json:"step_lower_limit"`             // 14: 下限値
	StepUpperLimit             int    `csv:"15" json:"step_upper_limit"`             // 15: 上限値
	StepValue                  int    `csv:"16" json:"step_value"`                   // 16: きざみ値
	StepScore                  int    `csv:"17" json:"step_score"`                   // 17: きざみ点数
	StepErrorCategory          string `csv:"18" json:"step_error_category"`          // 18: 上下限エラー処理
	ReductionCategory          string `csv:"19" json:"reduction_category"`           // 19: 減算行為区分
	ReductionTargetCategory    string `csv:"20" json:"reduction_target_category"`    // 20: 減算対象行為区分
	DispensingCategory         string `csv:"21" json:"dispensing_category"`          // 21: 調剤行為識別区分
	ComprehensiveCategory      string `csv:"22" json:"comprehensive_category"`       // 22: 包括識別区分
	Reserved23                 string `csv:"23" json:"reserved_23"`                  // 23: 予備
	Reserved24                 string `csv:"24" json:"reserved_24"`                  // 24: 予備
	Reserved25                 string `csv:"25" json:"reserved_25"`                  // 25: 予備
	Reserved26                 string `csv:"26" json:"reserved_26"`                  // 26: 予備
	Reserved27                 string `csv:"27" json:"reserved_27"`                  // 27: 予備
	DispensingType1            string `csv:"28" json:"dispensing_type_1"`            // 28: 調剤行為種類（1）
	DispensingType2            string `csv:"29" json:"dispensing_type_2"`            // 29: 調剤行為種類（2）
	ElderlyMedicalCategory     string `csv:"30" json:"elderly_medical_category"`     // 30: 後期高齢者適用区分
	FacilityStandard1          string `csv:"31" json:"facility_standard_1"`          // 31: 施設基準1
	FacilityStandard2          string `csv:"32" json:"facility_standard_2"`          // 32: 施設基準2
	FacilityStandard3          string `csv:"33" json:"facility_standard_3"`          // 33: 施設基準3
	FacilityStandard4          string `csv:"34" json:"facility_standard_4"`          // 34: 施設基準4
	FacilityStandard5          string `csv:"35" json:"facility_standard_5"`          // 35: 施設基準5
	FacilityStandard6          string `csv:"36" json:"facility_standard_6"`          // 36: 施設基準6
	FacilityStandard7          string `csv:"37" json:"facility_standard_7"`          // 37: 施設基準7
	FacilityStandard8          string `csv:"38" json:"facility_standard_8"`          // 38: 施設基準8
	FacilityStandard9          string `csv:"39" json:"facility_standard_9"`          // 39: 施設基準9
	FacilityStandard10         string `csv:"40" json:"facility_standard_10"`         // 40: 施設基準10
	ReceiptConflictCode        string `csv:"41" json:"receipt_conflict_code"`        // 41: レセプト単位背反区分コード
	PrescriptionConflictCode   string `csv:"42" json:"prescription_conflict_code"`   // 42: 処方箋受付回単位背反区分コード
	DispensingConflictCode     string `csv:"43" json:"dispensing_conflict_code"`     // 43: 調剤単位背反区分コード
	NarcoticToxicDrugCategory  string `csv:"44" json:"narcotic_toxic_drug_category"` // 44: 麻薬・毒薬・覚醒剤原料・向精神薬
	TimeAdditionCategory       string `csv:"45" json:"time_addition_category"`       // 45: 時間加算区分
	DosageForm                 string `csv:"46" json:"dosage_form"`                  // 46: 剤形
	ReceiptMaxTimes            int    `csv:"47" json:"receipt_max_times"`            // 47: レセプト単位上限回数
	ReceiptMaxTimesError       string `csv:"48" json:"receipt_max_times_error"`      // 48: レセプト単位上限回数エラー処理
	PrescriptionMaxTimes       int    `csv:"49" json:"prescription_max_times"`       // 49: 処方箋受付回単位上限回数
	PrescriptionMaxTimesError  string `csv:"50" json:"prescription_max_times_error"` // 50: 処方箋受付回単位上限回数エラー処理
	NoteAdditionCode           string `csv:"51" json:"note_addition_code"`           // 51: 注加算コード
	NoteAdditionSequence       string `csv:"52" json:"note_addition_sequence"`       // 52: 注加算通番
	LowerAge                   string `csv:"53" json:"lower_age"`                    // 53: 下限年齢
	UpperAge                   string `csv:"54" json:"upper_age"`                    // 54: 上限年齢
	MedicineManagementCategory string `csv:"55" json:"medicine_management_category"` // 55: 薬学管理料区分
	NotificationCategory1      string `csv:"56" json:"notification_category_1"`      // 56: 告示等識別区分（1）
	NotificationCategory2      string `csv:"57" json:"notification_category_2"`      // 57: 告示等識別区分（2）
	OldScoreType               string `csv:"58" json:"old_score_type"`               // 58: 旧点数点数識別
	OldScore                   int    `csv:"59" json:"old_score"`                    // 59: 旧点数
	UpdateDate                 string `csv:"60" json:"update_date"`                  // 60: 変更年月日
	ExpiryDate                 string `csv:"61" json:"expiry_date"`                  // 61: 廃止年月日
	PublicationOrder           int    `csv:"62" json:"publication_order"`            // 62: 公表順序番号
	SectionCode                string `csv:"63" json:"section_code"`                 // 63: コード表用番号
	RelatedSectionCode         string `csv:"64" json:"related_section_code"`         // 64: 告示・通知関連番号
	MigrationRelatedCode       string `csv:"65" json:"migration_related_code"`       // 65: 異動関連
}
