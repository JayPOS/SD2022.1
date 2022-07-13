package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	var nDigits int = 0
	var pi float64 = 0.0
	var option int = 0

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Printf("Digite o numero de passos que quer calcular do número PI! (tente não passar de 100000)\n\n")
	fmt.Printf("Digite aqui o número de passos: ")
	fmt.Scan(&nDigits)

	fmt.Println("\nQual a função que quer utilizar?")
	fmt.Println("1 - Leibniz")
	fmt.Printf("2 - Nilakantha\n\n")
	fmt.Print("Digite aqui sua opção: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		client.Call("Server.CalculatePiLeibniz", nDigits, &pi)
		fmt.Printf("O valor de Pi depois de %d passos usando a função de Leibniz é %f\n", nDigits, pi)
	case 2:
		client.Call("Server.CalculatePiNilakantha", nDigits, &pi)
		fmt.Printf("O valor de Pi depois de %d passos usando a função de Nilakantha é %f\n", nDigits, pi)
	default:
		fmt.Println("Digite uma opção válida! Tente novamente.")
	}
}
