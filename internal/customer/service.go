package customer

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(req CustomerRequest, tenantID uint) (*CustomerResponse, error)
	GetAll(tenantID uint) ([]CustomerResponse, error)
	GetByID(id, tenantID uint) (*CustomerResponse, error)
	Update(id uint, req CustomerRequest, tenantID uint) (*CustomerResponse, error)
	Delete(id, tenantID uint) error
}

type service struct {
	repo Repository
}

func (s *service) Create(req CustomerRequest, tenantID uint) (*CustomerResponse, error) {
	customer := &Customer{
		TenantID:    tenantID,
		Name:        req.Name,
		Email:       req.Email,
		Phone:       req.Phone,
		CompanyName: req.CompanyName,
		LeadID:      req.LeadID,
	}

	if err := s.repo.Create(customer); err != nil {
		return nil, err
	}

	return ToCustomerResponse(customer), nil
}

func (s *service) GetAll(tenantID uint) ([]CustomerResponse, error) {
	customers, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []CustomerResponse
	for _, c := range customers {
		responses = append(responses, *ToCustomerResponse(&c))
	}
	return responses, nil
}

func (s *service) GetByID(id, tenantID uint) (*CustomerResponse, error) {
	customer, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer tidak ditemukan")
		}
		return nil, err
	}
	return ToCustomerResponse(customer), nil
}

func (s *service) Update(id uint, req CustomerRequest, tenantID uint) (*CustomerResponse, error) {
	customer, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer tidak ditemukan")
		}
		return nil, err
	}

	customer.Name = req.Name
	customer.Email = req.Email
	customer.Phone = req.Phone
	customer.CompanyName = req.CompanyName
	customer.LeadID = req.LeadID

	if err := s.repo.Update(customer); err != nil {
		return nil, err
	}

	return ToCustomerResponse(customer), nil
}

func (s *service) Delete(id, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
