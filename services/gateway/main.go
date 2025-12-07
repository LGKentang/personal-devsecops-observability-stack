package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func proxyTo(target string, prefix string) *httputil.ReverseProxy {
    u, _ := url.Parse(target)
    rp := httputil.NewSingleHostReverseProxy(u)
    origDirector := rp.Director
    rp.Director = func(r *http.Request) {
        origDirector(r)
        r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
    }
    return rp
}

func main() {
    catalogProxy := proxyTo("http://coffee-catalog:8081", "/catalog")
    ordersProxy := proxyTo("http://coffee-orders:8082", "/orders")

    mux := http.NewServeMux()
    mux.HandleFunc("/catalog/", func(w http.ResponseWriter, r *http.Request) {
        catalogProxy.ServeHTTP(w, r)
    })
    mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
        ordersProxy.ServeHTTP(w, r)
    })
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    log.Println("gateway listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
