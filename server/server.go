package server

import (
	"Re/logic"
	"Re/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func PriceRequest(w http.ResponseWriter, r *http.Request) {

	ticker := r.URL.Query().Get("symbol")

	resp, err := http.Get(logic.MakeApiUrl(ticker))
	if err != nil {
		fmt.Println("Response error!", err)
		http.Error(w, "External API error", http.StatusServiceUnavailable)
		return // ВЫХОДИМ, чтобы не трогать resp
	}

	defer resp.Body.Close()

	var conv struct {
		Data struct {
			Price string `json:"price"` // Это "полка" для цены
		} `json:"data"` // Это "полка" для объекта data
	}

	json.NewDecoder(resp.Body).Decode(&conv)

	price := conv.Data.Price

	t := time.Now().Format(time.RFC3339)

	storage.WriteHistory(ticker, price, t)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(conv)

}

// отвечает за маршруты, парсинг JSON
