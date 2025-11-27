package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "fhg-backend up and running version from v0.0.1 to v0.0.2")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    addr := ":8080"
    log.Printf("starting fhg-backend on %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
