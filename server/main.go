package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type POSTStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func generateKey() string {
	title := ""
	adj, err := os.Open("./prefix/adj.txt")

	if err != nil {
		panic(err)
	}
	defer adj.Close()

	noun, err := os.Open("./prefix/noun.txt")

	if err != nil {
		panic(err)
	}
	defer noun.Close()

	var adjs []string
	var nouns []string

	scanner := bufio.NewScanner(adj)
	for scanner.Scan() {
		adjs = append(adjs, scanner.Text())
	}

	scannerNoun := bufio.NewScanner(noun)
	for scannerNoun.Scan() {
		nouns = append(nouns, scannerNoun.Text())
	}

	rand.Seed(time.Now().Unix())
	title += strings.Title(adjs[rand.Int()%len(adjs)])
	title += strings.Title(adjs[rand.Int()%len(adjs)])
	title += strings.Title(nouns[rand.Int()%len(nouns)])

	return title
}

func writeBody(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var t POSTStruct
	// err := decoder.Decode(&t)
	// if err != nil {
	// 	panic(err)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(generateKey())
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
		// log.Fatal("$PORT must be set")
	}

	// Generates new router for api to use
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v0/").Subrouter()

	api.HandleFunc("/writeURL/", writeBody).Methods(http.MethodPost)

	handler := cors.Default().Handler(api)

	fmt.Printf(port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
