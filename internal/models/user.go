package models

import (
	"database/sql"
	"errors"
	"sync"
	"time"
)

var mutex sync.Mutex

type User struct {
	Id       int
	Username string
	Password string
}

func (u *User) JoinIntoChat(chat *Chat, db *sql.DB) error {
	mutex.Lock()
	defer mutex.Unlock()
	if chat.Users[u] {
		return errors.New("you are already in the chat")
	}
	chat.Users[u] = true
	query := "INSERT INTO Chat_Users(user_id, chat_id) VALUES ($1, $2);"
	db.QueryRow(query, u.Id, chat.Id)
	return nil
}

func (u *User) SendMessage(m *Message, chat *Chat, db *sql.DB) error {
	m.SentOn = time.Now()
	chat.mutex.Lock()
	chat.Messages = append(chat.Messages, UserWithMessage{u: u, m: m})
	query := "INSERT INTO Messages(sender_id, msg, sent_on) VALUES ($1, $2, $3) RETURNING id;"
	if err := db.QueryRow(query, u.Id, m.Message, m.SentOn).Scan(&m.Id); err != nil {
		return err
	}
	query = "INSERT INTO Messages_Of_Chats(user_id, message_id) VALUES ($1, $2);"
	db.Query(query, u.Id, m.Id)
	chat.mutex.Unlock()
	return nil
}

func NewUser(username, password string, db *sql.DB) (u *User, e error) {
	if usernames[username] {
		return nil, errors.New("this username is already in use")
	}
	mutex.Lock()
	usernames[username] = true
	mutex.Unlock()
	query := "INSERT INTO Users(username, password) VALUES ($1, $2);"
	err := db.QueryRow(query, username, password).Scan(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) CreateChat(chatName string, db *sql.DB) (*Chat, error) {
	var chat Chat = Chat{
		CreatedBy: u,
		ChatName:  chatName,
	}
	query := "INSERT INTO Chats(chat_name, created_by) VALUES ($1, $2) RETURNING id;"
	if err := db.QueryRow(query, chat.ChatName, chat.CreatedBy.Id).Scan(&chat.Id); err != nil {
		return nil, err
	}
	return &chat, nil
}

// CREATE TABLE IF NOT EXISTS Messages (
//     id SERIAL PRIMARY KEY,
//     sender INTEGER REFERENCES Users(id),
//     msg TEXT,
//     sent_on TIMESTAMP
// );
