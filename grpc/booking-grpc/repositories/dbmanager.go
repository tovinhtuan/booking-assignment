package repositories

import (
	"booking-assignment/grpc/booking-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookingRepository interface{
	ReadBookingByID(ctx context.Context,customerID uuid.UUID)([]*models.Booking, error)
	ReadBookingByCode(ctx context.Context,code int64)(*models.Booking, error)
	CustomerBooking(ctx context.Context, model *models.Booking)(*models.Booking,error)
	CancelBooking(ctx context.Context, code int64) error
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager()(BookingRepository, error){
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=booking port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.Booking{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) CustomerBooking(ctx context.Context, model *models.Booking)(*models.Booking, error){
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (m *dbmanager) ReadBookingByID (ctx context.Context,customerID uuid.UUID)([]*models.Booking, error){
	booking := []*models.Booking{}
	if err := m.Where(&models.Booking{CustomerID: customerID}).Find(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}
func (m *dbmanager) ReadBookingByCode (ctx context.Context,code int64)(*models.Booking, error){
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Code: code}).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
func (m *dbmanager) CancelBooking(ctx context.Context, code int64) error{
	booking, err := m.ReadBookingByCode(ctx, code)
	if err != nil {
		return err
	}
	return m.Unscoped().Delete(&booking).Error
}