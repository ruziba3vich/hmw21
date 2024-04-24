package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "userRequests")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		PrintError(err)
	}
	defer db.Close()

	databases := []string{"users", "admins"}

	for _, dbName := range databases {
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

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := router.Run(address); err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	log.Fatal("Error :", err)
}
