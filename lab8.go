package main

import (
    "fmt"
    "net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)

    fmt.Println("Server listening on port 8000")
    http.ListenAndServe(":8000", mux)
}
