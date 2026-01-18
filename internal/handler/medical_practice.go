package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 医科診療行為マスター API ---

// 医科診療行為マスター 全件取得API
func (s *ServerImpl) GetMedicalPractices(c *gin.Context) {
	var practices []model.MedicalPractice
	if err := s.DB.Limit(1000).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

// 医科診療行為マスター 名称検索API
func (s *ServerImpl) GetMedicalPracticesSearchName(c *gin.Context, params api.GetMedicalPracticesSearchNameParams) {
	var practices []model.MedicalPractice
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("abbreviated_kanji_name LIKE ? OR abbreviated_kana_name LIKE ? OR basic_kanji_name LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

// 医科診療行為マスター コード検索API
func (s *ServerImpl) GetMedicalPracticesSearchCode(c *gin.Context, params api.GetMedicalPracticesSearchCodeParams) {
	var practices []model.MedicalPractice
	if err := s.DB.Where("medical_practice_code = ?", params.Q).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

// 医科診療行為マスター 包含関係取得API
func (s *ServerImpl) GetMedicalPracticesCodeInclusions(c *gin.Context, code string) {
	var inclusions []model.MedicalPracticeInclusion
	if err := s.DB.Where("comprehensive_practice_code = ?", code).Find(&inclusions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inclusions)
}

// 医科診療行為マスター 背反関係取得API
func (s *ServerImpl) GetMedicalPracticesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.MedicalPracticeConflict
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

// 医科診療行為マスター 補助関係取得API
func (s *ServerImpl) GetMedicalPracticesCodeSupports(c *gin.Context, code string) {
	var supports []model.MedicalPracticeSupport
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&supports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supports)
}

// 医科診療行為マスター 入院基本料関係検索API
func (s *ServerImpl) GetMedicalPracticesCodeInpatientFees(c *gin.Context, code string) {
	var fees []model.InpatientBasicFee
	if err := s.DB.Where("basic_fee_code = ?", code).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fees)
}

// 医科診療行為マスター 算定回数制限検索API
func (s *ServerImpl) GetMedicalPracticesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.CalculationCount
	if err := s.DB.Where("medical_practice_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}
