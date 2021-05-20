package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"ports-parser-service/api"
)

var (
	port = flag.Int("port", 8080, "the ports parser service port")
)

func main() {
	flag.Parse()
	log.Printf("Ports parser service up and listening on %d", *port)
	api.NewPortParserService(os.Getenv("PORT_MANAGER_ADDR")).Run(fmt.Sprintf(":%d", *port))
}
