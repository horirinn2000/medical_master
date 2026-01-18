package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// --- 特定器材マスター API ---

// 特定器材マスター 検索API
func (s *ServerImpl) GetDevices(c *gin.Context, params api.GetDevicesParams) {
	query := s.DB.Model(&model.SpecialMedicalDevice{})

	if params.Name != nil && *params.Name != "" {
		searchTerm := "%" + *params.Name + "%"
		query = query.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	if params.Code != nil && *params.Code != "" {
		query = query.Where("code = ?", *params.Code)
	}

	if params.DeviceCategory != nil && *params.DeviceCategory != "" {
		query = query.Where("device_category = ?", *params.DeviceCategory)
	}

	if params.OxygenCategory != nil && *params.OxygenCategory != "" {
		query = query.Where("oxygen_category = ?", *params.OxygenCategory)
	}

	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("update_date <= ? AND (discontinued_date = '99999999' OR discontinued_date >= ?)", today, today)
	}

	limit := 100
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	var devices []model.SpecialMedicalDevice
	if err := query.Limit(limit).Offset(offset).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// 特定器材マスター 名称検索API
func (s *ServerImpl) GetDevicesSearchName(c *gin.Context, params api.GetDevicesSearchNameParams) {
	query := s.DB.Model(&model.SpecialMedicalDevice{})

	searchTerm := "%" + params.Q + "%"
	query = query.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm)

	if params.DeviceCategory != nil && *params.DeviceCategory != "" {
		query = query.Where("device_category = ?", *params.DeviceCategory)
	}

	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("update_date <= ? AND (discontinued_date = '99999999' OR discontinued_date >= ?)", today, today)
	}

	limit := 100
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	var devices []model.SpecialMedicalDevice
	if err := query.Limit(limit).Offset(offset).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// 特定器材マスター コード検索API
func (s *ServerImpl) GetDevicesSearchCode(c *gin.Context, params api.GetDevicesSearchCodeParams) {
	query := s.DB.Model(&model.SpecialMedicalDevice{})

	query = query.Where("code = ?", params.Q)

	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("update_date <= ? AND (discontinued_date = '99999999' OR discontinued_date >= ?)", today, today)
	}

	var devices []model.SpecialMedicalDevice
	if err := query.Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}
