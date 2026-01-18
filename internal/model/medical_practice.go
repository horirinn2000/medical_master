package model

// MedicalPractice 医科診療行為マスター (s_*.csv)
// 全150項目を網羅 (R06rec3.pdf 201-204ページに準拠)
type MedicalPractice struct {
	ChangeCategory                 string  `csv:"1" json:"change_category"`                    // 1: 変更区分
	MasterType                     string  `csv:"2" json:"master_type"`                        // 2: マスター種別
	MedicalPracticeCode            string  `csv:"3" json:"medical_practice_code"`              // 3: 診療行為コード
	AbbreviatedKanjiNameLength     int     `csv:"4" json:"abbreviated_kanji_name_length"`      // 4: 省略漢字有効桁数
	AbbreviatedKanjiName           string  `csv:"5" json:"abbreviated_kanji_name"`             // 5: 省略漢字名称
	AbbreviatedKanaNameLength      int     `csv:"6" json:"abbreviated_kana_name_length"`       // 6: 省略カナ有効桁数
	AbbreviatedKanaName            string  `csv:"7" json:"abbreviated_kana_name"`              // 7: 省略カナ名称
	DataStandardCode               string  `csv:"8" json:"data_standard_code"`                 // 8: データ規格コード
	KanjiUnitNameLength            int     `csv:"9" json:"kanji_unit_name_length"`             // 9: 漢字有効桁数 (データ規格名)
	KanjiUnitName                  string  `csv:"10" json:"kanji_unit_name"`                   // 10: 漢字名称 (データ規格名)
	ScoreType                      string  `csv:"11" json:"score_type"`                        // 11: 点数識別
	Score                          float64 `csv:"12" json:"score"`                             // 12: 新又は現点数
	InpatientOutpatientCategory    string  `csv:"13" json:"inpatient_outpatient_category"`     // 13: 入外適用区分
	ElderlyMedicalCategory         string  `csv:"14" json:"elderly_medical_category"`          // 14: 後期高齢者医療適用区分
	AggregationCategoryOutpatient  string  `csv:"15" json:"aggregation_category_outpatient"`   // 15: 点数欄集計先識別（入院外）
	ComprehensiveTargetTest        string  `csv:"16" json:"comprehensive_target_test"`         // 16: 包括対象検査
	Reserved17                     string  `csv:"17" json:"reserved_17"`                       // 17: 予備
	DPCApplicabilityCategory       string  `csv:"18" json:"dpc_applicability_category"`        // 18: DPC適用区分
	HospitalClinicCategory         string  `csv:"19" json:"hospital_clinic_category"`          // 19: 病院・診療所区分
	SurgerySupportCategory         string  `csv:"20" json:"surgery_support_category"`          // 20: 画像等手術支援加算区分
	MedicalObservationCategory     string  `csv:"21" json:"medical_observation_category"`      // 21: 医療観察法対象区分
	NursingAddition                string  `csv:"22" json:"nursing_addition"`                  // 22: 看護加算
	AnesthesiaCategory             string  `csv:"23" json:"anesthesia_category"`               // 23: 麻酔識別区分
	Reserved24                     string  `csv:"24" json:"reserved_24"`                       // 24: 予備
	DiseaseRelatedCategory         string  `csv:"25" json:"disease_related_category"`          // 25: 傷病名関連区分
	Reserved26                     string  `csv:"26" json:"reserved_26"`                       // 26: 予備
	ActualDaysCategory             string  `csv:"27" json:"actual_days_category"`              // 27: 実日数
	DaysTimesCategory              string  `csv:"28" json:"days_times_category"`               // 28: 日数・回数
	MedicineRelatedCategory        string  `csv:"29" json:"medicine_related_category"`         // 29: 医薬品関連区分
	StepCalculationCategory        string  `csv:"30" json:"step_calculation_category"`         // 30: きざみ値計算識別
	StepLowerLimit                 float64 `csv:"31" json:"step_lower_limit"`                  // 31: 下限値
	StepUpperLimit                 float64 `csv:"32" json:"step_upper_limit"`                  // 32: 上限値
	StepValue                      float64 `csv:"33" json:"step_value"`                        // 33: きざみ値
	StepScore                      float64 `csv:"34" json:"step_score"`                        // 34: きざみ点数
	StepErrorCategory              string  `csv:"35" json:"step_error_category"`               // 35: 上下限エラー処理
	MaxTimes                       int     `csv:"36" json:"max_times"`                         // 36: 上限回数
	MaxTimesErrorCategory          string  `csv:"37" json:"max_times_error_category"`          // 37: 上限回数エラー処理
	NoteAdditionCode               string  `csv:"38" json:"note_addition_code"`                // 38: 注加算コード
	NoteAdditionSequence           string  `csv:"39" json:"note_addition_sequence"`            // 39: 注加算通番
	GeneralAgeCategory             string  `csv:"40" json:"general_age_category"`              // 40: 通則年齢
	LowerAge                       string  `csv:"41" json:"lower_age"`                         // 41: 下限年齢
	UpperAge                       string  `csv:"42" json:"upper_age"`                         // 42: 上限年齢
	TimeAdditionCategory           string  `csv:"43" json:"time_addition_category"`            // 43: 時間加算区分
	FacilityStandardCategory       string  `csv:"44" json:"facility_standard_category"`        // 44: 適合区分 (基準適合識別)
	FacilityStandardCode           string  `csv:"45" json:"facility_standard_code"`            // 45: 対象施設基準 (基準適合識別)
	TreatmentInfantAddition        string  `csv:"46" json:"treatment_infant_addition"`         // 46: 処置乳幼児加算区分
	VeryLowBirthWeightAddition     string  `csv:"47" json:"very_low_birth_weight_addition"`    // 47: 極低出生体重児加算区分
	InpatientReductionCategory     string  `csv:"48" json:"inpatient_reduction_category"`      // 48: 入院基本料等減算対象識別
	DonorCollectionCategory        string  `csv:"49" json:"donor_collection_category"`         // 49: ドナー分集計区分
	TestJudgmentCategory           string  `csv:"50" json:"test_judgment_category"`            // 50: 検査等実施判断区分
	TestJudgmentGroupCategory      string  `csv:"51" json:"test_judgment_group_category"`      // 51: 検査等実施判断グループ区分
	ReductionApplicability         string  `csv:"52" json:"reduction_applicability"`           // 52: 逓減対象区分
	SpinalEvokedPotentialCategory  string  `csv:"53" json:"spinal_evoked_potential_category"`  // 53: 脊髄誘発電位測定等加算区分
	NeckDissectionCategory         string  `csv:"54" json:"neck_dissection_category"`          // 54: 頸部郭清術併施加算区分
	AutoSutureCategory             string  `csv:"55" json:"auto_suture_category"`              // 55: 自動縫合器加算区分
	OutpatientManagementAddition   string  `csv:"56" json:"outpatient_management_addition"`    // 56: 外来管理加算区分
	OldScoreType                   string  `csv:"57" json:"old_score_type"`                    // 57: 点数識別 (旧)
	OldScore                       float64 `csv:"58" json:"old_score"`                         // 58: 旧点数
	KanjiNameUpdateFlag            string  `csv:"59" json:"kanji_name_update_flag"`            // 59: 漢字名称変更区分
	KanaNameUpdateFlag             string  `csv:"60" json:"kana_name_update_flag"`             // 60: カナ名称変更区分
	TestCommentCategory            string  `csv:"61" json:"test_comment_category"`             // 61: 検体検査コメント
	GeneralAdditionApplicability   string  `csv:"62" json:"general_addition_applicability"`    // 62: 通則加算所定点数対象区分
	ComprehensiveReductionCategory string  `csv:"63" json:"comprehensive_reduction_category"`  // 63: 包括逓減区分
	EndoscopicSurgeryAddition      string  `csv:"64" json:"endoscopic_surgery_addition"`       // 64: 超音波内視鏡加算区分
	Reserved65                     string  `csv:"65" json:"reserved_65"`                       // 65: 予備
	AggregationCategoryInpatient   string  `csv:"66" json:"aggregation_category_inpatient"`    // 66: 点数欄集計先識別（入院）
	AutoAnastomosisCategory        string  `csv:"67" json:"auto_anastomosis_category"`         // 67: 自動吻合器加算区分
	NotificationCategory1          string  `csv:"68" json:"notification_category1"`            // 68: 告示等識別区分（1）
	NotificationCategory2          string  `csv:"69" json:"notification_category2"`            // 69: 告示等識別区分（2）
	RegionAddition                 string  `csv:"70" json:"region_addition"`                   // 70: 地域加算
	BedCountCategory               string  `csv:"71" json:"bed_count_category"`                // 71: 病床数区分
	FacilityStandard1              string  `csv:"72" json:"facility_standard_1"`               // 72: 施設基準1
	FacilityStandard2              string  `csv:"73" json:"facility_standard_2"`               // 73: 施設基準2
	FacilityStandard3              string  `csv:"74" json:"facility_standard_3"`               // 74: 施設基準3
	FacilityStandard4              string  `csv:"75" json:"facility_standard_4"`               // 75: 施設基準4
	FacilityStandard5              string  `csv:"76" json:"facility_standard_5"`               // 76: 施設基準5
	FacilityStandard6              string  `csv:"77" json:"facility_standard_6"`               // 77: 施設基準6
	FacilityStandard7              string  `csv:"78" json:"facility_standard_7"`               // 78: 施設基準7
	FacilityStandard8              string  `csv:"79" json:"facility_standard_8"`               // 79: 施設基準8
	FacilityStandard9              string  `csv:"80" json:"facility_standard_9"`               // 80: 施設基準9
	FacilityStandard10             string  `csv:"81" json:"facility_standard_10"`              // 81: 施設基準10
	UltrasoundCoagulationCategory  string  `csv:"82" json:"ultrasound_coagulation_category"`   // 82: 超音波凝固切開装置等加算区分
	ShortStaySurgery               string  `csv:"83" json:"short_stay_surgery"`                // 83: 短期滞在手術
	DentalApplicability            string  `csv:"84" json:"dental_applicability"`              // 84: 歯科適用区分
	SectionCodeAlpha               string  `csv:"85" json:"section_code_alpha"`                // 85: コード表用番号（アルファベット部）
	NotificationRelatedAlpha       string  `csv:"86" json:"notification_related_alpha"`        // 86: 告示・通知関連番号（アルファベット部）
	UpdateDate                     string  `csv:"87" json:"update_date"`                       // 87: 変更年月日
	ExpiryDate                     string  `csv:"88" json:"expiry_date"`                       // 88: 廃止年月日
	PublicationOrder               string  `csv:"89" json:"publication_order"`                 // 89: 公表順序番号
	SectionChapter                 string  `csv:"90" json:"section_chapter"`                   // 90: 章 (コード表用番号)
	SectionPart                    string  `csv:"91" json:"section_part"`                      // 91: 部 (コード表用番号)
	SectionNumber                  string  `csv:"92" json:"section_number"`                    // 92: 区分番号 (コード表用番号)
	SectionSubNumber               string  `csv:"93" json:"section_sub_number"`                // 93: 枝番 (コード表用番号)
	SectionItemNumber              string  `csv:"94" json:"section_item_number"`               // 94: 項番 (コード表用番号)
	RelatedSectionChapter          string  `csv:"95" json:"related_section_chapter"`           // 95: 章 (告示・通知関連番号)
	RelatedSectionPart             string  `csv:"96" json:"related_section_part"`              // 96: 部 (告示・通知関連番号)
	RelatedSectionNumber           string  `csv:"97" json:"related_section_number"`            // 97: 区分番号 (告示・通知関連番号)
	RelatedSectionSubNumber        string  `csv:"98" json:"related_section_sub_number"`        // 98: 枝番 (告示・通知関連番号)
	RelatedSectionItemNumber       string  `csv:"99" json:"related_section_item_number"`       // 99: 項番 (告示・通知関連番号)
	AgeAdditionLower1              string  `csv:"100" json:"age_addition_lower_1"`             // 100: 下限年齢 (年齢加算1)
	AgeAdditionUpper1              string  `csv:"101" json:"age_addition_upper_1"`             // 101: 上限年齢 (年齢加算1)
	AgeAdditionCode1               string  `csv:"102" json:"age_addition_code_1"`              // 102: 注加算診療行為コード (年齢加算1)
	AgeAdditionLower2              string  `csv:"103" json:"age_addition_lower_2"`             // 103: 下限年齢 (年齢加算2)
	AgeAdditionUpper2              string  `csv:"104" json:"age_addition_upper_2"`             // 104: 上限年齢 (年齢加算2)
	AgeAdditionCode2               string  `csv:"105" json:"age_addition_code_2"`              // 105: 注加算診療行為コード (年齢加算2)
	AgeAdditionLower3              string  `csv:"106" json:"age_addition_lower_3"`             // 106: 下限年齢 (年齢加算3)
	AgeAdditionUpper3              string  `csv:"107" json:"age_addition_upper_3"`             // 107: 上限年齢 (年齢加算3)
	AgeAdditionCode3               string  `csv:"108" json:"age_addition_code_3"`              // 108: 注加算診療行為コード (年齢加算3)
	AgeAdditionLower4              string  `csv:"109" json:"age_addition_lower_4"`             // 109: 下限年齢 (年齢加算4)
	AgeAdditionUpper4              string  `csv:"110" json:"age_addition_upper_4"`             // 110: 上限年齢 (年齢加算4)
	AgeAdditionCode4               string  `csv:"111" json:"age_addition_code_4"`              // 111: 注加算診療行為コード (年齢加算4)
	MigrationRelatedCode           string  `csv:"112" json:"migration_related_code"`           // 112: 異動関連
	BasicKanjiName                 string  `csv:"113" json:"basic_kanji_name"`                 // 113: 基本漢字名称
	SinusSurgeryEndoscope          string  `csv:"114" json:"sinus_surgery_endoscope"`          // 114: 副鼻腔手術用内視鏡加算
	SinusSurgeryBoneTissueDevice   string  `csv:"115" json:"sinus_surgery_bone_tissue_device"` // 115: 副鼻腔手術用骨軟部組織切除機器加算
	LongTimeAnesthesia             string  `csv:"116" json:"long_time_anesthesia"`             // 116: 長時間麻酔管理加算
	PointTableNumber               string  `csv:"117" json:"point_table_number"`               // 117: 点数表区分番号
	MonitoringAddition             string  `csv:"118" json:"monitoring_addition"`              // 118: モニタリング加算
	FrozenTissueAddition           string  `csv:"119" json:"frozen_tissue_addition"`           // 119: 凍結保存同種組織加算
	MalignantTumorPathology        string  `csv:"120" json:"malignant_tumor_pathology"`        // 120: 悪性腫瘍病理組織標本加算
	ExternalFixationAddition       string  `csv:"121" json:"external_fixation_addition"`       // 121: 創外固定器加算
	UltrasoundCuttingAddition      string  `csv:"122" json:"ultrasound_cutting_addition"`      // 122: 超音波切削機器加算
	LeftAtrialAppendageClosure     string  `csv:"123" json:"left_atrial_appendage_closure"`    // 123: 左心耳閉鎖術併施加算
	OutpatientInfectionControl     string  `csv:"124" json:"outpatient_infection_control"`     // 124: 外来感染対策向上加算等
	OtolaryngologyInfantTreatment  string  `csv:"125" json:"otolaryngology_infant_treatment"`  // 125: 耳鼻咽喉科乳幼児処置加算
	OtolaryngologyAntimicrobial    string  `csv:"126" json:"otolaryngology_antimicrobial"`     // 126: 耳鼻咽喉科小児抗菌薬適正使用支援加算
	NegativePressureDressing       string  `csv:"127" json:"negative_pressure_dressing"`       // 127: 切開創局所陰圧閉鎖処置機器加算
	NursingStaffTreatmentImprove   string  `csv:"128" json:"nursing_staff_treatment_improve"`  // 128: 看護職員処遇改善評価料等
	OutpatientBaseUpEvaluation1    string  `csv:"129" json:"outpatient_base_up_evaluation_1"`  // 129: 外来・在宅ベースアップ評価料（1）
	OutpatientBaseUpEvaluation2    string  `csv:"130" json:"outpatient_base_up_evaluation_2"`  // 130: 外来・在宅ベースアップ評価料（2）
	RemanufacturedSingleUseDevice  string  `csv:"131" json:"remanufactured_single_use_device"` // 131: 再製造単回使用医療機器使用加算
	Reserved132                    string  `csv:"132" json:"reserved_132"`                     // 132: 予備
	Reserved133                    string  `csv:"133" json:"reserved_133"`                     // 133: 予備
	Reserved134                    string  `csv:"134" json:"reserved_134"`                     // 134: 予備
	Reserved135                    string  `csv:"135" json:"reserved_135"`                     // 135: 予備
	Reserved136                    string  `csv:"136" json:"reserved_136"`                     // 136: 予備
	Reserved137                    string  `csv:"137" json:"reserved_137"`                     // 137: 予備
	Reserved138                    string  `csv:"138" json:"reserved_138"`                     // 138: 予備
	Reserved139                    string  `csv:"139" json:"reserved_139"`                     // 139: 予備
	Reserved140                    string  `csv:"140" json:"reserved_140"`                     // 140: 予備
	Reserved141                    string  `csv:"141" json:"reserved_141"`                     // 141: 予備
	Reserved142                    string  `csv:"142" json:"reserved_142"`                     // 142: 予備
	Reserved143                    string  `csv:"143" json:"reserved_143"`                     // 143: 予備
	Reserved144                    string  `csv:"144" json:"reserved_144"`                     // 144: 予備
	Reserved145                    string  `csv:"145" json:"reserved_145"`                     // 145: 予備
	Reserved146                    string  `csv:"146" json:"reserved_146"`                     // 146: 予備
	Reserved147                    string  `csv:"147" json:"reserved_147"`                     // 147: 予備
	Reserved148                    string  `csv:"148" json:"reserved_148"`                     // 148: 予備
	Reserved149                    string  `csv:"149" json:"reserved_149"`                     // 149: 予備
	Reserved150                    string  `csv:"150" json:"reserved_150"`                     // 150: 予備
}

