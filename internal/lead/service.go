package lead

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(req LeadRequest, tenantID uint) (*LeadResponse, error)
	GetAll(tenantID uint) ([]LeadResponse, error)
	GetByID(id, tenantID uint) (*LeadResponse, error)
	GetByAssignedTo(userID, tenantID uint) ([]LeadResponse, error)
	Update(id uint, req LeadRequest, tenantID uint) (*LeadResponse, error)
	Delete(id, tenantID uint) error
}

type service struct {
	repo Repository
}

func (s *service) Create(req LeadRequest, tenantID uint) (*LeadResponse, error) {
	status := req.Status
	if status == "" {
		status = StatusNew
	}

	lead := &Lead{
		TenantID:   tenantID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Source:     req.Source,
		Status:     StatusEnum(status),
		AssignedTo: req.AssignedTo,
	}

	if err := s.repo.Create(lead); err != nil {
		return nil, err
	}

	return ToLeadResponse(lead), nil
}

func (s *service) GetAll(tenantID uint) ([]LeadResponse, error) {
	leads, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []LeadResponse
	for _, l := range leads {
		responses = append(responses, *ToLeadResponse(&l))
	}
	return responses, nil
}

func (s *service) GetByID(id, tenantID uint) (*LeadResponse, error) {
	lead, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lead tidak ditemukan")
		}
		return nil, err
	}
	return ToLeadResponse(lead), nil
}

func (s *service) GetByAssignedTo(userID, tenantID uint) ([]LeadResponse, error) {
	leads, err := s.repo.FindByAssignedTo(userID, tenantID)
	if err != nil {
		return nil, err
	}

	var responses []LeadResponse
	for _, l := range leads {
		responses = append(responses, *ToLeadResponse(&l))
	}
	return responses, nil
}

func (s *service) Update(id uint, req LeadRequest, tenantID uint) (*LeadResponse, error) {
	lead, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lead tidak ditemukan")
		}
		return nil, err
	}

	lead.Name = req.Name
	lead.Email = req.Email
	lead.Phone = req.Phone
	lead.Source = req.Source
	if req.Status != "" {
		lead.Status = StatusEnum(req.Status)
	}
	lead.AssignedTo = req.AssignedTo

	if err := s.repo.Update(lead); err != nil {
		return nil, err
	}

	return ToLeadResponse(lead), nil
}

func (s *service) Delete(id, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("lead tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
