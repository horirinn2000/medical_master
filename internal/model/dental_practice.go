package model

// DentalPractice 歯科診療行為マスター (h_*.csv)
type DentalPractice struct {
	ChangeCategory                  string  `csv:"1" json:"change_category"`                     // 変更区分
	MasterType                      string  `csv:"2" json:"master_type"`                         // マスター種別
	DentalPracticeCode              string  `csv:"3" json:"dental_practice_code"`                // 歯科診療行為コード
	SectionChar                     string  `csv:"4" json:"section_char"`                        // 区分
	SectionNumber                   string  `csv:"5" json:"section_number"`                      // 区分番号
	SectionBranchNumber             string  `csv:"6" json:"section_branch_number"`               // 枝番
	SectionItemNumber               string  `csv:"7" json:"section_item_number"`                 // 項番
	AdditionCode                    string  `csv:"8" json:"addition_code"`                       // 加算コード
	BasicKanjiName                  string  `csv:"9" json:"basic_kanji_name"`                    // 基本名称
	AbbreviatedKanjiName            string  `csv:"10" json:"abbreviated_kanji_name"`             // 省略名称
	ScoreType                       string  `csv:"11" json:"score_type"`                         // 点数等識別
	Score                           float64 `csv:"12" json:"score"`                              // 点数等
	OldScoreType                    string  `csv:"13" json:"old_score_type"`                     // 点数等識別（旧）
	OldScore                        float64 `csv:"14" json:"old_score"`                          // 旧点数等
	InpatientOutpatientCategory     string  `csv:"15" json:"inpatient_outpatient_category"`      // 入外適用区分
	ElderlyMedicalCategory          string  `csv:"16" json:"elderly_medical_category"`           // 後期高齢者医療適用区分
	TimeAdditionCategory            string  `csv:"17" json:"time_addition_category"`             // 時間加算区分
	HospitalClinicCategory          string  `csv:"18" json:"hospital_clinic_category"`           // 病院・診療所区分
	NursingAddition                 string  `csv:"19" json:"nursing_addition"`                   // 看護加算
	Reserved1                       string  `csv:"20" json:"reserved_1"`                         // 予備
	Reserved2                       string  `csv:"21" json:"reserved_2"`                         // 予備
	RegionAddition                  string  `csv:"22" json:"region_addition"`                    // 地域加算
	DiseaseRelatedCategory          string  `csv:"23" json:"disease_related_category"`           // 傷病名関連区分
	MedicineRelatedCategory         string  `csv:"24" json:"medicine_related_category"`          // 医薬品関連区分
	BedCountCategory                string  `csv:"25" json:"bed_count_category"`                 // 病床数区分
	NotificationCategory            string  `csv:"26" json:"notification_category"`              // 届出
	FutureVisitCategory             string  `csv:"27" json:"future_visit_category"`              // 未来院
	ShortStaySurgery                string  `csv:"28" json:"short_stay_surgery"`                 // 短期滞在手術
	SpecialNote                     string  `csv:"29" json:"special_note"`                       // 特記事項
	TestJudgmentCategory            string  `csv:"30" json:"test_judgment_category"`             // 検査等実施判断区分
	TestJudgmentGroupCategory       string  `csv:"31" json:"test_judgment_group_category"`       // 検査等実施判断グループ区分
	ReductionTargetCategory         string  `csv:"32" json:"reduction_target_category"`          // 逓減対象区分
	ComprehensiveReductionCategory  string  `csv:"33" json:"comprehensive_reduction_category"`   // 包括逓減区分
	ConformityCategory              string  `csv:"34" json:"conformity_category"`                // 適合区分
	TargetFacilityStandard          string  `csv:"35" json:"target_facility_standard"`           // 対象施設基準
	FacilityStandardCode1           string  `csv:"36" json:"facility_standard_code_1"`           // 施設基準①
	FacilityStandardCode2           string  `csv:"37" json:"facility_standard_code_2"`           // 施設基準②
	FacilityStandardCode3           string  `csv:"38" json:"facility_standard_code_3"`           // 施設基準③
	FacilityStandardCode4           string  `csv:"39" json:"facility_standard_code_4"`           // 施設基準④
	FacilityStandardCode5           string  `csv:"40" json:"facility_standard_code_5"`           // 施設基準⑤
	FacilityStandardCode6           string  `csv:"41" json:"facility_standard_code_6"`           // 施設基準⑥
	FacilityStandardCode7           string  `csv:"42" json:"facility_standard_code_7"`           // 施設基準⑦
	FacilityStandardCode8           string  `csv:"43" json:"facility_standard_code_8"`           // 施設基準⑧
	FacilityStandardCode9           string  `csv:"44" json:"facility_standard_code_9"`           // 施設基準⑨
	FacilityStandardCode10          string  `csv:"45" json:"facility_standard_code_10"`          // 施設基準⑩
	GeneralAdditionGroup            string  `csv:"46" json:"general_addition_group"`             // 通則加算グループ
	BasicAdditionGroup              string  `csv:"47" json:"basic_addition_group"`               // 基本加算グループ
	NoteAdditionGroup               string  `csv:"48" json:"note_addition_group"`                // 注加算グループ
	TechniqueMaterialAdditionGroup  string  `csv:"49" json:"technique_material_addition_group"`  // 手技・材料加算グループ
	CountLimitTableId               string  `csv:"50" json:"count_limit_table_id"`               // 算定回数限度テーブル関連識別
	StepTableId                     string  `csv:"51" json:"step_table_id"`                      // きざみテーブル関連識別
	AgeLimitTableId                 string  `csv:"52" json:"age_limit_table_id"`                 // 年齢制限テーブル関連識別
	ConflictTableId                 string  `csv:"53" json:"conflict_table_id"`                  // 併算定背反テーブル関連識別
	ActualDaysTableId               string  `csv:"54" json:"actual_days_table_id"`               // 実日数テーブル関連識別
	Reserved3                       string  `csv:"55" json:"reserved_3"`                         // 予備
	Reserved4                       string  `csv:"56" json:"reserved_4"`                         // 予備
	UpdateDate                      string  `csv:"57" json:"update_date"`                        // 変更年月日
	ExpiryDate                      string  `csv:"58" json:"expiry_date"`                        // 廃止年月日
	LongTermAnesthesiaAddition      string  `csv:"59" json:"long_term_anesthesia_addition"`      // 長時間麻酔管理加算
	MalignantTumorPathologyAddition string  `csv:"60" json:"malignant_tumor_pathology_addition"` // 悪性腫瘍病理組織標本加算
	NursingStaffTreatmentEvaluation string  `csv:"61" json:"nursing_staff_treatment_evaluation"` // 看護職員処遇改善評価料等
	DentalOutpatientBaseUp1         string  `csv:"62" json:"dental_outpatient_base_up_1"`        // 歯科外来・在宅ベースアップ評価料（１）
	DentalOutpatientBaseUp2         string  `csv:"63" json:"dental_outpatient_base_up_2"`        // 歯科外来・在宅ベースアップ評価料（２）
	Reserved5                       string  `csv:"64" json:"reserved_5"`                         // 予備４
	Reserved6                       string  `csv:"65" json:"reserved_6"`                         // 予備５
	PublicationOrder                string  `csv:"66" json:"publication_order"`                  // 公表順序番号
}

