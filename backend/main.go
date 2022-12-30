package main

import (
	"fmt"
	"log"
	"os"

	"github.com/okoge-kaz/taschola/backend/db"
	"github.com/okoge-kaz/taschola/backend/router"
)

func main() {
	// initialize DB connection
	dsn := db.DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	if err := db.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	// initialize engine
	engine := router.Init()

	// start server
	const port = 8000
	engine.Run(fmt.Sprintf(":%d", port))
}
