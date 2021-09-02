package requests

import "github.com/google/uuid"

type UpdateCustomerRequest struct {
	ID          uuid.UUID
	Name        string
	Address     string
	PhoneNumber string
	Email       string
}
type ChangePasswordRequest struct{
	ID          uuid.UUID
	Password    string
}