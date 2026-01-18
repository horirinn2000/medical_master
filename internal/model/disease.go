package model

// Disease 傷病名マスター (b_*.txt)
// 全46項目を網羅 (R06rec3.pdf 192-193ページに準拠)
type Disease struct {
	ChangeCategory                string `csv:"1" json:"change_category"`                   // 1: 変更区分
	MasterType                    string `csv:"2" json:"master_type"`                       // 2: マスター種別
	Code                          string `csv:"3" json:"code"`                              // 3: 傷病名コード
	MigrationCode                 string `csv:"4" json:"migration_code"`                    // 4: 移行先コード
	BasicNameLength               int    `csv:"5" json:"basic_name_length"`                 // 5: 傷病名基本名称桁数
	BasicName                     string `csv:"6" json:"basic_name"`                        // 6: 傷病名基本名称
	AbbreviatedNameLength         int    `csv:"7" json:"abbreviated_name_length"`           // 7: 傷病名省略名称桁数
	AbbreviatedName               string `csv:"8" json:"abbreviated_name"`                  // 8: 傷病名省略名称
	KanaNameLength                int    `csv:"9" json:"kana_name_length"`                  // 9: 傷病名カナ名称桁数
	KanaName                      string `csv:"10" json:"kana_name"`                        // 10: 傷病名カナ名称
	ManagementNumber              string `csv:"11" json:"management_number"`                // 11: 病名管理番号
	SelectionCategory             string `csv:"12" json:"selection_category"`               // 12: 採択区分
	ExchangeCode                  string `csv:"13" json:"exchange_code"`                    // 13: 病名交換用コード
	Reserved14                    string `csv:"14" json:"reserved_14"`                      // 14: 予備
	Reserved15                    string `csv:"15" json:"reserved_15"`                      // 15: 予備
	ICD10_1                       string `csv:"16" json:"icd10_1"`                          // 16: ICD-10-1 (2013)
	ICD10_2                       string `csv:"17" json:"icd10_2"`                          // 17: ICD-10-2 (2013)
	Reserved18                    string `csv:"18" json:"reserved_18"`                      // 18: 予備
	SingleUseForbiddenCategory    string `csv:"19" json:"single_use_forbidden_category"`    // 19: 単独使用禁止区分
	InsuranceOutCategory          string `csv:"20" json:"insurance_out_category"`           // 20: 保険請求外区分
	SpecificDiseaseCategory       string `csv:"21" json:"specific_disease_category"`        // 21: 特定疾患等対象区分
	ListingDate                   string `csv:"22" json:"listing_date"`                     // 22: 収載年月日
	UpdateDate                    string `csv:"23" json:"update_date"`                      // 23: 変更年月日
	ExpiryDate                    string `csv:"24" json:"expiry_date"`                      // 24: 廃止年月日
	BasicNameUpdateFlag           string `csv:"25" json:"basic_name_update_flag"`           // 25: 傷病名基本名称（変更情報）
	AbbreviatedNameUpdateFlag     string `csv:"26" json:"abbreviated_name_update_flag"`     // 26: 傷病名省略名称（変更情報）
	KanaNameUpdateFlag            string `csv:"27" json:"kana_name_update_flag"`            // 27: 傷病名カナ名称（変更情報）
	SelectionCategoryUpdateFlag   string `csv:"28" json:"selection_category_update_flag"`   // 28: 採択区分（変更情報）
	ExchangeCodeUpdateFlag        string `csv:"29" json:"exchange_code_update_flag"`        // 29: 病名交換用コード（変更情報）
	Reserved30                    string `csv:"30" json:"reserved_30"`                      // 30: 予備
	Reserved31                    string `csv:"31" json:"reserved_31"`                      // 31: 予備
	DentalAbbreviatedUpdateFlag   string `csv:"32" json:"dental_abbreviated_update_flag"`   // 32: 歯科傷病名省略名称（変更情報）
	IncurableDiseaseUpdateFlag    string `csv:"33" json:"incurable_disease_update_flag"`    // 33: 難病外来対象区分（変更情報）
	DentalSpecificUpdateFlag      string `csv:"34" json:"dental_specific_update_flag"`      // 34: 歯科特定疾患対象区分（変更情報）
	SingleUseForbiddenUpdateFlag  string `csv:"35" json:"single_use_forbidden_update_flag"` // 35: 単独使用禁止区分（変更情報）
	InsuranceOutUpdateFlag        string `csv:"36" json:"insurance_out_update_flag"`        // 36: 保険請求外区分（変更情報）
	SpecificDiseaseUpdateFlag     string `csv:"37" json:"specific_disease_update_flag"`     // 37: 特定疾患等対象区分（変更情報）
	MigrationManagementNumber     string `csv:"38" json:"migration_management_number"`      // 38: 移行先病名管理番号
	DentalAbbreviatedName         string `csv:"39" json:"dental_abbreviated_name"`          // 39: 歯科傷病名省略名称
	Reserved40                    string `csv:"40" json:"reserved_40"`                      // 40: 予備
	Reserved41                    string `csv:"41" json:"reserved_41"`                      // 41: 予備
	DentalAbbreviatedNameLength   int    `csv:"42" json:"dental_abbreviated_name_length"`   // 42: 歯科傷病名省略名称桁数
	IncurableDiseaseCategory      string `csv:"43" json:"incurable_disease_category"`       // 43: 難病外来対象区分
	DentalSpecificDiseaseCategory string `csv:"44" json:"dental_specific_disease_category"` // 44: 歯科特定疾患対象区分
	ICD10_1UpdateFlag             string `csv:"45" json:"icd10_1_update_flag"`              // 45: ICD-10-1 (2013)（変更情報）
	ICD10_2UpdateFlag             string `csv:"46" json:"icd10_2_update_flag"`              // 46: ICD-10-2 (2013)（変更情報）

	IsOld bool `gorm:"-" json:"is_old"` // 旧傷病名管理ファイル由来の場合
}
