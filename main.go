package main

import (
	"log"
	"github.com/thebigyovadiaz/api-twitter/handlers"
	"github.com/thebigyovadiaz/api-twitter/db"
)

func main() {
	if db.CheckingConnection() == 0 {
		log.Fatal("Not Connect to DB")
		return
	}

	handlers.Handlers()
}
