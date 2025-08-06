package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BitcoinPrices []struct {
	Date                       string  `json:"Date"`
	Price                      int     `json:"Price"`
	Open                       float64 `json:"Open"`
	High                       float64 `json:"High"`
	ChangePercentFromLastMonth float64 `json:"ChangePercentFromLastMonth"`
	Volume                     string  `json:"Volume"`
}

func main() {
	url := "https://api.sampleapis.com/bitcoin/historical_prices"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bitcoinPrices BitcoinPrices
	json.Unmarshal(responseBodyBytes, &bitcoinPrices)

	for _, record := range bitcoinPrices {
		fmt.Println(record)
	}

}
