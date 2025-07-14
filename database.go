package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var notas []NotaFiscal // variável global para armazenar as notas

var notasDetalhadas []NotaDetalhada

func LoadNotasFiscais() {
	//log.Printf("Notas carregadas: %+v\n", notas)
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

func LoadNotasDetalhadasCSV() {
	file, err := os.Open("notas.csv")
	if err != nil {
		log.Fatalf("Erro ao abrir notas.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // o separador é ponto e vírgula

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Erro ao ler CSV: %v", err)
	}

	for i, line := range lines {
		if i == 0 {
			continue // pula cabeçalho
		}

		if len(line) != 5 {
			log.Printf("Linha %d ignorada: formato inválido", i+1)
			continue
		}

		valor, err := strconv.ParseFloat(strings.ReplaceAll(line[4], ",", "."), 64)
		if err != nil || valor <= 0 {
			log.Printf("Linha %d ignorada: valor inválido", i+1)
			continue
		}

		_, err = time.Parse("2006-01-02", line[2])
		if err != nil {
			log.Printf("Linha %d ignorada: data inválida", i+1)
			continue
		}

		if !isCNPJValido(line[3]) {
			log.Printf("Linha %d ignorada: CNPJ inválido", i+1)
			continue
		}

		nota := NotaDetalhada{
			CodEmpresa:   line[0],
			NumeroNota:   line[1],
			DataEmissao:  line[2],
			CNPJEmitente: line[3],
			ValorTotal:   valor,
		}

		notasDetalhadas = append(notasDetalhadas, nota)
	}

	log.Printf("Notas detalhadas carregadas: %d", len(notasDetalhadas))
}

func isCNPJValido(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	return len(cnpj) == 14 && strings.IndexFunc(cnpj, func(r rune) bool {
		return r < '0' || r > '9'
	}) == -1
}
