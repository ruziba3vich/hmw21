package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-shafaq/gin"
	"github.com/ruziba3vich/hmw21/internal/models"
	"github.com/ruziba3vich/hmw21/internal/services"
)

func JoinChat(c *gin.Context, db *sql.DB) {
	var jchr models.JoinChatRequest
	err := services.JoinIntoChatService(&jchr.User, &jchr.Chat, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, nil)
}
