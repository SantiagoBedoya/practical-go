package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/SantiagoBedoya/gobank-users-api/users"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
)

const (
	createUserQuery      = "INSERT INTO users(first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	findUserByIDQuery    = "SELECT id, first_name, last_name, email, password FROM users WHERE id = $1"
	findUserByEmailQuery = "SELECT id, first_name, last_name, email, password FROM users WHERE email = $1"
)

type repository struct {
	db      *sql.DB
	timeout time.Duration
}

// NewPostgreSQLRepository return postgreSQL users.Repository implementation
func NewPostgreSQLRepository(db *sql.DB, timeout time.Duration) users.Repository {
	return &repository{
		db:      db,
		timeout: timeout,
	}
}

func (r *repository) Create(data *users.User) (*users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return nil, errors.Wrap(err, "postgreSQLRepository.Create")
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(ctx, data.FirstName, data.LastName, data.Email, data.Password).Scan(&data.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return nil, errors.Wrap(users.ErrUserAlreadyExists, "postgreSQLRepository.Create")
		}
		return nil, errors.Wrap(err, "postgreSQLRepository.Create")
	}
	return data, nil
}

func (r *repository) FindByEmail(data *users.User) (*users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(findUserByEmailQuery)
	if err != nil {
		return nil, errors.Wrap(err, "postgreSQLRepository.FindByID")
	}
	defer stmt.Close()
	var result users.User
	if err := stmt.QueryRowContext(ctx, data.Email).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(users.ErrUserNotFound, "postgreSQLRepository.FindByID")
		}
		return nil, errors.Wrap(err, "postgreSQLRepository.FindByID")
	}
	return &result, nil
}

func (r *repository) FindByID(userID string) (*users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(findUserByIDQuery)
	if err != nil {
		return nil, errors.Wrap(err, "postgreSQLRepository.FindByID")
	}
	defer stmt.Close()
	var result users.User
	if err := stmt.QueryRowContext(ctx, userID).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Email, &result.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(users.ErrUserNotFound, "postgreSQLRepository.FindByID")
		}
		return nil, errors.Wrap(err, "postgreSQLRepository.FindByID")
	}
	return &result, nil
}
