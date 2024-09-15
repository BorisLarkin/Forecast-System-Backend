package main

import (
	"log"
	"web/internal/api"
)

func main() {
	log.Println("Application started!")
	api.StartServer()
	log.Println("Application terminated!")
}
