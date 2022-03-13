package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
}

func main() {
	handleRequest()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
