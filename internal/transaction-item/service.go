package transactionItem

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for transactionItems.
type Service interface {
	Get(ctx context.Context, id string) (TransactionItem, error)
	Query(ctx context.Context, offset, limit int) ([]TransactionItem, error)
	QueryByTrip(ctx context.Context, tripId string) ([]TransactionItem, error)
	QueryByTransaction(ctx context.Context, transactionId string) ([]TransactionItem, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTransactionItemRequest) (TransactionItem, error)
	Update(ctx context.Context, id string, input UpdateTransactionItemRequest) (TransactionItem, error)
	Delete(ctx context.Context, id string) (TransactionItem, error)
}

// TransactionItem represents the data about an transactionItem.
type TransactionItem struct {
	entity.TransactionItem
}

// CreateTransactionItemRequest represents an transactionItem creation request.
type CreateTransactionItemRequest struct {
	TripId        string `json:"trip_id"`
	TransactionId string `json:"transaction_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         int64  `json:"price"`
}

// Validate validates the CreateTransactionItemRequest fields.
func (m CreateTransactionItemRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateTransactionItemRequest represents an transactionItem update request.
type UpdateTransactionItemRequest struct {
	TripId        string `json:"trip_id"`
	TransactionId string `json:"transaction_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         int64  `json:"price"`
}

// Validate validates the CreateTransactionItemRequest fields.
func (m UpdateTransactionItemRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Title, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new transactionItem service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the transactionItem with the specified the transactionItem ID.
func (s service) Get(ctx context.Context, id string) (TransactionItem, error) {
	transactionItem, err := s.repo.Get(ctx, id)
	if err != nil {
		return TransactionItem{}, err
	}
	return TransactionItem{transactionItem}, nil
}

// Create creates a new transactionItem.
func (s service) Create(ctx context.Context, req CreateTransactionItemRequest) (TransactionItem, error) {
	if err := req.Validate(); err != nil {
		return TransactionItem{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.TransactionItem{
		ID:            id,
		TripId:        req.TripId,
		TransactionId: req.TransactionId,
		Title:         req.Title,
		Description:   req.Description,
		Price:         req.Price,
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		return TransactionItem{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the transactionItem with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTransactionItemRequest) (TransactionItem, error) {
	if err := req.Validate(); err != nil {
		return TransactionItem{}, err
	}

	transactionItem, err := s.Get(ctx, id)
	if err != nil {
		return transactionItem, err
	}
	transactionItem.TripId = req.TripId
	transactionItem.TransactionId = req.TransactionId
	transactionItem.Title = req.Title
	transactionItem.Description = req.Description
	transactionItem.Price = req.Price
	transactionItem.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, transactionItem.TransactionItem); err != nil {
		return transactionItem, err
	}
	return transactionItem, nil
}

// Delete deletes the transactionItem with the specified ID.
func (s service) Delete(ctx context.Context, id string) (TransactionItem, error) {
	transactionItem, err := s.Get(ctx, id)
	if err != nil {
		return TransactionItem{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return TransactionItem{}, err
	}
	return transactionItem, nil
}

// Count returns the number of transactionItems.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the transactionItems with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]TransactionItem, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []TransactionItem{}
	for _, item := range items {
		result = append(result, TransactionItem{item})
	}
	return result, nil
}

// Query returns the transactionItems with the specified offset and limit.
func (s service) QueryByTrip(ctx context.Context, tripId string) ([]TransactionItem, error) {
	items, err := s.repo.QueryByTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TransactionItem{}
	for _, item := range items {
		result = append(result, TransactionItem{item})
	}
	return result, nil
}

// Query returns the transactionItems with the specified offset and limit.
func (s service) QueryByTransaction(ctx context.Context, tripId string) ([]TransactionItem, error) {
	items, err := s.repo.QueryByTransaction(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TransactionItem{}
	for _, item := range items {
		result = append(result, TransactionItem{item})
	}
	return result, nil
}
