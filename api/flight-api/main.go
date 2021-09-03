package main

import (
	"booking-assignment/api/flight-api/handlers"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	flightConn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	flightClient := pb.NewFPTFlightClient(flightConn)
	h := handlers.NewFlightHandler(flightClient)
	g := gin.Default()
	gr := g.Group("/v2/api")

	gr.POST("/createFlight", h.CreateFlight)
	gr.PUT("/updateFlight",h.UpdateFlight)
	gr.POST("/findFlight",h.FindFlight)
	gr.POST("/searchFlight", h.SearchFlight)
	http.ListenAndServe(":3333",g)
}