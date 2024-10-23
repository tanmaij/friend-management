package test

import (
	"database/sql"
	"os"
	"testing"
)

// LoadSqlFile reads the SQL file and executes it.
func LoadSqlFile(t *testing.T, tx *sql.DB, sqlFile string) {
	b, err := os.ReadFile(sqlFile)
	if err != nil {
		t.Fatalf("Failed to read SQL file %s: %v", sqlFile, err)
	}

	_, err = tx.Exec(string(b))
	if err != nil {
		t.Fatalf("Failed to execute SQL file %s: %v", sqlFile, err)
	}
}
