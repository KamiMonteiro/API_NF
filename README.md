# API RESTful em Go para Consulta de Itens de Nota Fiscal

Este projeto implementa uma API RESTful simples em Go (sem frameworks externos) para consulta de itens de nota fiscal e dados detalhados. O foco é simular uma integração comum entre sistemas fiscais e ERPs, usando dados estruturados em arquivos locais (JSON e CSV).

## 🛠 Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Biblioteca padrão**: `net/http`, `encoding/json`, `encoding/csv`
- **Dados simulados**: `itens.json` e `notas.csv`

## 🚀 Como Executar o Projeto

1. Certifique-se de ter o Go instalado (versão 1.20 ou superior).
2. Clone o repositório ou baixe os arquivos.
3. Garanta que os arquivos `itens.json` e `notas.csv` estejam na mesma pasta do projeto.
4. Execute o comando:

```bash
go run main.go database.go handlers.go models.go
```

A API estará disponível em: **http://localhost:8080**

## 📂 Estrutura dos Arquivos

- `itens.json`: contém as notas fiscais com seus respectivos itens
- `notas.csv`: contém dados fiscais detalhados (empresa, CNPJ, valor total, etc)
- `main.go`: inicia a API e registra as rotas
- `handlers.go`: contém os endpoints HTTP
- `database.go`: faz leitura dos arquivos JSON e CSV na inicialização
- `models.go`: define as structs (modelos de dados)

## 📌 Endpoints Disponíveis

### ✅ GET /notas/{numero}/itens

Consulta os itens de uma nota fiscal pelo número.

**Exemplo:**
```bash
GET /notas/12345/itens
```

**Retorno:**
```json
{
  "numero_nota": "12345",
  "itens": [
    {
      "codigo": "A01",
      "descricao": "Caneta",
      "quantidade": 10,
      "valor_unitario": 2.5
    }
  ]
}
```

**Regras:**
- Apenas itens com quantidade > 0 e valor_unitário > 0 são retornados.
- Se a nota não existir, retorna 404.

### ✅ GET /notas-detalhadas

Retorna todas as notas válidas lidas do arquivo `notas.csv`.

**Exemplo de resposta:**
```json
[
  {
    "CodEmpresa": "001",
    "NumeroNota": "12345",
    "DataEmissao": "2025-06-01",
    "CNPJEmitente": "12345678000195",
    "ValorTotal": 1500.75
  }
]
```

**Critérios de validação aplicados:**
- CNPJ deve ter 14 dígitos numéricos
- Data de emissão deve ser válida (formato YYYY-MM-DD)
- Valor total deve ser maior que zero

### ✅ GET /health

Verifica se a API está online.

**Retorno:**
```json
{
  "status": "ok",
  "message": "API rodando com sucesso"
}
```

## 🎁 Bônus Implementado

✔ Leitura e validação de notas a partir de um arquivo CSV (com validações de CNPJ, data e valor)  
✔ Endpoint de verificação de saúde `/health`  
✔ Cache: os dados dos arquivos são carregados uma única vez na inicialização  
✔ Separação clara entre modelos, handlers e leitura de arquivos  

## 🙋‍♀️ Autora

**Cami Monteiro**  
Estudante e desenvolvedora com foco em backend, APIs RESTful e sistemas de dados.
