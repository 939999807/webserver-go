package main

import (
	"fmt"
	"net/http"
)

var cfg Config

func main() {
	cfg = parseConfig()

	http.Handle("/css/", http.FileServer(http.Dir("res")))
	http.Handle("/js/", http.FileServer(http.Dir("res")))
	http.Handle("/html/", http.FileServer(http.Dir("res")))

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
