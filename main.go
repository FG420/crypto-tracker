package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/FG420/crypto-tracker/crypto"
	"github.com/FG420/crypto-tracker/types"
)

func main() {
	var listCoins types.Coins
	input := ""
	fmt.Print("Insert a coin symbol: ")
	fmt.Scan(&input)

	readJson, err := os.ReadFile("coins.json")
	if err != nil {
		log.Print("File not found")
		return
	}

	if err := json.Unmarshal(readJson, &listCoins); err != nil {
		log.Printf("Error during unmarshal file for %s: %v", readJson, err)
		return
	}

	coinName := strings.ToUpper(input)
	crypto.GetData(coinName)

	if slices.Contains(listCoins.Names, coinName) {
		log.Print("This Coin is already inside the coins.json file")
		log.Print(listCoins)
		return
	}

	addCoin := append(listCoins.Names, coinName)
	var listCoins2 types.Coins
	listCoins2.Names = addCoin

	content, err := json.Marshal(listCoins2)
	if err != nil {
		fmt.Println(err)
	}

	if err := os.WriteFile("coins.json", []byte(content), os.ModeAppend); err != nil {
		log.Print("Data not inserted")
		return
	}

	log.Print("Coin Added")
	log.Print(listCoins2)

}
