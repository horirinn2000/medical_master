package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// --- 傷病名マスター API ---

// applyDiseaseFilters は共通のフィルタリングとページネーションを適用します
func applyDiseaseFilters(db *gorm.DB, categoryStr string, validOnly *api.ValidOnlyQuery, limit *api.LimitQuery, offset *api.OffsetQuery) *gorm.DB {
	query := db

	// カテゴリーフィルター
	if categoryStr != "" {
		switch categoryStr {
		case "medical":
			query = query.Where("adoption_category = ?", "1")
		case "dental":
			query = query.Where("adoption_category = ?", "2")
		}
	}

	// 有効期間フィルター
	if validOnly != nil && *validOnly {
		query = query.Where("discontinued_date = ? OR discontinued_date > ?", "99999999", "20260119") // 本日の日付を想定
	}

	// ページネーション
	l := 100
	if limit != nil {
		l = *limit
	}
	o := 0
	if offset != nil {
		o = *offset
	}

	return query.Limit(l).Offset(o)
}

// 傷病名マスター 全件取得API
func (s *ServerImpl) GetDiseases(c *gin.Context, params api.GetDiseasesParams) {
	var diseases []model.Disease
	categoryStr := ""
	if params.Category != nil {
		categoryStr = string(*params.Category)
	}
	query := applyDiseaseFilters(s.DB, categoryStr, params.ValidOnly, params.Limit, params.Offset)

	if err := query.Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargetsOptimized(&diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター 名称検索API
func (s *ServerImpl) GetDiseasesSearchName(c *gin.Context, params api.GetDiseasesSearchNameParams) {
	var diseases []model.Disease
	searchTerm := "%" + params.Q + "%"
	query := s.DB.Where("name_kanji LIKE ? OR name_abbr LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm, searchTerm)
	categoryStr := ""
	if params.Category != nil {
		categoryStr = string(*params.Category)
	}
	query = applyDiseaseFilters(query, categoryStr, params.ValidOnly, params.Limit, params.Offset)

	if err := query.Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargetsOptimized(&diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター コード検索API
func (s *ServerImpl) GetDiseasesSearchCode(c *gin.Context, params api.GetDiseasesSearchCodeParams) {
	var diseases []model.Disease
	query := s.DB.Where("code = ?", params.Q)
	if params.ValidOnly != nil && *params.ValidOnly {
		query = query.Where("discontinued_date = ? OR discontinued_date > ?", "99999999", "20260119")
	}

	if err := query.Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargetsOptimized(&diseases)
	c.JSON(http.StatusOK, diseases)
}

// 傷病名マスター ICD10検索API
func (s *ServerImpl) GetDiseasesSearchIcd10(c *gin.Context, params api.GetDiseasesSearchIcd10Params) {
	var diseases []model.Disease
	query := s.DB.Where("icd10_1 = ? OR icd10_2 = ?", params.Q, params.Q)
	categoryStr := ""
	if params.Category != nil {
		categoryStr = string(*params.Category)
	}
	query = applyDiseaseFilters(query, categoryStr, params.ValidOnly, params.Limit, params.Offset)

	if err := query.Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargetsOptimized(&diseases)
	c.JSON(http.StatusOK, diseases)
}

// attachMigrationTargetsOptimized は一括取得によりN+1問題を回避します
func (s *ServerImpl) attachMigrationTargetsOptimized(diseases *[]model.Disease) {
	var migrationMgmtNums []string
	for _, d := range *diseases {
		if d.IsOld && d.MigrationManagementNumber != "" && d.MigrationManagementNumber != "00000000" {
			migrationMgmtNums = append(migrationMgmtNums, d.MigrationManagementNumber)
		}
	}

	if len(migrationMgmtNums) == 0 {
		return
	}

	var targets []model.Disease
	if err := s.DB.Where("disease_management_number IN ? AND is_old = ?", migrationMgmtNums, false).Find(&targets).Error; err != nil {
		return
	}

	targetMap := make(map[string]*model.Disease)
	for i := range targets {
		targetMap[targets[i].DiseaseManagementNumber] = &targets[i]
	}

	for i := range *diseases {
		d := &(*diseases)[i]
		if target, ok := targetMap[d.MigrationManagementNumber]; ok {
			d.MigrationTarget = target
		}
	}
}
