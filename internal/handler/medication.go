package handler

import (
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 調剤行為マスター API ---

// 手動定義（NG）
type GetMedicationsSearchNameParams struct {
	Q string `form:"q" json:"q"`
}
type GetMedicationsSearchCodeParams struct {
	Q string `form:"q" json:"q"`
}

// 調剤行為マスター 全件取得API
func (s *ServerImpl) GetMedications(c *gin.Context) {
	var medications []model.Medication
	if err := s.DB.Limit(1000).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

// 調剤行為マスター 名称検索API
func (s *ServerImpl) GetMedicationsSearchName(c *gin.Context, params GetMedicationsSearchNameParams) {
	var medications []model.Medication
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name_kanji LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm).
		Limit(100).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

// 調剤行為マスター コード検索API
func (s *ServerImpl) GetMedicationsSearchCode(c *gin.Context, params GetMedicationsSearchCodeParams) {
	var medications []model.Medication
	if err := s.DB.Where("code = ?", params.Q).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}
