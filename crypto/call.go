package crypto

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Coin struct {
	Name        string
	Buy         string
	Sell        string
	ChangeRate  string
	ChangePrice string
	High        string
	Low         string
}

var (
	baseApi  = "https://api.kucoin.com/api/v1/market/stats?symbol="
	baseApi2 = "-USDT"
)

func CallApi(coin string) {
	apiUrl := baseApi + coin + baseApi2
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Printf("Error during API call for %s: %v", coin, err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error during body reading for %s: %v", coin, err)
		res.Body.Close()
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error during unmarshal result for %s: %v", coin, err)
		return
	}

	data, _ := json.Marshal(result)
	log.Printf("The data is: %s", data)
}
