package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload" // Load .env file automatically
	"github.com/orlmonteverde/go-postgres-microblog/internal/data"
	"github.com/orlmonteverde/go-postgres-microblog/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	//connect to the database
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Panic(err)
	}

	// Start the server
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
	data.Close()
}
