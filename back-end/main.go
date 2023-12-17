package main

import (
	"database/sql"
	"e-identitet/tick-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

//liveness route
//liveness route with test
//get ticks for a company
//find specfic tic sms, bankid-auth, bankid-sign, SPAR

func main() {
	logger := log.Default()
	r := routes.SetupRouter()
	logger.Print("Server is up an running 🌱")
	initDatabase()
	r.Run(":80")

}
func initDatabase() {
	gotdotenvError := godotenv.Load() // Load Env file
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

	defer db.Close()

	//check connection to the database
	err = db.Ping()
	CheckError(err)

	logger.Printf("Whop Whop! Database connection established ⭐️")

	companyTable, companyTableErr := db.Query("SELECT * FROM comapny;")
	CheckError(companyTableErr)
	logger.Print("📖", companyTable)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
