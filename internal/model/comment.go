package model

// Comment コメントマスター (c_*.csv)
// 全30項目を網羅 (R06rec3.pdf 200ページに準拠)
type Comment struct {
	ChangeCategory    string `csv:"1" json:"change_category"`     // 1: 変更区分
	MasterType        string `csv:"2" json:"master_type"`         // 2: マスター種別 ('C'固定)
	Category          string `csv:"3" json:"category"`            // 3: 区分 ('8'固定)
	Pattern           string `csv:"4" json:"pattern"`             // 4: パターン
	SerialNo          string `csv:"5" json:"serial_no"`           // 5: 一連番号
	KanjiNameLength   int    `csv:"6" json:"kanji_name_length"`   // 6: 漢字有効桁数
	KanjiName         string `csv:"7" json:"kanji_name"`          // 7: 漢字名称
	KanaNameLength    int    `csv:"8" json:"kana_name_length"`    // 8: カナ有効桁数
	KanaName          string `csv:"9" json:"kana_name"`           // 9: カナ名称
	EditInfo1Pos      int    `csv:"10" json:"edit_info_1_pos"`    // 10: ①カラム位置
	EditInfo1Len      int    `csv:"11" json:"edit_info_1_len"`    // 11: ①桁数
	EditInfo2Pos      int    `csv:"12" json:"edit_info_2_pos"`    // 12: ②カラム位置
	EditInfo2Len      int    `csv:"13" json:"edit_info_2_len"`    // 13: ②桁数
	EditInfo3Pos      int    `csv:"14" json:"edit_info_3_pos"`    // 14: ③カラム位置
	EditInfo3Len      int    `csv:"15" json:"edit_info_3_len"`    // 15: ③桁数
	EditInfo4Pos      int    `csv:"16" json:"edit_info_4_pos"`    // 16: ④カラム位置
	EditInfo4Len      int    `csv:"17" json:"edit_info_4_len"`    // 17: ④桁数
	Reserved18        string `csv:"18" json:"reserved_18"`        // 18: 予備
	Reserved19        string `csv:"19" json:"reserved_19"`        // 19: 予備
	SelectionCategory string `csv:"20" json:"selection_category"` // 20: 選択式コメント識別
	UpdateDate        string `csv:"21" json:"update_date"`        // 21: 変更年月日
	ExpiryDate        string `csv:"22" json:"expiry_date"`        // 22: 廃止年月日
	Code              string `csv:"23" json:"code"`               // 23: コメントコード
	PublicationOrder  int    `csv:"24" json:"publication_order"`  // 24: 公表順序番号
	Reserved25        string `csv:"25" json:"reserved_25"`        // 25: 予備
	Reserved26        string `csv:"26" json:"reserved_26"`        // 26: 予備
	Reserved27        string `csv:"27" json:"reserved_27"`        // 27: 予備
	Reserved28        string `csv:"28" json:"reserved_28"`        // 28: 予備
	Reserved29        string `csv:"29" json:"reserved_29"`        // 29: 予備
	Reserved30        string `csv:"30" json:"reserved_30"`        // 30: 予備
}
