package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	var nDigits int = 0
	var pi float64 = 0.0

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	fmt.Println("Digite o numero de casas que quer calcular do número PI! (tente não passar de 100)")
	fmt.Scan(&nDigits)
	client.Call("Server.CalculatePI", nDigits, &pi)
	fmt.Println("pi is", pi)
}
