package trip

import (
	"net/http"
	"tribbie/internal/errors"
	"tribbie/pkg/log"
	"tribbie/pkg/pagination"

	routing "github.com/go-ozzo/ozzo-routing/v2"

	Transaction "tribbie/internal/transaction"
	TransactionExpenses "tribbie/internal/transaction-expenses"
	TransactionItem "tribbie/internal/transaction-item"
	TransactionPayment "tribbie/internal/transaction-payment"
	TripMember "tribbie/internal/trip-member"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(
	r *routing.RouteGroup,
	service Service,
	tripMemberService TripMember.Service,
	transactionService Transaction.Service,
	transactionItemService TransactionItem.Service,
	transactionExpenseservice TransactionExpenses.Service,
	transactionPaymentService TransactionPayment.Service,
	authHandler routing.Handler,
	logger log.Logger) {
	res := resource{service, tripMemberService, transactionService, transactionItemService, transactionExpenseservice, transactionPaymentService, logger}

	r.Get("/trips/<id>", res.get)
	r.Get("/trips/<id>/trip-members", res.queryMemberList)
	r.Get("/trips/<id>/transactions", res.queryTransactionList)
	r.Get("/trips/<id>/transaction-items", res.queryTransactionItemList)
	r.Get("/trips/<id>/transaction-expenses", res.queryTransactionExpensesList)
	r.Get("/trips/<id>/transaction-payments", res.queryTransactionPaymentList)
	r.Get("/trips", res.query)
	r.Post("/trips", res.create)
	r.Put("/trips/<id>", res.update)
	r.Delete("/trips/<id>", res.delete)
}

type resource struct {
	service                   Service
	tripMemberService         TripMember.Service
	transactionService        Transaction.Service
	transactionItemService    TransactionItem.Service
	TransactionExpenseservice TransactionExpenses.Service
	transactionPaymentService TransactionPayment.Service
	logger                    log.Logger
}

func (r resource) get(c *routing.Context) error {
	trip, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryMemberList(c *routing.Context) error {
	trip, err := r.tripMemberService.QueryByTrip(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryTransactionList(c *routing.Context) error {
	trip, err := r.transactionService.QueryByTrip(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryTransactionItemList(c *routing.Context) error {
	trip, err := r.transactionItemService.QueryByTrip(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryTransactionExpensesList(c *routing.Context) error {
	trip, err := r.TransactionExpenseservice.QueryByTrip(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) queryTransactionPaymentList(c *routing.Context) error {
	trip, err := r.transactionPaymentService.QueryByTrip(c.Request.Context(), c.Param("id"))
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
	trips, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = trips
	return c.Write(trips)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTripRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	trip, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(trip, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTripRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	trip, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(trip)
}

func (r resource) delete(c *routing.Context) error {
	trip, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(trip)
}
