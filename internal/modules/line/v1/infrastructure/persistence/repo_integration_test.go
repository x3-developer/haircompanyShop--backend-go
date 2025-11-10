//go:build integration
// +build integration

package persistence

import (
	"context"
	"serv_shop_haircompany/internal/modules/line/v1/domain"
	"serv_shop_haircompany/internal/shared/infrastructure/persistence"
	"testing"
)

func setupDB(t *testing.T) *persistence.Postgres {
	t.Helper()

	dsn := "postgres://test:test@localhost:5433/testdb?sslmode=disable"
	db, err := persistence.NewPostgresTest(dsn)
	if err != nil {
		t.Fatalf("failed to connect test db: %v", err)
	}
	return db
}

func TestRepo_Create_And_Exists(t *testing.T) {
	db := setupDB(t)
	repo := NewRepo(db)

	ctx := context.Background()

	model := &domain.Line{
		Name:  "TestLine",
		Color: "#112233",
	}

	created, err := repo.Create(ctx, model)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	if created.ID == 0 {
		t.Fatalf("expected ID > 0, got %d", created.ID)
	}

	exists, err := repo.ExistsByUniqueFields(ctx, created.Name)
	if err != nil {
		t.Fatalf("exists failed: %v", err)
	}
	if !exists {
		t.Fatalf("expected exists = true, got false")
	}
}
