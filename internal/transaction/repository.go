package transaction

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access transactions from the data source.
type Repository interface {
	// Get returns the transaction with the specified transaction ID.
	Get(ctx context.Context, id string) (entity.Transaction, error)
	// Count returns the number of transactions.
	Count(ctx context.Context) (int, error)
	// Query returns the list of transactions with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Transaction, error)
	// Query returns the list of transactions with the given offset and limit.
	QueryByTrip(ctx context.Context, tripId string) ([]entity.Transaction, error)
	// Create saves a new transaction in the storage.
	Create(ctx context.Context, transaction entity.Transaction) error
	// Update updates the transaction with given ID in the storage.
	Update(ctx context.Context, transaction entity.Transaction) error
	// Delete removes the transaction with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists transactions in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new transaction repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the transaction with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.With(ctx).Select().Model(id, &transaction)
	return transaction, err
}

// Create saves a new transaction record in the database.
// It returns the ID of the newly inserted transaction record.
func (r repository) Create(ctx context.Context, transaction entity.Transaction) error {
	return r.db.With(ctx).Model(&transaction).Insert()
}

// Update saves the changes to an transaction in the database.
func (r repository) Update(ctx context.Context, transaction entity.Transaction) error {
	return r.db.With(ctx).Model(&transaction).Update()
}

// Delete deletes an transaction with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	transaction, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&transaction).Delete()
}

// Count returns the number of the transaction records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("transaction").Row(&count)
	return count, err
}

// Query retrieves the transaction records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&transactions)
	return transactions, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTrip(ctx context.Context, tripId string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	q := r.db.DB().NewQuery("SELECT * FROM transaction WHERE trip_id='" + tripId + "'")
	err := q.All(&transactions)

	return transactions, err
}
