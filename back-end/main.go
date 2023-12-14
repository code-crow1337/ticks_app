package main

import (
	"e-identitet/tick-api/routes"
<<<<<<< HEAD
	"log"
=======
	"fmt"
>>>>>>> be1aaa1 (initial commit for the back-end endpoints)
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