// DentalPracticeAdditionRelation 加算対応テーブル (h-2, h-3, h-4, h-5)
type DentalPracticeAdditionRelation struct {
	ChangeCategory     string `csv:"1" json:"change_category"`      // 変更区分
	GroupNumber        string `csv:"2" json:"group_number"`         // グループ番号
	AdditionCode       string `csv:"3" json:"addition_code"`        // 加算コード
	DentalPracticeCode string `csv:"4" json:"dental_practice_code"` // 歯科診療行為コード
	BasicKanjiName     string `csv:"5" json:"basic_kanji_name"`     // 基本名称
	AdditionIdentifier string `csv:"6" json:"addition_identifier"`  // 加算識別
	UpdateDate         string `csv:"7" json:"update_date"`          // 変更年月日
	ExpiryDate         string `csv:"8" json:"expiry_date"`          // 廃止年月日
	Reserved           string `csv:"9" json:"reserved"`             // 予備
}

// DentalPracticeCalculationCount 算定回数限度テーブル (h-6)
type DentalPracticeCalculationCount struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 変更区分
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`   // 歯科診療行為コード
	SectionChar          string `csv:"3" json:"section_char"`           // 区分
	SectionNumber        string `csv:"4" json:"section_number"`         // 区分番号
	SectionBranchNumber  string `csv:"5" json:"section_branch_number"`  // 枝番
	SectionItemNumber    string `csv:"6" json:"section_item_number"`    // 項番
	AdditionCode         string `csv:"7" json:"addition_code"`          // 加算コード
	BasicKanjiName       string `csv:"8" json:"basic_kanji_name"`       // 基本名称
	AbbreviatedKanjiName string `csv:"9" json:"abbreviated_kanji_name"` // 省略名称
	UnitCode             string `csv:"10" json:"unit_code"`             // 算定単位
	CountLimit           int    `csv:"11" json:"count_limit"`           // 算定回数限度
	CountErrorTreatment  string `csv:"12" json:"count_error_treatment"` // 上限回数エラー処理
	UpdateDate           string `csv:"13" json:"update_date"`           // 変更年月日
	ExpiryDate           string `csv:"14" json:"expiry_date"`           // 廃止年月日
	Reserved             string `csv:"15" json:"reserved"`              // 予備
}

// DentalPracticeStep きざみテーブル (h-7)
type DentalPracticeStep struct {
	ChangeCategory       string  `csv:"1" json:"change_category"`        // 変更区分
	DentalPracticeCode   string  `csv:"2" json:"dental_practice_code"`   // 歯科診療行為コード
	SectionChar          string  `csv:"3" json:"section_char"`           // 区分
	SectionNumber        string  `csv:"4" json:"section_number"`         // 区分番号
	SectionBranchNumber  string  `csv:"5" json:"section_branch_number"`  // 枝番
	SectionItemNumber    string  `csv:"6" json:"section_item_number"`    // 項番
	AdditionCode         string  `csv:"7" json:"addition_code"`          // 加算コード
	BasicKanjiName       string  `csv:"8" json:"basic_kanji_name"`       // 基本名称
	AbbreviatedKanjiName string  `csv:"9" json:"abbreviated_kanji_name"` // 省略名称
	ScoreType            string  `csv:"10" json:"score_type"`            // 点数等識別
	Score                float64 `csv:"11" json:"score"`                 // 点数
	StepUnit             string  `csv:"12" json:"step_unit"`             // きざみ単位
	StepLowerLimit       float64 `csv:"13" json:"step_limit"`            // きざみ下限値
	StepUpperLimit       float64 `csv:"14" json:"step_upper_limit"`      // きざみ上限値
	StepValue            float64 `csv:"15" json:"step_value"`            // きざみ値
	StepScore            float64 `csv:"16" json:"step_score"`            // きざみ点数
	StepErrorTreatment   string  `csv:"17" json:"step_error_treatment"`  // きざみ上下限エラー処理
	UpdateDate           string  `csv:"18" json:"update_date"`           // 変更年月日
	ExpiryDate           string  `csv:"19" json:"expiry_date"`           // 廃止年月日
	Reserved             string  `csv:"20" json:"reserved"`              // 予備
}

// DentalPracticeAgeConstraint 年齢制限テーブル (h-8)
type DentalPracticeAgeConstraint struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 変更区分
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`   // 歯科診療行為コード
	SectionChar          string `csv:"3" json:"section_char"`           // 区分
	SectionNumber        string `csv:"4" json:"section_number"`         // 区分番号
	SectionBranchNumber  string `csv:"5" json:"section_branch_number"`  // 枝番
	SectionItemNumber    string `csv:"6" json:"section_item_number"`    // 項番
	AdditionCode         string `csv:"7" json:"addition_code"`          // 加算コード
	BasicKanjiName       string `csv:"8" json:"basic_kanji_name"`       // 基本名称
	AbbreviatedKanjiName string `csv:"9" json:"abbreviated_kanji_name"` // 省略名称
	LowerAge             string `csv:"10" json:"lower_age"`             // 下限年齢
	UpperAge             string `csv:"11" json:"upper_age"`             // 上限年齢
	UpdateDate           string `csv:"12" json:"update_date"`           // 変更年月日
	ExpiryDate           string `csv:"13" json:"expiry_date"`           // 廃止年月日
	Reserved             string `csv:"14" json:"reserved"`              // 予備
}

