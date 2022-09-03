package transactionExpenses

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

	r.Get("/transaction-expenses/<id>", res.get)
	r.Get("/transaction-expenses", res.query)

	r.Use(authHandler)

	r.Post("/transaction-expenses", res.create)
	r.Put("/transaction-expenses/<id>", res.update)
	r.Delete("/transaction-expenses/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	transactionExpenses, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionExpenses)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	transactionExpenses, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = transactionExpenses
	return c.Write(transactionExpenses)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTransactionExpensesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	transactionExpenses, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(transactionExpenses, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTransactionExpensesRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	transactionExpenses, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(transactionExpenses)
}

func (r resource) delete(c *routing.Context) error {
	transactionExpenses, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionExpenses)
}
