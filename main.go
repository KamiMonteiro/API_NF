package main

import (
	"log"
	"net/http"
)

func main() {
	InitDatadabe()

	http.HandleFunc("/notas/", GetItensNotaHandler)

	log.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
