package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/hmw21/internal/models"
	"github.com/ruziba3vich/hmw21/internal/services"
)

func SendMessage(c *gin.Context, db *sql.DB) {
	var uwm models.SendMessageRequest
	c.ShouldBindJSON(&uwm)

	err := services.SendMessage(&uwm.User, &uwm.Chat, db, &uwm.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}
