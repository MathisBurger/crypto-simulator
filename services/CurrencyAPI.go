package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CurrencyStruct struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUSD          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

type AllCurrencysResponse struct {
	Data      []CurrencyStruct `json:"data"`
	Timestamp int              `json:"timestamp"`
}

func GetAllCurrencys() AllCurrencysResponse {
	url := "https://api.coincap.io/v2/assets"
	httpClient := http.Client{Timeout: time.Second * 2}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := httpClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, _ := ioutil.ReadAll(res.Body)
	obj := AllCurrencysResponse{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println("resp code:", res.StatusCode)
		fmt.Println("response:", string(body))
		panic(err.Error())
	}
	return obj
}
