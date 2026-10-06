package main

import (
	"fmt"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome the web service, Health Check !!!")
	if err != nil {
		return
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Forge")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/health", healthcheck)

	http.HandleFunc("/", welcome)

	fmt.Println("Server is Listening on the port 8080")
	if err := http.ListenAndServe(":9990", nil); err != nil {
		fmt.Printf("Error in starting the server: %v\n", err)
	}
}
