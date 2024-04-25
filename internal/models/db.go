package models

import (
	"database/sql"
	"log"
	"time"
)

var usernames map[string]bool = map[string]bool{}
var chats map[string]*Chat = map[string]*Chat{}
var messages map[string]Message = map[string]Message{}
var users map[int]*User = map[int]*User{}

func UpdateUsernames(db *sql.DB) {
	query := "SELECT username FROM Users;"
	db.QueryRow(query).Scan(&usernames)
}

func UpdateChats(db *sql.DB) {
	query := "SELECT id, chat_name, created_by FROM Chats;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	chats = make(map[string]*Chat)

	for rows.Next() {
		var id int
		var chatName string
		var createdBy int
		if err := rows.Scan(&id, &chatName, &createdBy); err != nil {
			log.Fatal(err)
		}
		chat := &Chat{
			Id:        id,
			ChatName:  chatName,
			CreatedBy: users[createdBy],
		}
		chats[chatName] = chat
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func UpdateMessages(db *sql.DB) error {
	query := "SELECT chat_name, id, sender_id, message, sent_on FROM Messages JOIN Chats ON Messages.chat_id = Chats.id;"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var chatName string
		var id int
		var senderID int
		var messageText string
		var sentOn time.Time
		if err := rows.Scan(&chatName, &id, &senderID, &messageText, &sentOn); err != nil {
			return err
		}
		message := Message{
			Id:      id,
			Sender:  *users[senderID],
			Message: messageText,
			SentOn:  sentOn,
		}
		messages[chatName] = message
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
