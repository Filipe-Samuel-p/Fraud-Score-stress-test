package main

import (
	"fmt"
	"fraud-score/db"
	"fraud-score/internal/transaction"
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

	repo := &transaction.Repository{DB: db}
	service := &transaction.EngineService{Repo: repo}
	h := &transaction.TransactionHandler{Service: service}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /transaction", h.Transaction)

	log.Println("Servidor rodando na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Erro ao inicializar o servidor: %v", err)
	}

}
