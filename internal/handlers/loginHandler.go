package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/hmw21/internal/models"
	"github.com/ruziba3vich/hmw21/internal/services"
)

func LogInHandler(c *gin.Context, db *sql.DB) {
	var user models.User
	c.ShouldBindJSON(&user)
	newUser, err := services.LogInService(db, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}
