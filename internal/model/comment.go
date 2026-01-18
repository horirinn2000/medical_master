package model

import (
	"time"
)

// Comment コメントマスターモデル
// doc/R06rec1.pdf 200ページ（ファイルレイアウト：コメントマスター）に基づく
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UpdateCategory      string `gorm:"size:1" json:"update_category"`    // 1: 変更区分 (数字 1)
	MasterType          string `gorm:"size:1" json:"master_type"`        // 2: マスター種別 (英数 1) - "C" 固定
	Category            string `gorm:"size:1" json:"category"`           // 3: 区分 (数字 1) - "8" 固定
	Pattern             string `gorm:"size:2" json:"pattern"`            // 4: パターン (数字 2)
	SerialNo            string `gorm:"size:6" json:"serial_no"`          // 5: 一連番号 (数字 6)
	NameKanjiLen        int    `json:"name_kanji_len"`                   // 6: 漢字有効桁数 (数字 3)
	NameKanji           string `gorm:"size:300;index" json:"name_kanji"` // 7: 漢字名称 (漢字 300)
	NameKanaLen         int    `json:"name_kana_len"`                    // 8: カナ有効桁数 (数字 2)
	NameKana            string `gorm:"size:20;index" json:"name_kana"`   // 9: カナ名称 (英数カナ 20)
	ReceiptEditInfo1Pos int    `json:"receipt_edit_info_1_pos"`          // 10: レセプト編集情報①カラム位置
	ReceiptEditInfo1Len int    `json:"receipt_edit_info_1_len"`          // 11: レセプト編集情報①桁数
	ReceiptEditInfo2Pos int    `json:"receipt_edit_info_2_pos"`          // 12: レセプト編集情報②カラム位置
	ReceiptEditInfo2Len int    `json:"receipt_edit_info_2_len"`          // 13: レセプト編集情報②桁数
	ReceiptEditInfo3Pos int    `json:"receipt_edit_info_3_pos"`          // 14: レセプト編集情報③カラム位置
	ReceiptEditInfo3Len int    `json:"receipt_edit_info_3_len"`          // 15: レセプト編集情報③桁数
	ReceiptEditInfo4Pos int    `json:"receipt_edit_info_4_pos"`          // 16: レセプト編集情報④カラム位置
	ReceiptEditInfo4Len int    `json:"receipt_edit_info_4_len"`          // 17: レセプト編集情報④桁数
	Reserved1           string `gorm:"size:1" json:"reserved_1"`         // 18: 予備 (数字 1)
	Reserved2           string `gorm:"size:1" json:"reserved_2"`         // 19: 予備 (数字 1)
	SelectionCategory   string `gorm:"size:1" json:"selection_category"` // 20: 選択式コメント識別 (数字 1)
	UpdateDate          string `gorm:"size:8" json:"update_date"`        // 21: 変更年月日 (数字 8)
	DiscontinuedDate    string `gorm:"size:8" json:"discontinued_date"`  // 22: 廃止年月日 (数字 8)
	Code                string `gorm:"size:9;index" json:"code"`         // 23: コメントコード (数字 9)
	PublishOrder        int    `json:"publish_order"`                    // 24: 公表順序番号 (数字 9)
	Reserved3           string `gorm:"size:1" json:"reserved_3"`         // 25: 予備 (数字 1)
	Reserved4           string `gorm:"size:1" json:"reserved_4"`         // 26: 予備 (数字 1)
	Reserved5           string `gorm:"size:1" json:"reserved_5"`         // 27: 予備 (数字 1)
	Reserved6           string `gorm:"size:1" json:"reserved_6"`         // 28: 予備 (数字 1)
	Reserved7           string `gorm:"size:1" json:"reserved_7"`         // 29: 予備 (数字 1)
	Reserved8           string `gorm:"size:1" json:"reserved_8"`         // 30: 予備 (数字 1)
}
