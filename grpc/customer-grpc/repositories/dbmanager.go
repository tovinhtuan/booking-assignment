package repositories

import (
	"booking-assignment/grpc/customer-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	ReadCustomerByID(ctx context.Context, id uuid.UUID)(*models.Customer, error)
	ReadCustomerByName(ctx context.Context, name string)(*models.Customer, error)
	CreateCustomer(ctx context.Context, model *models.Customer)(*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer)(*models.Customer, error)
	ChangePassword(ctx context.Context, model *models.Customer)(*models.Customer, error)
}
type dbmanager struct {
	*gorm.DB
}

func NewDbManager() (CustomerRepository, error){
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=customer port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.Customer{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) ReadCustomerByID(ctx context.Context, id uuid.UUID)(*models.Customer, error){
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Id: id}).First(&customer).Error; err != nil{
		return nil, err
	}
	return &customer, nil
}
func (m *dbmanager) ReadCustomerByName(ctx context.Context, name string)(*models.Customer, error){
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Name: name}).First(&customer).Error; err != nil{
		return nil, err
	}
	return &customer, nil
}
func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Save(model).Error; err != nil{
		return nil, err
	}
	return model, nil
}
func (m *dbmanager) ChangePassword(ctx context.Context, model *models.Customer) (*models.Customer, error){
	if err := m.Save(model).Error; err != nil{
		return nil, err
	}
	return model, nil
}