// MedicalPracticeInclusion 包括テーブル (02包括テーブル.csv)
// 全7項目を網羅 (R06rec3.pdf 27ページ [仕様説明書はPDF 27, 201] 201ページは医科、207ページ以降に補助テーブルあり)
// PDF 27ページ (R06rec1.pdf) によれば包括・被包括テーブルは7項目
type MedicalPracticeInclusion struct {
	ChangeCategory            string `csv:"1" json:"change_category"`             // 1: 変更区分
	ComprehensiveGroupNumber  string `csv:"2" json:"comprehensive_group_number"`  // 2: グループ番号
	ComprehensivePracticeCode string `csv:"3" json:"comprehensive_practice_code"` // 3: 診療行為コード
	AbbreviatedKanjiName      string `csv:"4" json:"abbreviated_kanji_name"`      // 4: 診療行為省略名称
	SpecialCondition          string `csv:"5" json:"special_condition"`           // 5: 特例条件
	UpdateDate                string `csv:"6" json:"update_date"`                 // 6: 新設年月日
	ExpiryDate                string `csv:"7" json:"expiry_date"`                 // 7: 廃止年月日
}

// MedicalPracticeConflict 背反関連テーブル (03-*背反テーブル.csv)
// 全10項目を網羅 (R06rec3.pdf 28ページ)
type MedicalPracticeConflict struct {
	ChangeCategory        string `csv:"1" json:"change_category"`          // 1: 変更区分
	MedicalPracticeCode1  string `csv:"2" json:"medical_practice_code_1"`  // 2: 診療行為コード1
	AbbreviatedKanjiName1 string `csv:"3" json:"abbreviated_kanji_name_1"` // 3: 診療行為省略名称1
	MedicalPracticeCode2  string `csv:"4" json:"medical_practice_code_2"`  // 4: 診療行為コード2
	AbbreviatedKanjiName2 string `csv:"5" json:"abbreviated_kanji_name_2"` // 5: 診療行為省略名称2
	ConflictCategory      string `csv:"6" json:"conflict_category"`        // 6: 背反区分
	SpecialCondition      string `csv:"7" json:"special_condition"`        // 7: 特例条件
	Reserved8             string `csv:"8" json:"reserved_8"`               // 8: 予備
	UpdateDate            string `csv:"9" json:"update_date"`              // 9: 新設年月日
	ExpiryDate            string `csv:"10" json:"expiry_date"`             // 10: 廃止年月日
}

