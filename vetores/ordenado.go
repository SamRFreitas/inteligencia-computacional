package main

import "fmt"

func exibeVetor(vetorOrdenado [5]int) {
	fmt.Println(vetorOrdenado)
}

func buscaInserir(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if numero < elemento || numero == 0 {
			return index
		}
	}

	return -1
}

func buscaRemove(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if numero == elemento {
			return index
		}
	}

	return -1
}

func remover(vetorOrdenado [5]int, numero int) {

	var indiceDeRemocao int = buscaRemove(vetorOrdenado, numero)

	for index, _ := range vetorOrdenado {
		if indiceDeRemocao == index {
			vetorOrdenado[index] = vetorOrdenado[index+1]
			vetorOrdenado[index+1] = 0
		}
	}

	exibeVetor(vetorOrdenado)
}
