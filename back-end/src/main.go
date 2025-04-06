package main

import (
	"net/http"
	"fmt"
)

// http://localhost:4242/create-initial-payment
// http://localhost:4042/health

func main()  {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("health")
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	
}