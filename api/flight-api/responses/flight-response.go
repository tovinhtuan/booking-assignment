package responses

import "time"

type FlightResponse struct {
	ID            string    `json:"id" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	AvaliableSlot int64     `json:"avaliable_slot" binding:"required"`
}
type SearchFlightResponse struct {
	FlightResponses []*FlightResponse
}