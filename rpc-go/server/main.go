package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "rpc-go/server/proto"

	"google.golang.org/grpc"
)

func leibniz(steps uint64) float64 {
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

func nilakantha(steps uint64) float64 {
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

func (a *Server) CalculatePiLeibniz(nDigits uint64, pi *float64) error {
	fmt.Printf("Well, you should calculate %d steps pi here\n", nDigits)
	*pi = leibniz(nDigits)

	return nil
}

func (a *Server) CalculatePiNilakantha(nDigits uint64, pi *float64) error {
	fmt.Printf("Well, you should calculate %d steps pi here\n", nDigits)
	*pi = nilakantha(nDigits)

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	server := grpc.NewServer()
	pb.RegisterPiCalculusServer(server, &PiCalculus{})

	if err != nil {
		log.Fatal("Error registering server", err)
	}

	//log.Println("serving rpc on port 4040")

	server.Serve(listener)

	if err != nil {
		log.Fatal("error serving : ", err)
	}

	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM)
	<-interruptSignal

	server.GracefulStop()
	listener.Close()
}
