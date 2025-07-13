package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var notas []NotaFiscal // variável global para armazenar as notas

func LoadNotasFiscais() {
	file, err := os.Open("itens.json")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo JSON: %v", err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo JSON: %v", err)
	}

	err = json.Unmarshal(bytes, &notas)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %v", err)
	}

	log.Println("NOtas fiscais carregadas com sucesso")

}
