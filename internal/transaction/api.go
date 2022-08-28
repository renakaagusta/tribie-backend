package transaction

import (
	"net/http"
	"tribbie/internal/errors"
	"tribbie/pkg/log"
	"tribbie/pkg/pagination"

	routing "github.com/go-ozzo/ozzo-routing/v2"

	TransactionExpenses "tribbie/internal/transaction-expenses"
	TransactionItem "tribbie/internal/transaction-item"
	TransactionPayment "tribbie/internal/transaction-payment"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(
	r *routing.RouteGroup,
	service Service,
	transactionItemService TransactionItem.Service,
	transactionPaymentService TransactionPayment.Service,
	transactionExpensesService TransactionExpenses.Service,
	authHandler routing.Handler,
	logger log.Logger) {
	res := resource{service, transactionItemService, transactionPaymentService, transactionExpensesService, logger}

	r.Get("/transactions/<id>", res.get)
	r.Get("/transactions", res.query)
	r.Get("/transactions/<id>/transaction-items", res.queryItemList)
	r.Get("/transactions/<id>/transaction-expenses", res.queryExpensesList)
	r.Get("/transactions/<id>/transaction-payments", res.queryPaymentList)

	r.Use(authHandler)

	r.Post("/transactions", res.create)
	r.Put("/transactions/<id>", res.update)
	r.Delete("/transactions/<id>", res.delete)
}

type resource struct {
	service                    Service
	transactionItemService     TransactionItem.Service
	transactionPaymentService  TransactionPayment.Service
	transactionExpensesService TransactionExpenses.Service
	logger                     log.Logger
}

func (r resource) get(c *routing.Context) error {
	transaction, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transaction)
}

func (r resource) queryItemList(c *routing.Context) error {
	trip, err := r.transactionItemService.QueryByTransaction(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryExpensesList(c *routing.Context) error {
	trip, err := r.transactionExpensesService.QueryByTransaction(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryPaymentList(c *routing.Context) error {
	trip, err := r.transactionPaymentService.QueryByTransaction(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	transactions, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = transactions
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTransactionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	transaction, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(transaction, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTransactionRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	transaction, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(transaction)
}

func (r resource) delete(c *routing.Context) error {
	transaction, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(transaction)
}

func (r resource) queryTransactionItemList(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	transactions, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = transactions
	return c.Write(pages)
}
