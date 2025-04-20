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

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	var response []byte = []byte("Server is up and running!")

	_, err := writer.Write(response)

	if err != nil {
		fmt.Println(err)
	}
}
