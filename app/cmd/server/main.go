package main

import (
	"fmt"
	"fraud-score/db"
	"fraud-score/internal/handler"
	"fraud-score/internal/repository"
	"fraud-score/internal/scoring"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Hello, novo projeto")

	db, err := db.ConnectionDB()
	if err != nil {
		log.Fatalf("falha ao conectar no banco: %v", err)
	}
	defer db.Close()

	repo := &repository.Repository{DB: db}
	service := &scoring.EngineService{Repo: repo}
	h := &handler.TransactionHandler{Service: service}

	http.HandleFunc("POST /transaction", h.Transaction)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Servidor rodando")

}
