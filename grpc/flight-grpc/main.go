package main

import (
	"booking-assignment/grpc/flight-grpc/handlers"
	"booking-assignment/grpc/flight-grpc/repositories"
	"booking-assignment/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	flightRepository, err := repositories.NewDbManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewFlightHandler(flightRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFPTFlightServer(s,h)
	fmt.Println("Listen port 9001")
	s.Serve(listen)

}