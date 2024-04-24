package models

import "time"

type Message struct {
	Id      int
	Sender  User
	Message string
	SentOn  time.Time
}
