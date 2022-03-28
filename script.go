package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		exibiMenu()

		comando := opcao()

		switch comando {
		case 1:
			monitoramento()
			fmt.Println("")
		case 2:
			fmt.Println("***** LOGS *****")
		case 0:
			fmt.Println("Saiu")
			os.Exit(0)
		default:
			fmt.Println("Opção invalida")
			os.Exit(-1)
		}
	}
}

func exibiMenu() {
	fmt.Println("********** MENU ************")
	fmt.Println("1) Iniciar monitoramento")
	fmt.Println("2) Exibir os logs")
	fmt.Println("0) Sair do programa")
}

func opcao() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func monitoramento() {
	fmt.Println("********** MONITORANDO SITES ************")
	fmt.Println("")

	var op string

	fmt.Println("Digite o dominio do site (ex: google)")

	fmt.Scan(&op)

	site := "https://" + op + ".com.br"

	fmt.Println(site)

	response, _ := http.Get(site)

	switch response.StatusCode {
	case 200:
		fmt.Println("Site no ar:", response.StatusCode)

	case 400:
		fmt.Println("Site não encontrado, status:", response.StatusCode)

	case 403:
		fmt.Println("Sem permisão, status:", response.StatusCode)

	case 500:
		fmt.Println("Erro interno, status:", response.StatusCode)

	case 404:
		fmt.Println("Site não encontrado, status", response.StatusCode)
	}
}
