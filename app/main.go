package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-shafaq/defcase"
	gg "github.com/go-shafaq/gin"
	"github.com/ruziba3vich/hmw21/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	dcase := defcase.Get()
	dcase.SetCase("json", "*", defcase.Snak_case)
	dcase.SetCase("form", "*", defcase.Snak_case)

	gg.SetDefCase(dcase)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "userRequests")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		PrintError(err)
	}
	defer db.Close()

	dbNames := []string{"users", "admins"}

	for _, dbName := range dbNames {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		if err != nil {
			PrintError(err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			PrintError(err)
		}
	}

	router.POST("/login", func(c *gin.Context) {
		handlers.LogInHandler(c, db)
	})

	router.POST("/join-chat/id", func(c *gin.Context) {
		handlers.JoinChat(c, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	log.Fatal("Error :", err)
}