// DentalPracticeConflict 併算定背反テーブル (h-9)
type DentalPracticeConflict struct {
	ChangeCategory       string `csv:"1" json:"change_category"`              // 変更区分
	DentalPracticeCode   string `csv:"2" json:"source_code"`                  // 歯科診療行為コード
	SectionChar          string `csv:"3" json:"source_type"`                  // 区分
	SectionNumber        string `csv:"4" json:"source_section"`               // 区分番号
	SectionBranchNumber  string `csv:"5" json:"section_branch_number"`        // 枝番
	SectionItemNumber    string `csv:"6" json:"section_item_number"`          // 項番
	AdditionCode         string `csv:"7" json:"addition_code"`                // 加算コード
	BasicKanjiName       string `csv:"8" json:"source_name1"`                 // 基本名称
	AbbreviatedKanjiName string `csv:"9" json:"source_name2"`                 // 省略名称
	Conflict1_Flag       string `csv:"10" json:"conflict1_flag"`              // 背反1:算定可否
	Conflict1_Code       string `csv:"11" json:"conflict1_code"`              // 背反1:歯科診療行為コード
	Conflict1_Section    string `csv:"12" json:"conflict1_section_char"`      // 背反1:区分
	Conflict1_SecNum     string `csv:"13" json:"conflict1_section_number"`    // 背反1:区分番号
	Conflict1_Branch     string `csv:"14" json:"conflict1_branch_number"`     // 背反1:枝番
	Conflict1_Item       string `csv:"15" json:"conflict1_item_number"`       // 背反1:項番
	Conflict1_AddCode    string `csv:"16" json:"conflict1_addition_code"`     // 背反1:加算コード
	Conflict1_Name       string `csv:"17" json:"conflict1_name"`              // 背反1:基本名称
	Conflict1_Abbr       string `csv:"18" json:"conflict1_abbreviated_name"`  // 背反1:省略名称
	Conflict2_Flag       string `csv:"19" json:"conflict2_flag"`              // 背反2:算定可否
	Conflict2_Code       string `csv:"20" json:"conflict2_code"`              // 背反2:歯科診療行為コード
	Conflict2_Section    string `csv:"21" json:"conflict2_section_char"`      // 背反2:区分
	Conflict2_SecNum     string `csv:"22" json:"conflict2_section_number"`    // 背反2:区分番号
	Conflict2_Branch     string `csv:"23" json:"conflict2_branch_number"`     // 背反2:枝番
	Conflict2_Item       string `csv:"24" json:"conflict2_item_number"`       // 背反2:項番
	Conflict2_AddCode    string `csv:"25" json:"conflict2_addition_code"`     // 背反2:加算コード
	Conflict2_Name       string `csv:"26" json:"conflict2_name"`              // 背反2:基本名称
	Conflict2_Abbr       string `csv:"27" json:"conflict2_abbreviated_name"`  // 背反2:省略名称
	Conflict3_Flag       string `csv:"28" json:"conflict3_flag"`              // 背反3:算定可否
	Conflict3_Code       string `csv:"29" json:"conflict3_code"`              // 背反3:歯科診療行為コード
	Conflict3_Section    string `csv:"30" json:"conflict3_section_char"`      // 背反3:区分
	Conflict3_SecNum     string `csv:"31" json:"conflict3_section_number"`    // 背反3:区分番号
	Conflict3_Branch     string `csv:"32" json:"conflict3_branch_number"`     // 背反3:枝番
	Conflict3_Item       string `csv:"33" json:"conflict3_item_number"`       // 背反3:項番
	Conflict3_AddCode    string `csv:"34" json:"conflict3_addition_code"`     // 背反3:加算コード
	Conflict3_Name       string `csv:"35" json:"conflict3_name"`              // 背反3:基本名称
	Conflict3_Abbr       string `csv:"36" json:"conflict3_abbreviated_name"`  // 背反3:省略名称
	Conflict4_Flag       string `csv:"37" json:"conflict4_flag"`              // 背反4:算定可否
	Conflict4_Code       string `csv:"38" json:"conflict4_code"`              // 背反4:歯科診療行為コード
	Conflict4_Section    string `csv:"39" json:"conflict4_section_char"`      // 背反4:区分
	Conflict4_SecNum     string `csv:"40" json:"conflict4_section_number"`    // 背反4:区分番号
	Conflict4_Branch     string `csv:"41" json:"conflict4_branch_number"`     // 背反4:枝番
	Conflict4_Item       string `csv:"42" json:"conflict4_item_number"`       // 背反4:項番
	Conflict4_AddCode    string `csv:"43" json:"conflict4_addition_code"`     // 背反4:加算コード
	Conflict4_Name       string `csv:"44" json:"conflict4_name"`              // 背反4:基本名称
	Conflict4_Abbr       string `csv:"45" json:"conflict4_abbreviated_name"`  // 背反4:省略名称
	Conflict5_Flag       string `csv:"46" json:"conflict5_flag"`              // 背反5:算定可否
	Conflict5_Code       string `csv:"47" json:"conflict5_code"`              // 背反5:歯科診療行為コード
	Conflict5_Section    string `csv:"48" json:"conflict5_section_char"`      // 背反5:区分
	Conflict5_SecNum     string `csv:"49" json:"conflict5_section_number"`    // 背反5:区分番号
	Conflict5_Branch     string `csv:"50" json:"conflict5_branch_number"`     // 背反5:枝番
	Conflict5_Item       string `csv:"51" json:"conflict5_item_number"`       // 背反5:項番
	Conflict5_AddCode    string `csv:"52" json:"conflict5_addition_code"`     // 背反5:加算コード
	Conflict5_Name       string `csv:"53" json:"conflict5_name"`              // 背反5:基本名称
	Conflict5_Abbr       string `csv:"54" json:"conflict5_abbreviated_name"`  // 背反5:省略名称
	Conflict6_Flag       string `csv:"55" json:"conflict6_flag"`              // 背反6:算定可否
	Conflict6_Code       string `csv:"56" json:"conflict6_code"`              // 背反6:歯科診療行為コード
	Conflict6_Section    string `csv:"57" json:"conflict6_section_char"`      // 背反6:区分
	Conflict6_SecNum     string `csv:"58" json:"conflict6_section_number"`    // 背反6:区分番号
	Conflict6_Branch     string `csv:"59" json:"conflict6_branch_number"`     // 背反6:枝番
	Conflict6_Item       string `csv:"60" json:"conflict6_item_number"`       // 背反6:項番
	Conflict6_AddCode    string `csv:"61" json:"conflict6_addition_code"`     // 背反6:加算コード
	Conflict6_Name       string `csv:"62" json:"conflict6_name"`              // 背反6:基本名称
	Conflict6_Abbr       string `csv:"63" json:"conflict6_abbreviated_name"`  // 背反6:省略名称
	Conflict7_Flag       string `csv:"64" json:"conflict7_flag"`              // 背反7:算定可否
	Conflict7_Code       string `csv:"65" json:"conflict7_code"`              // 背反7:歯科診療行為コード
	Conflict7_Section    string `csv:"66" json:"conflict7_section_char"`      // 背反7:区分
	Conflict7_SecNum     string `csv:"67" json:"conflict7_section_number"`    // 背反7:区分番号
	Conflict7_Branch     string `csv:"68" json:"conflict7_branch_number"`     // 背反7:枝番
	Conflict7_Item       string `csv:"69" json:"conflict7_item_number"`       // 背反7:項番
	Conflict7_AddCode    string `csv:"70" json:"conflict7_addition_code"`     // 背反7:加算コード
	Conflict7_Name       string `csv:"71" json:"conflict7_name"`              // 背反7:基本名称
	Conflict7_Abbr       string `csv:"72" json:"conflict7_abbreviated_name"`  // 背反7:省略名称
	Conflict8_Flag       string `csv:"73" json:"conflict8_flag"`              // 背反8:算定可否
	Conflict8_Code       string `csv:"74" json:"conflict8_code"`              // 背反8:歯科診療行為コード
	Conflict8_Section    string `csv:"75" json:"conflict8_section_char"`      // 背反8:区分
	Conflict8_SecNum     string `csv:"76" json:"conflict8_section_number"`    // 背反8:区分番号
	Conflict8_Branch     string `csv:"77" json:"conflict8_branch_number"`     // 背反8:枝番
	Conflict8_Item       string `csv:"78" json:"conflict8_item_number"`       // 背反8:項番
	Conflict8_AddCode    string `csv:"79" json:"conflict8_addition_code"`     // 背反8:加算コード
	Conflict8_Name       string `csv:"80" json:"conflict8_name"`              // 背反8:基本名称
	Conflict8_Abbr       string `csv:"81" json:"conflict8_abbreviated_name"`  // 背反8:省略名称
	Conflict9_Flag       string `csv:"82" json:"conflict9_flag"`              // 背反9:算定可否
	Conflict9_Code       string `csv:"83" json:"conflict9_code"`              // 背反9:歯科診療行為コード
	Conflict9_Section    string `csv:"84" json:"conflict9_section_char"`      // 背反9:区分
	Conflict9_SecNum     string `csv:"85" json:"conflict9_section_number"`    // 背反9:区分番号
	Conflict9_Branch     string `csv:"86" json:"conflict9_branch_number"`     // 背反9:枝番
	Conflict9_Item       string `csv:"87" json:"conflict9_item_number"`       // 背反9:項番
	Conflict9_AddCode    string `csv:"88" json:"conflict9_addition_code"`     // 背反9:加算コード
	Conflict9_Name       string `csv:"89" json:"conflict9_name"`              // 背反9:基本名称
	Conflict9_Abbr       string `csv:"90" json:"conflict9_abbreviated_name"`  // 背反9:省略名称
	Conflict10_Flag      string `csv:"91" json:"conflict10_flag"`             // 背反10:算定可否
	Conflict10_Code      string `csv:"92" json:"conflict10_code"`             // 背反10:歯科診療行為コード
	Conflict10_Section   string `csv:"93" json:"conflict10_section_char"`     // 背反10:区分
	Conflict10_SecNum    string `csv:"94" json:"conflict10_section_number"`   // 背反10:区分番号
	Conflict10_Branch    string `csv:"95" json:"conflict10_branch_number"`    // 背反10:枝番
	Conflict10_Item      string `csv:"96" json:"conflict10_item_number"`      // 背反10:項番
	Conflict10_AddCode   string `csv:"97" json:"conflict10_addition_code"`    // 背反10:加算コード
	Conflict10_Name      string `csv:"98" json:"conflict10_name"`             // 背反10:基本名称
	Conflict10_Abbr      string `csv:"99" json:"conflict10_abbreviated_name"` // 背反10:省略名称
	UpdateDate           string `csv:"100" json:"update_date"`                // 変更年月日
	ExpiryDate           string `csv:"101" json:"expiry_date"`                // 廃止年月日
	Reserved1            string `csv:"102" json:"reserved_1"`                 // 予備
	Reserved2            string `csv:"103" json:"reserved_2"`                 // 予備
	Reserved3            string `csv:"104" json:"reserved_3"`                 // 予備
}

