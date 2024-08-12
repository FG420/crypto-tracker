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

	fmt.Print("WELCOME!")

	var listCoins types.Coins
	readJson, err := os.ReadFile("coins.json")
	if err != nil {
		log.Print("File not found")
		os.Create("coins.json")
		log.Print("File created")
	}

	if err := json.Unmarshal(readJson, &listCoins); err != nil {
		log.Printf("Error during unmarshal file for %s: %v", readJson, err)
		return
	}

	fmt.Printf("\n\nThese are the Crypto currencies currently saved: \n%s\n\nWhat do you wanna do?\n"+
		"[1] - Follow the inputed crypto currency market.\n"+
		"[2] - Insert a new crypto currency to follow in the file.\n"+
		"[3] - Delete the inputed crypto currency from the file.\n", listCoins.Names)

	input := ""
	// fmt.Print("Insert a coin symbol: ")
	fmt.Print("Select a choice: ")
	fmt.Scan(&input)

	var listCoins2 types.Coins

	switch input {
	case "1":
		coinName := inputName()
		crypto.GetData(coinName)

	case "2":
		coinName := inputName()
		crypto.CallApi(coinName)

		if slices.Contains(listCoins.Names, coinName) {
			log.Print("The file contains the coin inputed")
			log.Print(listCoins)
			return
		}

		addCoin := append(listCoins.Names, coinName)
		listCoins2.Names = addCoin

		content, err := json.Marshal(listCoins2)
		if err != nil {
			fmt.Println(err)
		}

		if err := os.WriteFile("coins.json", []byte(content), os.ModeAppend); err != nil {
			log.Print("Data not inserted")
			return
		}

		log.Print(listCoins2)

	case "3":
		coinName := inputName()
		if !slices.Contains(listCoins.Names, coinName) {
			log.Print("The file doesn't contain the coin inputed")
			log.Print(listCoins)
			return
		}

		index := slices.Index(listCoins.Names, coinName)
		del := slices.Delete(listCoins.Names, index, index+1)

		listCoins2.Names = del

		content, err := json.Marshal(listCoins2)
		if err != nil {
			fmt.Println(err)
		}

		if err := os.WriteFile("coins.json", []byte(content), os.ModeAppend); err != nil {
			log.Print("Data not inserted")
			return
		}

		log.Print(listCoins2)

	default:
		log.Print("Select one of the three choises!")
	}

}

func inputName() string {
	inputName := ""
	fmt.Print("Insert a coin symbol: ")
	fmt.Scan(&inputName)
	coinName := strings.ToUpper(inputName)
	return coinName
}
