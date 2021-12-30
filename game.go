package main

import (
	"game/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type counter struct {
	v int
}

var serverAddr = "localhost:9898"

func main() {
	log.Println("Starting")

	// server.ServerRun(serverAddr, action.ActionHandler)
	server.ServerRun(serverAddr)
	log.Println("Done!")
}
