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
	ChangeCategory       string `csv:"1" json:"change_category"`             // 変更区分
	DentalPracticeCode   string `csv:"2" json:"source_code"`                 // 歯科診療行為コード
	SectionChar          string `csv:"3" json:"source_type"`                 // 区分
	SectionNumber        string `csv:"4" json:"source_section"`              // 区分番号
	SectionBranchNumber  string `csv:"5" json:"section_branch_number"`       // 枝番
	SectionItemNumber    string `csv:"6" json:"section_item_number"`         // 項番
	AdditionCode         string `csv:"7" json:"addition_code"`               // 加算コード
	BasicKanjiName       string `csv:"8" json:"source_name1"`                // 基本名称
	AbbreviatedKanjiName string `csv:"9" json:"source_name2"`                // 省略名称
	Conflict1_Flag       string `csv:"10" json:"conflict1_flag"`             // 背反1:算定可否
	Conflict1_Code       string `csv:"11" json:"conflict1_code"`             // 背反1:歯科診療行為コード
	Conflict1_Section    string `csv:"12" json:"conflict1_section_char"`     // 背反1:区分
	Conflict1_SecNum     string `csv:"13" json:"conflict1_section_number"`   // 背反1:区分番号
	Conflict1_Branch     string `csv:"14" json:"conflict1_branch_number"`    // 背反1:枝番
	Conflict1_Item       string `csv:"15" json:"conflict1_item_number"`      // 背反1:項番
	Conflict1_AddCode    string `csv:"16" json:"conflict1_addition_code"`    // 背反1:加算コード
	Conflict1_Name       string `csv:"17" json:"conflict1_name"`             // 背反1:基本名称
	Conflict1_Abbr       string `csv:"18" json:"conflict1_abbreviated_name"` // 背反1:省略名称
	Conflict2_Flag       string `csv:"19" json:"conflict2_flag"`             // 背反2:算定可否
	Conflict2_Code       string `csv:"20" json:"conflict2_code"`             // 背反2:歯科診療行為コード
	Conflict2_Section    string `csv:"21" json:"conflict2_section_char"`     // 背反2:区分
	Conflict2_SecNum     string `csv:"22" json:"conflict2_section_number"`   // 背反2:区分番号
	Conflict2_Branch     string `csv:"23" json:"conflict2_branch_number"`    // 背反2:枝番
	Conflict2_Item       string `csv:"24" json:"conflict2_item_number"`      // 背反2:項番
	Conflict2_AddCode    string `csv:"25" json:"conflict2_addition_code"`    // 背反2:加算コード
	Conflict2_Name       string `csv:"26" json:"conflict2_name"`             // 背反2:基本名称
	Conflict2_Abbr       string `csv:"27" json:"conflict2_abbreviated_name"` // 背反2:省略名称
	Conflict3_Flag       string `csv:"28" json:"conflict3_flag"`
	Conflict3_Code       string `csv:"29" json:"conflict3_code"`
	Conflict3_Section    string `csv:"30" json:"conflict3_section_char"`
	Conflict3_SecNum     string `csv:"31" json:"conflict3_section_number"`
	Conflict3_Branch     string `csv:"32" json:"conflict3_branch_number"`
	Conflict3_Item       string `csv:"33" json:"conflict3_item_number"`
	Conflict3_AddCode    string `csv:"34" json:"conflict3_addition_code"`
	Conflict3_Name       string `csv:"35" json:"conflict3_name"`
	Conflict3_Abbr       string `csv:"36" json:"conflict3_abbreviated_name"`
	Conflict4_Flag       string `csv:"37" json:"conflict4_flag"`
	Conflict4_Code       string `csv:"38" json:"conflict4_code"`
	Conflict4_Section    string `csv:"39" json:"conflict4_section_char"`
	Conflict4_SecNum     string `csv:"40" json:"conflict4_section_number"`
	Conflict4_Branch     string `csv:"41" json:"conflict4_branch_number"`
	Conflict4_Item       string `csv:"42" json:"conflict4_item_number"`
	Conflict4_AddCode    string `csv:"43" json:"conflict4_addition_code"`
	Conflict4_Name       string `csv:"44" json:"conflict4_name"`
	Conflict4_Abbr       string `csv:"45" json:"conflict4_abbreviated_name"`
	Conflict5_Flag       string `csv:"46" json:"conflict5_flag"`
	Conflict5_Code       string `csv:"47" json:"conflict5_code"`
	Conflict5_Section    string `csv:"48" json:"conflict5_section_char"`
	Conflict5_SecNum     string `csv:"49" json:"conflict5_section_number"`
	Conflict5_Branch     string `csv:"50" json:"conflict5_branch_number"`
	Conflict5_Item       string `csv:"51" json:"conflict5_item_number"`
	Conflict5_AddCode    string `csv:"52" json:"conflict5_addition_code"`
	Conflict5_Name       string `csv:"53" json:"conflict5_name"`
	Conflict5_Abbr       string `csv:"54" json:"conflict5_abbreviated_name"`
	Conflict6_Flag       string `csv:"55" json:"conflict6_flag"`
	Conflict6_Code       string `csv:"56" json:"conflict6_code"`
	Conflict6_Section    string `csv:"57" json:"conflict6_section_char"`
	Conflict6_SecNum     string `csv:"58" json:"conflict6_section_number"`
	Conflict6_Branch     string `csv:"59" json:"conflict6_branch_number"`
	Conflict6_Item       string `csv:"60" json:"conflict6_item_number"`
	Conflict6_AddCode    string `csv:"61" json:"conflict6_addition_code"`
	Conflict6_Name       string `csv:"62" json:"conflict6_name"`
	Conflict6_Abbr       string `csv:"63" json:"conflict6_abbreviated_name"`
	Conflict7_Flag       string `csv:"64" json:"conflict7_flag"`
	Conflict7_Code       string `csv:"65" json:"conflict7_code"`
	Conflict7_Section    string `csv:"66" json:"conflict7_section_char"`
	Conflict7_SecNum     string `csv:"67" json:"conflict7_section_number"`
	Conflict7_Branch     string `csv:"68" json:"conflict7_branch_number"`
	Conflict7_Item       string `csv:"69" json:"conflict7_item_number"`
	Conflict7_AddCode    string `csv:"70" json:"conflict7_addition_code"`
	Conflict7_Name       string `csv:"71" json:"conflict7_name"`
	Conflict7_Abbr       string `csv:"72" json:"conflict7_abbreviated_name"`
	Conflict8_Flag       string `csv:"73" json:"conflict8_flag"`
	Conflict8_Code       string `csv:"74" json:"conflict8_code"`
	Conflict8_Section    string `csv:"75" json:"conflict8_section_char"`
	Conflict8_SecNum     string `csv:"76" json:"conflict8_section_number"`
	Conflict8_Branch     string `csv:"77" json:"conflict8_branch_number"`
	Conflict8_Item       string `csv:"78" json:"conflict8_item_number"`
	Conflict8_AddCode    string `csv:"79" json:"conflict8_addition_code"`
	Conflict8_Name       string `csv:"80" json:"conflict8_name"`
	Conflict8_Abbr       string `csv:"81" json:"conflict8_abbreviated_name"`
	Conflict9_Flag       string `csv:"82" json:"conflict9_flag"`
	Conflict9_Code       string `csv:"83" json:"conflict9_code"`
	Conflict9_Section    string `csv:"84" json:"conflict9_section_char"`
	Conflict9_SecNum     string `csv:"85" json:"conflict9_section_number"`
	Conflict9_Branch     string `csv:"86" json:"conflict9_branch_number"`
	Conflict9_Item       string `csv:"87" json:"conflict9_item_number"`
	Conflict9_AddCode    string `csv:"88" json:"conflict9_addition_code"`
	Conflict9_Name       string `csv:"89" json:"conflict9_name"`
	Conflict9_Abbr       string `csv:"90" json:"conflict9_abbreviated_name"`
	Conflict10_Flag      string `csv:"91" json:"conflict10_flag"`
	Conflict10_Code      string `csv:"92" json:"conflict10_code"`
	Conflict10_Section   string `csv:"93" json:"conflict10_section_char"`
	Conflict10_SecNum    string `csv:"94" json:"conflict10_section_number"`
	Conflict10_Branch    string `csv:"95" json:"conflict10_branch_number"`
	Conflict10_Item      string `csv:"96" json:"conflict10_item_number"`
	Conflict10_AddCode   string `csv:"97" json:"conflict10_addition_code"`
	Conflict10_Name      string `csv:"98" json:"conflict10_name"`
	Conflict10_Abbr      string `csv:"99" json:"conflict10_abbreviated_name"`
	UpdateDate           string `csv:"100" json:"update_date"` // 変更年月日
	ExpiryDate           string `csv:"101" json:"expiry_date"` // 廃止年月日
	Reserved1            string `csv:"102" json:"reserved_1"`  // 予備
	Reserved2            string `csv:"103" json:"reserved_2"`  // 予備
	Reserved3            string `csv:"104" json:"reserved_3"`  // 予備
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
	ChangeCategory      string `csv:"1" json:"change_category"`       // 変更区分
	MedicalPracticeCode string `csv:"2" json:"medical_practice_code"` // 診療行為コード
	SupportCode         string `csv:"3" json:"support_code"`          // 補助コード
	Name                string `csv:"4" json:"name"`                  // 名称
	Flag1               string `csv:"5" json:"flag1"`                 // フラグ1
	Flag2               string `csv:"6" json:"flag2"`                 // フラグ2
	Flag3               string `csv:"7" json:"flag3"`                 // フラグ3
	Flag4               string `csv:"8" json:"flag4"`                 // フラグ4
	Flag5               string `csv:"9" json:"flag5"`                 // フラグ5
	Flag6               string `csv:"10" json:"flag6"`                // フラグ6
	Flag7               string `csv:"11" json:"flag7"`                // フラグ7
	Flag8               string `csv:"12" json:"flag8"`                // フラグ8
	Flag9               string `csv:"13" json:"flag9"`                // フラグ9
	Flag10              string `csv:"14" json:"flag10"`               // フラグ10
	Flag11              string `csv:"15" json:"flag11"`               // フラグ11
	Flag12              string `csv:"16" json:"flag12"`               // フラグ12
	Flag13              string `csv:"17" json:"flag13"`               // フラグ13
	Flag14              string `csv:"18" json:"flag14"`               // フラグ14
	Flag15              string `csv:"19" json:"flag15"`               // フラグ15
	Flag16              string `csv:"20" json:"flag16"`               // フラグ16
	Flag17              string `csv:"21" json:"flag17"`               // フラグ17
	Flag18              string `csv:"22" json:"flag18"`               // フラグ18
	Flag19              string `csv:"23" json:"flag19"`               // フラグ19
	UpdateDate          string `csv:"24" json:"update_date"`          // 変更年月日
	ExpiryDate          string `csv:"25" json:"expiry_date"`          // 廃止年月日
}

