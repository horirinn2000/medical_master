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
	query := s.DB.Model(&model.Tooth{})

	// 名称検索
	searchTerm := "%" + params.Q + "%"
	query = query.Where("name LIKE ?", searchTerm)

	// 部位検索
	if params.Site != nil {
		switch *params.Site {
		case api.Upper:
			// 1001xx (上顎全体), 1003xx (右上顎), 1004xx (上顎前歯), 1005xx (左上顎), 101xxx (右側上顎), 102xxx (左側上顎)
			// コード体系に基づき、1001, 1003, 1004, 1005, 101, 102 で始まるものを上顎とする
			query = query.Where("code LIKE '1001%' OR code LIKE '1003%' OR code LIKE '1004%' OR code LIKE '1005%' OR code LIKE '101%' OR code LIKE '102%'")
		case api.Lower:
			// 1002xx (下顎全体), 1006xx (左下顎), 1007xx (下顎前歯), 1008xx (右下顎), 103xxx (左側下顎), 104xxx (右側下顎)
			// コード体系に基づき、1002, 1006, 1007, 1008, 103, 104 で始まるものを下顎とする
			query = query.Where("code LIKE '1002%' OR code LIKE '1006%' OR code LIKE '1007%' OR code LIKE '1008%' OR code LIKE '103%' OR code LIKE '104%'")
		}
	}

	// 有効なもののみ
	if params.ValidOnly != nil && *params.ValidOnly {
		query = query.Where("discontinued_date = '99999999'")
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

	var teeth []model.Tooth
	if err := query.Limit(limit).Offset(offset).Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}

// 歯式マスター コード検索API
func (s *ServerImpl) GetTeethSearchCode(c *gin.Context, params api.GetTeethSearchCodeParams) {
	query := s.DB.Model(&model.Tooth{}).Where("code = ?", params.Q)

	// 有効なもののみ
	if params.ValidOnly != nil && *params.ValidOnly {
		query = query.Where("discontinued_date = '99999999'")
	}

	var teeth []model.Tooth
	if err := query.Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}
