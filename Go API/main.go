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

// 1. Using net/http (Standard Library)



func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", GetUser).Methods("GET") // Route with parameter
    http.ListenAndServe(":8080", r)
}
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", GetUser).Methods("GET") // Route with parameter
    http.ListenAndServe(":8080", r)
}
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", GetUser).Methods("GET") // Route with parameter
    http.ListenAndServe(":8080", r)
}
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", GetUser).Methods("GET") // Route with parameter
    http.ListenAndServe(":8080", r)
}

package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", GetUser).Methods("GET") // Route with parameter
    http.ListenAndServe(":8080", r)
}
