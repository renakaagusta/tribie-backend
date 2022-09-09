package transactionPayment

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for transactionPayments.
type Service interface {
	Get(ctx context.Context, id string) (TransactionPayment, error)
	Query(ctx context.Context, offset, limit int) ([]TransactionPayment, error)
	QueryByTrip(ctx context.Context, tripId string) ([]TransactionPayment, error)
	QueryByTransaction(ctx context.Context, tripId string) ([]TransactionPayment, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTransactionPaymentRequest) (TransactionPayment, error)
	Update(ctx context.Context, id string, input UpdateTransactionPaymentRequest) (TransactionPayment, error)
	Delete(ctx context.Context, id string) (TransactionPayment, error)
}

// TransactionPayment represents the data about an transactionPayment.
type TransactionPayment struct {
	entity.TransactionPayment
}

// CreateTransactionPaymentRequest represents an transactionPayment creation request.
type CreateTransactionPaymentRequest struct {
	TripId        string `json:"trip_id"`
	TransactionId string `json:"transaction_id"`
	UserFromId    string `json:"user_from_id"`
	UserToId      string `json:"user_to_id"`
	Status        string `json:"status"`
	Nominal       int64  `json:"nominal"`
}

// Validate validates the CreateTransactionPaymentRequest fields.
func (m CreateTransactionPaymentRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Nominal, validation.Required),
	)
}

// UpdateTransactionPaymentRequest represents an transactionPayment update request.
type UpdateTransactionPaymentRequest struct {
	TripId        string `json:"trip_id"`
	TripMemberId  string `json:"trip_member_id"`
	TransactionId string `json:"transaction_id"`
	UserFromId    string `json:"user_from_id"`
	UserToId      string `json:"user_to_id"`
	Status        string `json:"status"`
	Nominal       int64  `json:"transaction_nominal"`
}

// Validate validates the CreateTransactionPaymentRequest fields.
func (m UpdateTransactionPaymentRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.TripId, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.TransactionId, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new transactionPayment service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the transactionPayment with the specified the transactionPayment ID.
func (s service) Get(ctx context.Context, id string) (TransactionPayment, error) {
	transactionPayment, err := s.repo.Get(ctx, id)
	if err != nil {
		return TransactionPayment{}, err
	}
	return TransactionPayment{transactionPayment}, nil
}

// Create creates a new transactionPayment.
func (s service) Create(ctx context.Context, req CreateTransactionPaymentRequest) (TransactionPayment, error) {
	if err := req.Validate(); err != nil {
		return TransactionPayment{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.TransactionPayment{
		ID:            id,
		TripId:        req.TripId,
		TripMemberId:  "",
		TransactionId: req.TransactionId,
		UserFromId:    req.UserFromId,
		UserToId:      req.UserToId,
		Nominal:       req.Nominal,
		Status:        req.Status,
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		return TransactionPayment{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the transactionPayment with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTransactionPaymentRequest) (TransactionPayment, error) {
	if err := req.Validate(); err != nil {
		return TransactionPayment{}, err
	}

	transactionPayment, err := s.Get(ctx, id)
	if err != nil {
		return transactionPayment, err
	}
	transactionPayment.TripId = req.TripId
	transactionPayment.TripMemberId = req.TripMemberId
	transactionPayment.TransactionId = req.TransactionId
	transactionPayment.UserFromId = req.UserFromId
	transactionPayment.UserToId = req.UserToId
	transactionPayment.Nominal = req.Nominal
	transactionPayment.Status = req.Status
	transactionPayment.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, transactionPayment.TransactionPayment); err != nil {
		return transactionPayment, err
	}
	return transactionPayment, nil
}

// Delete deletes the transactionPayment with the specified ID.
func (s service) Delete(ctx context.Context, id string) (TransactionPayment, error) {
	transactionPayment, err := s.Get(ctx, id)
	if err != nil {
		return TransactionPayment{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return TransactionPayment{}, err
	}
	return transactionPayment, nil
}

// Count returns the number of transactionPayments.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the transactionPayments with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]TransactionPayment, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []TransactionPayment{}
	for _, item := range items {
		result = append(result, TransactionPayment{item})
	}
	return result, nil
}

// Query returns the transactions with the specified offset and limit.
func (s service) QueryByTrip(ctx context.Context, tripId string) ([]TransactionPayment, error) {
	items, err := s.repo.QueryByTrip(ctx, tripId)
	if err != nil {
		return nil, err
	}
	result := []TransactionPayment{}
	for _, item := range items {
		result = append(result, TransactionPayment{item})
	}
	return result, nil
}

// Query returns the transactions with the specified offset and limit.
func (s service) QueryByTransaction(ctx context.Context, transactionId string) ([]TransactionPayment, error) {
	items, err := s.repo.QueryByTransaction(ctx, transactionId)
	if err != nil {
		return nil, err
	}
	result := []TransactionPayment{}
	for _, item := range items {
		result = append(result, TransactionPayment{item})
	}
	return result, nil
}
