package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func WriteHistory(ticker string, price string) {

	t := time.Now().Format(time.RFC3339)

	entry := HistoryStruct{
		Symbol: ticker,
		Price:  price,
		Time:   t,
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

	fmt.Println("DATA TO WRITE:", string(data)) // проверка выводит инфу для записи в консоль
	file.Write(data)

}

type HistoryStruct struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   string `json:"time"`
} //структура в которую мы собираем входящие данные для записи и данные истории для вывода

func ReadHistory() []HistoryStruct {

	data, err := os.ReadFile("History.txt")
	if err != nil {
		fmt.Println("Reading history error", err.Error())
	}

	lines := strings.Split(string(data), "\n") //раскидываем по строкам инфу из файла и собираем в одну переменную

	list := []HistoryStruct{} //массив из структур в которых хранятся наши данные

	for _, line := range lines { //в цикле прогоняем каждую строку

		if line == "" {
			continue
		} // Пропускаем пустые строки

		oneBox := HistoryStruct{} //структура в которую формируются данные одной строки для записи в массив
		err := json.Unmarshal([]byte(line), &oneBox)
		if err != nil {
			continue // Если одна строка битая — не ломаем всё остальное
		}

		list = append(list, oneBox) // собираем все вместе
	}

	return list

}

// хранилище данных
