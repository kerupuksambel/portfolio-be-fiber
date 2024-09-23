package main

import (
	"flag"
	"fmt"

	"github.com/kerupuksambel/portfolio-be/app"
	"github.com/kerupuksambel/portfolio-be/internal/database"
)

func main() {
	// Args read
	var (
		isMigration bool
	)
	flag.BoolVar(&isMigration, "migrate", false, "Include if you want to run migration (not forcing migration)")
	flag.Parse()

	// Initialize DB
	database.InitDB()

	// Initiate migration
	if isMigration {
		database.Migrate()
	}

	server := app.Init()

	server.Listen(":6060")
	fmt.Printf("Listening on port 6060")
}
