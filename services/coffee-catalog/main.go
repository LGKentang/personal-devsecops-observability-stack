package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Coffee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
	Roast  string `json:"roast"`
}

var (
	mu      sync.Mutex
	nextID  = 1
	coffees = []Coffee{
		{ID: 1, Name: "Colombian Supremo", Origin: "Colombia", Roast: "Medium"},
		{ID: 2, Name: "Ethiopian Yirgacheffe", Origin: "Ethiopia", Roast: "Light"},
	}
)

func listCoffees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func addCoffee(w http.ResponseWriter, r *http.Request) {
	var c Coffee
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	mu.Lock()
	c.ID = nextID
	nextID++
	coffees = append(coffees, c)
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func main() {
	// ensure nextID starts higher than seeded IDs
	nextID = 3
	mux := http.NewServeMux()
	mux.HandleFunc("/coffees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listCoffees(w, r)
		case http.MethodPost:
			addCoffee(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("coffee-catalog listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
