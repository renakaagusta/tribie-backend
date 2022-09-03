package user

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(ctx context.Context, id string) (UserDefault, error)
	GetByEmail(ctx context.Context, email string) (UserDefault, error)
	GetByAppleId(ctx context.Context, appleId string) (UserDefault, error)
	GetByDeviceId(ctx context.Context, deviceId string) (UserDefault, error)
	Query(ctx context.Context, offset, limit int) ([]UserDefault, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateUserRequest) (UserDefault, error)
	Update(ctx context.Context, id string, input UpdateUserRequest) (UserDefault, error)
	Delete(ctx context.Context, id string) (UserDefault, error)
}

// User represents the data about an user.
type UserDefault struct {
	entity.UserDefault
}

// CreateUserRequest represents an user creation request.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	AppleId  string `json:"apple_id"`
	DeviceId string `json:"device_id"`
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m)
}

// UpdateUserRequest represents an user update request.
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"UserDefault{"`
	Password string `json:"password"`
	AppleId  string `json:"apple_id"`
	DeviceId string `json:"device_id"`
}

// Validate validates the CreateUserRequest fields.
func (m UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&m)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the user with the specified the user ID.
func (s service) Get(ctx context.Context, id string) (UserDefault, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return UserDefault{}, err
	}
	return UserDefault{user}, nil
}

// Get returns the user with the specified the user ID.
func (s service) GetByEmail(ctx context.Context, email string) (UserDefault, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return UserDefault{}, err
	}
	return UserDefault{user}, nil
}

// Get returns the user with the specified the user ID.
func (s service) GetByAppleId(ctx context.Context, appleId string) (UserDefault, error) {
	user, err := s.repo.GetByAppleId(ctx, appleId)
	if err != nil {
		return UserDefault{}, err
	}
	return UserDefault{user}, nil
}

// Get returns the user with the specified the user ID.
func (s service) GetByDeviceId(ctx context.Context, deviceId string) (UserDefault, error) {
	user, err := s.repo.GetByDeviceId(ctx, deviceId)
	if err != nil {
		return UserDefault{}, err
	}
	return UserDefault{user}, nil
}

// Create creates a new user.
func (s service) Create(ctx context.Context, req CreateUserRequest) (UserDefault, error) {
	if err := req.Validate(); err != nil {
		return UserDefault{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.UserDefault{
		ID:        id,
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
		AppleId:   req.AppleId,
		DeviceId:  req.DeviceId,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return UserDefault{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the user with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateUserRequest) (UserDefault, error) {
	if err := req.Validate(); err != nil {
		return UserDefault{}, err
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		return user, err
	}
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password
	user.AppleId = req.AppleId
	user.DeviceId = req.DeviceId
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user.UserDefault); err != nil {
		return user, err
	}
	return user, nil
}

// Delete deletes the user with the specified ID.
func (s service) Delete(ctx context.Context, id string) (UserDefault, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return UserDefault{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return UserDefault{}, err
	}
	return user, nil
}

// Count returns the number of users.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the users with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]UserDefault, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []UserDefault{}
	for _, item := range items {
		result = append(result, UserDefault{item})
	}
	return result, nil
}
