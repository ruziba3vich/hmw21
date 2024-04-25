package services

import (
	"database/sql"

	"github.com/ruziba3vich/hmw21/internal/models"
)

func JoinIntoChatService(u *models.User, ch *models.Chat, db *sql.DB) error {
	return u.JoinIntoChat(ch, db)
}
