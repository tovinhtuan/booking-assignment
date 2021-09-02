package responses

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
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	LicenseID   string `json:"licenseID" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binhding:"required"`
	Password    string `json:"password" binding:"required"`
	Active      bool   `json:"active" binding:"required"`
}

