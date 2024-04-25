package services

import (
	"database/sql"

	"github.com/ruziba3vich/hmw21/internal/models"
)

func SendMessage(u *models.User, ch *models.Chat, db *sql.DB, m *models.Message) error {
	return u.SendMessage(m, ch, db)
}


