package main

import (
	"fmt"
	"log"
	"net"

	"github.com/63070028/backend-challenge/ex3/server/services"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterBeefSerivceServer(s, services.NewBeefSerivceServer())

	fmt.Println("gRPC server listening on port 50051")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
