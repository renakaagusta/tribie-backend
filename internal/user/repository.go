package user

import (
	"context"
	"tribbie/internal/entity"
	"tribbie/pkg/dbcontext"
	"tribbie/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.UserDefault, error)
	GetByEmail(ctx context.Context, email string) (entity.UserDefault, error)
	GetByAppleId(ctx context.Context, appleId string) (entity.UserDefault, error)
	GetByDeviceId(ctx context.Context, deviceId string) (entity.UserDefault, error)
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]entity.UserDefault, error)
	Create(ctx context.Context, user entity.UserDefault) error
	Update(ctx context.Context, user entity.UserDefault) error
	Delete(ctx context.Context, id string) error
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.UserDefault, error) {
	var user entity.UserDefault
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) GetByEmail(ctx context.Context, email string) (entity.UserDefault, error) {
	var user entity.UserDefault

	q := r.db.DB().NewQuery("SELECT * FROM user_default WHERE email='" + email + "'")
	err := q.One(&user)

	return user, err
}

func (r repository) GetByAppleId(ctx context.Context, appleId string) (entity.UserDefault, error) {
	var user entity.UserDefault

	q := r.db.DB().NewQuery("SELECT * FROM user_default WHERE apple_id='" + appleId + "'")
	err := q.One(&user)

	return user, err
}

func (r repository) GetByDeviceId(ctx context.Context, deviceId string) (entity.UserDefault, error) {
	var user entity.UserDefault

	q := r.db.DB().NewQuery("SELECT * FROM user_default WHERE device_id='" + deviceId + "'")
	err := q.One(&user)

	return user, err
}

func (r repository) Create(ctx context.Context, user entity.UserDefault) error {
	return r.db.With(ctx).Model(&user).Insert()
}

func (r repository) Update(ctx context.Context, user entity.UserDefault) error {
	return r.db.With(ctx).Model(&user).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	user, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("user_default").Row(&count)
	return count, err
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.UserDefault, error) {
	var users []entity.UserDefault

	q := r.db.DB().NewQuery("SELECT * FROM user_default")
	err := q.All(&users)

	return users, err
}
