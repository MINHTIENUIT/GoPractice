package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"example.com/model"	
)

var Articles []model.Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")	
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Articles = []model.Article{
		model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()
}