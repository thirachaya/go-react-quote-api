package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{1, "Stay hungry, stay foolish.", "Steve Jobs"},
	{2, "Simplicity is the ultimate sophistication.", "Leonardo da Vinci"},
	{3, "First, solve the problem. Then, write the code.", "John Johnson"},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	random := quotes[rand.Intn(len(quotes))]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(random)
}

func main() {

	http.HandleFunc("/quotes", getQuotes)
	http.HandleFunc("/quotes/random", getRandomQuote)

	http.ListenAndServe(":8080", nil)
}
