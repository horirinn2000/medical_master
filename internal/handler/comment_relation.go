package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- コメント関連テーブル API ---

// コメント関連テーブル 検索API
func (s *ServerImpl) GetCommentsRelated(c *gin.Context, params api.GetCommentsRelatedParams) {
	var relations []model.CommentRelation
	query := s.DB.Preload("Comment")

	if params.ActCode != nil && *params.ActCode != "" {
		query = query.Where("act_code = ?", *params.ActCode)
	}
	if params.CommentCode != nil && *params.CommentCode != "" {
		query = query.Where("comment_code = ?", *params.CommentCode)
	}
	if params.Section != nil && *params.Section != "" {
		query = query.Where("section = ?", *params.Section)
	}

	if err := query.Find(&relations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relations)
}
