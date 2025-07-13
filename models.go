package main

type ItemNota struct {
	ID            int     `jason:"id"`
	Descricao     string  `jason:"descricao"`
	Quantidade    int     `jason:"quantidade"`
	ValorUnitario float64 `jason:"valor_unitario"`
}
