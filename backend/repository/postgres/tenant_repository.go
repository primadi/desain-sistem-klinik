package postgres

import (
	"context"
	"errors"
	"fmt"
	"sistem-klinik-backend/entity"

	"github.com/primadi/lokstra/serviceapi"
)

// @Service "postgres-tenant-repository"
type PostgresTenantRepository struct {
	// @Inject "db_main"
	dbPool serviceapi.DbPool
}

func (s *PostgresTenantRepository) Create(ctx context.Context, tenant *entity.Tenant) error {
	fmt.Printf("Simulating DB Create: %v\n", tenant)
	return nil
}

func (s *PostgresTenantRepository) FindByID(ctx context.Context, id string) (*entity.Tenant, error) {
	if id == "not-found" {
		return nil, errors.New("tenant not found")
	}
	return &entity.Tenant{
		ID:      id,
		Name:    "Sample Clinic",
		Address: "Jl. Kesehatan No. 1",
	}, nil
}

func (s *PostgresTenantRepository) FindAll(ctx context.Context) ([]*entity.Tenant, error) {
	return []*entity.Tenant{
		{ID: "1", Name: "Clinic A", Address: "Addr A"},
		{ID: "2", Name: "Clinic B", Address: "Addr B"},
	}, nil
}
