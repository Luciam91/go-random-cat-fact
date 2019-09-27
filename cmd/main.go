package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type server struct{}

type factAPIResponse struct {
	Facts []fact `json:"all"`
}

type fact struct {
	ID   string `json:"_id"`
	Text string `json:"text"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/liveness", healthHandler())
	mux.Handle("/facts", factsHandler())

	server := &http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func getRandomFact(factList []byte) string {
	var facts = new(factAPIResponse)
	err := json.Unmarshal(factList, &facts)
	if err != nil {
		panic("Could not decode")
	}

	newFactList := facts.Facts

	rand.Seed(time.Now().UnixNano())

	chosen := newFactList[rand.Intn(len(newFactList)-1)]

	return chosen.Text
}

func factsHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		a, err := http.Get("https://cat-fact.herokuapp.com/facts")

		if err != nil {
			writer.WriteHeader(http.StatusForbidden)
		} else {
			data, _ := ioutil.ReadAll(a.Body)
			randomFact := getRandomFact([]byte(data))
			fmt.Fprintf(writer, string(randomFact))
		}
	})
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	})
}
