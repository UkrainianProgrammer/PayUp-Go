package main

import (
	"fmt"
	"net/http"
)

// http://localhost:4242/create-initial-payment
// http://localhost:4042/health

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint called!")
}
