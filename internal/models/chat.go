package models

import "sync"

type Chat struct {
	Id        int
	ChatName  string
	CreatedBy *User
	Users     map[*User]bool
	mutex     sync.Mutex
	Messages  []UserWithMessage
}

type UserWithMessage struct {
	u *User
	m *Message
}

func (ch *Chat) GetAllMessages() []UserWithMessage {
	return ch.Messages
}

/*
CREATE TABLE IF NOT EXISTS Chats (
    id SERIAL PRIMARY KEY,
    chat_name VARCHAR(64),
    created_by INTEGER REFERENCES Users(id)
);
*/
