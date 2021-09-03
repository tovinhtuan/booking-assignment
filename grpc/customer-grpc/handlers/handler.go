package handlers

import (
	"booking-assignment/grpc/customer-grpc/models"
	"booking-assignment/grpc/customer-grpc/repositories"
	"booking-assignment/pb"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	bookingClient    pb.FPTBookingClient
	pb.UnimplementedFPTCustomerServer
	CustomerRepository repositories.CustomerRepository
}

func NewCustomerHandler(bookingClient pb.FPTBookingClient, customerRepository repositories.CustomerRepository) (*CustomerHandler, error){
	return &CustomerHandler{
		bookingClient: bookingClient,
		CustomerRepository: customerRepository,
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer)(*pb.Customer, error){
	cRequest := &models.Customer{
		Id:          uuid.New(),
		Name:        in.Name,
		Address:     in.Address,
		LicenseID:   in.LicenseId,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
		Active:      false,
	}
	customer, err := h.CustomerRepository.CreateCustomer(ctx,cRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cResponse := &pb.Customer{
		Id:          customer.Id.String(),
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.LicenseID,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      false,
	}
	return cResponse, nil
}
func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer)(*pb.Customer, error){
	customer, err := h.CustomerRepository.ReadCustomerByID(ctx,uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Name != "" {
		customer.Name = in.Name
	}
	if in.Address != ""{
		customer.Address = in.Address
	}
	if in.PhoneNumber != ""{
		customer.PhoneNumber = in.PhoneNumber
	}
	if in.Email != ""{
		customer.Email = in.Email
	}
	newCustomer, err := h.CustomerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}
	cResponse := &pb.Customer{
		Id:          newCustomer.Id.String(),
		Name:        newCustomer.Name,
		Address:     newCustomer.Address,
		LicenseId:   newCustomer.LicenseID,
		PhoneNumber: newCustomer.PhoneNumber,
		Email:       newCustomer.Email,
		Password:    newCustomer.Password,
		Active:      false,
	}
	return cResponse, nil
}
func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest)(*pb.ChangePasswordResponse, error){
	customer, err := h.CustomerRepository.ReadCustomerByName(ctx,in.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.OldPassword == customer.Password {
		if in.NewPassword == customer.Password{
			return nil, status.Error(codes.AlreadyExists,"passwords is exist")
		}else{
			if in.NewPassword != ""{
				customer.Password = in.NewPassword
			}
		}

	}
	_, err = h.CustomerRepository.ChangePassword(ctx,customer)
	if err != nil {
		return nil, err
	}
	// cResponse := &pb.Customer{
	// 	Id:          newCustomer.Id.String(),
	// 	Name:        newCustomer.Name,
	// 	Address:     newCustomer.Address,
	// 	LicenseId:   newCustomer.LicenseID,
	// 	PhoneNumber: newCustomer.PhoneNumber,
	// 	Email:       newCustomer.Email,
	// 	Password:    newCustomer.Password,
	// 	Active:      false,
	// }
	cResponse := &pb.ChangePasswordResponse{
		SuccessChangePassword: true,
	}
	return cResponse, nil
}
func (h *CustomerHandler) FindCustomer(ctx context.Context, in *pb.FindCustomerRequest)(*pb.Customer, error){
	var (
		customer = &models.Customer{}
		err error
	)
	if in.Id == "" &&  in.Name == ""{
		return nil, status.Error(codes.InvalidArgument, "id or name is required")
	}
	if in.Id != "" {
		customer, err = h.CustomerRepository.ReadCustomerByID(ctx, uuid.MustParse(in.Id))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}
	if in.Name != "" {
		customer, err = h.CustomerRepository.ReadCustomerByName(ctx, in.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}
	cRes := &pb.Customer{
		Id:          customer.Id.String(),
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.LicenseID,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      false,
	}
	return cRes, nil
}
func (h *CustomerHandler) BookingHistory(ctx context.Context,in *pb.BookingHistoryRequest)(*pb.ViewBookingByIDResponse, error){
	if in.Id == "" {
		return nil, status.Error(codes.NotFound, "Id is required")
	}
	booking, err := h.bookingClient.ViewBookingByID(ctx, &pb.ViewBookingByIDRequest{
		CustomerId: in.Id,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer_id is not found")
			}
		} else {
			return nil, err
		}
	}
	return booking, nil
}