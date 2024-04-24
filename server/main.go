package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	fmt.Println("GRPC server listening on port 50051")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
