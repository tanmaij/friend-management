package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// Registry represents a Data Access Layer which return the repositories
type Registry interface {
	// Ping returns the friendship interface
	Ping(ctx context.Context) error
}

// New return a new Registry
func New(sqlDB *sql.DB) Registry {
	return &impl{sqlDB: sqlDB}
}

type impl struct {
	sqlDB *sql.DB
}

// Ping ping to the database connection
func (r impl) Ping(ctx context.Context) error {
	if err := r.sqlDB.PingContext(ctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
