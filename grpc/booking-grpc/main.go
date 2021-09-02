package main

import (
	"booking-assignment/grpc/booking-grpc/handlers"
	"booking-assignment/grpc/booking-grpc/repositories"
	"booking-assignment/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	customerConn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	flightConn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	customerClient := pb.NewFPTCustomerClient(customerConn)
	flightClient := pb.NewFPTFlightClient(flightConn)
	listen, err := net.Listen("tcp",":9002")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	bookingRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewBookingHandler(customerClient, flightClient, bookingRepository)
	if err != nil{
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFPTBookingServer(s,h)
	fmt.Println("Listen at port 9002")
	s.Serve(listen)
}