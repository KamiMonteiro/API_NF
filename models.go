package main

type Item struct {
	Codigo        string  `json:"codigo"`
	Descricao     string  `json:"descricao"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

type NotaFiscal struct {
	NumeroNota string `json:"numero_nota"`
	Itens      []Item `json:"itens"`
}

type NotaDetalhada struct {
	CodEmpresa   string  // Ex: 001
	NumeroNota   string  // Ex: 12345
	DataEmissao  string  // Ex: 2025-06-01
	CNPJEmitente string  // Ex: 12345678000195
	ValorTotal   float64 // Ex: 1500.75
}
