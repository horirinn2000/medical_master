package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 歯式マスター API ---

// 歯式マスター 全件取得API
func (s *ServerImpl) GetTeeth(c *gin.Context) {
	var teeth []model.Tooth
	if err := s.DB.Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}

// 歯式マスター 名称検索API
func (s *ServerImpl) GetTeethSearchName(c *gin.Context, params api.GetTeethSearchNameParams) {
	var teeth []model.Tooth
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name LIKE ?", searchTerm).Limit(100).Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}

// 歯式マスター コード検索API
func (s *ServerImpl) GetTeethSearchCode(c *gin.Context, params api.GetTeethSearchCodeParams) {
	var teeth []model.Tooth
	if err := s.DB.Where("code = ?", params.Q).Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}
