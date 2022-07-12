package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

func (a *Server) CalculatePI(nDigits int, pi *float64) error {
	fmt.Printf("Well, you should calculate %d digits pi here\n", nDigits)

	return nil
}

func main() {
	var server = new(Server)
	err := rpc.Register(server)

	if err != nil {
		log.Fatal("Error registering server", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Println("serving rpc on port 4040")

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving : ", err)
	}
}
