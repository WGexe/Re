package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type UserHandler struct {
	DB     string // В реальном коде здесь будет *sql.DB
	Logger string
}

123123
