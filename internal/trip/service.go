package trip

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"tribbie/internal/entity"
	"tribbie/pkg/log"
	"time"
)

// Service encapsulates usecase logic for trips.
type Service interface {
	Get(ctx context.Context, id string) (Trip, error)
	Query(ctx context.Context, offset, limit int) ([]Trip, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTripRequest) (Trip, error)
	Update(ctx context.Context, id string, input UpdateTripRequest) (Trip, error)
	Delete(ctx context.Context, id string) (Trip, error)
}

// Trip represents the data about an trip.
type Trip struct {
	entity.Trip
}

// CreateTripRequest represents an trip creation request.
type CreateTripRequest struct {
	Title      	string    `json:"title"`
	Description string    `json:"description"`
	Place	 	string    `json:"place"`
}

// Validate validates the CreateTripRequest fields.
func (m CreateTripRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Description, validation.Length(0, 128)),
		validation.Field(&m.Place, validation.Length(0, 128)),
	)
}

// UpdateTripRequest represents an trip update request.
type UpdateTripRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Place string `json:"place"`
}

// Validate validates the CreateTripRequest fields.
func (m UpdateTripRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new trip service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the trip with the specified the trip ID.
func (s service) Get(ctx context.Context, id string) (Trip, error) {
	trip, err := s.repo.Get(ctx, id)
	if err != nil {
		return Trip{}, err
	}
	return Trip{trip}, nil
}

// Create creates a new trip.
func (s service) Create(ctx context.Context, req CreateTripRequest) (Trip, error) {
	if err := req.Validate(); err != nil {
		return Trip{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Trip{
		ID:        id,
		Title:      req.Title,
		Description:      req.Description,
		Place:      req.Place,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return Trip{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the trip with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTripRequest) (Trip, error) {
	if err := req.Validate(); err != nil {
		return Trip{}, err
	}

	trip, err := s.Get(ctx, id)
	if err != nil {
		return trip, err
	}
	trip.Title = req.Title
	trip.Description = req.Description
	trip.Place = req.Place
	trip.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, trip.Trip); err != nil {
		return trip, err
	}
	return trip, nil
}

// Delete deletes the trip with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Trip, error) {
	trip, err := s.Get(ctx, id)
	if err != nil {
		return Trip{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Trip{}, err
	}
	return trip, nil
}

// Count returns the number of trips.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the trips with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Trip, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Trip{}
	for _, item := range items {
		result = append(result, Trip{item})
	}
	return result, nil
}
