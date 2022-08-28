package tripMember

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for tripMembers.
type Service interface {
	Get(ctx context.Context, id string) (TripMember, error)
	Query(ctx context.Context, offset, limit int) ([]TripMember, error)
	QueryByTrip(ctx context.Context, tripId string) ([]TripMember, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTripMemberRequest) (TripMember, error)
	Update(ctx context.Context, id string, input UpdateTripMemberRequest) (TripMember, error)
	Delete(ctx context.Context, id string) (TripMember, error)
}

// TripMember represents the data about an tripMember.
type TripMember struct {
	entity.TripMember
}

// CreateTripMemberRequest represents an tripMember creation request.
type CreateTripMemberRequest struct {
	TripId string `json:"trip_id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

// Validate validates the CreateTripMemberRequest fields.
func (m CreateTripMemberRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.UserId),
		validation.Field(&m.Name),
	)
}

// UpdateTripMemberRequest represents an tripMember update request.
type UpdateTripMemberRequest struct {
	TripId string `json:"trip_id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

// Validate validates the CreateTripMemberRequest fields.
func (m UpdateTripMemberRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.UserId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new tripMember service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the tripMember with the specified the tripMember ID.
func (s service) Get(ctx context.Context, id string) (TripMember, error) {
	tripMember, err := s.repo.Get(ctx, id)
	if err != nil {
		return TripMember{}, err
	}
	return TripMember{tripMember}, nil
}

// Create creates a new tripMember.
func (s service) Create(ctx context.Context, req CreateTripMemberRequest) (TripMember, error) {
	if err := req.Validate(); err != nil {
		return TripMember{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.TripMember{
		ID:        id,
		TripId:    req.TripId,
		UserId:    req.UserId,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return TripMember{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the tripMember with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTripMemberRequest) (TripMember, error) {
	if err := req.Validate(); err != nil {
		return TripMember{}, err
	}

	tripMember, err := s.Get(ctx, id)
	if err != nil {
		return tripMember, err
	}
	tripMember.TripId = req.TripId
	tripMember.UserId = req.UserId
	tripMember.Name = req.Name
	tripMember.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, tripMember.TripMember); err != nil {
		return tripMember, err
	}
	return tripMember, nil
}

// Delete deletes the tripMember with the specified ID.
func (s service) Delete(ctx context.Context, id string) (TripMember, error) {
	tripMember, err := s.Get(ctx, id)
	if err != nil {
		return TripMember{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return TripMember{}, err
	}
	return tripMember, nil
}

// Count returns the number of tripMembers.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the tripMembers with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]TripMember, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []TripMember{}
	for _, item := range items {
		result = append(result, TripMember{item})
	}
	return result, nil
}

// Query returns the transactions with the specified offset and limit.
func (s service) QueryByTrip(ctx context.Context, tripId string) ([]TripMember, error) {
	items, err := s.repo.QueryByTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TripMember{}
	for _, item := range items {
		result = append(result, TripMember{item})
	}
	return result, nil
}
