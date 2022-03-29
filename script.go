package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

//Constantes
const tempoMonitoramento = 3

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

	var tamanhoArray int
	sites := []string{}
	var nomeSite string
	var url string

	fmt.Println("Quantos sites deseja monitorar?")
	fmt.Scan(&tamanhoArray)

	for i := 0; i < tamanhoArray; i++ {
		fmt.Print("Digite o dominio do site (ex: google): ")
		fmt.Scan(&nomeSite)
		url = "https://" + nomeSite + ".com.br"
		sites = append(sites, url)
	}

	fmt.Println("")
	for j := 0; j < len(sites); j++ {
		response, _ := http.Get(sites[j])

		fmt.Println("Site:", sites[j], ",status:", response.StatusCode, ",id:", j+1)
		verificaStatus(response.StatusCode)
		fmt.Println("-------------------------------------")
		time.Sleep(tempoMonitoramento * time.Second)
	}
}

func verificaStatus(status int) {
	switch status {
	case 200:
		fmt.Println("Site no ar:", status)

	case 400:
		fmt.Println("Site não encontrado, status:", status)

	case 403:
		fmt.Println("Sem permisão, status:", status)

	case 500:
		fmt.Println("Erro interno, status:", status)

	case 404:
		fmt.Println("Site não encontrado, status", status)
	}
}
