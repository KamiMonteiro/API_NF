# API RESTful em Go para Consulta de Itens de Nota Fiscal

Este projeto implementa uma API RESTful simples em Go (sem frameworks externos) para consulta de itens de nota fiscal e dados detalhados. O foco é simular uma integração comum entre sistemas fiscais e ERPs, usando dados estruturados em arquivos locais (JSON e CSV).

---

## 🛠 Tecnologias Utilizadas

- Linguagem: Go (Golang)
- Biblioteca padrão: net/http, encoding/json, encoding/csv
- Dados simulados: `itens.json` e `notas.csv`

---

## 🚀 Como Executar o Projeto

1. Certifique-se de ter o Go instalado (versão 1.20 ou superior).
2. Clone o repositório ou baixe os arquivos.
3. Garanta que os arquivos `itens.json` e `notas.csv` estejam na mesma pasta do projeto.
4. Rode o comando:

```bash
go run main.go database.go handlers.go models.go
