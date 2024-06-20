package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users  = []User{}
	nextID = 1
	mu     sync.Mutex
)

type Response struct {
	Message string `json:"response"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	user.ID = nextID
	nextID++
	users = append(users, user)
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.ID > len(users) {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	users = append(users[:user.ID-1], users[user.ID:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := Response{Message: fmt.Sprintf("User %d is deleted", user.ID)}
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
