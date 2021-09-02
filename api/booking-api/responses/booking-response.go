package responses

import "time"

type BookingResponse struct {
	ID         string    `json:"id" binding:"required"`
	CustomerID string    `json:"customer_id" binding:"required"`
	FlightID   string    `json:"flight_id" binding:"required"`
	Code       int64     `json:"code" binding:"required"`
	Status     string    `json:"status" binding:"required"`
	BookedDate time.Time `json:"booked_date" binding:"required"`
}
type ViewBookingResponse struct {
	Code         int64   `json:"code" binding:"required"`
	CustomerID   string  `json:"customer_id" binding:"required"`
	NameCustomer string  `json:"name_customer" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	PhoneNumber  string  `json:"phone_number" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	FlightID     string  `json:"flight_id" binding:"required"`
	Status       string  `json:"status" binding:"required"`
	From         string  `json:"from" binding:"required"`
	To           string  `json:"to" binding:"required"`
	NameFlight   string  `json:"name_flight" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
	BookedDate   time.Time `json:"booked_date" binding:"required"`
}
