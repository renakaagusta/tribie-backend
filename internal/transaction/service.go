package transaction

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for transactions.
type Service interface {
	Get(ctx context.Context, id string) (Transaction, error)
	Query(ctx context.Context, offset, limit int) ([]Transaction, error)
	QueryByTrip(ctx context.Context, tripId string) ([]Transaction, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTransactionRequest) (Transaction, error)
	Update(ctx context.Context, id string, input UpdateTransactionRequest) (Transaction, error)
	Delete(ctx context.Context, id string) (Transaction, error)
}

// Transaction represents the data about an transaction.
type Transaction struct {
	entity.Transaction
}

// CreateTransactionRequest represents an transaction creation request.
type CreateTransactionRequest struct {
	TripId      string `json:"trip_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Validate validates the CreateTransactionRequest fields.
func (m CreateTransactionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Description, validation.Length(0, 128)),
	)
}

// UpdateTransactionRequest represents an transaction update request.
type UpdateTransactionRequest struct {
	TripId      string `json:"trip_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Validate validates the CreateTransactionRequest fields.
func (m UpdateTransactionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new transaction service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the transaction with the specified the transaction ID.
func (s service) Get(ctx context.Context, id string) (Transaction, error) {
	transaction, err := s.repo.Get(ctx, id)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{transaction}, nil
}

// Create creates a new transaction.
func (s service) Create(ctx context.Context, req CreateTransactionRequest) (Transaction, error) {
	if err := req.Validate(); err != nil {
		return Transaction{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Transaction{
		ID:          id,
		TripId:      req.TripId,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return Transaction{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the transaction with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTransactionRequest) (Transaction, error) {
	if err := req.Validate(); err != nil {
		return Transaction{}, err
	}

	transaction, err := s.Get(ctx, id)
	if err != nil {
		return transaction, err
	}
	transaction.TripId = req.TripId
	transaction.Title = req.Title
	transaction.Description = req.Description
	transaction.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, transaction.Transaction); err != nil {
		return transaction, err
	}
	return transaction, nil
}

// Delete deletes the transaction with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Transaction, error) {
	transaction, err := s.Get(ctx, id)
	if err != nil {
		return Transaction{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

// Count returns the number of transactions.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the transactions with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Transaction, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Transaction{}
	for _, item := range items {
		result = append(result, Transaction{item})
	}
	return result, nil
}

// Get returns the transaction with the specified the transaction ID.
func (s service) QueryByTrip(ctx context.Context, tripId string) ([]Transaction, error) {
	items, err := s.repo.QueryByTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []Transaction{}
	for _, item := range items {
		result = append(result, Transaction{item})
	}
	return result, nil
}
