package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:4242/create-initial-payment

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on port 4242")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint called!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK!")
}
