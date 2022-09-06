package transactionItem

import (
	"net/http"
	"tribbie/internal/errors"
	"tribbie/pkg/log"
	"tribbie/pkg/pagination"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/transaction-items/<id>", res.get)
	r.Get("/transaction-items", res.query)

	// r.Use(authHandler)

	r.Post("/transaction-items", res.create)
	r.Put("/transaction-items/<id>", res.update)
	r.Delete("/transaction-items/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	transactionItem, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionItem)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	transactionItems, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = transactionItems
	return c.Write(transactionItems)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTransactionItemRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	transactionItem, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(transactionItem, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTransactionItemRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	transactionItem, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(transactionItem)
}

func (r resource) delete(c *routing.Context) error {
	transactionItem, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionItem)
}
