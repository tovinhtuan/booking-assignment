package handlers

import (
	"booking-assignment/api/booking-api/requests"
	"booking-assignment/api/booking-api/responses"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler interface {
	CustomerBooking(c *gin.Context)
	ViewBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct{
	bookingClient pb.FPTBookingClient
}
func NewBookingHandler(bookingClient pb.FPTBookingClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

func (h * bookingHandler) CustomerBooking(c *gin.Context){
	req := requests.BookingRequest{}

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
	bReq := &pb.BookingRequest{
		CustomerId: req.CustomerID,
		FlightId:   req.FlightID,
	}
	fRes, err := h.bookingClient.CustomerBooking(c.Request.Context(),bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.BookingResponse{
		ID:            fRes.Id,
		CustomerID: fRes.CustomerId,
		FlightID: fRes.FlightId,
		Code:  fRes.Code,
		Status: fRes.Status,
		BookedDate: fRes.BookedDate.AsTime(),
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h * bookingHandler) ViewBooking(c *gin.Context){
	req := requests.ViewBookingRequest{}

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
	bReq := &pb.ViewBookingRequest{
		Code: req.Code,
	}
	fRes, err := h.bookingClient.ViewBooking(c.Request.Context(),bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.ViewBookingResponse{
		Code:         fRes.Code,
		CustomerID:   fRes.CustomerId,
		NameCustomer: fRes.NameCustomer,
		Address:      fRes.Address,
		PhoneNumber:  fRes.PhoneNumber,
		Email:        fRes.Email,
		FlightID:     fRes.FlightId,
		Status:       fRes.Status,
		From:         fRes.From,
		To:           fRes.To,
		NameFlight:   fRes.NameFlight,
		Date:         fRes.Date.AsTime(),
		BookedDate:   fRes.BookedDate.AsTime(),
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h * bookingHandler) CancelBooking(c *gin.Context){
	req := requests.CancelBookingRequest{}
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
	bReq := &pb.CancelBookingRequest{
		Code: req.Code,
	}
	_, err := h.bookingClient.CancelBooking(c.Request.Context(),bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": "cancel booking success",
	})
}