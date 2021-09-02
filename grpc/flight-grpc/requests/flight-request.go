package requests

import (
	"time"

	"github.com/google/uuid"
)

type UpdateFlightRequest struct {
	ID             uuid.UUID
	Name           string
	From           string
	To             string
	Date           time.Time
	Status         string
	Avaliable_slot int64
}
type SearchFlightRequest struct {
	Name string
	From string
	To   string
	Date time.Time
}
