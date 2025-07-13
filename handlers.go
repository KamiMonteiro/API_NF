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

	var notaID int
	err := DB.QueryRow("SELECT id FROM notas_fiscais WHERE numero = ?", numero).Scan(&notaID)
	if err != nil {
		http.Error(w, "Nota fiscais não encontrada", http.StatusNotFound)
		return
	}

	rows, err := DB.Query(
		`SELECT id, descricao, quantidade, valor_unitario
		FROM itens_nota WHERE nota_id = ?`, notaID)
	if err != nil {
		http.Error(w, "Erro ao buscar itens", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var itens []ItemNota
	for rows.Next() {
		var item ItemNota
		err := rows.Scan(&item.ID, &item.Descricao, &item.Quantidade, &item.ValorUnitario)
		if err != nil {
			http.Error(w, "Erro ao ler dados", http.StatusInternalServerError)
			return
		}

		itens = append(itens, item)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itens)
}
