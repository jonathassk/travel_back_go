package main

import (
	"github.com/jonathassk/travel_back_go/cmd/api"
	"github.com/jonathassk/travel_back_go/db"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dbRds, err := db.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewApiServer(":8080", dbRds)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
