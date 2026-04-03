package main

import (
	"Re/server"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Price", server.PriceRequest)
	http.HandleFunc("/History", server.HistoryRequest)
	http.HandleFunc("/Poller", server.Poller)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Smth goes wrong:", err.Error())
	}

}

// пусковик хоста, содержит все хендлеры
