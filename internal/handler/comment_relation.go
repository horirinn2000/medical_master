package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- コメント関連テーブル API ---

// コメント関連テーブル ActCode検索API
func (s *ServerImpl) GetCommentsRelated(c *gin.Context, params api.GetCommentsRelatedParams) {
	var relations []model.CommentRelation
	if err := s.DB.Where("act_code = ?", params.ActCode).Find(&relations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relations)
}
