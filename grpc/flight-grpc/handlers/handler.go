package handlers

import (
	"booking-assignment/grpc/flight-grpc/models"
	"booking-assignment/grpc/flight-grpc/repositories"
	"booking-assignment/grpc/flight-grpc/requests"

	// "booking-assignment/grpc/flight-grpc/requests"
	"booking-assignment/pb"
	"context"
	"database/sql"

	// "time"

	// "time"

	"github.com/google/uuid"
	// "github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	pb.UnimplementedFPTFlightServer
	FlightRepository repositories.FlightRepository
}

func NewFlightHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		FlightRepository: flightRepository,
	}, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	fRequest := &models.Flight{
		Id:             uuid.New(),
		Name:           in.Name,
		From:           in.From,
		To:             in.To,
		Date:           in.Date.AsTime(),
		Status:         in.Status,
		Avaliable_slot: in.AvaliableSlot,
	}
	flight, err := h.FlightRepository.CreateFlight(ctx, fRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	fResponse := &pb.Flight{
		Id:   flight.Id.String(),
		Name: flight.Name,
		From: flight.From,
		To:   flight.To,
		Date: &timestamppb.Timestamp{
			Seconds: flight.Date.Unix(),
			Nanos:   int32(flight.Date.Nanosecond()),
		},
		Status:        flight.Status,
		AvaliableSlot: flight.Avaliable_slot,
	}
	return fResponse, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	flight, err := h.FlightRepository.ReadFlightByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Name != "" {
		flight.Name = in.Name
	}
	if in.From != "" {
		flight.From = in.From
	}
	if in.To != "" {
		flight.To = in.To
	}
	if in.Status != "" {
		flight.Status = in.Status
	}
	if in.AvaliableSlot != 0 {
		flight.Avaliable_slot = in.AvaliableSlot
	}
	newFlight, err := h.FlightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}
	fResponse := &pb.Flight{
		Id:   newFlight.Id.String(),
		Name: newFlight.Name,
		From: newFlight.From,
		To:   newFlight.To,
		Date: &timestamppb.Timestamp{
			Seconds: newFlight.Date.Unix(),
			Nanos:   int32(newFlight.Date.Nanosecond()),
		},
		Status:        newFlight.Status,
		AvaliableSlot: newFlight.Avaliable_slot,
	}
	return fResponse, nil
}
func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	var (
		flights = []*models.Flight{}
		err     error
	)
	if in.From == "" && in.To == "" && in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "id or slut is required")
	}
	// if in.From != "" {
	// 	flights, err = h.FlightRepository.ReadFlightByFrom(ctx, in.From)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows{
	// 			return nil, status.Error(codes.NotFound, err.Error())
	// 		}
	// 		return nil, err
	// 	}
	// }
	// if in.To != "" {
	// 	flights, err = h.FlightRepository.ReadFlightByTo(ctx, in.To, flights)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows{
	// 			return nil, status.Error(codes.NotFound, err.Error())
	// 		}
	// 		return nil, err
	// 	}
	// }
	// if in.Name != "" {
	// 	flights, err = h.FlightRepository.ReadFlightByName(ctx, in.Name, flights)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows{
	// 			return nil, status.Error(codes.NotFound, err.Error())
	// 		}
	// 		return nil, err
	// 	}
	// }
	fReq := &requests.SearchFlightRequest{
		Name: in.Name,
		From: in.From,
		To:   in.To,
	}
	flights, err = h.FlightRepository.SearchFlight(ctx, fReq)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	var arrFlights []*pb.Flight
	for _, v := range flights {
		flightpb := &pb.Flight{
			Id:   v.Id.String(),
			Name: v.Name,
			From: v.From,
			To:   v.To,
			Date: &timestamppb.Timestamp{
				Seconds: v.Date.Unix(),
				Nanos:   int32(v.Date.Nanosecond())},
			Status:        v.Status,
			AvaliableSlot: v.Avaliable_slot,
		}
		arrFlights = append(arrFlights, flightpb)
	}
	fRes := &pb.SearchFlightResponse{
		Flights: arrFlights,
	}
	return fRes, nil
}

func (h *FlightHandler) FindFlight(ctx context.Context, in *pb.FindFlightRequest) (*pb.Flight, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	flight, err := h.FlightRepository.ReadFlightByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	fRes := &pb.Flight{
		Id:   flight.Id.String(),
		Name: flight.Name,
		From: flight.From,
		To:   flight.To,
		Date: &timestamppb.Timestamp{
			Seconds: flight.Date.Unix(),
			Nanos:   int32(flight.Date.Nanosecond())},
		Status:        flight.Status,
		AvaliableSlot: flight.Avaliable_slot,
	}
	return fRes, nil
}
