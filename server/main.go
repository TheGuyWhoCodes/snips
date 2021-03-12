package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

type POSTStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Markdown  bool `json:"markdown"`
}

type WriteStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Markdown  bool `json:"markdown"`
	Id    string
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

	title += strings.Title(adjs[rand.Int()%len(adjs)])
	title += strings.Title(adjs[rand.Int()%len(adjs)])
	title += strings.Title(nouns[rand.Int()%len(nouns)])

	return title
}

func writeBody(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t POSTStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(writeNewPaste(t))
}

func getPostInfo(w http.ResponseWriter, r *http.Request) {
	pathParams := r.URL.Query()["id"][0]

	fmt.Println(pathParams)

	ctx := context.Background()

	conf := &firebase.Config{
		DatabaseURL: "https://snips-ec210-default-rtdb.firebaseio.com/",
	}

	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	// The app only has access to public data as defined in the Security Rules
	ref := client.NewRef(pathParams)
	var data WriteStruct
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func writeNewPaste(post POSTStruct) WriteStruct {
	ctx := context.Background()

	// config to firebase database url
	conf := &firebase.Config{
		DatabaseURL: "https://snips-ec210-default-rtdb.firebaseio.com/",
	}

	// load in serviceAccount key
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// yikes, no access to the database! we should log and throw an error ASAP
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// first, lets generate a new key for our data
	id := generateKey();
	ref := client.NewRef(fmt.Sprint(id))

	output := WriteStruct{
		Title: post.Title,
		Body:  post.Body,
		Markdown:  post.Markdown,
		Id:    id,
	}
	ref.Set(ctx, output)

	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return output
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	// Generates new router for api to use
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v0/").Subrouter()

	api.HandleFunc("/writeBody/", writeBody).Methods(http.MethodPost)
	api.HandleFunc("/getPost/", getPostInfo).Methods(http.MethodGet)
	
	handler := cors.Default().Handler(api)

	fmt.Printf(port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
