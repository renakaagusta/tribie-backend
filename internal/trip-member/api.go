package tripMember

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"tribbie/internal/errors"
	"tribbie/pkg/log"
	"tribbie/pkg/pagination"
	"net/http"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/trip-members/<id>", res.get)
	r.Get("/trip-members", res.query)

	r.Use(authHandler)

	r.Post("/trip-members", res.create)
	r.Put("/trip-members/<id>", res.update)
	r.Delete("/trip-members/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	tripMember, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(tripMember)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	tripMembers, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = tripMembers
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTripMemberRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	tripMember, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(tripMember, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTripMemberRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	tripMember, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(tripMember)
}

func (r resource) delete(c *routing.Context) error {
	tripMember, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(tripMember)
}
