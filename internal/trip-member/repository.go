package tripMember

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

// Repository encapsulates the logic to access tripMembers from the data source.
type Repository interface {
	// Get returns the tripMember with the specified tripMember ID.
	Get(ctx context.Context, id string) (entity.TripMember, error)
	// Get returns the tripMember with the specified tripMember ID.
	QueryByTrip(ctx context.Context, tripId string) ([]entity.TripMember, error)
	// Count returns the number of tripMembers.
	Count(ctx context.Context) (int, error)
	// Query returns the list of tripMembers with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TripMember, error)
	// Create saves a new tripMember in the storage.
	Create(ctx context.Context, tripMember entity.TripMember) error
	// Update updates the tripMember with given ID in the storage.
	Update(ctx context.Context, tripMember entity.TripMember) error
	// Delete removes the tripMember with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists tripMembers in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new tripMember repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the tripMember with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TripMember, error) {
	var tripMember entity.TripMember
	err := r.db.With(ctx).Select().Model(id, &tripMember)
	return tripMember, err
}

// Create saves a new tripMember record in the database.
// It returns the ID of the newly inserted tripMember record.
func (r repository) Create(ctx context.Context, tripMember entity.TripMember) error {
	return r.db.With(ctx).Model(&tripMember).Insert()
}

// Update saves the changes to an tripMember in the database.
func (r repository) Update(ctx context.Context, tripMember entity.TripMember) error {
	return r.db.With(ctx).Model(&tripMember).Update()
}

// Delete deletes an tripMember with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	tripMember, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&tripMember).Delete()
}

// Count returns the number of the tripMember records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("trip_member").Row(&count)
	return count, err
}

// Query retrieves the tripMember records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TripMember, error) {
	var tripMembers []entity.TripMember
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&tripMembers)
	return tripMembers, err
}

// Get reads the tripMember with the specified Trip ID from the database.
func (r repository) QueryByTrip(ctx context.Context, tripId string) ([]entity.TripMember, error) {
	var tripMembers []entity.TripMember

	q := r.db.DB().NewQuery("SELECT * FROM trip_member WHERE trip_id='" + tripId + "'")
	err := q.All(&tripMembers)

	return tripMembers, err
}
