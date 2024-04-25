package services

import (
	"database/sql"

	"github.com/ruziba3vich/hmw21/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func LogInService(db *sql.DB, user models.User) (u *models.User, e error) {
	query := "SELECT COUNT(*) FROM your_table WHERE username = $1;"
	var count int
	db.QueryRow(query, user.Username).Scan(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		err = db.QueryRow("INSERT INTO Users(username, pwd) VALUES ($1, $2) RETURNING id", user.Username, hashedPassword).Scan(&user.Id)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
