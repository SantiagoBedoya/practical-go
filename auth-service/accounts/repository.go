package accounts

import (
	"context"
	"database/sql"
)

// Repository define interface for repositories
type Repository interface {
	GetConnection(context.Context) (*sql.DB, error)
	CloseConnection(context.Context) error
	FindByEmail(context.Context, string) (string, string, error)
	Create(context.Context, string, string) (string, error)
}
