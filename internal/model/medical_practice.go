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
	Reserved1                      string  `csv:"17"`  // 予備
	DPCApplicabilityCategory       string  `csv:"18"`  // DPC適用区分
	HospitalClinicCategory         string  `csv:"19"`  // 病院・診療所区分
	SurgerySupportCategory         string  `csv:"20"`  // 画像等手術支援加算
	MedicalObservationCategory     string  `csv:"21"`  // 医療観察法対象区分
	NursingAddition                string  `csv:"22"`  // 看護加算
	AnesthesiaCategory             string  `csv:"23"`  // 麻酔識別区分
	Reserved2                      string  `csv:"24"`  // 予備
	DiseaseRelatedCategory         string  `csv:"25"`  // 傷病名関連区分
	Reserved3                      string  `csv:"26"`  // 予備
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
	KanjiNameChangeCategory        string  `csv:"59"`  // 漢字名称変更区分
	KanaNameChangeCategory         string  `csv:"60"`  // カナ名称変更区分
	TestCommentCategory            string  `csv:"61"`  // 検体検査コメント
	GeneralAdditionApplicability   string  `csv:"62"`  // 通則加算所定点数対象区分
	ComprehensiveReductionCategory string  `csv:"63"`  // 包括逓減区分
	EndoscopicSurgeryAddition      string  `csv:"64"`  // 超音波内視鏡検査加算区分
	Reserved4                      string  `csv:"65"`  // 予備
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
	NotificationSectionChapter     string  `csv:"95"`  // 章（告示・通知）
	NotificationSectionPart        string  `csv:"96"`  // 部（告示・通知）
	NotificationSectionNumber      string  `csv:"97"`  // 区分番号（告示・通知）
	NotificationSectionSubNumber   string  `csv:"98"`  // 枝番（告示・通知）
	NotificationSectionItemNumber  string  `csv:"99"`  // 項番（告示・通知）
	AgeAddition1Lower              string  `csv:"100"` // 年齢加算1（下限年齢）
	AgeAddition1Upper              string  `csv:"101"` // 年齢加算1（上限年齢）
	AgeAddition1Code               string  `csv:"102"` // 年齢加算1（診療行為コード）
	AgeAddition2Lower              string  `csv:"103"` // 年齢加算2（下限年齢）
	AgeAddition2Upper              string  `csv:"104"` // 年齢加算2（上限年齢）
	AgeAddition2Code               string  `csv:"105"` // 年齢加算2（診療行為コード）
	AgeAddition3Lower              string  `csv:"106"` // 年齢加算3（下限年齢）
	AgeAddition3Upper              string  `csv:"107"` // 年齢加算3（上限年齢）
	AgeAddition3Code               string  `csv:"108"` // 年齢加算3（診療行為コード）
	AgeAddition4Lower              string  `csv:"109"` // 年齢加算4（下限年齢）
	AgeAddition4Upper              string  `csv:"110"` // 年齢加算4（上限年齢）
	AgeAddition4Code               string  `csv:"111"` // 年齢加算4（診療行為コード）
	MigrationRelated               string  `csv:"112"` // 異動関連
	BasicKanjiName                 string  `csv:"113"` // 基本漢字名称
	SinusSurgeryEndoscopy          string  `csv:"114"` // 副鼻腔手術用内視鏡加算
	SinusSurgeryBoneCutting        string  `csv:"115"` // 副鼻腔手術用骨軟部組織切除機器加算
	LongTermAnesthesia             string  `csv:"116"` // 長時間麻酔管理加算
	PointTableNumber               string  `csv:"117"` // 点数表区分番号
	MonitoringAddition             string  `csv:"118"` // モニタリング加算
	FrozenTissueAddition           string  `csv:"119"` // 凍結保存同種組織加算
	MalignantTumorAddition         string  `csv:"120"` // 悪性腫瘍病理組織標本加算
	ExternalFixatorAddition        string  `csv:"121"` // 創外固定器加算
	UltrasoundCuttingAddition      string  `csv:"122"` // 超音波切削機器加算
	LeftAtrialAppendageClosure     string  `csv:"123"` // 左心耳閉鎖術併施区分
	OutpatientInfectionControl     string  `csv:"124"` // 外来感染対策向上加算等
	EntInfantTreatment             string  `csv:"125"` // 耳鼻咽喉科乳幼児処置加算
	EntAntimicrobialSupport        string  `csv:"126"` // 耳鼻咽喉科小児抗菌薬適正使用支援加算
	NegativePressureClosure        string  `csv:"127"` // 切開創局所陰圧閉鎖処置機器加算
	NursingStaffTreatment          string  `csv:"128"` // 看護職員処遇改善評価料等
	OutpatientBaseUp1              string  `csv:"129"` // 外来・在宅ベースアップ評価料（１）
	OutpatientBaseUp2              string  `csv:"130"` // 外来・在宅ベースアップ評価料（２）
	RemanufacturedDevice           string  `csv:"131"` // 再製造単回使用医療機器使用加算
	Reserved5                      string  `csv:"132"` // 予備
	Reserved6                      string  `csv:"133"` // 予備
	Reserved7                      string  `csv:"134"` // 予備
	Reserved8                      string  `csv:"135"` // 予備
	Reserved9                      string  `csv:"136"` // 予備
	Reserved10                     string  `csv:"137"` // 予備
	Reserved11                     string  `csv:"138"` // 予備
	Reserved12                     string  `csv:"139"` // 予備
	Reserved13                     string  `csv:"140"` // 予備
	Reserved14                     string  `csv:"141"` // 予備
	Reserved15                     string  `csv:"142"` // 予備
	Reserved16                     string  `csv:"143"` // 予備
	Reserved17                     string  `csv:"144"` // 予備
	Reserved18                     string  `csv:"145"` // 予備
	Reserved19                     string  `csv:"146"` // 予備
	Reserved20                     string  `csv:"147"` // 予備
	Reserved21                     string  `csv:"148"` // 予備
	Reserved22                     string  `csv:"149"` // 予備
	Reserved23                     string  `csv:"150"` // 予備
}

