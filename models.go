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
