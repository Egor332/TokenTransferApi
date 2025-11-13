package testutil

import (
	"context"
	"os"
	"testing"

	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) (*gorm.DB, func()) {
	ctx := context.Background()

	pgContainer, err := tcpostgres.RunContainer(ctx,
		tcpostgres.WithDatabase("testdb"),
		tcpostgres.WithUsername("testuser"),
		tcpostgres.WithPassword("testpass"),
	)
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	db, err := gorm.Open(gormpostgres.Open(connStr), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	schema, err := os.ReadFile("test/TestDatabaseInit.sql")
	if err != nil {
		t.Fatalf("failed to read schema: %v", err)
	}

	err = db.Exec(string(schema)).Error
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	cleanup := func() {
		_ = pgContainer.Terminate(ctx)
	}

	return db, cleanup
}
