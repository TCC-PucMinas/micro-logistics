package main

import (
	"fmt"
	"micro-logistic/communicate"
	"micro-logistic/controller"
	"net"

	"google.golang.org/grpc"
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
	communicate.RegisterCarryCommunicateServer(grpcServer, &controller.CarryServer{})
	communicate.RegisterDepositCommunicateServer(grpcServer, &controller.DepositServer{})
	communicate.RegisterTruckCommunicateServer(grpcServer, &controller.TruckServer{})
	communicate.RegisterDriverCommunicateServer(grpcServer, &controller.DriverServer{})
	communicate.RegisterCourierCommunicateServer(grpcServer, &controller.CourierServer{})
	communicate.RegisterCourierRouteCommunicateServer(grpcServer, &controller.CourierRoutes{})

	fmt.Printf("[x] - Server logistic listen http://localhost:%v\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err.Error())
	}
}
