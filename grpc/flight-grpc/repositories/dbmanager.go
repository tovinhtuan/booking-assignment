package repositories

import (
	"booking-assignment/grpc/flight-grpc/models"
	"booking-assignment/grpc/flight-grpc/requests"
	"context"
	// "time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FlightRepository interface {
	ReadFlightByID(ctx context.Context, id uuid.UUID) (*models.Flight, error)
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	SearchFlight(ctx context.Context, model *requests.SearchFlightRequest) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDbManager() (FlightRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=flight port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.Flight{},
	)
	if err1 != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}
func (m *dbmanager) ReadFlightByID(ctx context.Context, id uuid.UUID) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{Id: id}).First(&flight).Error; err != nil {
		return nil, err
	}
	return &flight, nil
}
func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Save(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) SearchFlight(ctx context.Context, model *requests.SearchFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	model.From += "%"
	model.To += "%"
	model.Name += "%"
	if err := m.Where("flights.from LIKE ? AND flights.to LIKE ? AND flights.name LIKE ?", model.From, model.To, model.Name).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}
