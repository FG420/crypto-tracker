package crypto

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/FG420/crypto-tracker/types"
)

var (
	baseApi  = "https://api.kucoin.com/api/v1/market/stats?symbol="
	baseApi2 = "-USDT"
)

func CallApi(coinName string) {
	apiUrl := baseApi + coinName + baseApi2
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Printf("Error during API call for %s: %v", coinName, err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error during body reading for %s: %v", coinName, err)
		res.Body.Close()
		return
	}

	var coin types.Coin
	if err := json.Unmarshal(body, &coin); err != nil {
		log.Printf("Error during unmarshal result for %s: %v", coinName, err)
		return
	}

	log.Print(coin)

}
