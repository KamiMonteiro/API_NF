package main

type Item struct {
	ID            int     `jason:"id"`
	Descricao     string  `jason:"descricao"`
	Quantidade    int     `jason:"quantidade"`
	ValorUnitario float64 `jason:"valor_unitario"`
}

type NotaFiscal struct {
	NumeroNota string `json:"numero_nota"`
	Item       []Item `json:"itens"`
}
