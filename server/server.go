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

	symbol := r.URL.Query().Get("symbol")

	price, err := logic.TakePrice(symbol)
	if err != nil {
		fmt.Println("Response error!", err)
		http.Error(w, "External API error", http.StatusServiceUnavailable)
		return // ВЫХОДИМ, чтобы не трогать resp
	}

	storage.WriteHistory(symbol, price)

	w.Header().Set("Content-Type", "application/json") // Обязательно устанавливаем тип отдаваемой инфы!

	json.NewEncoder(w).Encode(price) //отдаем инфу

}

func HistoryRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Обязательно устанавливаем тип отдаваемой инфы!
	json.NewEncoder(w).Encode(storage.ReadHistory())   // Кодируем джейсон
}

func Poller(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	// price, err := logic.TakePrice(symbol)
	// if err == nil {
	// 	storage.WriteHistory(symbol, price)
	// }

	go func(s string) {
		timer := time.NewTicker(1 * time.Minute) // Создаем тикер, который срабатывает каждые 1 минуту

		defer timer.Stop()
		// Важно остановить тикер, чтобы освободить ресурсы

		for range timer.C {
			price, err := logic.TakePrice(s)
			if err == nil {
				storage.WriteHistory(s, price)
			}
		}
	}(symbol)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // Код 202: Принято в обработку
	json.NewEncoder(w).Encode(map[string]string{"status": "monitoring_started", "symbol": symbol})
}

// отвечает за маршруты, парсинг JSON
