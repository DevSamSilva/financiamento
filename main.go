package main

import "fmt"

func main() {

	var salario float32
	var financiemento float32

	fmt.Println("Digite o seu salário: ")
	fmt.Scanf("%f", &salario)

	fmt.Println("Digite o seu financiamento: ")
	fmt.Scanf("%f", &financiemento)

	if financiemento <= (5 * salario) {
		fmt.Println("Financiamento Concedido")
	} else {
		var conjuge float32
		fmt.Println("DIgite o salário do seu cônjuge: ")
		fmt.Scanf("%f", &conjuge)
		if financiemento <= (5 * (salario + conjuge)) {
			fmt.Println("Financiamento Concedido")
		} else {
			fmt.Println("Financiamento Negado")
		}
	}

	fmt.Println("Obrigado por nos consultar.")

}
