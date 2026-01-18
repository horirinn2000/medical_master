package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 病棟マスター API ---

// 病棟マスター 全件取得API
func (s *ServerImpl) GetWards(c *gin.Context) {
	var wards []model.Ward
	if err := s.DB.Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

// 病棟マスター 名称検索API
func (s *ServerImpl) GetWardsSearchName(c *gin.Context, params api.GetWardsSearchNameParams) {
	var wards []model.Ward
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(1000).Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

// 病棟マスター コード検索API
func (s *ServerImpl) GetWardsSearchCode(c *gin.Context, params api.GetWardsSearchCodeParams) {
	var wards []model.Ward
	if err := s.DB.Where("code = ?", params.Q).Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}
