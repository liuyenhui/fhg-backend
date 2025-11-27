package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "fhg-backend up and running")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    addr := ":8080"
    log.Printf("starting fhg-backend on %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
