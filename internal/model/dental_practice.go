package model

// DentalPractice 歯科診療行為マスター (h_*.csv)
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
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	Column3             string `csv:"3"`  // 予備
	Column4             string `csv:"4"`  // 予備
	Column5             string `csv:"5"`  // 予備
	Column6             string `csv:"6"`  // 予備
	Column7             string `csv:"7"`  // 予備
	Name1               string `csv:"8"`  // 名称1
	Name2               string `csv:"9"`  // 名称2
	CountLimit          int    `csv:"10"` // 限度回数
	Column11            string `csv:"11"` // 予備
	Column12            string `csv:"12"` // 予備
	UpdateDate          string `csv:"13"` // 変更年月日
	ExpiryDate          string `csv:"14"` // 廃止年月日
	Column15            string `csv:"15"` // 予備
}

// DentalPracticeStep きざみテーブル (h-7)
type DentalPracticeStep struct {
	ChangeCategory      string  `csv:"1"`  // 変更区分
	MedicalPracticeCode string  `csv:"2"`  // 診療行為コード
	Column3             string  `csv:"3"`  // 予備
	Column4             string  `csv:"4"`  // 予備
	Column5             string  `csv:"5"`  // 予備
	Column6             string  `csv:"6"`  // 予備
	Column7             string  `csv:"7"`  // 予備
	Name1               string  `csv:"8"`  // 名称1
	Name2               string  `csv:"9"`  // 名称2
	Column10            string  `csv:"10"` // 予備
	StepLimit           float64 `csv:"11"` // 制限値
	Column12            string  `csv:"12"` // 予備
	Column13            string  `csv:"13"` // 予備
	Column14            string  `csv:"14"` // 予備
	Column15            string  `csv:"15"` // 予備
	StepScore           float64 `csv:"16"` // 加算点数
	Column17            string  `csv:"17"` // 予備
	UpdateDate          string  `csv:"18"` // 変更年月日
	ExpiryDate          string  `csv:"19"` // 廃止年月日
	Column20            string  `csv:"20"` // 予備
}

// DentalPracticeAgeConstraint 年齢制限テーブル (h-8)
type DentalPracticeAgeConstraint struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	Column3             string `csv:"3"`  // 予備
	Column4             string `csv:"4"`  // 予備
	Column5             string `csv:"5"`  // 予備
	Column6             string `csv:"6"`  // 予備
	Column7             string `csv:"7"`  // 予備
	Name1               string `csv:"8"`  // 名称1
	Name2               string `csv:"9"`  // 名称2
	LowerAge            string `csv:"10"` // 下限年齢
	UpperAge            string `csv:"11"` // 上限年齢
	UpdateDate          string `csv:"12"` // 変更年月日
	ExpiryDate          string `csv:"13"` // 廃止年月日
	Column14            string `csv:"14"` // 予備
}

