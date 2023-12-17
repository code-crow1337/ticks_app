package data

import (
	"database/sql"
	"e-identitet/tick-api/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	gotdotenvError := godotenv.Load() // Load Env file
	CheckError := utils.CheckError
	CheckError(gotdotenvError)

	logger := log.Default()
	logger.Printf("Connecting to the database....")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	userDb := os.Getenv("DB_USER")
	passwordDb := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userDb, passwordDb, databaseName)

	//creates a connection to the database
	db, err := sql.Open("postgres", connectionString)
	CheckError(err)

	//check connection to the database
	err = db.Ping()
	CheckError(err)
	return db
}
