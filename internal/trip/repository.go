package trip

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access trips from the data source.
type Repository interface {
	// Get returns the trip with the specified trip ID.
	Get(ctx context.Context, id string) (entity.Trip, error)
	// Count returns the number of trips.
	Count(ctx context.Context) (int, error)
	// Query returns the list of trips with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Trip, error)
	// Create saves a new trip in the storage.
	Create(ctx context.Context, trip entity.Trip) error
	// Update updates the trip with given ID in the storage.
	Update(ctx context.Context, trip entity.Trip) error
	// Delete removes the trip with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists trips in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new trip repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the trip with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Trip, error) {
	var trip entity.Trip
	err := r.db.With(ctx).Select().Model(id, &trip)
	return trip, err
}

// Create saves a new trip record in the database.
// It returns the ID of the newly inserted trip record.
func (r repository) Create(ctx context.Context, trip entity.Trip) error {
	return r.db.With(ctx).Model(&trip).Insert()
}

// Update saves the changes to an trip in the database.
func (r repository) Update(ctx context.Context, trip entity.Trip) error {
	return r.db.With(ctx).Model(&trip).Update()
}

// Delete deletes an trip with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	trip, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&trip).Delete()
}

// Count returns the number of the trip records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("trip").Row(&count)
	return count, err
}

// Query retrieves the trip records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Trip, error) {
	var trips []entity.Trip
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&trips)
	return trips, err
}
