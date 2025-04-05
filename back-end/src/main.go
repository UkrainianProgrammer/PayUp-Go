package main

import (
	"net/http"
	"fmt"
)

// http://localhost:4242/create-initial-payment

func main()  {
	http.HandleFunc("/create-payment-intent")
	http.HandleFunc("health")
}