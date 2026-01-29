package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"mempool-monitor/providers"
	"mempool-monitor/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	color.Yellow("Starting mempool monitor...")
	poolProvider := providers.GetLocalGethProvider("")
	err = poolProvider.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to provider: %v", err)
	}
	color.Green("Subscribed to provider successfully!")

	for tx := range poolProvider.GetObserver() {
		go func(tx *types.Transaction) {
			fmt.Println("New tx:")
			color.Green("Hash: %s", tx.Hash)
			color.Green("From: %s", tx.From)
			color.Green("To: %s", tx.To)
		}(tx)

	}
}
