package main

import (
	"fmt"

	"github.com/FG420/crypto-tracker/crypto"
)

func main() {
	coin := ""
	fmt.Print("Insert a coin symbol: ")
	fmt.Scan(&coin)

	crypto.CallApi(coin)
}