// DentalPracticeActualDays 実日数関連テーブル (h-10)
type DentalPracticeActualDays struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 変更区分
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`   // 歯科診療行為コード
	SectionChar          string `csv:"3" json:"section_char"`           // 区分
	SectionNumber        string `csv:"4" json:"section_number"`         // 区分番号
	SectionBranchNumber  string `csv:"5" json:"section_branch_number"`  // 枝番
	SectionItemNumber    string `csv:"6" json:"section_item_number"`    // 項番
	AdditionCode         string `csv:"7" json:"addition_code"`          // 加算コード
	BasicKanjiName       string `csv:"8" json:"basic_kanji_name"`       // 基本名称
	AbbreviatedKanjiName string `csv:"9" json:"abbreviated_kanji_name"` // 省略名称
	ActualDays           string `csv:"10" json:"days_limit"`            // 実日数 (API: days_limit)
	DaysTimes            string `csv:"11" json:"days_times"`            // 日数・回数
	UpdateDate           string `csv:"12" json:"update_date"`           // 変更年月日
	ExpiryDate           string `csv:"13" json:"expiry_date"`           // 廃止年月日
	Reserved             string `csv:"14" json:"reserved"`              // 予備
}

// DentalPracticeSupport 補助マスターテーブル (01補助マスターテーブル（歯科）.csv)
type DentalPracticeSupport struct {
	ChangeCategory               string `csv:"1" json:"change_category"`                  // 変更区分
	MedicalPracticeCode          string `csv:"2" json:"medical_practice_code"`            // 診療行為コード
	AdditionCode                 string `csv:"3" json:"addition_code"`                    // 加算コード
	AbbreviatedKanjiName         string `csv:"4" json:"abbreviated_kanji_name"`           // 診療行為省略名称
	InclusionUnit1               string `csv:"5" json:"inclusion_unit_1"`                 // 包括単位①
	InclusionGroup1              string `csv:"6" json:"inclusion_group_1"`                // グループ番号①
	InclusionUnit2               string `csv:"7" json:"inclusion_unit_2"`                 // 包括単位②
	InclusionGroup2              string `csv:"8" json:"inclusion_group_2"`                // グループ番号②
	InclusionUnit3               string `csv:"9" json:"inclusion_unit_3"`                 // 包括単位③
	InclusionGroup3              string `csv:"10" json:"inclusion_group_3"`               // グループ番号③
	ConflictPerDay               string `csv:"11" json:"conflict_per_day"`                // 背反関連識別（１日につき）
	ConflictSameMonth            string `csv:"12" json:"conflict_same_month"`             // 背反関連識別（同一月内）
	ConflictSimultaneous         string `csv:"13" json:"conflict_simultaneous"`           // 背反関連識別（同時）
	ConflictSameSiteSimultaneous string `csv:"14" json:"conflict_same_site_simultaneous"` // 背反関連識別（同一部位同時）
	ConflictPerWeek              string `csv:"15" json:"conflict_per_week"`               // 背反関連識別（１週間につき）
	ConflictReserved             string `csv:"16" json:"conflict_reserved"`               // 背反関連識別（予備）
	CalculationCountRelation     string `csv:"17" json:"calculation_count_relation"`      // 算定回数関連
	Reserved1                    string `csv:"18" json:"reserved_1"`                      // 予備
	Reserved2                    string `csv:"19" json:"reserved_2"`                      // 予備
	Reserved3                    string `csv:"20" json:"reserved_3"`                      // 予備
	Reserved4                    string `csv:"21" json:"reserved_4"`                      // 予備
	Reserved5                    string `csv:"22" json:"reserved_5"`                      // 予備
	Reserved6                    string `csv:"23" json:"reserved_6"`                      // 予備
	UpdateDate                   string `csv:"24" json:"update_date"`                     // 変更年月日
	ExpiryDate                   string `csv:"25" json:"expiry_date"`                     // 廃止年月日
}

// DentalPracticeInclusion 包括テーブル (02包括テーブル（歯科）.csv)
type DentalPracticeInclusion struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 変更区分
	GroupNumber          string `csv:"2" json:"group_number"`           // グループ番号
	IncludedPracticeCode string `csv:"3" json:"included_practice_code"` // 包括対象診療行為コード
	AdditionCode         string `csv:"4" json:"addition_code"`          // 加算コード
	AbbreviatedKanjiName string `csv:"5" json:"abbreviated_kanji_name"` // 診療行為省略名称
	SpecialCondition     string `csv:"6" json:"special_condition"`      // 特例条件
	NewDate              string `csv:"7" json:"new_date"`               // 新設年月日
	ExpiryDate           string `csv:"8" json:"expiry_date"`            // 廃止年月日
}

// DentalPracticeConflictDetail 背反テーブル (03-*背反テーブル*（歯科）.csv)
type DentalPracticeConflictDetail struct {
	ChangeCategory        string `csv:"1" json:"change_category"`          // 変更区分
	MedicalPracticeCode1  string `csv:"2" json:"medical_practice_code_1"`  // 診療行為コード1
	AdditionCode1         string `csv:"3" json:"addition_code_1"`          // 加算コード1
	AbbreviatedKanjiName1 string `csv:"4" json:"abbreviated_kanji_name_1"` // 診療行為省略名称1
	MedicalPracticeCode2  string `csv:"5" json:"medical_practice_code_2"`  // 診療行為コード2
	AdditionCode2         string `csv:"6" json:"addition_code_2"`          // 加算コード2
	AbbreviatedKanjiName2 string `csv:"7" json:"abbreviated_kanji_name_2"` // 診療行為省略名称2
	ConflictCategory      string `csv:"8" json:"conflict_category"`        // 背反区分
	SpecialCondition      string `csv:"9" json:"special_condition"`        // 特例条件
	Reserved              string `csv:"10" json:"reserved"`                // 予備
	NewDate               string `csv:"11" json:"new_date"`                // 新設年月日
	ExpiryDate            string `csv:"12" json:"expiry_date"`             // 廃止年月日
}

// DentalPracticeCalculationCountLimit 算定回数テーブル (04算定回数テーブル（歯科）.csv)
type DentalPracticeCalculationCountLimit struct {
	ChangeCategory         string `csv:"1" json:"change_category"`         // 変更区分
	MedicalPracticeCode    string `csv:"2" json:"medical_practice_code"`   // 診療行為コード
	AdditionCode           string `csv:"3" json:"addition_code"`           // 加算コード
	AbbreviatedKanjiName   string `csv:"4" json:"abbreviated_kanji_name"`  // 診療行為省略名称
	CalculationRequirement string `csv:"5" json:"calculation_requirement"` // 算定要件
	UnitCode               string `csv:"6" json:"unit_code"`               // 算定単位コード
	UnitName               string `csv:"7" json:"unit_name"`               // 算定単位名称
	CountLimit             int    `csv:"8" json:"count_limit"`             // 算定回数
	SpecialCondition       string `csv:"9" json:"special_condition"`       // 特例条件
	Reserved               string `csv:"10" json:"reserved"`               // 予備
	NewDate                string `csv:"11" json:"new_date"`               // 新設年月日
	ExpiryDate             string `csv:"12" json:"expiry_date"`            // 廃止年月日
}
