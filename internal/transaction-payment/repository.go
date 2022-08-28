package transactionPayment

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access transactionPayments from the data source.
type Repository interface {
	// Get returns the transactionPayment with the specified transactionPayment ID.
	Get(ctx context.Context, id string) (entity.TransactionPayment, error)
	// Count returns the number of transactionPayments.
	Count(ctx context.Context) (int, error)
	// Query returns the list of transactionPayments with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TransactionPayment, error)
	// Query returns the list of transactionPayments with the given offset and limit.
	QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionPayment, error)
	// Query returns the list of transactionPayments with the given offset and limit.
	QueryByTransaction(ctx context.Context, tripId string) ([]entity.TransactionPayment, error)
	// Create saves a new transactionPayment in the storage.
	Create(ctx context.Context, transactionPayment entity.TransactionPayment) error
	// Update updates the transactionPayment with given ID in the storage.
	Update(ctx context.Context, transactionPayment entity.TransactionPayment) error
	// Delete removes the transactionPayment with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists transactionPayments in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new transactionPayment repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the transactionPayment with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TransactionPayment, error) {
	var transactionPayment entity.TransactionPayment
	err := r.db.With(ctx).Select().Model(id, &transactionPayment)
	return transactionPayment, err
}

// Create saves a new transactionPayment record in the database.
// It returns the ID of the newly inserted transactionPayment record.
func (r repository) Create(ctx context.Context, transactionPayment entity.TransactionPayment) error {
	return r.db.With(ctx).Model(&transactionPayment).Insert()
}

// Update saves the changes to an transactionPayment in the database.
func (r repository) Update(ctx context.Context, transactionPayment entity.TransactionPayment) error {
	return r.db.With(ctx).Model(&transactionPayment).Update()
}

// Delete deletes an transactionPayment with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	transactionPayment, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&transactionPayment).Delete()
}

// Count returns the number of the transactionPayment records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("transaction_payment").Row(&count)
	return count, err
}

// Query retrieves the transactionPayment records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TransactionPayment, error) {
	var transactionPayments []entity.TransactionPayment
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&transactionPayments)
	return transactionPayments, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionPayment, error) {
	var tripMembers []entity.TransactionPayment

	q := r.db.DB().NewQuery("SELECT * FROM transaction_payment WHERE trip_id='" + tripId + "'")
	err := q.All(&tripMembers)

	return tripMembers, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTransaction(ctx context.Context, transactionId string) ([]entity.TransactionPayment, error) {
	var tripMembers []entity.TransactionPayment

	q := r.db.DB().NewQuery("SELECT * FROM transaction_payment WHERE transaction_id='" + transactionId + "'")
	err := q.All(&tripMembers)

	return tripMembers, err
}
