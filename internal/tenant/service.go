package tenant

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	GetAll() ([]TenantResponse, error)
	GetByID(id uint) (*TenantResponse, error)
	Update(id uint, req TenantRequest) (*TenantResponse, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

// Delete implements Service.
func (s *service) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("tenant tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id)
}

// GetAll implements Service.
func (s *service) GetAll() ([]TenantResponse, error) {
	tenants, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []TenantResponse
	for _, t := range tenants {
		responses = append(responses, *ToTenantResponse(&t))
	}
	return responses, nil
}

// GetByID implements Service.
func (s *service) GetByID(id uint) (*TenantResponse, error) {
	tenant, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tenant tidak ditemukan")
		}
		return nil, err
	}
	return ToTenantResponse(tenant), nil
}

// Update implements Service.
func (s *service) Update(id uint, req TenantRequest) (*TenantResponse, error) {
	tenant, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tenant tidak ditemukan")
		}
		return nil, err
	}

	// Check if domain already exists (for other tenant)
	existingTenant, err := s.repo.FindByDomain(req.Domain)
	if err == nil && existingTenant.ID != id {
		return nil, errors.New("domain sudah digunakan")
	}

	tenant.Name = req.Name
	tenant.Domain = req.Domain

	if err := s.repo.Update(tenant); err != nil {
		return nil, err
	}

	return ToTenantResponse(tenant), nil
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
