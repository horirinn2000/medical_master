package model

// MedicalPractice 医科診療行為マスター (s_*.csv)
type MedicalPractice struct {
	ChangeCategory                 string  `csv:"1"`   // 変更区分
	MasterType                     string  `csv:"2"`   // マスター種別
	MedicalPracticeCode            string  `csv:"3"`   // 診療行為コード
	AbbreviatedKanjiNameLength     int     `csv:"4"`   // 省略漢字有効桁数
	AbbreviatedKanjiName           string  `csv:"5"`   // 省略漢字名称
	AbbreviatedKanaNameLength      int     `csv:"6"`   // 省略カナ有効桁数
	AbbreviatedKanaName            string  `csv:"7"`   // 省略カナ名称
	DataStandardCode               string  `csv:"8"`   // データ規格コード
	KanjiUnitNameLength            int     `csv:"9"`   // 漢字名称有効桁数（単位）
	KanjiUnitName                  string  `csv:"10"`  // 漢字名称（単位）
	ScoreType                      string  `csv:"11"`  // 点数識別
	Score                          float64 `csv:"12"`  // 新又は現点数
	InpatientOutpatientCategory    string  `csv:"13"`  // 入外適用区分
	ElderlyMedicalCategory         string  `csv:"14"`  // 後期高齢者医療適用区分
	AggregationCategoryOutpatient  string  `csv:"15"`  // 点数欄集計先識別（入院外）
	ComprehensiveTargetTest        string  `csv:"16"`  // 包括対象検査
	DPCApplicabilityCategory       string  `csv:"18"`  // DPC適用区分
	HospitalClinicCategory         string  `csv:"19"`  // 病院・診療所区分
	SurgerySupportCategory         string  `csv:"20"`  // 画像等手術支援加算
	MedicalObservationCategory     string  `csv:"21"`  // 医療観察法対象区分
	NursingAddition                string  `csv:"22"`  // 看護加算
	AnesthesiaCategory             string  `csv:"23"`  // 麻酔識別区分
	DiseaseRelatedCategory         string  `csv:"25"`  // 傷病名関連区分
	ActualDaysCategory             string  `csv:"27"`  // 実日数
	DaysTimesCategory              string  `csv:"28"`  // 日数・回数
	MedicineRelatedCategory        string  `csv:"29"`  // 医薬品関連区分
	StepCalculationCategory        string  `csv:"30"`  // きざみ値計算識別
	StepLowerLimit                 float64 `csv:"31"`  // 下限値
	StepUpperLimit                 float64 `csv:"32"`  // 上限値
	StepValue                      float64 `csv:"33"`  // きざみ値
	StepScore                      float64 `csv:"34"`  // きざみ点数
	StepErrorCategory              string  `csv:"35"`  // 上下限エラー処理
	MaxTimes                       int     `csv:"36"`  // 上限回数
	MaxTimesErrorCategory          string  `csv:"37"`  // 上限回数エラー処理
	NoteAdditionCode               string  `csv:"38"`  // 注加算コード
	NoteAdditionSequence           string  `csv:"39"`  // 注加算通番
	GeneralAgeCategory             string  `csv:"40"`  // 通則年齢
	LowerAge                       string  `csv:"41"`  // 下限年齢
	UpperAge                       string  `csv:"42"`  // 上限年齢
	TimeAdditionCategory           string  `csv:"43"`  // 時間加算区分
	FacilityStandardCategory       string  `csv:"44"`  // 適合区分
	FacilityStandardCode           string  `csv:"45"`  // 対象施設基準
	TreatmentInfantAddition        string  `csv:"46"`  // 処置乳幼児加算区分
	VeryLowBirthWeightAddition     string  `csv:"47"`  // 極低出生体重児加算区分
	InpatientReductionCategory     string  `csv:"48"`  // 入院基本料等減算対象識別
	DonorCollectionCategory        string  `csv:"49"`  // ドナー分集計区分
	TestJudgmentCategory           string  `csv:"50"`  // 検査等実施判断区分
	TestJudgmentGroupCategory      string  `csv:"51"`  // 検査等実施判断グループ区分
	ReductionApplicability         string  `csv:"52"`  // 逓減対象区分
	SpinalEvokedPotentialCategory  string  `csv:"53"`  // 脊髄誘発電位測定等加算区分
	NeckDissectionCategory         string  `csv:"54"`  // 頸部郭清術併施加算区分
	AutoSutureCategory             string  `csv:"55"`  // 自動縫合器加算区分
	OutpatientManagementAddition   string  `csv:"56"`  // 外来管理加算区分
	OldScoreType                   string  `csv:"57"`  // 点数識別（旧）
	OldScore                       float64 `csv:"58"`  // 旧点数
	TestCommentCategory            string  `csv:"61"`  // 検体検査コメント
	GeneralAdditionApplicability   string  `csv:"62"`  // 通則加算所定点数対象区分
	ComprehensiveReductionCategory string  `csv:"63"`  // 包括逓減区分
	EndoscopicSurgeryAddition      string  `csv:"64"`  // 超音波内視鏡検査加算区分
	AggregationCategoryInpatient   string  `csv:"66"`  // 点数欄集計先識別（入院）
	AutoAnastomosisCategory        string  `csv:"67"`  // 自動吻合器加算区分
	NotificationCategory1          string  `csv:"68"`  // 告示等識別区分（1）
	NotificationCategory2          string  `csv:"69"`  // 告示等識別区分（2）
	RegionAddition                 string  `csv:"70"`  // 地域加算
	BedCountCategory               string  `csv:"71"`  // 病床数区分
	FacilityStandard1              string  `csv:"72"`  // 施設基準1
	FacilityStandard2              string  `csv:"73"`  // 施設基準2
	FacilityStandard3              string  `csv:"74"`  // 施設基準3
	FacilityStandard4              string  `csv:"75"`  // 施設基準4
	FacilityStandard5              string  `csv:"76"`  // 施設基準5
	FacilityStandard6              string  `csv:"77"`  // 施設基準6
	FacilityStandard7              string  `csv:"78"`  // 施設基準7
	FacilityStandard8              string  `csv:"79"`  // 施設基準8
	FacilityStandard9              string  `csv:"80"`  // 施設基準9
	FacilityStandard10             string  `csv:"81"`  // 施設基準10
	UltrasoundCoagulationCategory  string  `csv:"82"`  // 超音波凝固切開装置等加算区分
	ShortStaySurgery               string  `csv:"83"`  // 短期滞在手術
	DentalApplicability            string  `csv:"84"`  // 歯科適用区分
	SectionCodeAlpha               string  `csv:"85"`  // コード表用番号（アルファベット部）
	NotificationRelatedAlpha       string  `csv:"86"`  // 告示・通知関連番号（アルファベット部）
	UpdateDate                     string  `csv:"87"`  // 変更年月日
	ExpiryDate                     string  `csv:"88"`  // 廃止年月日
	PublicationOrder               string  `csv:"89"`  // 公表順序番号
	SectionChapter                 string  `csv:"90"`  // 章
	SectionPart                    string  `csv:"91"`  // 部
	SectionNumber                  string  `csv:"92"`  // 区分番号
	SectionSubNumber               string  `csv:"93"`  // 枝番
	SectionItemNumber              string  `csv:"94"`  // 項番
	BasicKanjiName                 string  `csv:"113"` // 基本漢字名称
	PointTableNumber               string  `csv:"117"` // 点数表区分番号
}

