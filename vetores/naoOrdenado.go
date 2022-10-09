package main
import "fmt"
func iniciarVetor(vetor [6]int) {
	fmt.Println("Iniciar Vetor")
	imprimeVetor(vetor)
	for index, _ := range vetor {
		vetor[index] = -1
	}
	imprimeVetor(vetor)
	fmt.Println("------------------------------")
}

func imprimeVetor(vetor [6]int) {
	fmt.Println(vetor)
	fmt.Println()
}

func buscaLinear(vetor [6]int, numero int) int {
	for index, num := range vetor {
		if (num == numero) {
			return index 
		} 
	}

	return -1
}

func removerNumero(vetor [6]int, numero int) {
	fmt.Printf("Remover número %d Vetor: ", numero)
	imprimeVetor(vetor)

	var resultadoBuscaLinear int = buscaLinear(vetor, numero)

	if(resultadoBuscaLinear != -1){
		vetor[resultadoBuscaLinear] = -1
	} else {
		fmt.Println("Número não existente no vetor")
	}

	imprimeVetor(vetor)
	fmt.Println("------------------------------")
	fmt.Println()
}

func inserirNumero(vetor [6]int, numero int) {
	fmt.Printf("Inserir número %d no Vetor: ", numero)
	imprimeVetor(vetor)

	var temEspaco bool = false

	if(buscaLinear(vetor, numero) == -1){

		for index, num := range vetor {
			if(num == -1) {
				vetor[index] = numero
				temEspaco = true
				break
			}
		}

	} else{
		fmt.Println("Número já existente no vetor")
	}

	if (!temEspaco && buscaLinear(vetor, numero) == -1) {
		fmt.Println("Vetor Sem Espaço")
	}

	imprimeVetor(vetor)
	fmt.Println("------------------------------")
	fmt.Println()
}