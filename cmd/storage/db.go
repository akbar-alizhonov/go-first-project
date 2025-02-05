package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres",fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPort, dbPassword, dbName))
	if err != nil {
		panic(err.Error())
	}
	
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Succesfully connected to db")
}

func GetDB() *sql.DB {
	return db
}