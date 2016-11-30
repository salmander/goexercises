package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

type Response struct {
	Code   int    `json:code`
	Result string `json:result`
}

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a response back to the client
// 2. Is the Response struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	httpCode := 200
	result := fmt.Sprintf("%s %s", "hello", name)

	if name == "" {
		httpCode = 400
		result = "Please provide your name by appending '?name' param e.g '?name=Salman'"
	}

	// We use the standard libraries WriteString function to write back to the ResponseWriter interface
	// See docs above
	// Build the response struct
	res := Response{
		Code:   httpCode,
		Result: result,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	// Set JSON header
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.Write(json)

	//io.WriteString(w, fmt.Sprintf("%s %s", "hello", r.FormValue("name")))
}

func main() {
	// Default port
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Add the function that is going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	fmt.Println("Listening on port:", port)
	fmt.Println("To change the port, run the program with an argument e.g 'go run main.go 8080'")

	http.ListenAndServe(":"+port, nil)

}
