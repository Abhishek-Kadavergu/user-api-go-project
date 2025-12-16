package repository

import (
	"context"
	"time"

	db "user-api/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(pool),
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob string,
) (db.User, error) {

	parsedTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	pgDate := pgtype.Date{
		Time:  parsedTime,
		Valid: true,
	}

	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  pgDate,
	})
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) List(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}

func (r *UserRepository) Update(
	ctx context.Context,
	id int32,
	name string,
	dob string,
) (db.User, error) {

	parsedTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	pgDate := pgtype.Date{
		Time:  parsedTime,
		Valid: true,
	}

	return r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  pgDate,
	})
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
