package model

// DentalPractice 歯科診療行為マスター・基本テーブル (h_*.csv)
// 全66項目を網羅 (R06rec3.pdf 205-206ページに準拠)
type DentalPractice struct {
	ChangeCategory                 string  `csv:"1" json:"change_category"`                   // 1: 変更区分
	MasterType                     string  `csv:"2" json:"master_type"`                       // 2: マスター種別
	DentalPracticeCode             string  `csv:"3" json:"dental_practice_code"`              // 3: 歯科診療行為コード
	SectionCode                    string  `csv:"4" json:"section_code"`                      // 4: 区分 (告示番号)
	SectionNumber                  string  `csv:"5" json:"section_number"`                    // 5: 区分番号 (告示番号)
	SectionSubNumber               string  `csv:"6" json:"section_sub_number"`                // 6: 枝番 (告示番号)
	SectionItemNumber              string  `csv:"7" json:"section_item_number"`               // 7: 項番 (告示番号)
	AdditionCode                   string  `csv:"8" json:"addition_code"`                     // 8: 加算コード
	BasicKanjiName                 string  `csv:"9" json:"basic_kanji_name"`                  // 9: 基本名称
	AbbreviatedKanjiName           string  `csv:"10" json:"abbreviated_kanji_name"`           // 10: 省略名称
	ScoreType                      string  `csv:"11" json:"score_type"`                       // 11: 点数等識別
	Score                          float64 `csv:"12" json:"score"`                            // 12: 点数等
	OldScoreType                   string  `csv:"13" json:"old_score_type"`                   // 13: 旧点数等識別
	OldScore                       float64 `csv:"14" json:"old_score"`                        // 14: 旧点数等
	InpatientOutpatientCategory    string  `csv:"15" json:"inpatient_outpatient_category"`    // 15: 入外適用区分
	ElderlyMedicalCategory         string  `csv:"16" json:"elderly_medical_category"`         // 16: 後期高齢者医療適用区分
	TimeAdditionCategory           string  `csv:"17" json:"time_addition_category"`           // 17: 時間加算区分
	HospitalClinicCategory         string  `csv:"18" json:"hospital_clinic_category"`         // 18: 病院・診療所区分
	NursingAddition                string  `csv:"19" json:"nursing_addition"`                 // 19: 看護加算
	Reserved20                     string  `csv:"20" json:"reserved_20"`                      // 20: 予備
	Reserved21                     string  `csv:"21" json:"reserved_21"`                      // 21: 予備
	RegionAddition                 string  `csv:"22" json:"region_addition"`                  // 22: 地域加算
	DiseaseRelatedCategory         string  `csv:"23" json:"disease_related_category"`         // 23: 傷病名関連区分
	MedicineRelatedCategory        string  `csv:"24" json:"medicine_related_category"`        // 24: 医薬品関連区分
	BedCountCategory               string  `csv:"25" json:"bed_count_category"`               // 25: 病床数区分
	NotificationCategory           string  `csv:"26" json:"notification_category"`            // 26: 届出
	FutureVisitCategory            string  `csv:"27" json:"future_visit_category"`            // 27: 未来院
	ShortStaySurgery               string  `csv:"28" json:"short_stay_surgery"`               // 28: 短期滞在手術
	SpecialNoteCategory            string  `csv:"29" json:"special_note_category"`            // 29: 特記事項
	TestJudgmentCategory           string  `csv:"30" json:"test_judgment_category"`           // 30: 検査等実施判断区分
	TestJudgmentGroupCategory      string  `csv:"31" json:"test_judgment_group_category"`     // 31: 検査等実施判断グループ区分
	ReductionApplicability         string  `csv:"32" json:"reduction_applicability"`          // 32: 逓減対象区分
	ComprehensiveReductionCategory string  `csv:"33" json:"comprehensive_reduction_category"` // 33: 包括逓減区分
	FacilityStandardCategory       string  `csv:"34" json:"facility_standard_category"`       // 34: 適合区分 (基準適合識別)
	FacilityStandardCode           string  `csv:"35" json:"facility_standard_code"`           // 35: 対象施設基準 (基準適合識別)
	FacilityStandard1              string  `csv:"36" json:"facility_standard_1"`              // 36: 施設基準1
	FacilityStandard2              string  `csv:"37" json:"facility_standard_2"`              // 37: 施設基準2
	FacilityStandard3              string  `csv:"38" json:"facility_standard_3"`              // 38: 施設基準3
	FacilityStandard4              string  `csv:"39" json:"facility_standard_4"`              // 39: 施設基準4
	FacilityStandard5              string  `csv:"40" json:"facility_standard_5"`              // 40: 施設基準5
	FacilityStandard6              string  `csv:"41" json:"facility_standard_6"`              // 41: 施設基準6
	FacilityStandard7              string  `csv:"42" json:"facility_standard_7"`              // 42: 施設基準7
	FacilityStandard8              string  `csv:"43" json:"facility_standard_8"`              // 43: 施設基準8
	FacilityStandard9              string  `csv:"44" json:"facility_standard_9"`              // 44: 施設基準9
	FacilityStandard10             string  `csv:"45" json:"facility_standard_10"`             // 45: 施設基準10
	GeneralAdditionGroup           string  `csv:"46" json:"general_addition_group"`           // 46: 通則加算グループ
	BasicAdditionGroup             string  `csv:"47" json:"basic_addition_group"`             // 47: 基本加算グループ
	NoteAdditionGroup              string  `csv:"48" json:"note_addition_group"`              // 48: 注加算グループ
	ProcedureMaterialGroup         string  `csv:"49" json:"procedure_material_group"`         // 49: 手技・材料加算グループ
	CalculationCountTableRelated   string  `csv:"50" json:"calculation_count_table_related"`  // 50: 算定回数限度テーブル関連識別
	StepTableRelated               string  `csv:"51" json:"step_table_related"`               // 51: きざみテーブル関連識別
	AgeLimitTableRelated           string  `csv:"52" json:"age_limit_table_related"`          // 52: 年齢制限テーブル関連識別
	ConflictTableRelated           string  `csv:"53" json:"conflict_table_related"`           // 53: 併算定背反テーブル関連識別
	ActualDaysTableRelated         string  `csv:"54" json:"actual_days_table_related"`        // 54: 実日数テーブル関連識別
	Reserved55                     string  `csv:"55" json:"reserved_55"`                      // 55: 予備
	Reserved56                     string  `csv:"56" json:"reserved_56"`                      // 56: 予備
	UpdateDate                     string  `csv:"57" json:"update_date"`                      // 57: 変更年月日
	ExpiryDate                     string  `csv:"58" json:"expiry_date"`                      // 58: 廃止年月日
	LongTimeAnesthesia             string  `csv:"59" json:"long_time_anesthesia"`             // 59: 長時間麻酔管理加算
	MalignantTumorPathology        string  `csv:"60" json:"malignant_tumor_pathology"`        // 60: 悪性腫瘍病理組織標本加算
	NursingStaffTreatmentImprove   string  `csv:"61" json:"nursing_staff_treatment_improve"`  // 61: 看護職員処遇改善評価料等
	DentalBaseUpEvaluation1        string  `csv:"62" json:"dental_base_up_evaluation_1"`      // 62: 歯科外来・在宅ベースアップ評価料（1）
	DentalBaseUpEvaluation2        string  `csv:"63" json:"dental_base_up_evaluation_2"`      // 63: 歯科外来・在宅ベースアップ評価料（2）
	Reserved64                     string  `csv:"64" json:"reserved_64"`                      // 64: 予備4
	Reserved65                     string  `csv:"65" json:"reserved_65"`                      // 65: 予備5
	PublicationOrder               string  `csv:"66" json:"publication_order"`                // 66: 公表順序番号
}

