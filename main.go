package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nurmanhabib/go-rest-api/domain/repository/dummy"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
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
	articleRepository := dummy.NewArticleRepository()

	res := Result {Code: 200, Data: articleRepository.All(), Message: "Successfully retrieved article data"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(js)
}

func findArticle(w http.ResponseWriter, r *http.Request) {
	articleRepository := dummy.NewArticleRepository()

	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["ID"])

	res := Result {Code: 200, Data: articleRepository.FindByID(ID), Message: "Successfully retrieved article data"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(js)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	res := Result {Code: http.StatusNotFound, Message: "Not Found"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	w.Write(js)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	res := Result {Code: http.StatusMethodNotAllowed, Message: "Method Not Allowed"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	w.Write(js)
}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)

	r.HandleFunc("/", getIndex).Methods(http.MethodGet)
	r.HandleFunc("/ping", getPing).Methods(http.MethodGet)
	r.HandleFunc("/articles", getArticles).Methods(http.MethodGet)
	r.HandleFunc("/article/{ID:[0-9]+}", findArticle).Methods(http.MethodGet)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
