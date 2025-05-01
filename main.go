package main

import "fmt"

func CapturaInform(sal, fin *float32) {

	fmt.Println("Digite o seu salário: ")
	fmt.Scan(&*sal)

	fmt.Println("Digite o seu financiamento: ")
	fmt.Scan(&*fin)

}

func AnaliseInform(finan, sala float32) {
	if finan <= (5 * (sala)) {
		fmt.Println("Financiamento Concedido")
	} else {
		var conjuge float32
		fmt.Println("Digite o salário do seu cônjuge: ")
		fmt.Scanf("%f", &conjuge)
		if finan <= (5 * (sala + conjuge)) {
			fmt.Println("Financiamento Concedido")
		} else {
			fmt.Println("Financiamento Negado")
		}
	}
}

func Agradecimento() {
	fmt.Println("Obrigado por nos consultar.")
}

func main() {

	var salario float32
	var financiamento float32

	CapturaInform(&salario, &financiamento)

	AnaliseInform(financiamento, salario)

	Agradecimento()

}
