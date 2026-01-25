package dto

type CreateTenantRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
}

type TenantResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetTenantRequest struct {
	ID string `path:"id" validate:"required"`
}
