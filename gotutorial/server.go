package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Local server will be running on http://localhost:4242
	// http.HandleFunc takes an endpoint e.g. create-payment-intent and a handler for it
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)
	
	// Run the server and catch/log any error
	log.Println("Listening on http://localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint create-payment-intent was called")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ok")
}