// DentalPracticeConflict 併算定背反テーブル (h-9)
type DentalPracticeConflict struct {
	ChangeCategory string `csv:"1"`   // 変更区分
	SourceCode     string `csv:"2"`   // 診療行為コード(元)
	SourceType     string `csv:"3"`   // 種類(元)
	SourceSection  string `csv:"4"`   // 区分(元)
	Column5        string `csv:"5"`   // 予備
	Column6        string `csv:"6"`   // 予備
	Column7        string `csv:"7"`   // 予備
	SourceName1    string `csv:"8"`   // 名称(元)1
	SourceName2    string `csv:"9"`   // 名称(元)2
	ConflictFlag   string `csv:"10"`  // フラグ
	TargetCode     string `csv:"11"`  // 診療行為コード(先)
	TargetType     string `csv:"12"`  // 種類(先)
	TargetSection  string `csv:"13"`  // 区分(先)
	Column14       string `csv:"14"`  // 予備
	Column15       string `csv:"15"`  // 予備
	Column16       string `csv:"16"`  // 予備
	TargetName1    string `csv:"17"`  // 名称(先)1
	TargetName2    string `csv:"18"`  // 名称(先)2
	Column19       string `csv:"19"`  // 予備
	Column20       string `csv:"20"`  // 予備
	Column21       string `csv:"21"`  // 予備
	Column22       string `csv:"22"`  // 予備
	Column23       string `csv:"23"`  // 予備
	Column24       string `csv:"24"`  // 予備
	Column25       string `csv:"25"`  // 予備
	Column26       string `csv:"26"`  // 予備
	Column27       string `csv:"27"`  // 予備
	Column28       string `csv:"28"`  // 予備
	Column29       string `csv:"29"`  // 予備
	Column30       string `csv:"30"`  // 予備
	Column31       string `csv:"31"`  // 予備
	Column32       string `csv:"32"`  // 予備
	Column33       string `csv:"33"`  // 予備
	Column34       string `csv:"34"`  // 予備
	Column35       string `csv:"35"`  // 予備
	Column36       string `csv:"36"`  // 予備
	Column37       string `csv:"37"`  // 予備
	Column38       string `csv:"38"`  // 予備
	Column39       string `csv:"39"`  // 予備
	Column40       string `csv:"40"`  // 予備
	Column41       string `csv:"41"`  // 予備
	Column42       string `csv:"42"`  // 予備
	Column43       string `csv:"43"`  // 予備
	Column44       string `csv:"44"`  // 予備
	Column45       string `csv:"45"`  // 予備
	Column46       string `csv:"46"`  // 予備
	Column47       string `csv:"47"`  // 予備
	Column48       string `csv:"48"`  // 予備
	Column49       string `csv:"49"`  // 予備
	Column50       string `csv:"50"`  // 予備
	Column51       string `csv:"51"`  // 予備
	Column52       string `csv:"52"`  // 予備
	Column53       string `csv:"53"`  // 予備
	Column54       string `csv:"54"`  // 予備
	Column55       string `csv:"55"`  // 予備
	Column56       string `csv:"56"`  // 予備
	Column57       string `csv:"57"`  // 予備
	Column58       string `csv:"58"`  // 予備
	Column59       string `csv:"59"`  // 予備
	Column60       string `csv:"60"`  // 予備
	Column61       string `csv:"61"`  // 予備
	Column62       string `csv:"62"`  // 予備
	Column63       string `csv:"63"`  // 予備
	Column64       string `csv:"64"`  // 予備
	Column65       string `csv:"65"`  // 予備
	Column66       string `csv:"66"`  // 予備
	Column67       string `csv:"67"`  // 予備
	Column68       string `csv:"68"`  // 予備
	Column69       string `csv:"69"`  // 予備
	Column70       string `csv:"70"`  // 予備
	Column71       string `csv:"71"`  // 予備
	Column72       string `csv:"72"`  // 予備
	Column73       string `csv:"73"`  // 予備
	Column74       string `csv:"74"`  // 予備
	Column75       string `csv:"75"`  // 予備
	Column76       string `csv:"76"`  // 予備
	Column77       string `csv:"77"`  // 予備
	Column78       string `csv:"78"`  // 予備
	Column79       string `csv:"79"`  // 予備
	Column80       string `csv:"80"`  // 予備
	Column81       string `csv:"81"`  // 予備
	Column82       string `csv:"82"`  // 予備
	Column83       string `csv:"83"`  // 予備
	Column84       string `csv:"84"`  // 予備
	Column85       string `csv:"85"`  // 予備
	Column86       string `csv:"86"`  // 予備
	Column87       string `csv:"87"`  // 予備
	Column88       string `csv:"88"`  // 予備
	Column89       string `csv:"89"`  // 予備
	Column90       string `csv:"90"`  // 予備
	Column91       string `csv:"91"`  // 予備
	Column92       string `csv:"92"`  // 予備
	Column93       string `csv:"93"`  // 予備
	Column94       string `csv:"94"`  // 予備
	Column95       string `csv:"95"`  // 予備
	Column96       string `csv:"96"`  // 予備
	Column97       string `csv:"97"`  // 予備
	Column98       string `csv:"98"`  // 予備
	Column99       string `csv:"99"`  // 予備
	UpdateDate     string `csv:"100"` // 変更年月日
	ExpiryDate     string `csv:"101"` // 廃止年月日
	Column102      string `csv:"102"` // 予備
	Column103      string `csv:"103"` // 予備
	Column104      string `csv:"104"` // 予備
}

// DentalPracticeActualDays 実日数関連テーブル (h-10)
type DentalPracticeActualDays struct {
	ChangeCategory      string `csv:"1"`  // 変更区分
	MedicalPracticeCode string `csv:"2"`  // 診療行為コード
	Column3             string `csv:"3"`  // 予備
	Column4             string `csv:"4"`  // 予備
	Column5             string `csv:"5"`  // 予備
	Column6             string `csv:"6"`  // 予備
	Column7             string `csv:"7"`  // 予備
	Name1               string `csv:"8"`  // 名称1
	Name2               string `csv:"9"`  // 名称2
	DaysLimit           string `csv:"10"` // 日数制限
	Column11            string `csv:"11"` // 予備
	UpdateDate          string `csv:"12"` // 変更年月日
	ExpiryDate          string `csv:"13"` // 廃止年月日
	Column14            string `csv:"14"` // 予備
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
