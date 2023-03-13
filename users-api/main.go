package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	port := ":5600"
	log.Printf("Server started on Port: %s", port)

	users := []User{
		{Name: "Januário Arruda Xavier", Email: "januarioarruda@gmail.com"},
		{Name: "Beatriz Bianca dos Santos", Email: "beatriz.santos@gmail.com"},
		{Name: "Gabriela Alcantara", Email: "gabriela.alcantara@example.com"},
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now())
		log.Println("Request Recebida")
		if err := json.NewEncoder(w).Encode(&users); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(port, nil)
}
