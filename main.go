package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FG420/crypto-tracker/crypto"
)

type Coins struct {
	Names []string
}

func main() {
	var coins Coins
	input := ""
	fmt.Print("Insert a coin symbol: ")
	fmt.Scan(&input)

	jCoins, err := os.ReadFile("coins.json")
	if err != nil {
		log.Print("File not found")
		return
	}

	if err := json.Unmarshal(jCoins, &coins); err != nil {
		log.Printf("Error during unmarshal file for %s: %v", jCoins, err)
		return
	}

	// log.Print(append(coins.Names, "GUGU"))
	log.Print(coins.Names)

	coin := strings.ToUpper(input)
	crypto.CallApi(coin)
}
