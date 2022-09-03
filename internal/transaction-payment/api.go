package transactionPayment

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

	r.Get("/transaction-payments/<id>", res.get)
	r.Get("/transaction-payments", res.query)

	r.Use(authHandler)

	r.Post("/transaction-payments", res.create)
	r.Put("/transaction-payments/<id>", res.update)
	r.Delete("/transaction-payments/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	transactionPayment, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionPayment)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	transactionPayments, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = transactionPayments
	return c.Write(transactionPayments)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTransactionPaymentRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	transactionPayment, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(transactionPayment, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTransactionPaymentRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	transactionPayment, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(transactionPayment)
}

func (r resource) delete(c *routing.Context) error {
	transactionPayment, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transactionPayment)
}
