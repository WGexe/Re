package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func WriteHistory(ticker string, price string, time string) {

	entry := HistoryEntry{
		Symbol: ticker,
		Price:  price,
		Time:   time,
	} // структура которую будем писать в файл

	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("marshalling error", err.Error())
	} // маршаллинг в байты для записи

	data = append(data, '\n') // Добавляем байт новой строки в конец

	file, err := os.OpenFile("History.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Writing history error", err.Error())
	}
	defer file.Close()

	fmt.Println("DATA TO WRITE:", string(data))
	file.Write(data)

}

type HistoryEntry struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   string `json:"time"`
}

func ShowHistory(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("History.txt")
	if err != nil {
		fmt.Println("Reading history error", err.Error())
	}

	lines := strings.Split(string(data), "\n")

	list := []HistoryEntry{}

	for _, line := range lines {

		if line == "" {
			continue
		} // Пропускаем пустые строки

		oneBox := HistoryEntry{}
		err := json.Unmarshal([]byte(line), &oneBox)
		if err != nil {
			continue // Если одна строка битая — не ломаем всё остальное
		}

		list = append(list, oneBox)
	}

	w.Header().Set("Content-Type", "application/json") // Обязательно!
	json.NewEncoder(w).Encode(list)                    // Всё. Одной строкой.

}

// хранилище данных
