package requests

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	LicenseID   string `json:"licenseID" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binhding:"required"`
	Password    string `json:"password" binding:"required"`
	Active      bool   `json:"active" binding:"required"`
}

type UpdateCustomerRequest struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binhding:"required"`
}

type ChangePasswordRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
