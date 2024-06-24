package main

import (
	"encoding/json"
	"net/http"
)

type Produto struct {
	ID         string  `json:"id"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

var produtos = []Produto{
	{ID: "1", Nome: "Produto 1", Descricao: "A descrição do Produto 1", Preco: 10.0, Quantidade: 100},
	{ID: "2", Nome: "Produto 2", Descricao: "A descrição do Produto 2", Preco: 20.0, Quantidade: 200},
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func main() {
	http.HandleFunc("/produtos", getProdutos)
	http.ListenAndServe(":8080", nil)
}
