package logic

import (
	"encoding/json"
	"net/http"
)

func MakeApiUrl(ticker string) string {
	apiURL := "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=" + ticker + "-USDT"
	return apiURL
	//собираем URL в формате для конкретной площадки

}

func TakePrice(symbol string) (string, error) {

	TickerURL := MakeApiUrl(symbol) //подставляем наш тикер в создатель url

	ticker, err := http.Get(TickerURL) // get запрос на api kucoin, в ответ получаем кодированный джейсон с тикером

	if err != nil {
		return "", err // Просто передаем плохую новость наверх
	}

	defer ticker.Body.Close()

	var ConvertedTiker struct {
		Data struct {
			Price string `json:"price"` // Это "полка" для цены
		} `json:"data"` // Это "полка" для объекта data
	} //структура по форме ответов kucoin, для временного храниения декодированной из джейсона информации

	if err := json.NewDecoder(ticker.Body).Decode(&ConvertedTiker); err != nil {
		return "", err // Если биржа прислала "кривой" ответ — выходим
	} // декодировали джейсон и записали во временную структуру

	price := ConvertedTiker.Data.Price //вытащили поле price из декодированного джейсона

	return price, err
}

// расчеты, походы во внешние API
