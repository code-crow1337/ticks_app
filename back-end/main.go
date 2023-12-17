package main

import (
	"e-identitet/tick-api/routes"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	logger := log.Default()
	r := routes.SetupRouter()
	logger.Print("Server is up an running 🌱")
	r.Run(":80")

}
