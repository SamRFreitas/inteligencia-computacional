package vetores

import "fmt"

func IniciarVetor(vetor [6]int) {

	for index, _ := range vetor {
		vetor[index] = -1
	}

}

func BuscaLinear(vetor [6]int, numero int) int {

	for index, num := range vetor {
		if num == numero {
			return index
		}
	}

	return -1
}

func RemoverNumero(vetor [6]int, numero int) {
	fmt.Printf("Remover número %d Vetor: ", numero)

	var resultadoBuscaLinear int = BuscaLinear(vetor, numero)

	if resultadoBuscaLinear != -1 {
		vetor[resultadoBuscaLinear] = -1
	} else {
		fmt.Println("Número não existente no vetor")
	}

	fmt.Println("------------------------------")
	fmt.Println()
}

func InserirNumero(vetor [6]int, numero int) {
	fmt.Printf("Inserir número %d no Vetor: ", numero)

	var temEspaco bool = false

	if BuscaLinear(vetor, numero) == -1 {

		for index, num := range vetor {
			if num == -1 {
				vetor[index] = numero
				temEspaco = true
				break
			}
		}

	} else {
		fmt.Println("Número já existente no vetor")
	}

	if !temEspaco && BuscaLinear(vetor, numero) == -1 {
		fmt.Println("Vetor Sem Espaço")
	}

	fmt.Println("------------------------------")
	fmt.Println()
}
