package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"micro-logistic/communicate"
	"micro-logistic/controller"
)

func main() {

	// port := os.Getenv("PORT")
	port := 4000
	host := fmt.Sprintf("0.0.0.0:%v", port)

	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	communicate.RegisterLogisticCommunicateServer(grpcServer, &controller.LogisticServer{})

	fmt.Printf("[x] - Server logistic listen http://localhost:%v", port)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err.Error())
	}
}
