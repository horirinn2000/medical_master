package model

// DentalPractice 歯科診療行為マスター (h_*.csv)
// 医科(MedicalPractice)と類似しているが、項目数が異なるため個別に定義
type DentalPractice struct {
	ChangeCategory                string  `csv:"1"`  // 変更区分
	MasterType                    string  `csv:"2"`  // マスター種別
	MedicalPracticeCode           string  `csv:"3"`  // 診療行為コード
	AbbreviatedKanjiNameLength    int     `csv:"4"`  // 省略漢字有効桁数
	AbbreviatedKanjiName          string  `csv:"5"`  // 省略漢字名称
	AbbreviatedKanaNameLength     int     `csv:"6"`  // 省略カナ有効桁数
	AbbreviatedKanaName           string  `csv:"7"`  // 省略カナ名称
	DataStandardCode              string  `csv:"8"`  // データ規格コード
	KanjiUnitNameLength           int     `csv:"9"`  // 漢字名称有効桁数（単位）
	KanjiUnitName                 string  `csv:"10"` // 漢字名称（単位）
	ScoreType                     string  `csv:"11"` // 点数識別
	Score                         float64 `csv:"12"` // 新又は現点数
	InpatientOutpatientCategory   string  `csv:"13"` // 入外適用区分
	ElderlyMedicalCategory        string  `csv:"14"` // 後期高齢者医療適用区分
	AggregationCategoryOutpatient string  `csv:"15"` // 点数欄集計先識別（入院外）
	ComprehensiveTargetTest       string  `csv:"16"` // 包括対象検査
	DPCApplicabilityCategory      string  `csv:"18"` // DPC適用区分
	HospitalClinicCategory        string  `csv:"19"` // 病院・診療所区分
	SurgerySupportCategory        string  `csv:"20"` // 画像等手術支援加算
	MedicalObservationCategory    string  `csv:"21"` // 医療観察法対象区分
	NursingAddition               string  `csv:"22"` // 看護加算
	AnesthesiaCategory            string  `csv:"23"` // 麻酔識別区分
	DiseaseRelatedCategory        string  `csv:"25"` // 傷病名関連区分
	ActualDaysCategory            string  `csv:"27"` // 実日数
	DaysTimesCategory             string  `csv:"28"` // 日数・回数
	MedicineRelatedCategory       string  `csv:"29"` // 医薬品関連区分
	StepCalculationCategory       string  `csv:"30"` // きざみ値計算識別
	StepLowerLimit                float64 `csv:"31"` // 下限値
	StepUpperLimit                float64 `csv:"32"` // 上限値
	StepValue                     float64 `csv:"33"` // きざみ値
	StepScore                     float64 `csv:"34"` // きざみ点数
	StepErrorCategory             string  `csv:"35"` // 上下限エラー処理
	MaxTimes                      int     `csv:"36"` // 上限回数
	MaxTimesErrorCategory         string  `csv:"37"` // 上限回数エラー処理
	NoteAdditionCode              string  `csv:"38"` // 注加算コード
	NoteAdditionSequence          string  `csv:"39"` // 注加算通番
	GeneralAgeCategory            string  `csv:"40"` // 通則年齢
	LowerAge                      string  `csv:"41"` // 下限年齢
	UpperAge                      string  `csv:"42"` // 上限年齢
	TimeAdditionCategory          string  `csv:"43"` // 時間加算区分
	FacilityStandardCategory      string  `csv:"44"` // 適合区分
	FacilityStandardCode          string  `csv:"45"` // 対象施設基準
	TreatmentInfantAddition       string  `csv:"46"` // 処置乳幼児加算区分
	VeryLowBirthWeightAddition    string  `csv:"47"` // 極低出生体重児加算区分
	InpatientReductionCategory    string  `csv:"48"` // 入院基本料等減算対象識別
	DonorCollectionCategory       string  `csv:"49"` // ドナー分集計区分
	TestJudgmentCategory          string  `csv:"50"` // 検査等実施判断区分
	TestJudgmentGroupCategory     string  `csv:"51"` // 検査等実施判断グループ区分
	ReductionApplicability        string  `csv:"52"` // 逓減対象区分
	SpinalEvokedPotentialCategory string  `csv:"53"` // 脊髄誘発電位測定等加算区分
	NeckDissectionCategory        string  `csv:"54"` // 頸部郭清術併施加算区分
	AutoSutureCategory            string  `csv:"55"` // 自動縫合器加算区分
	OutpatientManagementAddition  string  `csv:"56"` // 外来管理加算区分
	UpdateDate                    string  `csv:"57"` // 変更年月日
	ExpiryDate                    string  `csv:"58"` // 廃止年月日
	PublicationOrder              string  `csv:"59"` // 公表順序番号
	SectionChapter                string  `csv:"60"` // 章
	SectionPart                   string  `csv:"61"` // 部
	SectionNumber                 string  `csv:"62"` // 区分番号
	SectionSubNumber              string  `csv:"63"` // 枝番
	SectionItemNumber             string  `csv:"64"` // 項番
	AggregationCategoryInpatient  string  `csv:"66"` // 点数欄集計先識別（入院）
}

