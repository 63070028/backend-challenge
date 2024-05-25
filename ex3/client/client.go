package main

import (
	"log"

	"github.com/63070028/backend-challenge/ex3/client/services"
	"google.golang.org/grpc"
)

func main() {

	opts := []grpc.DialOption{grpc.WithInsecure()}
	bc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer bc.Close()

	beefSerivceClient := services.NewBeefSerivceClient(bc)
	beefService := services.NewBeefService(beefSerivceClient)

	err = beefService.GetBeef()

	if err != nil {
		log.Fatal(err)
	}

}
