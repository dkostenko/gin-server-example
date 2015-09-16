package main

import (
	"log"

	apipkg "github.com/dkostenko/gin-server-example/api"
)

func main() {
	router := apipkg.New()

	log.Printf("type: %T, value: %v, readable value: %#v", router, router, router)
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
