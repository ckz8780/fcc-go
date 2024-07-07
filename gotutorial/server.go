package main

import (
    "encoding/json"
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

    // Callers of this endpoint must conform to this struct
    var req struct {
        ProductId string `json:"product_id"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
        Address1  string `json:"address_1"`
        Address2  string `json:"address_2"`
        City      string `json:"city"`
        State     string `json:"state"`
        Zip       string `json:"zip"`
        Country   string `json:"country"`
    }

    // Decode the submitted body into the struct (& is a pointer)
    err := json.NewDecoder(r.Body).Decode(&req)

    // Error will be due to invalid request data
    if err != nil {
        // Handy multiline strings with `backticks`
        errorText := `You must submit request data in the following format: 
        {
            "product_id": "test product id",
            "first_name": "John",
            "last_name": "Doe",
            "address_1": "123 Avenue",
            "address_2": "Building 1",
            "city": "Los Angeles",
            "state": "California",
            "zip": "99999",
            "country": "United States"
        }`
        http.Error(w, errorText, http.StatusInternalServerError)
        return
    }

    // If we get here, we can access any values of the struct
    fmt.Println(req.FirstName)
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
