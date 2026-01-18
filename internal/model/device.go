package model

// SpecialMedicalDevice 特定器材マスター (t_*.csv)
// 全38項目を網羅 (R06rec3.pdf 198-199ページに準拠)
type SpecialMedicalDevice struct {
	ChangeCategory               string  `csv:"1" json:"change_category"`                 // 1: 変更区分
	MasterType                   string  `csv:"2" json:"master_type"`                     // 2: マスター種別
	Code                         string  `csv:"3" json:"code"`                            // 3: 特定器材コード
	KanjiNameLength              int     `csv:"4" json:"kanji_name_length"`               // 4: 漢字有効桁数 (特定器材名・規格名)
	KanjiName                    string  `csv:"5" json:"kanji_name"`                      // 5: 漢字名称 (特定器材名・規格名)
	KanaNameLength               int     `csv:"6" json:"kana_name_length"`                // 6: カナ有効桁数 (特定器材名・規格名)
	KanaName                     string  `csv:"7" json:"kana_name"`                       // 7: カナ名称 (特定器材名・規格名)
	UnitCode                     string  `csv:"8" json:"unit_code"`                       // 8: コード (単位)
	UnitKanjiNameLength          int     `csv:"9" json:"unit_kanji_name_length"`          // 9: 漢字有効桁数 (単位)
	UnitKanjiName                string  `csv:"10" json:"unit_kanji_name"`                // 10: 漢字名称 (単位)
	PriceType                    string  `csv:"11" json:"price_type"`                     // 11: 金額種別 (新又は現金額)
	Price                        float64 `csv:"12" json:"price"`                          // 12: 新又は現金額
	Reserved13                   string  `csv:"13" json:"reserved_13"`                    // 13: 予備
	AgeAdditionCategory          string  `csv:"14" json:"age_addition_category"`          // 14: 年齢加算区分
	MinAge                       string  `csv:"15" json:"min_age"`                        // 15: 下限年齢 (上下限年齢)
	MaxAge                       string  `csv:"16" json:"max_age"`                        // 16: 上限年齢 (上下限年齢)
	OldPriceType                 string  `csv:"17" json:"old_price_type"`                 // 17: 金額種別 (旧金額)
	OldPrice                     float64 `csv:"18" json:"old_price"`                      // 18: 旧金額
	KanjiNameUpdateFlag          string  `csv:"19" json:"kanji_name_update_flag"`         // 19: 漢字名称変更区分
	KanaNameUpdateFlag           string  `csv:"20" json:"kana_name_update_flag"`          // 20: カナ名称変更区分
	OxygenCategory               string  `csv:"21" json:"oxygen_category"`                // 21: 酸素等区分
	DeviceCategory               string  `csv:"22" json:"device_category"`                // 22: 特定器材種別
	MaxPriceFlag                 string  `csv:"23" json:"max_price_flag"`                 // 23: 上限価格
	MaxPoints                    int     `csv:"24" json:"max_points"`                     // 24: 上限点数
	Reserved25                   string  `csv:"25" json:"reserved_25"`                    // 25: 予備
	PublicationOrder             int     `csv:"26" json:"publication_order"`              // 26: 公表順序番号
	RelatedCode                  string  `csv:"27" json:"related_code"`                   // 27: 廃止・新設関連
	UpdateDate                   string  `csv:"28" json:"update_date"`                    // 28: 変更年月日
	TransitionalDate             string  `csv:"29" json:"transitional_date"`              // 29: 経過措置年月日
	DiscontinuedDate             string  `csv:"30" json:"discontinued_date"`              // 30: 廃止年月日
	AppendixNumber               string  `csv:"31" json:"appendix_number"`                // 31: 別表番号 (告示番号)
	SectionNumber                string  `csv:"32" json:"section_number"`                 // 32: 区分番号 (告示番号)
	DPCApplicabilityCategory     string  `csv:"33" json:"dpc_applicability_category"`     // 33: DPC適用区分
	Reserved34                   string  `csv:"34" json:"reserved_34"`                    // 34: 予備
	Reserved35                   string  `csv:"35" json:"reserved_35"`                    // 35: 予備
	Reserved36                   string  `csv:"36" json:"reserved_36"`                    // 36: 予備
	BasicKanjiName               string  `csv:"37" json:"basic_kanji_name"`               // 37: 基本漢字名称
	RemanufacturedDeviceCategory string  `csv:"38" json:"remanufactured_device_category"` // 38: 再製造単回使用医療機器
}
