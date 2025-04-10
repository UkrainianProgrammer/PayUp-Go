package main

import (
	"fmt"
	"net/http"
)

// http://localhost:4242/create-initial-payment

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	http.ListenAndServe("localhost:4242", nil)
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint called!")
}