// DentalPracticeSupport 歯科診療行為マスター補助マスターテーブル (01補助マスターテーブル（歯科）.csv)
type DentalPracticeSupport struct {
	ChangeCategory           string `csv:"1" json:"change_category"`
	DentalPracticeCode       string `csv:"2" json:"dental_practice_code"`
	AdditionCode             string `csv:"3" json:"addition_code"`
	AbbreviatedKanjiName     string `csv:"4" json:"abbreviated_kanji_name"`
	ComprehensiveUnit1       string `csv:"5" json:"comprehensive_unit_1"`
	ComprehensiveGroup1      string `csv:"6" json:"comprehensive_group_1"`
	ComprehensiveUnit2       string `csv:"7" json:"comprehensive_unit_2"`
	ComprehensiveGroup2      string `csv:"8" json:"comprehensive_group_2"`
	ComprehensiveUnit3       string `csv:"9" json:"comprehensive_unit_3"`
	ComprehensiveGroup3      string `csv:"10" json:"comprehensive_group_3"`
	ConflictDaily            string `csv:"11" json:"conflict_daily"`
	ConflictMonthly          string `csv:"12" json:"conflict_monthly"`
	ConflictSimultaneous     string `csv:"13" json:"conflict_simultaneous"`
	ConflictPartSimultaneous string `csv:"14" json:"conflict_part_simultaneous"`
	ConflictWeekly           string `csv:"15" json:"conflict_weekly"`
	Reserved16               string `csv:"16" json:"reserved_16"`
	CalculationCountRelated  string `csv:"17" json:"calculation_count_related"`
	Reserved18               string `csv:"18" json:"reserved_18"`
	Reserved19               string `csv:"19" json:"reserved_19"`
	Reserved20               string `csv:"20" json:"reserved_20"`
	Reserved21               string `csv:"21" json:"reserved_21"`
	Reserved22               string `csv:"22" json:"reserved_22"`
	Reserved23               string `csv:"23" json:"reserved_23"`
	UpdateDate               string `csv:"24" json:"update_date"`
	ExpiryDate               string `csv:"25" json:"expiry_date"`
}

