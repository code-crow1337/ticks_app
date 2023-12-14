package main

import (
	"e-identitet/tick-api/routes"
	"fmt"
)

//liveness route
//liveness route with test
//get ticks for a company
//find specfic tic sms, bankid-auth, bankid-sign, SPAR

func main() {
	fmt.Println("Main running")
	//liveness route
	r := routes.SetupRouter()
	r.Run(":8080")
}
