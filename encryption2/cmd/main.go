package main

import (
	"fmt"
	"net/http"
	"os"
	"enc/handlers"
)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/es", handlers.EncodingString)
	mux.HandleFunc("/ds", handlers.DecodingString)
	mux.HandleFunc("/ef", handlers.EncodingFile)
	mux.HandleFunc("/df", handlers.DecodingFile)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}