// DentalPracticeInclusion 包括テーブル (02包括テーブル（歯科）.csv)
type DentalPracticeInclusion struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 変更区分
	GroupNumber          string `csv:"2" json:"group_number"`           // グループ番号
	IncludedPracticeCode string `csv:"3" json:"included_practice_code"` // 包括対象診療行為コード
	AdditionCode         string `csv:"4" json:"addition_code"`          // 加算コード
	Name                 string `csv:"5" json:"name"`                   // 名称
	SpecialCondition     string `csv:"6" json:"special_condition"`      // 特例条件
	UpdateDate           string `csv:"7" json:"update_date"`            // 変更年月日
	ExpiryDate           string `csv:"8" json:"expiry_date"`            // 廃止年月日
}

// DentalPracticeConflictDetail 背反テーブル (03-*背反テーブル*（歯科）.csv)
type DentalPracticeConflictDetail struct {
	ChangeCategory       string `csv:"1" json:"change_category"`         // 変更区分
	MedicalPracticeCode1 string `csv:"2" json:"medical_practice_code_1"` // 診療行為コード1
	AdditionCode1        string `csv:"3" json:"addition_code_1"`         // 加算コード1
	Name1                string `csv:"4" json:"name_1"`                  // 名称1
	MedicalPracticeCode2 string `csv:"5" json:"medical_practice_code_2"` // 診療行為コード2
	AdditionCode2        string `csv:"6" json:"addition_code_2"`         // 加算コード2
	Name2                string `csv:"7" json:"name_2"`                  // 名称2
	ConflictCategory     string `csv:"8" json:"conflict_category"`       // 背反区分
	SpecialCondition     string `csv:"9" json:"special_condition"`       // 特例条件
	Reserved             string `csv:"10" json:"reserved"`               // 予備
	UpdateDate           string `csv:"11" json:"update_date"`            // 変更年月日
	ExpiryDate           string `csv:"12" json:"expiry_date"`            // 廃止年月日
}

// DentalPracticeCalculationCountLimit 算定回数テーブル (04算定回数テーブル（歯科）.csv)
type DentalPracticeCalculationCountLimit struct {
	ChangeCategory      string `csv:"1" json:"change_category"`       // 変更区分
	MedicalPracticeCode string `csv:"2" json:"medical_practice_code"` // 診療行為コード
	AdditionCode        string `csv:"3" json:"addition_code"`         // 加算コード
	Name                string `csv:"4" json:"name"`                  // 名称
	UnitCode            string `csv:"5" json:"unit_code"`             // 算定単位コード
	UnitName            string `csv:"6" json:"unit_name"`             // 算定単位名称
	CountLimit          int    `csv:"7" json:"count_limit"`           // 算定回数限度
	SpecialCondition    string `csv:"8" json:"special_condition"`     // 特例条件
	Reserved1           string `csv:"9" json:"reserved_1"`            // 予備
	Reserved2           string `csv:"10" json:"reserved_2"`           // 予備
	UpdateDate          string `csv:"11" json:"update_date"`          // 変更年月日
	ExpiryDate          string `csv:"12" json:"expiry_date"`          // 廃止年月日
}
