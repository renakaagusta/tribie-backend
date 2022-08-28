package transactionExpenses

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access transactionExpenses from the data source.
type Repository interface {
	// Get returns the transactionExpenses with the specified transactionExpenses ID.
	Get(ctx context.Context, id string) (entity.TransactionExpenses, error)
	// Count returns the number of transactionExpenses.
	Count(ctx context.Context) (int, error)
	// Query returns the list of transactionExpenses with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TransactionExpenses, error)
	// Query returns the list of transactionExpenses with the given offset and limit.
	QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionExpenses, error)
	// Query returns the list of transactionExpenses with the given offset and limit.
	QueryByTransaction(ctx context.Context, tripId string) ([]entity.TransactionExpenses, error)
	// Create saves a new transactionExpenses in the storage.
	Create(ctx context.Context, transactionExpenses entity.TransactionExpenses) error
	// Update updates the transactionExpenses with given ID in the storage.
	Update(ctx context.Context, transactionExpenses entity.TransactionExpenses) error
	// Delete removes the transactionExpenses with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists transactionExpenses in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new transactionExpenses repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the transactionExpenses with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TransactionExpenses, error) {
	var transactionExpenses entity.TransactionExpenses
	err := r.db.With(ctx).Select().Model(id, &transactionExpenses)
	return transactionExpenses, err
}

// Create saves a new transactionExpenses record in the database.
// It returns the ID of the newly inserted transactionExpenses record.
func (r repository) Create(ctx context.Context, transactionExpenses entity.TransactionExpenses) error {
	return r.db.With(ctx).Model(&transactionExpenses).Insert()
}

// Update saves the changes to an transactionExpenses in the database.
func (r repository) Update(ctx context.Context, transactionExpenses entity.TransactionExpenses) error {
	return r.db.With(ctx).Model(&transactionExpenses).Update()
}

// Delete deletes an transactionExpenses with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	transactionExpenses, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&transactionExpenses).Delete()
}

// Count returns the number of the transactionExpenses records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("transaction_expenses").Row(&count)
	return count, err
}

// Query retrieves the transactionExpenses records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TransactionExpenses, error) {
	var transactionExpenses []entity.TransactionExpenses
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&transactionExpenses)
	return transactionExpenses, err
}

// Get reads the TransactionExpenses with the specified Trip ID from the database.
func (r repository) QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionExpenses, error) {
	var TransactionExpenses []entity.TransactionExpenses

	q := r.db.DB().NewQuery("SELECT * FROM transaction_expenses WHERE trip_id='" + tripId + "'")
	err := q.All(&TransactionExpenses)

	return TransactionExpenses, err
}

// Get reads the TransactionExpenses with the specified Trip ID from the database.
func (r repository) QueryByTransaction(ctx context.Context, tripId string) ([]entity.TransactionExpenses, error) {
	var TransactionExpenses []entity.TransactionExpenses

	q := r.db.DB().NewQuery("SELECT * FROM transaction_expenses WHERE transaction_id='" + tripId + "'")
	err := q.All(&TransactionExpenses)

	return TransactionExpenses, err
}
