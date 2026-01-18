package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 特定器材マスター API ---

// 特定器材マスター 全件取得API
func (s *ServerImpl) GetDevices(c *gin.Context) {
	var devices []model.SpecialMedicalDevice
	if err := s.DB.Limit(1000).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// 特定器材マスター 名称検索API
func (s *ServerImpl) GetDevicesSearchName(c *gin.Context, params api.GetDevicesSearchNameParams) {
	var devices []model.SpecialMedicalDevice
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// 特定器材マスター コード検索API
func (s *ServerImpl) GetDevicesSearchCode(c *gin.Context, params api.GetDevicesSearchCodeParams) {
	var devices []model.SpecialMedicalDevice
	if err := s.DB.Where("code = ?", params.Q).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}
