package service

import (
	"context"
	"database/sql"
)

type Transactor interface {
	NewTxContext(ctx context.Context) context.Context
	InTransaction(ctx context.Context, txFunc func(ctx context.Context) error) error
	RunTransaction(ctx context.Context) error
	Update(context.Context, string, ...any) (int, error)
	GetAll(context.Context, string, ...any) (*sql.Rows, error)
	MigrateSchema() error
}
