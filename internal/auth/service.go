package auth

import (
	"context"
	"time"
	"tribbie/internal/entity"
	"tribbie/pkg/log"

	"github.com/dgrijalva/jwt-go"
)

// Service encapsulates the authentication logic.
type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type RegisterRequest struct {
	Email       string `json:"email"`
	Username    string `json:"email"`
	Password    string `json:"password"`
	AppleId     string `json:"apple_id"`
	DeviceId    string `json:"device_id"`
	Description string `json:"description"`
	UserPaidId  string `json:"user_paid_id"`
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetEmail() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{signingKey, tokenExpiration, logger}
}

func (s service) Login(ctx context.Context, email, password string) (string, error) {
	identity := s.authenticate(ctx, email, password)
	return s.generateJWT(identity)
}

func (s service) authenticate(ctx context.Context, email, password string) Identity {
	return entity.UserDefault{ID: email, Email: email}
}

func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    identity.GetID(),
		"email": identity.GetEmail(),
		"exp":   time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
