package main

import (
	"booking-assignment/grpc/customer-grpc/handlers"
	"booking-assignment/grpc/customer-grpc/repositories"
	"booking-assignment/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	bookingConn, err := grpc.Dial(":9002",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	bookingClient := pb.NewFPTBookingClient(bookingConn)
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	customerRepository, err := repositories.NewDbManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewCustomerHandler(bookingClient,customerRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFPTCustomerServer(s,h)
	fmt.Println("listen port 9000")
	s.Serve(listen)
}