package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "fhg-backend update and running version chart update v0.1.9+ ")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	addr := ":8080"
	log.Printf("starting fhg-backend on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
