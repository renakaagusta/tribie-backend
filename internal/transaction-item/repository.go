package transactionItem

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access transactionItems from the data source.
type Repository interface {
	// Get returns the transactionItem with the specified transactionItem ID.
	Get(ctx context.Context, id string) (entity.TransactionItem, error)
	// Count returns the number of transactionItems.
	Count(ctx context.Context) (int, error)
	// Query returns the list of transactionItems with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TransactionItem, error)
	// Query returns the list of transactionItems with the given offset and limit.
	QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionItem, error)
	// Query returns the list of transactionItems with the given offset and limit.
	QueryByTransaction(ctx context.Context, transactionId string) ([]entity.TransactionItem, error)
	// Create saves a new transactionItem in the storage.
	Create(ctx context.Context, transactionItem entity.TransactionItem) error
	// Update updates the transactionItem with given ID in the storage.
	Update(ctx context.Context, transactionItem entity.TransactionItem) error
	// Delete removes the transactionItem with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists transactionItems in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new transactionItem repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the transactionItem with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TransactionItem, error) {
	var transactionItem entity.TransactionItem
	err := r.db.With(ctx).Select().Model(id, &transactionItem)
	return transactionItem, err
}

// Create saves a new transactionItem record in the database.
// It returns the ID of the newly inserted transactionItem record.
func (r repository) Create(ctx context.Context, transactionItem entity.TransactionItem) error {
	return r.db.With(ctx).Model(&transactionItem).Insert()
}

// Update saves the changes to an transactionItem in the database.
func (r repository) Update(ctx context.Context, transactionItem entity.TransactionItem) error {
	return r.db.With(ctx).Model(&transactionItem).Update()
}

// Delete deletes an transactionItem with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	transactionItem, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&transactionItem).Delete()
}

// Count returns the number of the transactionItem records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("transaction_item").Row(&count)
	return count, err
}

// Query retrieves the transactionItem records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TransactionItem, error) {
	var transactionItems []entity.TransactionItem
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&transactionItems)
	return transactionItems, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTrip(ctx context.Context, tripId string) ([]entity.TransactionItem, error) {
	var transactionItem []entity.TransactionItem

	q := r.db.DB().NewQuery("SELECT * FROM transaction_item WHERE trip_id='" + tripId + "'")
	err := q.All(&transactionItem)

	return transactionItem, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTransaction(ctx context.Context, transactionId string) ([]entity.TransactionItem, error) {
	var tripMembers []entity.TransactionItem

	q := r.db.DB().NewQuery("SELECT * FROM transaction_item WHERE transaction_id='" + transactionId + "'")
	err := q.All(&tripMembers)

	return tripMembers, err
}
