package requests

import "github.com/google/uuid"

type BookingRequest struct {
	CustomerID uuid.UUID
	FlightID   uuid.UUID
}
type ViewBookingRequest struct{
	Code int64
}

type CancelBookingRequest struct{
	Code int64
}
type ViewBookingByIDRequest struct{
	CustomerID uuid.UUID
}