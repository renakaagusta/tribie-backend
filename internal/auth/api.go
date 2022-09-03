package auth

import (
	"tribbie/internal/errors"
	"tribbie/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"

	User "tribbie/internal/user"
)

type LoginResponse struct {
	user  User.UserDefault
	token string
}

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, userService User.Service, logger log.Logger) {
	rg.Post("/login", login(service, userService, logger))
	rg.Post("/login/apple", loginByApple(service, userService, logger))
	rg.Post("/login/device", loginByDevice(service, userService, logger))
	rg.Post("/register", register(service, userService, logger))
}

// login returns a handler that handles user login request.
func login(service Service, userService User.Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		user, err := userService.GetByEmail(c.Request.Context(), req.Email)
		if err != nil {
			return err
		}

		token, err := service.Login(c.Request.Context(), user.Email, user.Password)
		if err != nil {
			return err
		}

		return c.Write(struct {
			User  User.UserDefault `json:"user"`
			Token string           `json:"token"`
		}{user, token})
	}
}

// login returns a handler that handles user login request.
func loginByApple(service Service, userService User.Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req struct {
			AppleId string `json:"apple_id"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		user, err := userService.GetByAppleId(c.Request.Context(), req.AppleId)
		if err != nil {
			return err
		}

		token, err := service.Login(c.Request.Context(), user.Email, user.Password)
		if err != nil {
			return err
		}

		return c.Write(struct {
			User  User.UserDefault `json:"user"`
			Token string           `json:"token"`
		}{user, token})
	}
}

// login returns a handler that handles user login request.
func loginByDevice(service Service, userService User.Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req struct {
			DeviceId string `json:"device_id"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		user, err := userService.GetByDeviceId(c.Request.Context(), req.DeviceId)
		if err != nil {
			return err
		}

		token, err := service.Login(c.Request.Context(), user.Email, user.Password)
		if err != nil {
			return err
		}

		return c.Write(struct {
			User  User.UserDefault `json:"user"`
			Token string           `json:"token"`
		}{user, token})
	}
}

// register returns a handler that handles user login request.
func register(service Service, userService User.Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req User.CreateUserRequest

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		trip, err := userService.Create(c.Request.Context(), req)
		if err != nil {
			return err
		}

		return c.Write(trip)
	}
}
