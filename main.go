package main

import (
	"fmt"
	"net"
	"os"

	"github.com/TCC-PucMinas/micro-logistics/communicate"
	"github.com/TCC-PucMinas/micro-logistics/controller"
	"github.com/TCC-PucMinas/micro-logistics/service"
	"google.golang.org/grpc"
)

func main() {

	port := os.Getenv("PORT")
	// port := 4000
	host := fmt.Sprintf("0.0.0.0:%v", port)

	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic(err)
	}

	calc := service.Calculate{}

	calc.CalculateRoute()

	grpcServer := grpc.NewServer()
	communicate.RegisterLogisticCommunicateServer(grpcServer, &controller.LogisticServer{})

	fmt.Println("[x] - Server register listen...")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err.Error())
	}
}
