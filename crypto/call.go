package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/FG420/crypto-tracker/types"
)

var (
	baseApi    = "https://api.kucoin.com/api/v1/market/stats?symbol="
	baseApi2   = "-USDT"
	coinsNames = []string{"BTC", "DOGE", "SOL", "ETH"}
)

func callApi(coinName string) *types.Coin {
	apiUrl := baseApi + coinName + baseApi2
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Printf("Error during API call for %s: %v", coinName, err)

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error during body reading for %s: %v", coinName, err)
		res.Body.Close()
	}

	var coin types.Coin
	if err := json.Unmarshal(body, &coin); err != nil {
		log.Printf("Error during unmarshal result for %s: %v", coinName, err)
	}

	return &coin
}

func LoadCoins() (types.Coins, error) {
	var listCoins types.Coins
	readJson, err := os.ReadFile("coins.json")
	if err != nil {
		log.Print("File not found, creating a new one.")
		listCoins.Names = coinsNames
		if err := SaveCoins(listCoins); err != nil {
			return listCoins, fmt.Errorf("failed to create file: %v", err)
		}
		log.Print("File created")
	} else if err := json.Unmarshal(readJson, &listCoins); err != nil {
		return listCoins, fmt.Errorf("error during unmarshal: %v", err)
	}
	return listCoins, nil
}

func SaveCoins(listCoins types.Coins) error {
	content, err := json.Marshal(listCoins)
	if err != nil {
		return fmt.Errorf("failed to marshal coins: %v", err)
	}
	if err := os.WriteFile("coins.json", content, 0644); err != nil {
		return fmt.Errorf("failed to write coins to file: %v", err)
	}
	return nil
}

func FollowCryptoMarket() {
	coinName := inputName()
	fmt.Print("\nPress the backspace key to terminate this action")

	exec.Command("stty", "-F", "/dev/tty", "-echo", "cbreak", "min", "1").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo", "-cbreak").Run()

	var b []byte = make([]byte, 1)
	stopChan := make(chan bool)

	go func() {
		for {
			select {
			case <-stopChan:
				return
			default:
				getData(coinName)
				time.Sleep(3 * time.Second)
			}
		}
	}()

	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			log.Fatalf("Error reading from stdin: %v", err)
		}

		if b[0] == 127 {
			fmt.Print("\nAction terminated!")
			stopChan <- true
			break
		}
	}
}

func AddNewCrypto(listCoins *types.Coins) {
	coinName := inputName()
	coin := callApi(coinName)
	if coin.Data.Buy == "" {
		log.Print("Error! The data for this crypto coin cannot be retrieved because the coin doesn't exist!")
		return
	}

	if slices.Contains(listCoins.Names, coinName) {
		log.Print("The file already contains the coin inputted")
		log.Print(listCoins)
		return
	}

	listCoins.Names = append(listCoins.Names, coinName)
	if err := SaveCoins(*listCoins); err != nil {
		log.Printf("Failed to add coin: %v", err)
	} else {
		log.Printf("Coin added: %s", coinName)
	}
}

func DeleteCrypto(listCoins *types.Coins) {
	coinName := inputName()
	if !slices.Contains(listCoins.Names, coinName) {
		log.Print("The file doesn't contain the coin inputted")
		log.Print(listCoins)
		return
	}

	index := slices.Index(listCoins.Names, coinName)
	listCoins.Names = slices.Delete(listCoins.Names, index, index+1)

	if err := SaveCoins(*listCoins); err != nil {
		log.Printf("Failed to delete coin: %v", err)
	} else {
		log.Printf("Coin deleted: %s", coinName)
	}
}

func SaveCryptoStock(listCoins types.Coins) {
	filename := "crypto-stock.json"
	var arrayCoins []types.Coin

	for _, coin := range listCoins.Names {
		coins := callApi(coin)
		arrayCoins = append(arrayCoins, *coins)
	}

	content, err := json.Marshal(arrayCoins)
	if err != nil {
		log.Print("Something went wrong during marshalling!")
		return
	}

	if err := os.WriteFile(filename, content, 0644); err != nil {
		log.Print("Something went wrong during file writing!")
	} else {
		fmt.Printf("\nFile created successfully!")
	}
}

func inputName() string {
	fmt.Print("\nInsert a coin symbol: ")
	var inputName string
	fmt.Scan(&inputName)
	return strings.ToUpper(inputName)
}

func getData(coinName string) {
	coin := callApi(coinName)

	time := time.Unix(coin.Data.Time, 0)

	fmt.Printf("\n                RESULT"+
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
		"\n____________________________________\n",
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
