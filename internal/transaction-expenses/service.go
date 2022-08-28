package transactionExpenses

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for transactionExpenses.
type Service interface {
	Get(ctx context.Context, id string) (TransactionExpenses, error)
	Query(ctx context.Context, offset, limit int) ([]TransactionExpenses, error)
	QueryByTrip(ctx context.Context, tripId string) ([]TransactionExpenses, error)
	QueryByTransaction(ctx context.Context, transactionId string) ([]TransactionExpenses, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTransactionExpensesRequest) (TransactionExpenses, error)
	Update(ctx context.Context, id string, input UpdateTransactionExpensesRequest) (TransactionExpenses, error)
	Delete(ctx context.Context, id string) (TransactionExpenses, error)
}

// TransactionExpenses represents the data about an transactionExpenses.
type TransactionExpenses struct {
	entity.TransactionExpenses
}

// CreateTransactionExpensesRequest represents an transactionExpenses creation request.
type CreateTransactionExpensesRequest struct {
	TripId        string `json:"trip_id"`
	TripMemberId  string `json:"trip_member_id"`
	TransactionId string `json:"transaction_id"`
	ItemId        string `json:"item_id"`
	Quantity      int64  `json:"quantity"`
}

// Validate validates the CreateTransactionExpensesRequest fields.
func (m CreateTransactionExpensesRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TripMemberId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.ItemId, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateTransactionExpensesRequest represents an transactionExpenses update request.
type UpdateTransactionExpensesRequest struct {
	TripId        string `json:"trip_id"`
	TripMemberId  string `json:"trip_member_id"`
	TransactionId string `json:"transaction_id"`
	ItemId        string `json:"item_d"`
	Quantity      int64  `json:"quantity"`
}

// Validate validates the CreateTransactionExpensesRequest fields.
func (m UpdateTransactionExpensesRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.ItemId, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new transactionExpenses service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the transactionExpenses with the specified the transactionExpenses ID.
func (s service) Get(ctx context.Context, id string) (TransactionExpenses, error) {
	transactionExpenses, err := s.repo.Get(ctx, id)
	if err != nil {
		return TransactionExpenses{}, err
	}
	return TransactionExpenses{transactionExpenses}, nil
}

// Create creates a new transactionExpenses.
func (s service) Create(ctx context.Context, req CreateTransactionExpensesRequest) (TransactionExpenses, error) {
	if err := req.Validate(); err != nil {
		return TransactionExpenses{}, err
	}
	id := entity.GenerateID()
	now := time.Now()

	err := s.repo.Create(ctx, entity.TransactionExpenses{
		ID:            id,
		TripId:        req.TripId,
		TripMemberId:  req.TripMemberId,
		TransactionId: req.TransactionId,
		ItemId:        req.ItemId,
		Quantity:      req.Quantity,
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		return TransactionExpenses{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the transactionExpenses with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTransactionExpensesRequest) (TransactionExpenses, error) {
	if err := req.Validate(); err != nil {
		return TransactionExpenses{}, err
	}

	transactionExpenses, err := s.Get(ctx, id)
	if err != nil {
		return transactionExpenses, err
	}
	transactionExpenses.TripId = req.TripId
	transactionExpenses.TripMemberId = req.TripMemberId
	transactionExpenses.TransactionId = req.TransactionId
	transactionExpenses.ItemId = req.ItemId
	transactionExpenses.Quantity = req.Quantity
	transactionExpenses.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, transactionExpenses.TransactionExpenses); err != nil {
		return transactionExpenses, err
	}
	return transactionExpenses, nil
}

// Delete deletes the transactionExpenses with the specified ID.
func (s service) Delete(ctx context.Context, id string) (TransactionExpenses, error) {
	transactionExpenses, err := s.Get(ctx, id)
	if err != nil {
		return TransactionExpenses{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return TransactionExpenses{}, err
	}
	return transactionExpenses, nil
}

// Count returns the number of transactionExpenses.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the transactionExpenses with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]TransactionExpenses, error) {
	expenses, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []TransactionExpenses{}
	for _, expenses := range expenses {
		result = append(result, TransactionExpenses{expenses})
	}
	return result, nil
}

// Query returns the transactions with the specified offset and limit.
func (s service) QueryByTrip(ctx context.Context, tripId string) ([]TransactionExpenses, error) {
	items, err := s.repo.QueryByTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TransactionExpenses{}
	for _, item := range items {
		result = append(result, TransactionExpenses{item})
	}
	return result, nil
}

// Query returns the transactions with the specified offset and limit.
func (s service) QueryByTransaction(ctx context.Context, tripId string) ([]TransactionExpenses, error) {
	items, err := s.repo.QueryByTransaction(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TransactionExpenses{}
	for _, item := range items {
		result = append(result, TransactionExpenses{item})
	}
	return result, nil
}
