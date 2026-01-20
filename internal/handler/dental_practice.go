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

// 歯科加算関係検索API (h-2, h-3, h-4, h-5)
func (s *ServerImpl) GetDentalPracticesCodeAdditions(c *gin.Context, code string) {
	var additions []model.DentalPracticeAdditionRelation
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&additions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, additions)
}

// 歯科算定回数制限マスター検索API (h-6)
func (s *ServerImpl) GetDentalPracticesCodeCalculationCountsMaster(c *gin.Context, code string) {
	var counts []model.DentalPracticeCalculationCount
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}

// 歯科きざみ検索API (h-7)
func (s *ServerImpl) GetDentalPracticesCodeSteps(c *gin.Context, code string) {
	var steps []model.DentalPracticeStep
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&steps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, steps)
}

// 歯科年齢制限検索API (h-8)
func (s *ServerImpl) GetDentalPracticesCodeAgeConstraints(c *gin.Context, code string) {
	var constraints []model.DentalPracticeAgeConstraint
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&constraints).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, constraints)
}

// 歯科併算定背反検索API (h-9)
func (s *ServerImpl) GetDentalPracticesCodeConflictsMaster(c *gin.Context, code string) {
	var conflicts []model.DentalPracticeConflict
	if err := s.DB.Where("source_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

// 歯科実日数検索API (h-10)
func (s *ServerImpl) GetDentalPracticesCodeActualDays(c *gin.Context, code string) {
	var actualDays []model.DentalPracticeActualDays
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&actualDays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actualDays)
}
