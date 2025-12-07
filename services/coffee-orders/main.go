package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Order struct {
    ID       int `json:"id"`
    CoffeeID int `json:"coffee_id"`
    Quantity int `json:"quantity"`
}

var (
    mu     sync.Mutex
    nextID = 1
    orders = []Order{}
)

func listOrders(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
    var in Order
    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, "invalid payload", http.StatusBadRequest)
        return
    }

    // validate coffee exists by asking catalog
    catalogURL := os.Getenv("COFFEE_CATALOG_URL")
    if catalogURL == "" {
        catalogURL = "http://localhost:8081"
    }
    resp, err := http.Get(catalogURL + "/coffees")
    if err != nil {
        http.Error(w, "could not validate coffee", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    // quick search for ID in catalog list
    if !bytes.Contains(body, []byte("\"id\":"+jsonNumber(in.CoffeeID))) {
        http.Error(w, "coffee not found", http.StatusBadRequest)
        return
    }

    mu.Lock()
    in.ID = nextID
    nextID++
    orders = append(orders, in)
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(in)
}

// helper to convert int to bytes like JSON number (no spaces)
func jsonNumber(n int) string {
    return fmt.Sprintf("%d", n)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            listOrders(w, r)
        case http.MethodPost:
            createOrder(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    srv := &http.Server{
        Addr:         ":8082",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    log.Printf("coffee-orders listening on %s", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}
