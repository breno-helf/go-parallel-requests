/*
HTTP Server example
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	log.Println("Handler started", counter)
	defer log.Println("Log ended")

	fmt.Println("Answering the request")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "ABACAXI")
	case "POST":
		fmt.Println(r.PostFormValue("id"))
	}
}
