package main

import (
	"booking-assignment/api/booking-api/handlers"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	bookingConn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	bookingClient := pb.NewFPTBookingClient(bookingConn)
	h := handlers.NewBookingHandler(bookingClient)
	g := gin.Default()
	gr := g.Group("/v3/api")

	gr.POST("/booking", h.CustomerBooking)
	gr.POST("/viewBooking",h.ViewBooking)
	gr.POST("/cancelBooking", h.CancelBooking)
	http.ListenAndServe(":3333", g)
}