package main

import "fmt"

func main() {
	// vetor := [5]int{2, 3, 7, 9, 0}
	// vetor := [5]int{2, 3, 7, 9, 10}
	vetor := [5]int{5, 0, 0, 0, 0}

	exibeVetor(vetor)

	fmt.Println(buscaInserir(vetor, 4))

	fmt.Println(buscaRemove(vetor, 10))

	//Tratar quando nao existe esse numero no vetor
	// remover(vetor, 7)

	inserir(vetor, 2)
}

func exibeVetor(vetorOrdenado [5]int) {
	fmt.Println(vetorOrdenado)
}

func buscaInserir(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if numero < elemento || elemento == 0 {
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

	var indiceDeRemocao int = buscaInserir(vetorOrdenado, numero)

	vetorOrdenado[indiceDeRemocao] = 0

	for index, elemento := range vetorOrdenado {

		if index != len(vetorOrdenado)-1 {

			if elemento < vetorOrdenado[index+1] && index >= indiceDeRemocao {
				vetorOrdenado[index] = vetorOrdenado[index+1]
				vetorOrdenado[index+1] = 0
			}

		}

	}

	exibeVetor(vetorOrdenado)
}

func inserir(vetorOrdenado [5]int, numero int) {

	var indiceDeInsercao int = buscaInserir(vetorOrdenado, numero)

	// for index, elemento := range vetorOrdenado {

	// 	if index != len(vetorOrdenado)-1 {

	// 		if indiceDeInsercao == index {
	// 			if elemento > vetorOrdenado[index+1] {

	// 			}
	// 		}

	// 	}

	// }

	exibeVetor(vetorOrdenado)
}
