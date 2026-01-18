package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// --- 傷病名マスター API ---

// 傷病名マスター 全件取得API
func (s *ServerImpl) GetDiseases(c *gin.Context) {
	var diseases []model.Disease
	if err := s.DB.Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	attachMigrationTargets(s.DB, &diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター 名称検索API
func (s *ServerImpl) GetDiseasesSearchName(c *gin.Context, params api.GetDiseasesSearchNameParams) {
	var diseases []model.Disease
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name_kanji LIKE ? OR name_abbr LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	attachMigrationTargets(s.DB, &diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター コード検索API
func (s *ServerImpl) GetDiseasesSearchCode(c *gin.Context, params api.GetDiseasesSearchCodeParams) {
	var diseases []model.Disease
	if err := s.DB.Where("code = ?", params.Q).Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	attachMigrationTargets(s.DB, &diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター ICD10検索API
func (s *ServerImpl) GetDiseasesSearchIcd10(c *gin.Context, params api.GetDiseasesSearchIcd10Params) {
	var diseases []model.Disease
	if err := s.DB.Where("icd10_1 = ? OR icd10_2 = ?", params.Q, params.Q).
		Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	attachMigrationTargets(s.DB, &diseases)
	c.JSON(http.StatusOK, diseases)
}

// 旧傷病名の新傷病名取得
func attachMigrationTargets(db *gorm.DB, diseases *[]model.Disease) {
	for i := range *diseases {
		d := &(*diseases)[i]
		if d.IsOld && d.MigrationManagementNumber != "" && d.MigrationManagementNumber != "00000000" {
			var target model.Disease
			if err := db.Where("disease_management_number = ? AND is_old = ?", d.MigrationManagementNumber, false).First(&target).Error; err == nil {
				d.MigrationTarget = &target
			}
		}
	}
}
