// 1. Using net/http (Standard Library)
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler) // Route
    http.ListenAndServe(":8080", nil)       // Start server
}





