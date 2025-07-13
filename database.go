package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatadabe() {
	var err error
	DB, err = sql.Open("sqlite3", "./notas.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()

	//Inserir exemplos

	DB.Exec("INSERT OR IGNORE INTO notas_fiscais (numero) VALUES (?)", "123456")
	DB.Exec("INSERT INTO itens_nota (nota_id, descricao, quantidade, valor_unitario) VALUES (?,?,?,?)", 1, "Produto A", 2, 10.0)
}

func createTables() {
	createNotas := `
	CREATE TABLE IF NOT EXISTS notas_fiscais (
	id INYEGER PRIMARY KEY AUTOINCREMENT, 
	numero TEXT NOT NULL UNIQUE
	);`

	createItens := `
	CREATE TABLE IF NOT EXIST itens_nota(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nota_id INTEGER,
	descricao TEXT,
	quantidade INTEGER,
	valor_unitario REAL,
	FOREING KEY(nota_id) REFERENCES notas_fiscais(id)
	);`

	_, err := DB.Exec(createNotas)
	if err != nil {
		log.Fatal("Erro criando tabela notas_fiscais:", err)
	}

	_, err = DB.Exec(createItens)
	if err != nil {
		log.Fatal("Erro criando tabela itens_notas:", err)
	}

}
