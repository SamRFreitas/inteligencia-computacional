package vetores

import "fmt"

func ExibeVetor(vetorOrdenado [5]int) {
	fmt.Println(vetorOrdenado)
}

func BuscaIndiceDeInsercao(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if numero < elemento || elemento == 0 {

			return index
		}
	}

	return -1
}

func BuscaIndiceDeRemocao(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if numero == elemento {
			return index
		}
	}

	return -1
}

func Remover(vetorOrdenado [5]int, numero int) {

	var indiceDeRemocao int = BuscaIndiceDeRemocao(vetorOrdenado, numero)

	vetorOrdenado[indiceDeRemocao] = 0

	for index, elemento := range vetorOrdenado {

		if index != len(vetorOrdenado)-1 {

			if elemento < vetorOrdenado[index+1] && index >= indiceDeRemocao {
				vetorOrdenado[index] = vetorOrdenado[index+1]
				vetorOrdenado[index+1] = 0
			}

		}

	}

	ExibeVetor(vetorOrdenado)
}

func ContarQuantidadeDeNumeroInseridos(vetorOrdenado [5]int) int {
	var quantidadeDeNumerosInseridosNoVetor int = 0

	for _, element := range vetorOrdenado {
		if element > 0 {
			quantidadeDeNumerosInseridosNoVetor++
		}
	}

	return quantidadeDeNumerosInseridosNoVetor
}

func VerificarSeTemEspacoNoVetor(quantidadeDeNumerosInseridosNoVetor int, vetorOrdenado [5]int) bool {

	if quantidadeDeNumerosInseridosNoVetor == len(vetorOrdenado) {
		return false
	} else {
		return true
	}
}

func Inserir(vetorOrdenado [5]int, numero int) [5]int {

	var indiceDeInsercao int = BuscaIndiceDeInsercao(vetorOrdenado, numero)

	var i int = len(vetorOrdenado) - 1

	for i > 0 {

		if i >= indiceDeInsercao {
			// fmt.Println(vetorOrdenado[i])
			vetorOrdenado[i] = vetorOrdenado[i-1]
		}

		i--
	}

	vetorOrdenado[indiceDeInsercao] = numero

	ExibeVetor(vetorOrdenado)

	return vetorOrdenado
}
