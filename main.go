package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
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

var produtoMutex = &sync.Mutex{}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProdutos(w, r)
	case "POST":
		postProduto(w, r)
	case "PUT":
		putProduto(w, r)
	case "PATCH":
		patchProduto(w, r)
	case "DELETE":
		deleteProduto(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func postProduto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var novoProduto Produto

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &novoProduto)

	produtoMutex.Lock()
	produtos = append(produtos, novoProduto)
	produtoMutex.Unlock()

	json.NewEncoder(w).Encode(novoProduto)
}

func putProduto(w http.ResponseWriter, r *http.Request) {
	// Implementação simples para atualizar um produto existente
	w.Header().Set("Content-Type", "application/json")

	var produtoAtualizado Produto

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &produtoAtualizado)

	produtoMutex.Lock()
	for i, produto := range produtos {
		if produto.ID == produtoAtualizado.ID {
			produtos[i] = produtoAtualizado
		}
	}
	produtoMutex.Unlock()

	json.NewEncoder(w).Encode(produtoAtualizado)
}

func patchProduto(w http.ResponseWriter, r *http.Request) {
	// Implementação simples para modificar parcialmente um produto existente
	w.Header().Set("Content-Type", "application/json")

	var produtoAtualizado Produto

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &produtoAtualizado)

	produtoMutex.Lock()
	for i, produto := range produtos {
		if produto.ID == produtoAtualizado.ID {
			if produtoAtualizado.Nome != "" {
				produtos[i].Nome = produtoAtualizado.Nome
			}
			if produtoAtualizado.Descricao != "" {
				produtos[i].Descricao = produtoAtualizado.Descricao
			}
			if produtoAtualizado.Preco != 0 {
				produtos[i].Preco = produtoAtualizado.Preco
			}
			if produtoAtualizado.Quantidade != 0 {
				produtos[i].Quantidade = produtoAtualizado.Quantidade
			}
		}
	}
	produtoMutex.Unlock()

	json.NewEncoder(w).Encode(produtoAtualizado)
}

func deleteProduto(w http.ResponseWriter, r *http.Request) {
	// Implementação simples para deletar um produto existente
	w.Header().Set("Content-Type", "application/json")

	var produtoParaDeletar Produto

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &produtoParaDeletar)

	produtoMutex.Lock()
	for i, produto := range produtos {
		if produto.ID == produtoParaDeletar.ID {
			produtos = append(produtos[:i], produtos[i+1:]...)
			break
		}
	}
	produtoMutex.Unlock()

	json.NewEncoder(w).Encode(produtos)
}

func main() {
	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}
