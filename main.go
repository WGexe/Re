package main

import (
	"Re/server"
	storage "Re/storage"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Price", server.PriceRequest)
	http.HandleFunc("/History", storage.ShowHistory)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Smth goes wrong:", err.Error())
	}

}

// пусковик хоста, содержит все хендлеры
