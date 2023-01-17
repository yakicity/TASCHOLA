package main

import (
	"fmt"
	"log"
	"os"

	"taschola/db"
	"taschola/router"

)

//	@title			TASCHOLA API
//	@version		1.0
//	@description	TASCHOLA API Server
//	@license		MIT
//	@host			localhost:8000
//	@BasePath		/
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
		log.Fatal("Failed to connect to DB")
	}

	// initialize engine
	engine := router.Init()

	// start server
	const port = 8000
	engine.Run(fmt.Sprintf(":%d", port))
}
