package main

import (
	"e-identitet/tick-api/routes"
	"log"
)

//liveness route
//liveness route with test
//get ticks for a company
//find specfic tic sms, bankid-auth, bankid-sign, SPAR

func main() {
	logger := log.Default()
	logger.Print("Server is up an running 🌱")
	r := routes.SetupRouter()
	r.Run(":80")
}
