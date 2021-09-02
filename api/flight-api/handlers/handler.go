package handlers

import (
	"booking-assignment/api/flight-api/requests"
	"booking-assignment/api/flight-api/responses"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	FindFlight(c *gin.Context)
}

type flightHandler struct{
	flightClient pb.FPTFlightClient
}
func NewFlightHandler(flightClient pb.FPTFlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}
func (h *flightHandler) CreateFlight(c *gin.Context){
	req := requests.CreateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}
			
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status" : http.StatusText(http.StatusBadRequest),
				"error"  : errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusText(http.StatusBadRequest),
			"error" : err.Error(),
		})
		return
	}
	fReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          &timestamppb.Timestamp{
			Seconds: req.Date.Unix(),
			Nanos:   int32(req.Date.Nanosecond()),
		},
		Status:        req.Status,
		AvaliableSlot: req.AvaliableSlot,
	}
	fRes, err := h.flightClient.CreateFlight(c.Request.Context(),fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.FlightResponse{
		ID:            fRes.Id,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvaliableSlot: fRes.AvaliableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) UpdateFlight (c *gin.Context){
	req := requests.UpdateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}
			
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status" : http.StatusText(http.StatusBadRequest),
				"error"  : errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusText(http.StatusBadRequest),
			"error" : err.Error(),
		})
		return
	}
	fReq := &pb.Flight{
		Id:	           req.ID,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          &timestamppb.Timestamp{
			Seconds: req.Date.Unix(),
			Nanos:   int32(req.Date.Nanosecond()),
		},
		Status:        req.Status,
		AvaliableSlot: req.AvaliableSlot,
	}
	fRes, err := h.flightClient.UpdateFlight(c.Request.Context(),fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.FlightResponse{
		ID:            fRes.Id,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvaliableSlot: fRes.AvaliableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *flightHandler) FindFlight (c *gin.Context){
	req := requests.FindFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}
			
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status" : http.StatusText(http.StatusBadRequest),
				"error"  : errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusText(http.StatusBadRequest),
			"error" : err.Error(),
		})
		return
	}
	fReq := &pb.FindFlightRequest{
		Id:	           req.ID,
	}
	fRes, err := h.flightClient.FindFlight(c.Request.Context(),fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.FlightResponse{
		ID:            fRes.Id,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvaliableSlot: fRes.AvaliableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}