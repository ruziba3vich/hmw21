package models

var usernames map[string]bool = map[string]bool{}

type User struct {
	Username string
	Password string
}

func NewUser(username, password string) *User {
	usernames[username] = true
	return &User{
		Username: username,
		Password: password,
	}
}
