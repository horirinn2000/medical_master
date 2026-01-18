package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// --- 調剤行為マスター API ---

// 調剤行為マスター 全件取得API
func (s *ServerImpl) GetMedications(c *gin.Context, params api.GetMedicationsParams) {
	var medications []model.Medication
	query := s.DB

	// 有効期間フィルタリング
	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("discontinued_date >= ?", today)
	}

	// ページネーション
	limit := 100
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	if err := query.Limit(limit).Offset(offset).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

// 調剤行為マスター 名称検索API
func (s *ServerImpl) GetMedicationsSearchName(c *gin.Context, params api.GetMedicationsSearchNameParams) {
	var medications []model.Medication
	query := s.DB

	searchTerm := "%" + params.Q + "%"
	query = query.Where("name_kanji LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm)

	// 有効期間フィルタリング
	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("discontinued_date >= ?", today)
	}

	// ページネーション
	limit := 100
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	if err := query.Limit(limit).Offset(offset).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

// 調剤行為マスター コード検索API
func (s *ServerImpl) GetMedicationsSearchCode(c *gin.Context, params api.GetMedicationsSearchCodeParams) {
	var medications []model.Medication
	query := s.DB

	query = query.Where("code = ?", params.Q)

	// 有効期間フィルタリング
	if params.ValidOnly != nil && *params.ValidOnly {
		today := time.Now().Format("20060102")
		query = query.Where("discontinued_date >= ?", today)
	}

	if err := query.Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}