// DentalPracticeInclusion 歯科包括・被包括テーブル (02包括テーブル（歯科）.csv)
type DentalPracticeInclusion struct {
	ChangeCategory       string `csv:"1" json:"change_category"`
	GroupNumber          string `csv:"2" json:"group_number"`
	DentalPracticeCode   string `csv:"3" json:"dental_practice_code"`
	AdditionCode         string `csv:"4" json:"addition_code"`
	AbbreviatedKanjiName string `csv:"5" json:"abbreviated_kanji_name"`
	SpecialCondition     string `csv:"6" json:"special_condition"`
	UpdateDate           string `csv:"7" json:"update_date"`
	ExpiryDate           string `csv:"8" json:"expiry_date"`
}

// DentalPracticeConflict 歯科背反関連テーブル (03-*背反テーブル（歯科）.csv)
type DentalPracticeConflict struct {
	ChangeCategory        string `csv:"1" json:"change_category"`
	DentalPracticeCode1   string `csv:"2" json:"dental_practice_code_1"`
	AdditionCode1         string `csv:"3" json:"addition_code_1"`
	AbbreviatedKanjiName1 string `csv:"4" json:"abbreviated_kanji_name_1"`
	DentalPracticeCode2   string `csv:"5" json:"dental_practice_code_2"`
	AdditionCode2         string `csv:"6" json:"addition_code_2"`
	AbbreviatedKanjiName2 string `csv:"7" json:"abbreviated_kanji_name_2"`
	ConflictCategory      string `csv:"8" json:"conflict_category"`
	SpecialCondition      string `csv:"9" json:"special_condition"`
	Reserved10            string `csv:"10" json:"reserved_10"`
	UpdateDate            string `csv:"11" json:"update_date"`
	ExpiryDate            string `csv:"12" json:"expiry_date"`
}

