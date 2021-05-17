package main

import (
	"log"
	"os"
	"ports-parser-service/api"
)

func main() {
	log.Println("Ports parser service up and listening on 8080")
	api.NewPortParserService(os.Getenv("PORT_MANAGER_ADDR")).Run(":8080")
}
