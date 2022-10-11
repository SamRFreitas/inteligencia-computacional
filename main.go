package main

import "fmt"

func main() {
	vetor := [5]int{0, 0, 0, 0, 0}
	// vetor := [5]int{2, 4, 0, 0, 0}
	// vetor := [5]int{2, 3, 7, 9, 10}
	// vetor := [5]int{5, 0, 0, 0, 0}

	exibeVetor(vetor)



	// fmt.Println(buscaInserir(vetor, 2))

	// fmt.Println(buscaRemove(vetor, 10))

	//Tratar quando nao existe esse numero no vetor
	// remover(vetor, 10)

	// var quantidadeDeNumerosInseridosNoVetor int = contarQuantidadeDeNumeroInseridos(vetor)

	// fmt.Println(quantidadeDeNumerosInseridosNoVetor)

	// fmt.Println(verificarSeTemEspacoNoVetor(quantidadeDeNumerosInseridosNoVetor, vetor))

	vetor = inserir(vetor, 27)

	vetor = inserir(vetor, 6)

	vetor = inserir(vetor, 13)

	vetor = inserir(vetor, 2)

	vetor = inserir(vetor, 1)
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

	var indiceDeRemocao int = buscaRemove(vetorOrdenado, numero)

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

func contarQuantidadeDeNumeroInseridos(vetorOrdenado [5]int) int {
	var quantidadeDeNumerosInseridosNoVetor int = 0

	for _, element := range vetorOrdenado {
		if (element > 0) {
			quantidadeDeNumerosInseridosNoVetor ++
		}
	}

	return quantidadeDeNumerosInseridosNoVetor
}

func verificarSeTemEspacoNoVetor(quantidadeDeNumerosInseridosNoVetor int, vetorOrdenado [5]int) bool {

	if quantidadeDeNumerosInseridosNoVetor == len(vetorOrdenado) {
		return false
	} else {
		return true
	}
}

func inserir(vetorOrdenado [5]int, numero int) [5]int{

	var indiceDeInsercao int = buscaInserir(vetorOrdenado, numero)
	
	var i int = len(vetorOrdenado) - 1

	for i > 0 {

		if i >= indiceDeInsercao {
			// fmt.Println(vetorOrdenado[i])
			vetorOrdenado[i] = vetorOrdenado[i-1]
		}

		i--
	}
	

	vetorOrdenado[indiceDeInsercao] = numero

	exibeVetor(vetorOrdenado)

	return vetorOrdenado
}
