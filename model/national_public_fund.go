package model

import (
	"gorm.io/gorm"
)

// NationalPublicFund 国公費マスタ
// doc/kouhi/kokkouhi_koumoku_240329.pdf に基づく定義
type NationalPublicFund struct {
	gorm.Model
	PriorityOrder                int     `gorm:"index;comment:1:公費適用順位"`
	LawNumber                    string  `gorm:"index;comment:2:法別番号"`
	IssuingUnit                  string  `gorm:"comment:3-1:実施機関番号の付番単位"`
	IssuingOrganizationNumber    string  `gorm:"comment:3-2:国が指定する実施機関番号"`
	Department                   string  `gorm:"comment:4:担当部署"`
	SystemStartDate              string  `gorm:"comment:5-1:制度開始日"`
	ProjectStartDate             string  `gorm:"comment:5-2:事業開始日"`
	EndDate                      string  `gorm:"comment:6:制度/事業終了予定日"`
	LegalBasis                   string  `gorm:"comment:7:根拠法律"`
	ProjectName                  string  `gorm:"comment:8:給付事業名"`
	HasTargetDisease             bool    `gorm:"comment:9-1:対象疾病の有無"`
	InpatientOnly                bool    `gorm:"comment:9-2:入院のみ"`
	OutpatientOnly               bool    `gorm:"comment:9-3:外来のみ"`
	BothInpatientOutpatient      bool    `gorm:"comment:9-4:入外両方"`
	TargetDiseaseNote            string  `gorm:"comment:9-5:備考(対象疾病)"`
	FoodExpensesTarget           bool    `gorm:"comment:9-6:食事療養費の給付対象区分"`
	LivingExpensesTarget         bool    `gorm:"comment:9-7:生活療養費の給付対象区分"`
	ExpensesNote                 string  `gorm:"comment:9-8:備考(療養費)"`
	HasOtherLawPriority          bool    `gorm:"comment:9-9:他法優先規定の有無"`
	HasLocalFundPriority         bool    `gorm:"comment:9-10:地単公費優先適用の有無"`
	PriorityNote                 string  `gorm:"comment:9-11:備考(優先規定)"`
	InsurancePublicFundPriority  int     `gorm:"comment:10-1:保険・公費の優先度(0:保険優先, 1:公費優先)"`
	PersonalBurdenRatio          float64 `gorm:"comment:10-2:本人負担割合"`
	InsuranceRatio3              float64 `gorm:"comment:11-1:保険負担割合(3割時)"`
	PublicFundRatio3             float64 `gorm:"comment:11-2:公費負担割合(3割時)"`
	PersonalBurdenRatio3         float64 `gorm:"comment:11-3:本人負担割合(3割時)"`
	InsuranceRatio2              float64 `gorm:"comment:11-4:保険負担割合(2割時)"`
	PublicFundRatio2             float64 `gorm:"comment:11-5:公費負担割合(2割時)"`
	PersonalBurdenRatio2         float64 `gorm:"comment:11-6:本人負担割合(2割時)"`
	InsuranceRatio1              float64 `gorm:"comment:11-7:保険負担割合(1割時)"`
	PublicFundRatio1             float64 `gorm:"comment:11-8:公費負担割合(1割時)"`
	PersonalBurdenRatio1         float64 `gorm:"comment:11-9:本人負担割合(1割時)"`
	ApplicableToWelfareRecipient bool    `gorm:"comment:12-1:生保受給者への適用有無"`
	MedicalAssistanceCombination bool    `gorm:"comment:12-2:医療扶助との併用"`
	PublicFundRatioWelfare       float64 `gorm:"comment:12-3:公費負担割合(生保併用時)"`
	MedicalAssistanceRatio       float64 `gorm:"comment:12-4:医療扶助負担割合"`
	UninsuredApplicability       string  `gorm:"comment:13:無保険者への適用可否及び割合"`
	HasIncomeCap                 bool    `gorm:"comment:14-1:所得額による上限の有無"`
	IncomeCapCondition           string  `gorm:"comment:14-2:上限額・自己負担額の条件"`
	HasDailyCalculation          bool    `gorm:"comment:14-3:日割計算の有無"`
	DailyCalculationMethod       string  `gorm:"comment:14-4:日割計算の計算方法"`
	CapCalculationNote           string  `gorm:"comment:14-5:上限額計算上の留意点"`
	OtherCapCondition            string  `gorm:"comment:14-6:所得以外の条件による上限"`
	HasCapManagementBook         bool    `gorm:"comment:15-1:上限額管理票の有無"`
	CapManagementBookType        string  `gorm:"comment:15-2:上限額管理票の媒体種類・提示方法"`
	CapManagementBookNote        string  `gorm:"comment:15-3:特記事項(上限額管理票)"`
	HasCountManagementBook       bool    `gorm:"comment:16-1:上限回数又は該当回数管理票の有無"`
	CountManagementBookType      string  `gorm:"comment:16-2:上限回数又は該当回数管理票の媒体種類・提示方法"`
	CountManagementBookNote      string  `gorm:"comment:16-3:特記事項(上限回数管理票)"`
	CanBenefitsInKind            bool    `gorm:"comment:17-1:現物給付の可否"`
	HasPartialInKindProhibition  bool    `gorm:"comment:17-2:一部不可のケースの有無"`
	PartialInKindProhibitionNote string  `gorm:"comment:17-3:一部不可の内容"`
	HighMedicalExpensePriority   bool    `gorm:"comment:18-1:高額療養費における公費優先"`
	SpecificBenefitTarget        bool    `gorm:"comment:18-2:特定給付対象療養"`
	SpecificDiseaseBenefitTarget bool    `gorm:"comment:18-3:特定疾病給付対象療養"`
	NormalHighMedicalExpense     bool    `gorm:"comment:18-4:通常の高額療養費負担"`
}
