package main

import (
	"net/http"
)

func handleRequest(writer http.ResponseWriter, request *http.Request) {}

func main() {
	http.HandleFunc("/generate", handleRequest)
	http.ListenAndServe(":8080", nil)
}

type simpleQRCode struct {
	Content string
	Size    int
}
