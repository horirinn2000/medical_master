package model

import "time"

type Disease struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 管理用フラグ
	IsOld           bool     `gorm:"index" json:"is_old"`
	MigrationTarget *Disease `gorm:"-" json:"migration_target,omitempty"`

	// 傷病名マスター仕様 (46カラム)
	UpdateCategory                string `json:"update_category"`                          // 1: 変更区分
	MasterType                    string `json:"master_type"`                              // 2: マスター種別
	Code                          string `gorm:"index" json:"code"`                        // 3: 傷病名コード
	RelatedCode                   string `json:"related_code"`                             // 4: 移行先コード
	NameKanjiLen                  int    `json:"name_kanji_len"`                           // 5: 傷病名基本名称桁数
	NameKanji                     string `gorm:"index" json:"name_kanji"`                  // 6: 傷病名基本名称
	NameAbbrLen                   int    `json:"name_abbr_len"`                            // 7: 傷病名省略名称桁数
	NameAbbr                      string `gorm:"index" json:"name_abbr"`                   // 8: 傷病名省略名称
	NameKanaLen                   int    `json:"name_kana_len"`                            // 9: 傷病名カナ名称桁数
	NameKana                      string `gorm:"index" json:"name_kana"`                   // 10: 傷病名カナ名称
	DiseaseManagementNumber       string `gorm:"index" json:"disease_management_number"`   // 11: 病名管理番号
	AdoptionCategory              string `json:"adoption_category"`                        // 12: 採択区分
	DiseaseExchangeCode           string `json:"disease_exchange_code"`                    // 13: 病名交換用コード
	Reserved1                     string `json:"reserved_1"`                               // 14: 予備
	Reserved2                     string `json:"reserved_2"`                               // 15: 予備
	Icd10_1                       string `gorm:"index" json:"icd10_1"`                     // 16: ICD-10-1
	Icd10_2                       string `gorm:"index" json:"icd10_2"`                     // 17: ICD-10-2
	Reserved3                     string `json:"reserved_3"`                               // 18: 予備
	SingleUseForbiddenCategory    string `json:"single_use_forbidden_category"`            // 19: 単独使用禁止区分
	InsuranceOutCategory          string `json:"insurance_out_category"`                   // 20: 保険請求外区分
	SpecificDiseaseCategory       string `json:"specific_disease_category"`                // 21: 特定疾患等対象区分
	ImportDate                    string `json:"import_date"`                              // 22: 収載年月日
	UpdateDate                    string `json:"update_date"`                              // 23: 変更年月日
	DiscontinuedDate              string `json:"discontinued_date"`                        // 24: 廃止年月日
	NameKanjiUpdateFlag           string `json:"name_kanji_update_flag"`                   // 25: 傷病名基本名称（変更情報）
	NameAbbrUpdateFlag            string `json:"name_abbr_update_flag"`                    // 26: 傷病名省略名称（変更情報）
	NameKanaUpdateFlag            string `json:"name_kana_update_flag"`                    // 27: 傷病名カナ名称（変更情報）
	AdoptionUpdateFlag            string `json:"adoption_update_flag"`                     // 28: 採択区分（変更情報）
	DiseaseExchangeUpdateFlag     string `json:"disease_exchange_update_flag"`             // 29: 病名交換用コード（変更情報）
	Reserved4                     string `json:"reserved_4"`                               // 30: 予備
	Reserved5                     string `json:"reserved_5"`                               // 31: 予備
	DentalNameAbbrUpdateFlag      string `json:"dental_name_abbr_update_flag"`             // 32: 歯科傷病名省略名称（変更情報）
	IncurableDiseaseUpdateFlag    string `json:"incurable_disease_update_flag"`            // 33: 難病外来対象区分（変更情報）
	DentalSpecificUpdateFlag      string `json:"dental_specific_update_flag"`              // 34: 歯科特定疾患対象区分（変更情報）
	SingleUseForbiddenUpdateFlag  string `json:"single_use_forbidden_update_flag"`         // 35: 単独使用禁止区分（変更情報）
	InsuranceOutUpdateFlag        string `json:"insurance_out_update_flag"`                // 36: 保険請求外区分（変更情報）
	SpecificDiseaseUpdateFlag     string `json:"specific_disease_update_flag"`             // 37: 特定疾患等対象区分（変更情報）
	MigrationManagementNumber     string `gorm:"index" json:"migration_management_number"` // 38: 移行先病名管理番号
	DentalNameAbbr                string `json:"dental_name_abbr"`                         // 39: 歯科傷病名省略名称
	Reserved6                     string `json:"reserved_6"`                               // 40: 予備
	Reserved7                     string `json:"reserved_7"`                               // 41: 予備
	DentalNameAbbrLen             int    `json:"dental_name_abbr_len"`                     // 42: 歯科傷病名省略名称桁数
	IncurableDiseaseCategory      string `json:"incurable_disease_category"`               // 43: 難病外来対象区分
	DentalSpecificDiseaseCategory string `json:"dental_specific_disease_category"`         // 44: 歯科特定疾患対象区分
	Icd10_1UpdateFlag             string `json:"icd10_1_update_flag"`                      // 45: ICD-10-1（変更情報）
	Icd10_2UpdateFlag             string `json:"icd10_2_update_flag"`                      // 46: ICD-10-2（変更情報）
}
