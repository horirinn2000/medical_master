package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 歯科診療行為マスター API ---

// 歯科診療行為マスター 全件取得API
func (s *ServerImpl) GetDentalPracticesSearchName(c *gin.Context, params api.GetDentalPracticesSearchNameParams) {
	var practices []model.DentalPractice
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("abbreviated_kanji_name LIKE ?", searchTerm).
		Limit(100).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

// 歯科診療行為マスター コード検索API
func (s *ServerImpl) GetDentalPracticesSearchCode(c *gin.Context, params api.GetDentalPracticesSearchCodeParams) {
	var practices []model.DentalPractice
	if err := s.DB.Where("medical_practice_code = ?", params.Q).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

// 歯科診療行為マスター 包括関係検索API
func (s *ServerImpl) GetDentalPracticesCodeInclusions(c *gin.Context, code string) {
	var inclusions []model.DentalPracticeInclusion
	if err := s.DB.Where("comprehensive_practice_code = ?", code).Find(&inclusions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inclusions)
}

// 歯科診療行為マスター 背反関係検索API
func (s *ServerImpl) GetDentalPracticesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.DentalPracticeConflictDetail
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

// 歯科診療行為マスター 補助関係検索API
func (s *ServerImpl) GetDentalPracticesCodeSupports(c *gin.Context, code string) {
	var supports []model.DentalPracticeSupport
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&supports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supports)
}

// 歯科診療行為マスター 算定回数制限検索API
func (s *ServerImpl) GetDentalPracticesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.DentalPracticeCalculationCountLimit
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}
