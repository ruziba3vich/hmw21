package services

import "github.com/ruziba3vich/hmw21/internal/models"

func LoadMessagesFromChatService(ch *models.Chat) []models.UserWithMessage {
	return ch.GetAllMessages()
}
