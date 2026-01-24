package main

import (
	"fmt"
	"github.com/fatih/color"
	"mempool-monitor/providers"
)

func main() {
	poolProvider := providers.GetLocalGethProvider("")
	err := poolProvider.Subscribe()
	if err != nil {
		panic(err)
	}
	for tx := range poolProvider.GetObserver() {
		fmt.Println("New tx:")
		color.Green("Hash: %s", tx.Hash)
		color.Green("From: %s", tx.From)
		color.Green("To: %s", tx.To)
	}
}
