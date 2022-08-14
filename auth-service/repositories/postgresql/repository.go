package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SantiagoBedoya/auth-service/accounts"
	"github.com/go-kit/kit/log"

	// It is the library for pgx
	"github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	findByEmailQuery = `SELECT id, password FROM accounts WHERE email = $1`
	createQuery      = `INSERT INTO accounts(email, password) VALUES ($1, $2) RETURNING id`
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func newPostgreSQLClient(ctx context.Context, dataSource string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

// NewRepository creates and implements an instance of accounts.Repository
func NewRepository(ctx context.Context, dataSource string, logger log.Logger) accounts.Repository {
	db, err := newPostgreSQLClient(ctx, dataSource)
	if err != nil {
		panic(err)
	}
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r repository) GetConnection(_ context.Context) (*sql.DB, error) {
	return r.db, nil
}
func (r repository) CloseConnection(_ context.Context) error {
	r.logger.Log("method", "CloseConnection", "message", "closing connection")
	return r.db.Close()
}
func (r repository) FindByEmail(ctx context.Context, email string) (string, string, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	stmt, err := r.db.Prepare(findByEmailQuery)
	if err != nil {
		r.logger.Log("method", "FindByEmail", "err", err.Error())
		return "", "", err
	}
	defer stmt.Close()
	var userID, userPassword string
	if err := stmt.QueryRowContext(c, email).Scan(&userID, &userPassword); err != nil {
		r.logger.Log("method", "FindByEmail", "err", err.Error())
		return "", "", err
	}
	return userID, userPassword, nil
}
func (r repository) Create(ctx context.Context, email string, password string) (string, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	stmt, err := r.db.Prepare(createQuery)
	if err != nil {
		r.logger.Log("method", "Create", "err", err.Error())
		return "", err
	}
	defer stmt.Close()
	var id string
	if err := stmt.QueryRowContext(c, email, password).Scan(&id); err != nil {
		r.logger.Log("method", "Create", "err", err.Error())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", accounts.ErrUserAlreadyExists
		}
		return "", err
	}
	return id, nil
}
