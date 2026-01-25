package repository

import (
	"context"
	"sistem-klinik-backend/entity"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *entity.Tenant) error
	FindByID(ctx context.Context, id string) (*entity.Tenant, error)
	FindAll(ctx context.Context) ([]*entity.Tenant, error)
}
