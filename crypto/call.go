package crypto

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/FG420/crypto-tracker/types"
)

var (
	baseApi  = "https://api.kucoin.com/api/v1/market/stats?symbol="
	baseApi2 = "-USDT"
)

func CallApi(coinName string) *types.Coin {
	apiUrl := baseApi + coinName + baseApi2
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Panicf("Error during API call for %s: %v", coinName, err)

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicf("Error during body reading for %s: %v", coinName, err)
		res.Body.Close()
	}

	var coin types.Coin
	if err := json.Unmarshal(body, &coin); err != nil {
		log.Panicf("Error during unmarshal result for %s: %v", coinName, err)
	}

	if coin.Data.Buy == "" {
		log.Panic("Error! The data for this crypto coin is not possible to retrive because the coin doesn't exist!")

	}

	return &coin
}

func GetData(coinName string) {
	coin := CallApi(coinName)

	time := time.Unix(coin.Data.Time, 0)

	log.Printf("\n                RESULT"+
		"\n____________________________________"+
		"\n- Searched Coin: %s"+
		"\n- Average Price: %s"+
		"\n- Last: %s"+
		"\n- Buy: %s"+
		"\n- Sell: %s"+
		"\n- High: %s"+
		"\n- Low: %s"+
		"\n- Taker Fee Rate: %s"+
		"\n- Taker Coefficent: %s"+
		"\n- Maker Fee Rate: %s"+
		"\n- Maker Coefficent: %s"+
		"\n- Change Rate: %s"+
		"\n- Change Price: %s"+
		"\n- Volume: %s"+
		"\n- Volume Value: %s"+
		"\n- Time: %s"+
		"\n____________________________________",
		coin.Data.Symbol,
		coin.Data.AveragePrice,
		coin.Data.Last,
		coin.Data.Buy,
		coin.Data.Sell,
		coin.Data.High,
		coin.Data.Low,
		coin.Data.TakerFeeRate,
		coin.Data.TakerCoefficient,
		coin.Data.MakerFeeRate,
		coin.Data.MakerCoefficient,
		coin.Data.ChangeRate,
		coin.Data.ChangePrice,
		coin.Data.Vol,
		coin.Data.VolValue,
		time)
}
