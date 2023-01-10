package repository

import (
	"context"
	"database/sql"

	"github.com/SantiagoBedoya/auth-ms/internal/config"
	"github.com/SantiagoBedoya/auth-ms/internal/model"
	"github.com/SantiagoBedoya/auth-ms/internal/service"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	storeQuery       = "INSERT INTO users(username, email, password) VALUES ($1, $2, $3)"
	findByEmailQuery = "SELECT id, username, email, password FROM users WHERE email = $1"
)

type postgresRepo struct {
	client *sql.DB
}

func newPostgresClient(uri string) (*sql.DB, error) {
	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgresRepo(cfg *config.Config) (service.Repository, error) {
	client, err := newPostgresClient(cfg.DbURI)
	if err != nil {
		return nil, err
	}
	return &postgresRepo{client}, nil
}

func (r postgresRepo) Store(ctx context.Context, doc *model.User) error {
	stmt, err := r.client.PrepareContext(ctx, storeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, doc.Username, doc.Email, doc.Password)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			if pgErr.Code == "23505" {
				return service.ErrEmailUsernameInUse
			}
		}
		return err
	}
	return nil
}

func (r postgresRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	stmt, err := r.client.PrepareContext(ctx, findByEmailQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var result model.User
	if err := stmt.QueryRowContext(ctx, email).Scan(&result.ID, &result.Username, &result.Email, &result.Password); err != nil {
		return nil, err
	}
	return &result, nil
}
