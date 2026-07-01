package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatalf("Error connecting to AMQP: %v", err)
	}
	defer connection.Close()

	fmt.Println("Connection to AMQP successful")

	// wait for close signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Shutdown signal received, program shutting down")
}
