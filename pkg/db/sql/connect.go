package sql

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// DriverName represents sql driver names
type DriverName string

const (
	Postgres DriverName = "postgres"
)

// ToString converts a DriverName to a string
func (driverName DriverName) ToString() string {
	return string(driverName)
}

// ConnectionOption holds configuration options for database connections
type ConnectionOption struct {
	MaxOpenConnections int // Maximum number of open connections
	MaxIdleConnections int // Maximum number of idle connections
	ConnMaxLifetime    int // Maximum lifetime for a connection in seconds
	ConnMaxIdleTime    int // Maximum idle time for a connection in seconds
}

// ConnectDB returns the instance of sql.DB which represents a pool of zero or more underlying connections
func ConnectDB(driverName DriverName, dbUrl string, options ConnectionOption) (*sql.DB, error) {
	if dbUrl == "" {
		return nil, ErrDBUrlIsEmpty
	}

	db, err := sql.Open(driverName.ToString(), dbUrl)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Set connection pool parameters using the provided options
	db.SetMaxOpenConns(options.MaxOpenConnections) // Set maximum number of open connections
	db.SetMaxIdleConns(options.MaxIdleConnections) // Set maximum number of idle connections

	// Set connection max lifetime and idle time if they are greater than zero
	if options.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(options.ConnMaxLifetime) * time.Second)
	}
	if options.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(options.ConnMaxIdleTime) * time.Second)
	}

	return db, nil
}
