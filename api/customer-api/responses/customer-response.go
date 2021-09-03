package responses

import "time"

type CustomerResponse struct {
	ID          string
	Name        string `json:"name"`
	Address     string `json:"address"`
	LicenseID   string `json:"licenseID"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
type UpdateCustomerResponse struct {
	ID          string
	Name        string `json:"name"`
	Address     string `json:"address"`
	LicenseID   string `json:"licenseID"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
}
type BookingResponse struct {
	Code         int64  `json:"code"`
	CustomerID   string `json:"customer_id"`
	NameCustomer string `json:"name_customer"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	FlightID     string `json:"flight_id"`
	Status       string `json:"status"`
	From         string `json:"from"`
	To           string `json:"to"`
	NameFlight   string `json:"name_flight"`
	Date         time.Time `json:"date"`
	BookedDate   time.Time `json:"booked_date"`
}
type BookingHistoryResponse struct {
	BookingResponses []*BookingResponse
}

type ChangePasswordResponse struct {
	SuccessChangePassword bool `json:"success_change_password"`
}