// DentalPracticeAdditionRelation 加算対応テーブル (h-2, h-3, h-4, h-5)
type DentalPracticeAdditionRelation struct {
	ChangeCategory       string `csv:"1"` // 変更区分
	SectionNumber        string `csv:"2"` // 区分番号
	AdditionCode         string `csv:"3"` // 加算コード
	MedicalPracticeCode  string `csv:"4"` // 診療行為コード
	AbbreviatedKanjiName string `csv:"5"` // 省略漢字名称
	AdditionCategory     string `csv:"6"` // 加算識別
	UpdateDate           string `csv:"7"` // 変更年月日
	ExpiryDate           string `csv:"8"` // 廃止年月日
	SearchPriority       string `csv:"9"` // 検索優先順位
}

// DentalPracticeCalculationCount 算定回数限度テーブル (h-6)
type DentalPracticeCalculationCount struct {
	ChangeCategory       string `csv:"1"` // 変更区分
	MedicalPracticeCode  string `csv:"2"` // 診療行為コード
	AbbreviatedKanjiName string `csv:"3"` // 省略漢字名称
	CalculationCategory  string `csv:"4"` // 算定区分
	CountLimit           int    `csv:"5"` // 限度回数
	UpdateDate           string `csv:"6"` // 変更年月日
	ExpiryDate           string `csv:"7"` // 廃止年月日
}

// DentalPracticeStep きざみテーブル (h-7)
type DentalPracticeStep struct {
	ChangeCategory       string  `csv:"1"` // 変更区分
	MedicalPracticeCode  string  `csv:"2"` // 診療行為コード
	AbbreviatedKanjiName string  `csv:"3"` // 省略漢字名称
	StepLowerLimit       float64 `csv:"4"` // 下限値
	StepUpperLimit       float64 `csv:"5"` // 上限値
	StepValue            float64 `csv:"6"` // きざみ値
	StepScore            float64 `csv:"7"` // 加算点数
	UpdateDate           string  `csv:"8"` // 変更年月日
	ExpiryDate           string  `csv:"9"` // 廃止年月日
}

// DentalPracticeAgeConstraint 年齢制限テーブル (h-8)
type DentalPracticeAgeConstraint struct {
	ChangeCategory       string `csv:"1"` // 変更区分
	MedicalPracticeCode  string `csv:"2"` // 診療行為コード
	AbbreviatedKanjiName string `csv:"3"` // 省略漢字名称
	LowerAge             int    `csv:"4"` // 下限年齢
	UpperAge             int    `csv:"5"` // 上限年齢
	UpdateDate           string `csv:"6"` // 変更年月日
	ExpiryDate           string `csv:"7"` // 廃止年月日
}

// DentalPracticeConflict 併算定背反テーブル (h-9)
type DentalPracticeConflict struct {
	ChangeCategory      string `csv:"1"` // 変更区分
	MedicalPracticeCode string `csv:"2"` // 診療行為コード
	ConflictCode        string `csv:"3"` // 併算定不可コード
	UpdateDate          string `csv:"4"` // 変更年月日
	ExpiryDate          string `csv:"5"` // 廃止年月日
}

// DentalPracticeActualDays 実日数関連テーブル (h-10)
type DentalPracticeActualDays struct {
	ChangeCategory       string `csv:"1"` // 変更区分
	MedicalPracticeCode  string `csv:"2"` // 診療行為コード
	AbbreviatedKanjiName string `csv:"3"` // 省略漢字名称
	UpdateDate           string `csv:"4"` // 変更年月日
	ExpiryDate           string `csv:"5"` // 廃止年月日
}

// DentalPracticeSupport 補助マスターテーブル (01補助マスターテーブル（歯科）.csv)
type DentalPracticeSupport struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	SupportCode         string `csv:"3"`  // 補助コード
	SupportInfo         string `csv:"4"`  // 補助情報
	UpdateDate          string `csv:"24"` // 変更年月日
	ExpiryDate          string `csv:"25"` // 廃止年月日
}

// DentalPracticeInclusion 包括テーブル (02包括テーブル（歯科）.csv)
type DentalPracticeInclusion struct {
	ChangeCategory            string `csv:"1"` // 変更区分
	ComprehensivePracticeCode string `csv:"2"` // 包括診療行為コード
	IncludedPracticeCode      string `csv:"3"` // 包括対象診療行為コード
	UpdateDate                string `csv:"7"` // 変更年月日
	ExpiryDate                string `csv:"8"` // 廃止年月日
}

// DentalPracticeConflictDetail 背反テーブル (03-*背反テーブル*（歯科）.csv)
type DentalPracticeConflictDetail struct {
	ChangeCategory       string `csv:"1"`  // 変更区分
	MedicalPracticeCode  string `csv:"2"`  // 診療行為コード
	ConflictPracticeCode string `csv:"4"`  // 背反診療行為コード
	UpdateDate           string `csv:"11"` // 変更年月日
	ExpiryDate           string `csv:"12"` // 廃止年月日
}

// DentalPracticeCalculationCountLimit 算定回数テーブル (04算定回数テーブル（歯科）.csv)
type DentalPracticeCalculationCountLimit struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	CountLimit          int    `csv:"8"`  // 算定回数限度
	UpdateDate          string `csv:"11"` // 変更年月日
	ExpiryDate          string `csv:"12"` // 廃止年月日
}