// MedicalPracticeSupport 医科診療行為マスター補助マスターテーブル (01補助マスターテーブル.csv)
// 全27項目を網羅 (R06rec3.pdf 25-26ページ)
type MedicalPracticeSupport struct {
	ChangeCategory           string `csv:"1" json:"change_category"`             // 1: 変更区分
	MedicalPracticeCode      string `csv:"2" json:"medical_practice_code"`       // 2: 診療行為コード
	AbbreviatedKanjiName     string `csv:"3" json:"abbreviated_kanji_name"`      // 3: 診療行為省略名称
	ComprehensiveUnit1       string `csv:"4" json:"comprehensive_unit_1"`        // 4: 包括単位1
	ComprehensiveGroup1      string `csv:"5" json:"comprehensive_group_1"`       // 5: グループ番号1
	ComprehensiveUnit2       string `csv:"6" json:"comprehensive_unit_2"`        // 6: 包括単位2
	ComprehensiveGroup2      string `csv:"7" json:"comprehensive_group_2"`       // 7: グループ番号2
	ComprehensiveUnit3       string `csv:"8" json:"comprehensive_unit_3"`        // 8: 包括単位3
	ComprehensiveGroup3      string `csv:"9" json:"comprehensive_group_3"`       // 9: グループ番号3
	ConflictDaily            string `csv:"10" json:"conflict_daily"`             // 10: 1日につき (背反関連識別)
	ConflictMonthly          string `csv:"11" json:"conflict_monthly"`           // 11: 同一月内 (背反関連識別)
	ConflictSimultaneous     string `csv:"12" json:"conflict_simultaneous"`      // 12: 同時 (背反関連識別)
	ConflictWeekly           string `csv:"13" json:"conflict_weekly"`            // 13: 1週間につき (背反関連識別)
	Reserved14               string `csv:"14" json:"reserved_14"`                // 14: 予備
	Reserved15               string `csv:"15" json:"reserved_15"`                // 15: 予備
	Reserved16               string `csv:"16" json:"reserved_16"`                // 16: 予備
	Reserved17               string `csv:"17" json:"reserved_17"`                // 17: 予備
	Reserved18               string `csv:"18" json:"reserved_18"`                // 18: 予備
	Reserved19               string `csv:"19" json:"reserved_19"`                // 19: 予備
	InpatientBasicFeeGroup   string `csv:"20" json:"inpatient_basic_fee_group"`  // 20: 入院基本料識別
	CalculationCountAddition string `csv:"21" json:"calculation_count_addition"` // 21: 算定回数関連
	Reserved22               string `csv:"22" json:"reserved_22"`                // 22: 予備
	Reserved23               string `csv:"23" json:"reserved_23"`                // 23: 予備
	Reserved24               string `csv:"24" json:"reserved_24"`                // 24: 予備
	Reserved25               string `csv:"25" json:"reserved_25"`                // 25: 予備
	UpdateDate               string `csv:"26" json:"update_date"`                // 26: 新設年月日
	ExpiryDate               string `csv:"27" json:"expiry_date"`                // 27: 廃止年月日
}

