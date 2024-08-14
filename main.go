package main

import (
	"fmt"
	"log"

	"github.com/FG420/crypto-tracker/crypto"
)

func main() {
	log.Print("WELCOME!")

	for {
		listCoins, err := crypto.LoadCoins()
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Printf("\n\nThese are the Crypto currencies currently saved: \n%s\n\nWhat do you want to do?\n"+
			"[1] - Follow the inputted crypto currency market.\n"+
			"[2] - Insert a new crypto currency to follow in the file.\n"+
			"[3] - Delete the inputted crypto currency from the file.\n"+
			"[4] - Print in a new .json file the current stock market for all the cryptos saved in the .json file.\n"+
			"[0] - Exit.\n", listCoins.Names)

		input := ""
		fmt.Print("Select a choice: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			crypto.FollowCryptoMarket()
		case "2":
			crypto.AddNewCrypto(&listCoins)
		case "3":
			crypto.DeleteCrypto(&listCoins)
		case "4":
			crypto.SaveCryptoStock(listCoins)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			log.Print("Select one of the choices!")
		}
	}
}
