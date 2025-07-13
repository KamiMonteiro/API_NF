package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetItensNotaHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 || parts[1] != "notas" || parts[3] != "itens" {
		http.Error(w, "Rota inválida", http.StatusNotFound)
		return
	}

	numero := parts[2]

	// Buscar a nota fiscal pelo número
	var notaEncontrada *NotaFiscal
	for i := range notas {
		if notas[i].NumeroNota == numero {
			notaEncontrada = &notas[i]
			break
		}
	}

	if notaEncontrada == nil {
		http.Error(w, "Nota fiscal não encontrada", http.StatusNotFound)
		return
	}

	// Filtrar itens com quantidade e valor > 0
	var itensValidos []Item
	for _, item := range notaEncontrada.Itens {
		if item.Quantidade > 0 && item.ValorUnitario > 0 {
			itensValidos = append(itensValidos, item)
		}
	}

	// Montar resposta no formato exigido
	resposta := struct {
		NumeroNota string `json:"numero_nota"`
		Itens      []Item `json:"itens"`
	}{
		NumeroNota: notaEncontrada.NumeroNota,
		Itens:      itensValidos,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}
