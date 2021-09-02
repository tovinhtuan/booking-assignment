package handlers

import (
	"booking-assignment/grpc/booking-grpc/models"
	"booking-assignment/grpc/booking-grpc/repositories"
	"database/sql"

	// "booking-assignment/grpc/booking-grpc/responses"
	"booking-assignment/pb"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type BookingHandler struct {
	customerClient    pb.FPTCustomerClient
	flightClient      pb.FPTFlightClient
	bookingRepository repositories.BookingRepository
	pb.UnimplementedFPTBookingServer
}

func NewBookingHandler(customerClient pb.FPTCustomerClient, flightClient pb.FPTFlightClient, bookingRepository repositories.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		customerClient:    customerClient,
		flightClient:      flightClient,
		bookingRepository: bookingRepository,
	}, nil
}
func (h *BookingHandler) CustomerBooking(ctx context.Context, in *pb.BookingRequest) (*pb.Booking, error) {
	if in.CustomerId == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_id is required")
	}
	if in.FlightId == "" {
		return nil, status.Error(codes.InvalidArgument, "flight_id is required")
	}
	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindCustomerRequest{
		Id: in.CustomerId,
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
	flight, err := h.flightClient.FindFlight(ctx, &pb.FindFlightRequest{
		Id: in.FlightId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "flight_id is not found")
			}
		} else {
			return nil, err
		}
	}
	if flight.AvaliableSlot == 0 {
		return nil, status.Error(codes.InvalidArgument, "full slot!")
	}
	booking := &models.Booking{
		Id:         uuid.New(),
		CustomerID: uuid.MustParse(customer.Id),
		FlightID:   uuid.MustParse(flight.Id),
		Code:       time.Now().UnixNano(),
		Status:     "Successful",
		BookedDate: time.Now(),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  gorm.DeletedAt{},
	}
	res, err := h.bookingRepository.CustomerBooking(ctx, booking)
	if err != nil {
		return nil, err
	}
	bRes := &pb.Booking{}
	err = copier.Copy(&bRes, &res)
	if err != nil {
		return nil, err
	}
	return bRes, nil
}

func (h *BookingHandler) ViewBooking(ctx context.Context, in *pb.ViewBookingRequest) (*pb.ViewBookingResponse, error){
	if in.Code == 0 {
		return nil, status.Error(codes.InvalidArgument, "Code is required")
	}
	viewbook , err := h.bookingRepository.ReadBookingByCode(ctx,in.Code)
	if err != nil {
		return nil, err
	}
	customer, err := h.customerClient.FindCustomer(ctx, &pb.FindCustomerRequest{
		Id: viewbook.CustomerID.String(),
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer_id of viewbook is not found")
			}
		} else {
			return nil, err
		}
	}
	flight, err := h.flightClient.FindFlight(ctx, &pb.FindFlightRequest{
		Id: viewbook.FlightID.String(),
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "flight_id of viewbook is not found")
			}
		} else {
			return nil, err
		}
	}
	view_booking := &pb.ViewBookingResponse{
		Code:         in.Code,
		CustomerId:   viewbook.CustomerID.String(),
		NameCustomer: customer.Name,
		Address:      customer.Address,
		PhoneNumber:  customer.PhoneNumber,
		Email:        customer.Email,
		FlightId:     viewbook.FlightID.String(),
		Status:       flight.Status,
		From:         flight.From,
		To:           flight.To,
		NameFlight:   flight.Name,
		Date:         flight.Date,
		BookedDate:   &timestamppb.Timestamp{
			Seconds: viewbook.BookedDate.Unix(), 
			Nanos: int32(viewbook.BookedDate.Nanosecond())},
	}

	return view_booking, nil
}
func (h *BookingHandler) CancelBooking (ctx context.Context, in *pb.CancelBookingRequest)(*pb.Empty, error){
	if err := h.bookingRepository.CancelBooking(ctx,in.Code); err != nil{
		if err == sql.ErrNoRows{
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &pb.Empty{}, nil
}
func (h *BookingHandler) ViewBookingByID (ctx context.Context, in *pb.ViewBookingByIDRequest)(*pb.ViewBookingByIDResponse, error){
	booking, err := h.bookingRepository.ReadBookingByID(ctx,uuid.MustParse(in.CustomerId))
	if err != nil {
		panic(err)
	}
	customer, err := h.customerClient.FindCustomer(ctx,&pb.FindCustomerRequest{
		Id:          in.CustomerId,
	})
	if err != nil {
		panic(err)
	}
	flights := []*pb.Flight{}
	for _, i := range booking{
		flight, err := h.flightClient.FindFlight(ctx, &pb.FindFlightRequest{
			Id: i.FlightID.String(),
		})
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	view_bookings := []*pb.ViewBookingResponse{}
	for j, i := range flights {
		view_booking := &pb.ViewBookingResponse{
			Code:         booking[j].Code,
			CustomerId:   customer.Id,
			NameCustomer: customer.Name,
			Address:      customer.Address,
			PhoneNumber:  customer.PhoneNumber,
			Email:        customer.Email,
			FlightId:     i.Id,
			Status:       i.Status,
			From:         i.From,
			To:           i.To,
			NameFlight:   i.Name,
			Date:         i.Date,
			BookedDate:   &timestamppb.Timestamp{
				Seconds: booking[j].BookedDate.Unix(), 
				Nanos: int32(booking[j].BookedDate.Nanosecond())},
		}
		view_bookings = append(view_bookings, view_booking)
	}
	// booking_response := &pb.ViewBookingByIDResponse{
	// 	ViewBookingResponses: []*pb.ViewBookingResponse{},
	// }
	// err = copier.CopyWithOption(&booking_response.ViewBookingResponses,&booking,copier.Option{
	// 	IgnoreEmpty: true,
	// 	DeepCopy: true,
	// })
	// if err != nil{
	// 	return nil, err
	// }
	booking_response := &pb.ViewBookingByIDResponse{
		ViewBookingResponses: view_bookings,
	}
	return booking_response, nil
}