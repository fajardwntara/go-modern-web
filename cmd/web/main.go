package main

import (
	"net/http"

	"github.com/fajardwntara/go-modern-web/pkg/handlers"
)

const PORT_NUMBER = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