// InpatientBasicFee 入院基本料テーブル (04入院基本料テーブル.csv)
// 全8項目を網羅 (R06rec3.pdf 29ページ)
type InpatientBasicFee struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 1: 変更区分
	BasicFeeGroup        string `csv:"2" json:"basic_fee_group"`        // 2: グループ番号
	MedicalPracticeCode  string `csv:"3" json:"medical_practice_code"`  // 3: 診療行為コード
	AbbreviatedKanjiName string `csv:"4" json:"abbreviated_kanji_name"` // 4: 診療行為省略名称
	AdditionIdentifier   string `csv:"5" json:"addition_identifier"`    // 5: 加算識別
	Reserved6            string `csv:"6" json:"reserved_6"`             // 6: 予備
	UpdateDate           string `csv:"7" json:"update_date"`            // 7: 新設年月日
	ExpiryDate           string `csv:"8" json:"expiry_date"`            // 8: 廃止年月日
}

// CalculationCount 算定回数テーブル (05算定回数テーブル.csv)
// 全14項目を網羅 (R06rec3.pdf 30ページ)
type CalculationCount struct {
	ChangeCategory       string `csv:"1" json:"change_category"`        // 1: 変更区分
	MedicalPracticeCode  string `csv:"2" json:"medical_practice_code"`  // 3: 診療行為コード
	AbbreviatedKanjiName string `csv:"3" json:"abbreviated_kanji_name"` // 3: 診療行為省略名称
	UnitCode             string `csv:"4" json:"unit_code"`              // 4: 算定単位コード
	UnitName             string `csv:"5" json:"unit_name"`              // 5: 算定単位名称
	MaxTimes             int    `csv:"6" json:"max_times"`              // 6: 算定回数
	SpecialCondition     string `csv:"7" json:"special_condition"`      // 7: 特例条件
	Reserved8            string `csv:"8" json:"reserved_8"`             // 8: 予備
	Reserved9            string `csv:"9" json:"reserved_9"`             // 9: 予備
	Reserved10           string `csv:"10" json:"reserved_10"`           // 10: 予備
	Reserved11           string `csv:"11" json:"reserved_11"`           // 11: 予備
	Reserved12           string `csv:"12" json:"reserved_12"`           // 12: 予備
	UpdateDate           string `csv:"13" json:"update_date"`           // 13: 新設年月日
	ExpiryDate           string `csv:"14" json:"expiry_date"`           // 14: 廃止年月日
}
