package main

import (
	"booking-assignment/api/customer-api/handlers"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	customerConn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	customerClient := pb.NewFPTCustomerClient(customerConn)

	h := handlers.NewCustomerHandler(customerClient)
	g := gin.Default()
	gr := g.Group("/v1/api")
	gr.POST("/createCustomer", h.CreateCustomer)
	gr.PUT("/updateCustomer", h.UpdateCustomer)
	gr.PUT("/changePassword", h.ChangePassword)
	gr.POST("/bookingHistory", h.BookingHistory)
	http.ListenAndServe(":3333", g)

}
