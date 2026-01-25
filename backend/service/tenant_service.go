package service

import (
	"sistem-klinik-backend/dto"
	"sistem-klinik-backend/entity"
	"sistem-klinik-backend/repository"

	"github.com/google/uuid"
	"github.com/primadi/lokstra"
)

// @EndpointService "tenant-service", "${api-tenant-prefix:/api/tenants", ["recovery", "request_logger"]
type TenantService struct {
	// @Inject "@store.tenant-repository"
	TenantRepo repository.TenantRepository
}

// @Route "POST /"
func (s *TenantService) Create(c *lokstra.RequestContext, req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	tenant := &entity.Tenant{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Address: req.Address,
	}

	if err := s.TenantRepo.Create(c, tenant); err != nil {
		return nil, err
	}

	return &dto.TenantResponse{
		ID:      tenant.ID,
		Name:    tenant.Name,
		Address: tenant.Address,
	}, nil
}

// @Route "GET /:id"
func (s *TenantService) Get(c *lokstra.RequestContext, req *dto.GetTenantRequest) (*dto.TenantResponse, error) {
	tenant, err := s.TenantRepo.FindByID(c, req.ID)
	if err != nil {
		return nil, err
	}

	return &dto.TenantResponse{
		ID:      tenant.ID,
		Name:    tenant.Name,
		Address: tenant.Address,
	}, nil
}

// @Route "GET /"
func (s *TenantService) List(c *lokstra.RequestContext) ([]*dto.TenantResponse, error) {
	tenants, err := s.TenantRepo.FindAll(c)
	if err != nil {
		return nil, err
	}

	var response []*dto.TenantResponse
	for _, t := range tenants {
		response = append(response, &dto.TenantResponse{
			ID:      t.ID,
			Name:    t.Name,
			Address: t.Address,
		})
	}

	return response, nil
}
