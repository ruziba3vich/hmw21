package models

type SendMessageRequest struct {
	User    User
	Chat    Chat
	Message Message
}
