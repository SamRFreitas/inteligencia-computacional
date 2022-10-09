package main
import "fmt"

func main () {
	vetor := [5]int{2, 3, 7, 9, 0}
	
	exibeVetor(vetor)

	fmt.Println(buscaInserir(vetor, 4))

	fmt.Println(buscaRemove(vetor, 3))
}

func exibeVetor(vetorOrdenado [5]int) {
	fmt.Println(vetorOrdenado)
}

func buscaInserir(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if (numero < elemento || numero == 0) {
			return index 
		} 
	}

	return -1
}

func buscaRemove(vetorOrdenado [5]int, numero int) int {
	for index, elemento := range vetorOrdenado {
		if (numero == elemento) {
			return index 
		} 
	}

	return -1
}