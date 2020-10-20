package main

import (
	"net/http"
)

func main() {
	hh := handlers.NewHello()

	http.ListenAndServe(":9090", nil)
}
