package model

import "time"

// VisitingNursingFee 訪問看護療養費マスター (ア 基本テーブル)
// csv/r/r_*.csv に対応
type VisitingNursingFee struct {
	ID                         uint    `gorm:"primaryKey"`
	ChangeCategory             string  `gorm:"size:1"`             // 1: 変更区分
	MasterType                 string  `gorm:"size:1"`             // 2: マスター種別 ('R'固定)
	VisitingNursingFeeCode     string  `gorm:"size:9"`             // 3: 訪問看護療養費コード
	NotificationSection        string  `gorm:"size:2"`             // 4: 区分番号
	NotificationBranch         string  `gorm:"size:1"`             // 5: 区分番号の枝番
	NotificationItem           string  `gorm:"size:2"`             // 6: 項番
	BasicName                  string  `gorm:"size:200"`           // 7: 基本名称
	AbbreviatedNameLength      int     `gorm:"type:smallint"`      // 8: 省略名称有効桁数
	AbbreviatedName            string  `gorm:"size:128"`           // 9: 省略名称
	AbbreviatedKanaNameLength  int     `gorm:"type:smallint"`      // 10: 省略カナ名称有効桁数
	AbbreviatedKanaName        string  `gorm:"size:20"`            // 11: 省略カナ名称
	DataStandardCode           string  `gorm:"size:3"`             // 12: データ規格コード
	UnitKanjiNameLength        int     `gorm:"type:smallint"`      // 13: 漢字名称有効桁数（単位）
	UnitKanjiName              string  `gorm:"size:12"`            // 14: 漢字名称（単位）
	PriceCategory              string  `gorm:"size:1"`             // 15: 金額識別
	Price                      float64 `gorm:"type:numeric(10,2)"` // 16: 新又は現金額
	OldPriceCategory           string  `gorm:"size:1"`             // 17: 金額識別（旧）
	OldPrice                   float64 `gorm:"type:numeric(10,2)"` // 18: 旧金額
	StepCalculationCategory    string  `gorm:"size:1"`             // 19: きざみ値計算識別
	StepLowerLimit             int64   `gorm:"type:bigint"`        // 20: 下限値
	StepUpperLimit             int64   `gorm:"type:bigint"`        // 21: 上限値
	StepValue                  int64   `gorm:"type:bigint"`        // 22: きざみ値
	StepPrice                  float64 `gorm:"type:numeric(10,2)"` // 23: きざみ金額
	StepErrorCategory          string  `gorm:"size:1"`             // 24: 上下限エラー処理
	LowerAge                   string  `gorm:"size:2"`             // 25: 下限年齢
	UpperAge                   string  `gorm:"size:2"`             // 26: 上限年齢
	ElderlyMedicalCategory     string  `gorm:"size:1"`             // 27: 後期高齢者医療適用区分
	MedicalObservationCategory string  `gorm:"size:1"`             // 28: 医療観察法対象区分
	JobCategory1               string  `gorm:"size:2"`             // 29: 職種等コード1
	JobCategory2               string  `gorm:"size:2"`             // 30: 職種等コード2
	JobCategory3               string  `gorm:"size:2"`             // 31: 職種等コード3
	JobCategory4               string  `gorm:"size:2"`             // 32: 職種等コード4
	JobCategory5               string  `gorm:"size:2"`             // 33: 職種等コード5
	JobCategory6               string  `gorm:"size:2"`             // 34: 職種等コード6
	JobCategory7               string  `gorm:"size:2"`             // 35: 職種等コード7
	JobCategory8               string  `gorm:"size:2"`             // 36: 職種等コード8
	JobCategory9               string  `gorm:"size:2"`             // 37: 職種等コード9
	JobCategory10              string  `gorm:"size:2"`             // 38: 職種等コード10
	JobCategory11              string  `gorm:"size:2"`             // 39: 職種等コード11
	JobCategory12              string  `gorm:"size:2"`             // 40: 職種等コード12
	JobCategory13              string  `gorm:"size:2"`             // 41: 職種等コード13
	JobCategory14              string  `gorm:"size:2"`             // 42: 職種等コード14
	JobCategory15              string  `gorm:"size:2"`             // 43: 職種等コード15
	VisitTimesCategory         string  `gorm:"size:1"`             // 44: 実施回数区分
	NursingInstructionCategory string  `gorm:"size:1"`             // 45: 訪問看護指示区分
	SpecialInstructionCategory string  `gorm:"size:1"`             // 46: 特別訪問看護指示区分
	SoloAdditionCategory       string  `gorm:"size:1"`             // 47: 加算単独算定区分
	AdditionGroup              string  `gorm:"size:4"`             // 48: 加算グループ
	FacilityStandardGroup      string  `gorm:"size:4"`             // 49: 施設基準グループ
	AdditionTableRelated       string  `gorm:"size:1"`             // 50: 基本・加算対応テーブル関連識別
	MaxTimesTableRelated       string  `gorm:"size:1"`             // 51: 算定回数限度テーブル関連識別
	ConflictTableRelated       string  `gorm:"size:1"`             // 52: 併算定背反テーブル関連識別
	ReceiptDisplaySection      string  `gorm:"size:2"`             // 53: レセプト表示欄
	ReceiptDisplayItem         string  `gorm:"size:2"`             // 54: レセプト表示項
	ReceiptDisplaySerial       string  `gorm:"size:3"`             // 55: レセプト表示連番
	ReceiptSymbol1             string  `gorm:"size:1"`             // 56: レセプト表示用記号1
	ReceiptSymbol2             string  `gorm:"size:1"`             // 57: レセプト表示用記号2
	ReceiptSymbol3             string  `gorm:"size:1"`             // 58: レセプト表示用記号3
	ReceiptSymbol4             string  `gorm:"size:1"`             // 59: レセプト表示用記号4
	ReceiptSymbol5             string  `gorm:"size:1"`             // 60: レセプト表示用記号5
	ReceiptSymbol6             string  `gorm:"size:1"`             // 61: レセプト表示用記号6
	ReceiptSymbol7             string  `gorm:"size:1"`             // 62: レセプト表示用記号7
	ReceiptSymbol8             string  `gorm:"size:1"`             // 63: レセプト表示用記号8
	ReceiptSymbol9             string  `gorm:"size:1"`             // 64: レセプト表示用記号9
	Reserved1                  string  `gorm:"size:1"`             // 65: 予備
	PublicationOrder           int     `gorm:"type:integer"`       // 66: 公表順序番号
	VisitingNursingType        string  `gorm:"size:2"`             // 67: 訪問看護療養費種類
	Reserved2                  string  `gorm:"size:2"`             // 68: 予備
	Reserved3                  string  `gorm:"size:2"`             // 69: 予備
	Reserved4                  string  `gorm:"size:2"`             // 70: 予備
	UpdateDate                 string  `gorm:"size:8"`             // 71: 変更年月日
	ExpiryDate                 string  `gorm:"size:8"`             // 72: 廃止年月日
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

// VisitingNursingAddition 訪問看護療養費基本・基本加算対応テーブル
// csv/r/r2_*.csv に対応
type VisitingNursingAddition struct {
	ID                     uint   `gorm:"primaryKey"`
	ChangeCategory         string `gorm:"size:1"`   // 1: 変更区分
	GroupNumber            string `gorm:"size:4"`   // 2: グループ番号
	VisitingNursingFeeCode string `gorm:"size:9"`   // 3: 訪問看護療養費コード（加算項目）
	AbbreviatedName        string `gorm:"size:128"` // 4: 省略名称
	AdditionIdentifier     string `gorm:"size:2"`   // 5: 加算識別
	UpdateDate             string `gorm:"size:8"`   // 6: 変更年月日
	ExpiryDate             string `gorm:"size:8"`   // 7: 廃止年月日
	Reserved1              string `gorm:"size:3"`   // 8: 予備
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

// VisitingNursingCalculationCount 訪問看護療養費算定回数限度テーブル
// csv/r/r3_*.csv に対応
type VisitingNursingCalculationCount struct {
	ID                     uint   `gorm:"primaryKey"`
	ChangeCategory         string `gorm:"size:1"`        // 1: 変更区分
	VisitingNursingFeeCode string `gorm:"size:9"`        // 2: 訪問看護療養費コード
	AbbreviatedName        string `gorm:"size:128"`      // 3: 省略名称
	UnitCode               string `gorm:"size:3"`        // 4: 算定単位
	UnitName               string `gorm:"size:12"`       // 5: 算定単位名称
	MaxTimes               int    `gorm:"type:smallint"` // 6: 上限回数
	MaxTimesErrorCategory  string `gorm:"size:1"`        // 7: 上限回数エラー処理
	UpdateDate             string `gorm:"size:8"`        // 8: 変更年月日
	ExpiryDate             string `gorm:"size:8"`        // 9: 廃止年月日
	Reserved1              string `gorm:"size:3"`        // 10: 予備
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

// VisitingNursingConflict 訪問看護療養費併算定背反テーブル
// csv/r/r4_*.csv に対応
type VisitingNursingConflict struct {
	ID                     uint   `gorm:"primaryKey"`
	ChangeCategory         string `gorm:"size:1"`   // 1: 変更区分
	VisitingNursingFeeCode string `gorm:"size:9"`   // 2: 訪問看護療養費コード1
	AbbreviatedName1       string `gorm:"size:128"` // 3: 省略名称1
	ConflictCategory       string `gorm:"size:1"`   // 4: 背反区分
	ConflictFeeCode        string `gorm:"size:9"`   // 5: 訪問看護療養費コード2
	AbbreviatedName2       string `gorm:"size:128"` // 6: 省略名称2
	ConflictUnit           string `gorm:"size:1"`   // 7: 背反単位
	SpecialCondition       string `gorm:"size:1"`   // 8: 特例条件
	UpdateDate             string `gorm:"size:8"`   // 9: 変更年月日
	ExpiryDate             string `gorm:"size:8"`   // 10: 廃止年月日
	Reserved1              string `gorm:"size:3"`   // 11: 予備
	Reserved2              string `gorm:"size:3"`   // 12: 予備
	Reserved3              string `gorm:"size:3"`   // 13: 予備
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

// VisitingNursingFacilityStandard 訪問看護療養費施設基準テーブル
// csv/r/r5_*.csv に対応
type VisitingNursingFacilityStandard struct {
	ID                     uint   `gorm:"primaryKey"`
	ChangeCategory         string `gorm:"size:1"`   // 1: 変更区分
	GroupNumber            string `gorm:"size:4"`   // 2: グループ番号
	VisitingNursingFeeCode string `gorm:"size:9"`   // 3: 訪問看護療養費コード
	AbbreviatedName        string `gorm:"size:128"` // 4: 省略名称
	FacilityStandardCode   string `gorm:"size:4"`   // 5: 施設基準コード
	FacilityIdentifier     string `gorm:"size:2"`   // 6: 施設基準識別
	UpdateDate             string `gorm:"size:8"`   // 7: 変更年月日
	ExpiryDate             string `gorm:"size:8"`   // 8: 廃止年月日
	Reserved1              string `gorm:"size:3"`   // 9: 予備
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
