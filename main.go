package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	krakenAPIURL = "https://api.kraken.com/0/public/Ticker?pair=BTCEUR,BTCUSD,BTCCHF"
)

type (
	TickerResponse struct {
		Result map[string]struct {
			C []string `json:"c"`
		} `json:"result"`
	}

	Price struct {
		Pair   string `json:"pair"`
		Amount string `json:"amount"`
	}
)

var (
	prices = make(map[string]string)
	mu     sync.RWMutex
)

func fetchTickers() {
	for {
		resp, err := http.Get(krakenAPIURL)
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			time.Sleep(1 * time.Minute)
			continue
		}
		defer resp.Body.Close()

		var tickerResponse TickerResponse
		if err := json.NewDecoder(resp.Body).Decode(&tickerResponse); err != nil {
			fmt.Printf("Error decoding response: %v\n", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		mu.Lock()
		prices["BTC/EUR"] = tickerResponse.Result["XXBTZEUR"].C[0]
		prices["BTC/USD"] = tickerResponse.Result["XXBTZUSD"].C[0]
		prices["BTC/CHF"] = tickerResponse.Result["XBTCHF"].C[0]
		mu.Unlock()

		time.Sleep(1 * time.Minute)
	}
}

func getLTPHandler(w http.ResponseWriter, _ *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	ltp := make([]Price, 0, len(prices))
	for key, value := range prices {
		ltp = append(ltp, Price{Pair: key, Amount: value})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]Price{"ltp": ltp})
}

func main() {
	go fetchTickers()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/ltp", getLTPHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
