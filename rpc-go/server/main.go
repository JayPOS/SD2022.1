package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func leibniz(steps int) float64 {
	var pi_value float64 = 0.0
	var d float64 = 1
	for i := 0; i < steps; i++ {
		if i%2 == 0 {
			pi_value += 4 / d
		} else {
			pi_value -= 4 / d
		}
		d += 2
	}
	return pi_value
}

func nilakantha(steps int) float64 {
	var pi_value float64 = 3.0
	var d float64 = 2
	for i := 0; i < steps; i++ {
		if i%2 == 0 {
			pi_value += 4 / ((d) * (d + 1) * (d + 2))
		} else {
			pi_value -= 4 / ((d) * (d + 1) * (d + 2))
		}
		d += 2
	}
	return pi_value
}

type Server struct {
}

func (a *Server) CalculatePiLeibniz(nDigits int, pi *float64) error {
	fmt.Printf("Well, you should calculate %d steps pi here\n", nDigits)
	*pi = leibniz(nDigits)

	return nil
}

func (a *Server) CalculatePiNilakantha(nDigits int, pi *float64) error {
	fmt.Printf("Well, you should calculate %d steps pi here\n", nDigits)
	*pi = nilakantha(nDigits)

	return nil
}

func (a *Server) StopServer(){
	
	server.GracefulStop()
    listener.Close()

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
