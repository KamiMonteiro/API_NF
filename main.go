package main

import (
	"log"
	"net/http"
)

func main() {
	LoadNotasFiscais() //carrega as notas do arquivo json

	http.HandleFunc("/notas/", GetItensNotaHandler)

	http.HandleFunc("/health", HealthCheckHandler)

	log.Println("API rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