// DentalPracticeCalculationCount 歯科算定回数テーブル (04算定回数テーブル（歯科）.csv)
type DentalPracticeCalculationCount struct {
	ChangeCategory       string `csv:"1" json:"change_category"`
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`
	AdditionCode         string `csv:"3" json:"addition_code"`
	AbbreviatedKanjiName string `csv:"4" json:"abbreviated_kanji_name"`
	RequirementCode      string `csv:"5" json:"requirement_code"`
	UnitCode             string `csv:"6" json:"unit_code"`
	UnitName             string `csv:"7" json:"unit_name"`
	MaxTimes             int    `csv:"8" json:"max_times"`
	SpecialCondition     string `csv:"9" json:"special_condition"`
	Reserved10           string `csv:"10" json:"reserved_10"`
	UpdateDate           string `csv:"11" json:"update_date"`
	ExpiryDate           string `csv:"12" json:"expiry_date"`
}

// DentalPracticeAdditionRelation 基本・通則加算/基本加算/注加算対応テーブル (イ, ウ, エ)
// 全9項目 (R06rec3.pdf 207-209ページ)
type DentalPracticeAdditionRelation struct {
	ChangeCategory     string `csv:"1" json:"change_category"`
	GroupNumber        string `csv:"2" json:"group_number"`
	AdditionCode       string `csv:"3" json:"addition_code"`
	DentalPracticeCode string `csv:"4" json:"dental_practice_code"`
	BasicName          string `csv:"5" json:"basic_name"`
	AdditionIdentifier string `csv:"6" json:"addition_identifier"`
	UpdateDate         string `csv:"7" json:"update_date"`
	ExpiryDate         string `csv:"8" json:"expiry_date"`
	Reserved9          string `csv:"9" json:"reserved_9"`
}

// DentalPracticeProcedureMaterial 手技・材料加算対応テーブル (オ)
// 全9項目 (R06rec3.pdf 210ページ)
type DentalPracticeProcedureMaterial struct {
	ChangeCategory     string `csv:"1" json:"change_category"`
	GroupNumber        string `csv:"2" json:"group_number"`
	AdditionCode       string `csv:"3" json:"addition_code"`
	DentalPracticeCode string `csv:"4" json:"dental_practice_code"`
	BasicName          string `csv:"5" json:"basic_name"`
	AdditionIdentifier string `csv:"6" json:"addition_identifier"`
	UpdateDate         string `csv:"7" json:"update_date"`
	ExpiryDate         string `csv:"8" json:"expiry_date"`
	Reserved9          string `csv:"9" json:"reserved_9"`
}

// DentalPracticeCalculationCountLimit 算定回数限度テーブル (カ)
// 全15項目 (R06rec3.pdf 211ページ)
type DentalPracticeCalculationCountLimit struct {
	ChangeCategory        string `csv:"1" json:"change_category"`
	DentalPracticeCode    string `csv:"2" json:"dental_practice_code"`
	SectionCode           string `csv:"3" json:"section_code"`
	SectionNumber         string `csv:"4" json:"section_number"`
	SectionSubNumber      string `csv:"5" json:"section_sub_number"`
	SectionItemNumber     string `csv:"6" json:"section_item_number"`
	AdditionCode          string `csv:"7" json:"addition_code"`
	BasicName             string `csv:"8" json:"basic_name"`
	AbbreviatedKanjiName  string `csv:"9" json:"abbreviated_kanji_name"`
	CalculationUnit       string `csv:"10" json:"calculation_unit"`
	CalculationCountLimit int    `csv:"11" json:"calculation_count_limit"`
	MaxTimesErrorCategory string `csv:"12" json:"max_times_error_category"`
	UpdateDate            string `csv:"13" json:"update_date"`
	ExpiryDate            string `csv:"14" json:"expiry_date"`
	Reserved15            string `csv:"15" json:"reserved_15"`
}

// DentalPracticeStep きざみテーブル (キ)
// 全20項目 (R06rec3.pdf 212ページ)
type DentalPracticeStep struct {
	ChangeCategory       string  `csv:"1" json:"change_category"`
	DentalPracticeCode   string  `csv:"2" json:"dental_practice_code"`
	SectionCode          string  `csv:"3" json:"section_code"`
	SectionNumber        string  `csv:"4" json:"section_number"`
	SectionSubNumber     string  `csv:"5" json:"section_sub_number"`
	SectionItemNumber    string  `csv:"6" json:"section_item_number"`
	AdditionCode         string  `csv:"7" json:"addition_code"`
	BasicName            string  `csv:"8" json:"basic_name"`
	AbbreviatedKanjiName string  `csv:"9" json:"abbreviated_kanji_name"`
	ScoreType            string  `csv:"10" json:"score_type"`
	Score                float64 `csv:"11" json:"score"`
	StepUnit             string  `csv:"12" json:"step_unit"`
	StepLowerLimit       float64 `csv:"13" json:"step_lower_limit"`
	StepUpperLimit       float64 `csv:"14" json:"step_upper_limit"`
	StepValue            float64 `csv:"15" json:"step_value"`
	StepScore            float64 `csv:"16" json:"step_score"`
	StepErrorCategory    string  `csv:"17" json:"step_error_category"`
	UpdateDate           string  `csv:"18" json:"update_date"`
	ExpiryDate           string  `csv:"19" json:"expiry_date"`
	Reserved20           string  `csv:"20" json:"reserved_20"`
}

// DentalPracticeAgeConstraint 年齢制限テーブル (ク)
// 全14項目 (R06rec3.pdf 213ページ)
type DentalPracticeAgeConstraint struct {
	ChangeCategory       string `csv:"1" json:"change_category"`
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`
	SectionCode          string `csv:"3" json:"section_code"`
	SectionNumber        string `csv:"4" json:"section_number"`
	SectionSubNumber     string `csv:"5" json:"section_sub_number"`
	SectionItemNumber    string `csv:"6" json:"section_item_number"`
	AdditionCode         string `csv:"7" json:"addition_code"`
	BasicName            string `csv:"8" json:"basic_name"`
	AbbreviatedKanjiName string `csv:"9" json:"abbreviated_kanji_name"`
	LowerAge             string `csv:"10" json:"lower_age"`
	UpperAge             string `csv:"11" json:"upper_age"`
	UpdateDate           string `csv:"12" json:"update_date"`
	ExpiryDate           string `csv:"13" json:"expiry_date"`
	Reserved14           string `csv:"14" json:"reserved_14"`
}

