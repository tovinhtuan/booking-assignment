package responses

import (
	"time"

	"github.com/google/uuid"
)

type ViewBookingResponse struct {
	Code         int64
	CustomerId   uuid.UUID
	NameCustomer string
	Address      string
	PhoneNumber  string
	Email        string
	FlightId     uuid.UUID
	Status       string
	From         string
	To           string
	NameFlight   string
	Date         time.Time
	BookedDate   time.Time
}
