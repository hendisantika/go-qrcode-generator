package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/generate", handleRequest)
	http.ListenAndServe(":8080", nil)
}
