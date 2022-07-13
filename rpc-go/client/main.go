package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func StepsNumber() {
	var nDigits int = 0

	fmt.Printf("Digite o numero de passos que quer calcular do número PI! (tente não passar de 100000)\n\n")
	fmt.Printf("Digite aqui o número de passos: ")
	fmt.Scan(&nDigits)

	return nDigits
}

func main() {

	var pi float64 = 0.0
	var option int = 0

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Println("\nQual a função que quer utilizar?")
	fmt.Println("1 - Leibniz")
	fmt.Println("2 - Nilakantha")
	fmt.Println("0 - Fechar servidor\n")
	fmt.Print("Digite aqui sua opção: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		nDigits := StepsNumber()
		client.Call("Server.CalculatePiLeibniz", nDigits, &pi)
		fmt.Printf("O valor de Pi depois de %d passos usando a função de Leibniz é %f\n", nDigits, pi)
	case 2:
		nDigits := StepsNumber()
		client.Call("Server.CalculatePiNilakantha", nDigits, &pi)
		fmt.Printf("O valor de Pi depois de %d passos usando a função de Nilakantha é %f\n", nDigits, pi)
	case 0:
		client.Call("Server.StopServer")
		fmt.Printf("O servidor foi encerrado")
	default:
		fmt.Println("Digite uma opção válida! Tente novamente.")
	}
}
