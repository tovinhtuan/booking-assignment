package requests

type BookingRequest struct {
	CustomerID string `json:"customer_id" binding:"required"`
	FlightID string `json:"flight_id" binding:"required"`
}
type ViewBookingRequest struct {
	Code int64   `json:"code" binding:"required"`
}
type CancelBookingRequest struct {
	Code int64   `json:"code" binding:"required"`
}