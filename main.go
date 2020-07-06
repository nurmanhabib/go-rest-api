package main

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
)

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	data := "Welcome"

	res := Result {Code: 200, Data: data, Message: "Welcome"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(js)
}

func getPing(w http.ResponseWriter, r *http.Request) {

	data := map[string]string {
		"timestamp": time.Now().String(),
	}

	res := Result {Code: 200, Data: data, Message: "PONG!"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(js)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article {
		Article{ID: 101, Title: "Article 1", Author: "John Doe", Content: "Sample content article 1"},
		Article{ID: 102, Title: "Article 2", Author: "Tarjono", Content: "Sample content article 2"},
		Article{ID: 103, Title: "Article 3", Author: "Tara Basro", Content: "Sample content article 3"},
	}

	res := Result {Code: 200, Data: articles, Message: "Successfully retrieved article data"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(js)
}

func main() {
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/ping", getPing)
	http.HandleFunc("/articles", getArticles)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