// DentalPracticeConflictDetail 併算定背反テーブル (ケ)
// 全104項目 (R06rec3.pdf 214ページ)
// 10-99項は 背反1-10 (算定可否, コード, 告示番号x4, 加算コード, 基本名称, 省略名称) = 9項目 x 10回 = 90項目
type DentalPracticeConflictDetail struct {
	ChangeCategory       string                         `csv:"1" json:"change_category"`
	DentalPracticeCode   string                         `csv:"2" json:"dental_practice_code"`
	SectionCode          string                         `csv:"3" json:"section_code"`
	SectionNumber        string                         `csv:"4" json:"section_number"`
	SectionSubNumber     string                         `csv:"5" json:"section_sub_number"`
	SectionItemNumber    string                         `csv:"6" json:"section_item_number"`
	AdditionCode         string                         `csv:"7" json:"addition_code"`
	BasicName            string                         `csv:"8" json:"basic_name"`
	AbbreviatedKanjiName string                         `csv:"9" json:"abbreviated_kanji_name"`
	Conflicts            [10]DentalPracticeConflictItem `json:"conflicts"`
	UpdateDate           string                         `csv:"100" json:"update_date"`
	ExpiryDate           string                         `csv:"101" json:"expiry_date"`
	Reserved102          string                         `csv:"102" json:"reserved_102"`
	Reserved103          string                         `csv:"103" json:"reserved_103"`
	Reserved104          string                         `csv:"104" json:"reserved_104"`
}

type DentalPracticeConflictItem struct {
	ScoreCondition       string `json:"score_condition"` // 算定可否
	DentalPracticeCode   string `json:"dental_practice_code"`
	SectionCode          string `json:"section_code"`
	SectionNumber        string `json:"section_number"`
	SectionSubNumber     string `json:"section_sub_number"`
	SectionItemNumber    string `json:"section_item_number"`
	AdditionCode         string `json:"addition_code"`
	BasicName            string `json:"basic_name"`
	AbbreviatedKanjiName string `json:"abbreviated_kanji_name"`
}

// DentalPracticeActualDays 実日数関連テーブル (コ)
// 全14項目 (R06rec3.pdf 215ページ)
type DentalPracticeActualDays struct {
	ChangeCategory       string `csv:"1" json:"change_category"`
	DentalPracticeCode   string `csv:"2" json:"dental_practice_code"`
	SectionCode          string `csv:"3" json:"section_code"`
	SectionNumber        string `csv:"4" json:"section_number"`
	SectionSubNumber     string `csv:"5" json:"section_sub_number"`
	SectionItemNumber    string `csv:"6" json:"section_item_number"`
	AdditionCode         string `csv:"7" json:"addition_code"`
	BasicName            string `csv:"8" json:"basic_name"`
	AbbreviatedKanjiName string `csv:"9" json:"abbreviated_kanji_name"`
	ActualDays           string `csv:"10" json:"actual_days"`
	DaysTimes            string `csv:"11" json:"days_times"`
	UpdateDate           string `csv:"12" json:"update_date"`
	ExpiryDate           string `csv:"13" json:"expiry_date"`
	Reserved14           string `csv:"14" json:"reserved_14"`
}
