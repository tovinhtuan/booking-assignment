package handlers

import (
	"booking-assignment/api/customer-api/requests"
	"booking-assignment/api/customer-api/responses"
	"booking-assignment/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
	BookingHistory(c*gin.Context)
}

type customerHandler struct {
	customerClient pb.FPTCustomerClient	
}

func NewCustomerHandler(customerClient pb.FPTCustomerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}
func (h *customerHandler) CreateCustomer(c *gin.Context){
	req := requests.CreateCustomerRequest{}

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
	cReq := &pb.Customer{
		Name:        req.Name,
		Address:     req.Address,
		LicenseId:   req.LicenseID,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		Active:      req.Active,
	}
	cRes, err := h.customerClient.CreateCustomer(c.Request.Context(),cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.CustomerResponse{
		ID:          cRes.Id,
		Name:        cRes.Name,
		Address:     cRes.Address,
		LicenseID:   cRes.LicenseId,
		PhoneNumber: cRes.PhoneNumber,
		Email:       cRes.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *customerHandler)UpdateCustomer(c *gin.Context){
	
	req := requests.UpdateCustomerRequest{}
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
	cReq := &pb.Customer{
		Id:          req.ID,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
	cRes, err := h.customerClient.UpdateCustomer(c.Request.Context(),cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.UpdateCustomerResponse{
		ID:          cRes.Id,
		Name:        cRes.Name,
		Address:     cRes.Address,
		LicenseID:   cRes.LicenseId,
		PhoneNumber: cRes.PhoneNumber,
		Email:       cRes.Email,
		Password:    cRes.Password,
		Active:      false,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *customerHandler) ChangePassword (c *gin.Context){
	req := requests.ChangePasswordRequest{}
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
	cReq := &pb.ChangePasswordRequest{
		Name:        req.Name,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	_, err := h.customerClient.ChangePassword(c.Request.Context(),cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.ChangePasswordResponse{
		SuccessChangePassword: true,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *customerHandler) BookingHistory (c *gin.Context){
	req := requests.BookingHistoryRequest{}
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
	cReq := &pb.BookingHistoryRequest{
		Id: req.ID,
	}
	cRes , err := h.customerClient.BookingHistory(c.Request.Context(), cReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	var dtos []*responses.BookingResponse
	var dto *responses.BookingResponse
	for _, v := range cRes.ViewBookingResponses{
		dto = &responses.BookingResponse{
			Code:         v.Code,
			CustomerID:   v.CustomerId,
			NameCustomer: v.NameCustomer,
			Address:      v.Address,
			PhoneNumber:  v.PhoneNumber,
			Email:        v.Email,
			FlightID:     v.FlightId,
			Status:       v.Status,
			From:         v.From,
			To:           v.To,
			NameFlight:   v.NameFlight,
			Date:         v.Date.AsTime(),
			BookedDate:   v.BookedDate.AsTime(),
		}
		dtos = append(dtos,dto)
	}
	dtoss := &responses.BookingHistoryResponse{
		BookingResponses: dtos,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtoss,
	})
}