package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/hmw21/internal/models"
	"github.com/ruziba3vich/hmw21/internal/services"
)

func SendMessage(c *gin.Context, db *sql.DB) {
	var uwm models.UserWithMessage
	c.ShouldBindJSON(&uwm)

	wrr := services.SendMessage(&uwm.User.id, &ch, db)
}
