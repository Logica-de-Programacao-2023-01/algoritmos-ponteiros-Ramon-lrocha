package main

import "fmt"

type produto struct {
	nome              string
	preco             float64
	quantidadeestoque int
}

func main() {
	produto := produto{
		nome:              "calça",
		preco:             56.78,
		quantidadeestoque: 50,
	}
	fmt.Println("preço antigo", produto.preco)
	novopreco := 72.65
	atualizarpreco(&produto, novopreco)

	fmt.Println("preço novo", produto.preco)

}
func atualizarpreco(produto *produto, novopreco float64) {
	produto.preco = novopreco
}