// MedicalPracticeSupport 補助マスターテーブル (01補助マスターテーブル.csv)
type MedicalPracticeSupport struct {
	ChangeCategory            string `csv:"1"`  // 変更区分
	MedicalPracticeCode       string `csv:"2"`  // 診療行為コード
	AbbreviatedKanjiName      string `csv:"3"`  // 診療行為省略名称
	InclusionUnit1            string `csv:"4"`  // 包括単位①
	InclusionGroup1           string `csv:"5"`  // グループ番号①
	InclusionUnit2            string `csv:"6"`  // 包括単位②
	InclusionGroup2           string `csv:"7"`  // グループ番号②
	InclusionUnit3            string `csv:"8"`  // 包括単位③
	InclusionGroup3           string `csv:"9"`  // グループ番号③
	ConflictPerDay            string `csv:"10"` // 背反関連識別（1日につき）
	ConflictPerMonth          string `csv:"11"` // 背反関連識別（同一月内）
	ConflictSimultaneous      string `csv:"12"` // 背反関連識別（同時）
	ConflictPerWeek           string `csv:"13"` // 背反関連識別（1週間につき）
	Reserved1                 string `csv:"14"` // 予備
	Reserved2                 string `csv:"15"` // 予備
	Reserved3                 string `csv:"16"` // 予備
	Reserved4                 string `csv:"17"` // 予備
	Reserved5                 string `csv:"18"` // 予備
	Reserved6                 string `csv:"19"` // 予備
	InpatientBasicFeeCategory string `csv:"20"` // 入院基本料識別
	CalculationCountCategory  string `csv:"21"` // 算定回数関連
	Reserved7                 string `csv:"22"` // 予備
	Reserved8                 string `csv:"23"` // 予備
	Reserved9                 string `csv:"24"` // 予備
	Reserved10                string `csv:"25"` // 予備
	NewDate                   string `csv:"26"` // 新設年月日
	ExpiryDate                string `csv:"27"` // 廃止年月日
}

// MedicalPracticeInclusion 包括テーブル (02包括テーブル.csv)
type MedicalPracticeInclusion struct {
	ChangeCategory       string `csv:"1"` // 変更区分
	GroupNumber          string `csv:"2"` // グループ番号
	IncludedPracticeCode string `csv:"3"` // 診療行為コード
	IncludedPracticeName string `csv:"4"` // 診療行為省略名称
	SpecialCondition     string `csv:"5"` // 特例条件
	NewDate              string `csv:"6"` // 新設年月日
	ExpiryDate           string `csv:"7"` // 廃止年月日
}

// MedicalPracticeConflict 背反テーブル (03-*背反テーブル.csv)
type MedicalPracticeConflict struct {
	ChangeCategory       string `csv:"1"`  // 変更区分
	MedicalPracticeCode1 string `csv:"2"`  // 診療行為コード①
	PracticeName1        string `csv:"3"`  // 診療行為省略名称①
	MedicalPracticeCode2 string `csv:"4"`  // 診療行為コード②
	PracticeName2        string `csv:"5"`  // 診療行為省略名称②
	ConflictCategory     string `csv:"6"`  // 背反区分
	SpecialCondition     string `csv:"7"`  // 特例条件
	Reserved1            string `csv:"8"`  // 予備
	NewDate              string `csv:"9"`  // 新設年月日
	ExpiryDate           string `csv:"10"` // 廃止年月日
}

// InpatientBasicFee 入院基本料テーブル (04入院基本料テーブル.csv)
type InpatientBasicFee struct {
	ChangeCategory      string `csv:"1"` // 変更区分
	GroupNumber         string `csv:"2"` // グループ番号
	MedicalPracticeCode string `csv:"3"` // 診療行為コード
	PracticeName        string `csv:"4"` // 診療行為省略名称
	AdditionIdentifier  string `csv:"5"` // 加算識別
	Reserved1           string `csv:"6"` // 予備
	NewDate             string `csv:"7"` // 新設年月日
	ExpiryDate          string `csv:"8"` // 廃止年月日
}

// CalculationCount 算定回数テーブル (05算定回数テーブル.csv)
type CalculationCount struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	PracticeName        string `csv:"3"`  // 診療行為省略名称
	UnitCode            string `csv:"4"`  // 算定単位コード
	UnitName            string `csv:"5"`  // 算定単位名称
	CountLimit          int    `csv:"6"`  // 算定回数
	SpecialCondition    string `csv:"7"`  // 特例条件
	Reserved1           string `csv:"8"`  // 予備
	Reserved2           string `csv:"9"`  // 予備
	Reserved3           string `csv:"10"` // 予備
	Reserved4           string `csv:"11"` // 予備
	Reserved5           string `csv:"12"` // 予備
	NewDate             string `csv:"13"` // 新設年月日
	ExpiryDate          string `csv:"14"` // 廃止年月日
}
