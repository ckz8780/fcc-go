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
	if r.Method != "POST" {
		// Convenience http error handlers auto-respond to the request w/ text/status codes
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Response must be bytes, so we convert a string to an array ("slice") of bytes
	var response []byte = []byte("Server is up and running")
	
	// The := operator is a shortcut for inferring var types
	// the .Write() function returns (int, error), so rather
	// than use `var theInt int, var err error` we can simply
	// use `_, err :=`. The first return value is the response
	// length
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}