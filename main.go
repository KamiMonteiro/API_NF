package main

import (
	"log"
	"net/http"
)

func main() {
	InitDatadabe()

	http.HandleFunc("/notas/", GetItensNotasHandler)

	log.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServer(":8080", nil))

}