// MedicalPracticeInclusion 包括テーブル (02包括テーブル.csv)
type MedicalPracticeInclusion struct {
	ChangeCategory            string `csv:"1"` // 変更区分
	ComprehensivePracticeCode string `csv:"2"` // 包括診療行為コード（A000001など）
	IncludedPracticeCode      string `csv:"3"` // 包括対象診療行為コード
	UpdateDate                string `csv:"6"` // 変更年月日
	ExpiryDate                string `csv:"7"` // 廃止年月日
}

// MedicalPracticeConflict 背反テーブル (03-*背反テーブル.csv)
type MedicalPracticeConflict struct {
	ChangeCategory       string `csv:"1"`   // 変更区分
	MedicalPracticeCode  string `csv:"2"`   // 診療行為コード
	ConflictPracticeCode string `csv:"12"`  // 背反診療行為コード (背反1)
	UpdateDate           string `csv:"100"` // 変更年月日
	ExpiryDate           string `csv:"101"` // 廃止年月日
}

// MedicalPracticeSupport 補助マスターテーブル (01補助マスターテーブル.csv)
type MedicalPracticeSupport struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	SupportInfo         string `csv:"4"`  // 補助情報
	UpdateDate          string `csv:"26"` // 変更年月日
	ExpiryDate          string `csv:"27"` // 廃止年月日
}

// InpatientBasicFee 入院基本料テーブル (04入院基本料テーブル.csv)
type InpatientBasicFee struct {
	ChangeCategory      string `csv:"1"` // 変更区分
	BasicFeeCode        string `csv:"2"` // 入院基本料コード
	MedicalPracticeCode string `csv:"3"` // 診療行為コード
	UpdateDate          string `csv:"7"` // 変更年月日
	ExpiryDate          string `csv:"8"` // 廃止年月日
}

// CalculationCount 算定回数テーブル (05算定回数テーブル.csv)
type CalculationCount struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	CountLimit          int    `csv:"6"`  // 算定回数限度
	UpdateDate          string `csv:"13"` // 変更年月日
	ExpiryDate          string `csv:"14"` // 廃止年月日
}
