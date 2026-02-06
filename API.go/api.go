 package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
MODEL
*/
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
REPOSITORY (data layer)
*/
type UserRepository struct{}

func (r *UserRepository) GetAll() []User {
	return []User{
		{ID: 1, Name: "Bishnu", Email: "bishnu@gmail.com"},
		{ID: 2, Name: "Alex", Email: "alex@gmail.com"},
	}
}

/*
SERVICE (business logic)
*/
type UserService struct {
	repo *UserRepository
}

func (s *UserService) GetUsers() []User {
	return s.repo.GetAll()
}

/*
HANDLER (controller / HTTP layer)
*/
type UserHandler struct {
	service *UserService
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

/*
MAIN
*/
func main() {
	// Manual dependency wiring
	repo := &UserRepository{}
	service := &UserService{repo: repo}
	handler := &UserHandler{service: service}

	http.HandleFunc("/users", handler.GetUsers)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
