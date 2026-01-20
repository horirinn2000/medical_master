package model

import (
	"time"
)

// Ward 病棟マスター (k.csv)
type Ward struct {
	ID                         uint   `gorm:"primaryKey" json:"id"`
	UpdateCategory             string `json:"update_category"`                // 1: 変更区分
	MasterType                 string `json:"master_type"`                    // 2: マスター種別
	Code                       string `gorm:"index" json:"code"`              // 3: 診療行為コード
	NameKanjiLen               int    `json:"name_kanji_len"`                 // 4: 省略漢字有効桁数
	NameKanji                  string `gorm:"index" json:"name_kanji"`        // 5: 省略漢字名称
	NameKanaLen                int    `json:"name_kana_len"`                  // 6: 省略カナ有効桁数
	NameKana                   string `json:"name_kana"`                      // 7: 省略カナ名称
	DataStandardCode           string `json:"data_standard_code"`             // 8: データ規格コード
	KanjiUnitNameLen           int    `json:"kanji_unit_name_len"`            // 9: 漢字有効桁数（単位）
	KanjiUnitName              string `json:"kanji_unit_name"`                // 10: 漢字名称（単位）
	PointCategory              string `json:"point_category"`                 // 11: 点数識別
	Point                      string `json:"point"`                          // 12: 新又は現点数 (小数点含むためstring推奨だが、計算に使うならfloat/decimal。既存はfloat64だがCSVママならstring) -> PDF "0.00" なので string または float。既存に合わせてfloatにするか、精度重視でstringにするか。ここでは既存コードが parseFloat しているので float64 にする。
	InOuterCategory            string `json:"in_outer_category"`              // 13: 入外適用区分
	ElderlyCategory            string `json:"elderly_category"`               // 14: 後期高齢者医療適用区分
	OutpatientAggregation      string `json:"outpatient_aggregation"`         // 15: 点数欄集計先識別（入院外）
	ComprehensiveTest          string `json:"comprehensive_test"`             // 16: 包括対象検査
	Reserved1                  string `json:"reserved_1"`                     // 17: 予備
	DpcCategory                string `json:"dpc_category"`                   // 18: DPC適用区分
	InstitutionCategory        string `json:"institution_category"`           // 19: 病院・診療所区分
	ImageSurgerySupport        string `json:"image_surgery_support"`          // 20: 画像等手術支援加算
	MedicalObservation         string `json:"medical_observation"`            // 21: 医療観察法対象区分
	NursingAddition            string `json:"nursing_addition"`               // 22: 看護加算
	AnesthesiaCategory         string `json:"anesthesia_category"`            // 23: 麻酔識別区分
	Reserved2                  string `json:"reserved_2"`                     // 24: 予備
	DiseaseRelated             string `json:"disease_related"`                // 25: 傷病名関連区分
	Reserved3                  string `json:"reserved_3"`                     // 26: 予備
	ActualDays                 string `json:"actual_days"`                    // 27: 実日数
	DaysTimes                  string `json:"days_times"`                     // 28: 日数・回数
	MedicineRelated            string `json:"medicine_related"`               // 29: 医薬品関連区分
	StepCalculation            string `json:"step_calculation"`               // 30: きざみ値計算識別
	StepLowerLimit             string `json:"step_lower_limit"`               // 31: 下限値
	StepUpperLimit             string `json:"step_upper_limit"`               // 32: 上限値
	StepValue                  string `json:"step_value"`                     // 33: きざみ値
	StepPoint                  string `json:"step_point"`                     // 34: きざみ点数
	StepErrorTreatment         string `json:"step_error_treatment"`           // 35: 上下限エラー処理
	MaxTimes                   string `json:"max_times"`                      // 36: 上限回数
	MaxTimesErrorTreatment     string `json:"max_times_error_treatment"`      // 37: 上限回数エラー処理
	NoteAdditionCode           string `json:"note_addition_code"`             // 38: 注加算コード
	NoteAdditionOrder          string `json:"note_addition_order"`            // 39: 注加算通番
	GeneralRuleAge             string `json:"general_rule_age"`               // 40: 通則年齢
	LowerAge                   string `json:"lower_age"`                      // 41: 下限年齢
	UpperAge                   string `json:"upper_age"`                      // 42: 上限年齢
	TimeAddition               string `json:"time_addition"`                  // 43: 時間加算区分
	StandardConformity         string `json:"standard_conformity"`            // 44: 適合区分
	TargetFacilityStandard     string `json:"target_facility_standard"`       // 45: 対象施設基準
	InfantTreatmentAddition    string `json:"infant_treatment_addition"`      // 46: 処置乳幼児加算区分
	VeryLowBirthWeightAddition string `json:"very_low_birth_weight_addition"` // 47: 極低出生体重児加算区分
	InpatientReduction         string `json:"inpatient_reduction"`            // 48: 入院基本料等減算対象識別
	DonorAggregation           string `json:"donor_aggregation"`              // 49: ドナー分集計区分
	TestJudgment               string `json:"test_judgment"`                  // 50: 検査等実施判断区分
	TestJudgmentGroup          string `json:"test_judgment_group"`            // 51: 検査等実施判断グループ区分
	ReductionTarget            string `json:"reduction_target"`               // 52: 逓減対象区分
	SpinalEvokedPotential      string `json:"spinal_evoked_potential"`        // 53: 脊髄誘発電位測定等加算区分
	NeckDissection             string `json:"neck_dissection"`                // 54: 頸部郭清術併施加算区分
	AutoSuture                 string `json:"auto_suture"`                    // 55: 自動縫合器加算区分
	OutpatientManagement       string `json:"outpatient_management"`          // 56: 外来管理加算区分
	OldPointCategory           string `json:"old_point_category"`             // 57: 点数識別（旧）
	OldPoint                   string `json:"old_point"`                      // 58: 旧点数
	NameKanjiUpdate            string `json:"name_kanji_update"`              // 59: 漢字名称変更区分
	NameKanaUpdate             string `json:"name_kana_update"`               // 60: カナ名称変更区分
	SpecimenTestComment        string `json:"specimen_test_comment"`          // 61: 検体検査コメント
	GeneralRuleAddition        string `json:"general_rule_addition"`          // 62: 通則加算所定点数対象区分
	ComprehensiveReduction     string `json:"comprehensive_reduction"`        // 63: 包括逓減区分
	UltrasoundEndoscopy        string `json:"ultrasound_endoscopy"`           // 64: 超音波内視鏡検査加算区分
	Reserved4                  string `json:"reserved_4"`                     // 65: 予備
	InpatientAggregation       string `json:"inpatient_aggregation"`          // 66: 点数欄集計先識別（入院）
	AutoAnastomosis            string `json:"auto_anastomosis"`               // 67: 自動吻合器加算区分
	NotificationCategory1      string `json:"notification_category_1"`        // 68: 告示等識別区分（１）
	NotificationCategory2      string `json:"notification_category_2"`        // 69: 告示等識別区分（２）
	RegionAddition             string `json:"region_addition"`                // 70: 地域加算
	BedCountCategory           string `json:"bed_count_category"`             // 71: 病床数区分
	FacilityStandard1          string `json:"facility_standard_1"`            // 72: 施設基準1
	FacilityStandard2          string `json:"facility_standard_2"`            // 73: 施設基準2
	FacilityStandard3          string `json:"facility_standard_3"`            // 74: 施設基準3
	FacilityStandard4          string `json:"facility_standard_4"`            // 75: 施設基準4
	FacilityStandard5          string `json:"facility_standard_5"`            // 76: 施設基準5
	FacilityStandard6          string `json:"facility_standard_6"`            // 77: 施設基準6
	FacilityStandard7          string `json:"facility_standard_7"`            // 78: 施設基準7
	FacilityStandard8          string `json:"facility_standard_8"`            // 79: 施設基準8
	FacilityStandard9          string `json:"facility_standard_9"`            // 80: 施設基準9
	FacilityStandard10         string `json:"facility_standard_10"`           // 81: 施設基準10
	UltrasoundCoagulation      string `json:"ultrasound_coagulation"`         // 82: 超音波凝固切開装置等加算区分
	ShortStaySurgery           string `json:"short_stay_surgery"`             // 83: 短期滞在手術
	DentalCategory             string `json:"dental_category"`                // 84: 歯科適用区分
	CodeTableAlpha             string `json:"code_table_alpha"`               // 85: コード表用番号（アルファベット部）
	NotificationAlpha          string `json:"notification_alpha"`             // 86: 告示・通知関連番号（アルファベット部）
	UpdateDate                 string `json:"update_date"`                    // 87: 変更年月日
	DiscontinuedDate           string `json:"discontinued_date"`              // 88: 廃止年月日
	PublishOrder               string `json:"publish_order"`                  // 89: 公表順序番号
	Chapter                    string `json:"chapter"`                        // 90: 章
	Section                    string `json:"section"`                        // 91: 部
	Subsection                 string `json:"subsection"`                     // 92: 区分番号
	BranchNumber               string `json:"branch_number"`                  // 93: 枝番
	ItemNumber                 string `json:"item_number"`                    // 94: 項番
	NotificationChapter        string `json:"notification_chapter"`           // 95: 章（告示・通知）
	NotificationSection        string `json:"notification_section"`           // 96: 部（告示・通知）
	NotificationSubsection     string `json:"notification_subsection"`        // 97: 区分番号（告示・通知）
	NotificationBranch         string `json:"notification_branch"`            // 98: 枝番（告示・通知）
	NotificationItem           string `json:"notification_item"`              // 99: 項番（告示・通知）
	AgeAddition1Lower          string `json:"age_addition_1_lower"`           // 100: 年齢加算1（下限年齢）
	AgeAddition1Upper          string `json:"age_addition_1_upper"`           // 101: 年齢加算1（上限年齢）
	AgeAddition1Code           string `json:"age_addition_1_code"`            // 102: 年齢加算1（診療行為コード）
	AgeAddition2Lower          string `json:"age_addition_2_lower"`           // 103: 年齢加算2（下限年齢）
	AgeAddition2Upper          string `json:"age_addition_2_upper"`           // 104: 年齢加算2（上限年齢）
	AgeAddition2Code           string `json:"age_addition_2_code"`            // 105: 年齢加算2（診療行為コード）
	AgeAddition3Lower          string `json:"age_addition_3_lower"`           // 106: 年齢加算3（下限年齢）
	AgeAddition3Upper          string `json:"age_addition_3_upper"`           // 107: 年齢加算3（上限年齢）
	AgeAddition3Code           string `json:"age_addition_3_code"`            // 108: 年齢加算3（診療行為コード）
	AgeAddition4Lower          string `json:"age_addition_4_lower"`           // 109: 年齢加算4（下限年齢）
	AgeAddition4Upper          string `json:"age_addition_4_upper"`           // 110: 年齢加算4（上限年齢）
	AgeAddition4Code           string `json:"age_addition_4_code"`            // 111: 年齢加算4（診療行為コード）
	MigrationRelated           string `json:"migration_related"`              // 112: 異動関連
	Fullname                   string `json:"fullname"`                       // 113: 基本漢字名称
	SinusSurgeryEndoscopy      string `json:"sinus_surgery_endoscopy"`        // 114: 副鼻腔手術用内視鏡加算
	SinusSurgeryBoneCutting    string `json:"sinus_surgery_bone_cutting"`     // 115: 副鼻腔手術用骨軟部組織切除機器加算
	LongTermAnesthesia         string `json:"long_term_anesthesia"`           // 116: 長時間麻酔管理加算
	PointTableNumber           string `json:"point_table_number"`             // 117: 点数表区分番号
	MonitoringAddition         string `json:"monitoring_addition"`            // 118: モニタリング加算
	FrozenTissueAddition       string `json:"frozen_tissue_addition"`         // 119: 凍結保存同種組織加算
	MalignantTumorAddition     string `json:"malignant_tumor_addition"`       // 120: 悪性腫瘍病理組織標本加算
	ExternalFixatorAddition    string `json:"external_fixator_addition"`      // 121: 創外固定器加算
	UltrasoundCuttingAddition  string `json:"ultrasound_cutting_addition"`    // 122: 超音波切削機器加算
	LeftAtrialAppendageClosure string `json:"left_atrial_appendage_closure"`  // 123: 左心耳閉鎖術併施区分
	OutpatientInfectionControl string `json:"outpatient_infection_control"`   // 124: 外来感染対策向上加算等
	EntInfantTreatment         string `json:"ent_infant_treatment"`           // 125: 耳鼻咽喉科乳幼児処置加算
	EntAntimicrobialSupport    string `json:"ent_antimicrobial_support"`      // 126: 耳鼻咽喉科小児抗菌薬適正使用支援加算
	NegativePressureClosure    string `json:"negative_pressure_closure"`      // 127: 切開創局所陰圧閉鎖処置機器加算
	NursingStaffTreatment      string `json:"nursing_staff_treatment"`        // 128: 看護職員処遇改善評価料等
	OutpatientBaseUp1          string `json:"outpatient_base_up_1"`           // 129: 外来・在宅ベースアップ評価料（１）
	OutpatientBaseUp2          string `json:"outpatient_base_up_2"`           // 130: 外来・在宅ベースアップ評価料（２）
	RemanufacturedDevice       string `json:"remanufactured_device"`          // 131: 再製造単回使用医療機器使用加算
	Reserved5                  string `json:"reserved_5"`                     // 132: 予備
	Reserved6                  string `json:"reserved_6"`                     // 133: 予備
	Reserved7                  string `json:"reserved_7"`                     // 134: 予備
	Reserved8                  string `json:"reserved_8"`                     // 135: 予備
	Reserved9                  string `json:"reserved_9"`                     // 136: 予備
	Reserved10                 string `json:"reserved_10"`                    // 137: 予備
	Reserved11                 string `json:"reserved_11"`                    // 138: 予備
	Reserved12                 string `json:"reserved_12"`                    // 139: 予備
	Reserved13                 string `json:"reserved_13"`                    // 140: 予備
	Reserved14                 string `json:"reserved_14"`                    // 141: 予備
	Reserved15                 string `json:"reserved_15"`                    // 142: 予備
	Reserved16                 string `json:"reserved_16"`                    // 143: 予備
	Reserved17                 string `json:"reserved_17"`                    // 144: 予備
	Reserved18                 string `json:"reserved_18"`                    // 145: 予備
	Reserved19                 string `json:"reserved_19"`                    // 146: 予備
	Reserved20                 string `json:"reserved_20"`                    // 147: 予備
	Reserved21                 string `json:"reserved_21"`                    // 148: 予備
	Reserved22                 string `json:"reserved_22"`                    // 149: 予備
	Reserved23                 string `json:"reserved_23"`                    // 150: 予備

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
