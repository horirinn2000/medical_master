package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- コメントマスター API ---

// コメントマスター 全件取得API
func (s *ServerImpl) GetComments(c *gin.Context) {
	var comments []model.Comment
	if err := s.DB.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// コメントマスター 名称検索API
func (s *ServerImpl) GetCommentsSearchName(c *gin.Context, params api.GetCommentsSearchNameParams) {
	var comments []model.Comment
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("name_kanji LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm).
		Limit(100).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// コメントマスター コード検索API
func (s *ServerImpl) GetCommentsSearchCode(c *gin.Context, params api.GetCommentsSearchCodeParams) {
	var comments []model.Comment
	if err := s.DB.Where("code = ?", params.Q).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}
