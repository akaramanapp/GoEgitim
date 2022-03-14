package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Articles []Article

// create type article struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
}

func main() {
	handleRequest()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
