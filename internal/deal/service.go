package deal

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(req DealRequest, tenantID uint) (*DealResponse, error)
	GetAll(tenantID uint) ([]DealResponse, error)
	GetByID(id, tenantID uint) (*DealResponse, error)
	GetByAssignedTo(userID, tenantID uint) ([]DealResponse, error)
	GetByStageID(stageID, tenantID uint) ([]DealResponse, error)
	Update(id uint, req DealRequest, tenantID uint) (*DealResponse, error)
	Delete(id, tenantID uint) error
}

type service struct {
	repo Repository
}

func (s *service) Create(req DealRequest, tenantID uint) (*DealResponse, error) {
	deal := &Deal{
		TenantID:          tenantID,
		CustomerID:        req.CustomerID,
		Title:             req.Title,
		Value:             req.Value,
		StageID:           req.StageID,
		Probability:       req.Probability,
		ExpectedCloseDate: req.ExpectedCloseDate,
		AssignedTo:        req.AssignedTo,
	}

	if err := s.repo.Create(deal); err != nil {
		return nil, err
	}

	return ToDealResponse(deal), nil
}

func (s *service) GetAll(tenantID uint) ([]DealResponse, error) {
	deals, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []DealResponse
	for _, d := range deals {
		responses = append(responses, *ToDealResponse(&d))
	}
	return responses, nil
}

func (s *service) GetByID(id, tenantID uint) (*DealResponse, error) {
	deal, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("deal tidak ditemukan")
		}
		return nil, err
	}
	return ToDealResponse(deal), nil
}

func (s *service) GetByAssignedTo(userID, tenantID uint) ([]DealResponse, error) {
	deals, err := s.repo.FindByAssignedTo(userID, tenantID)
	if err != nil {
		return nil, err
	}

	var responses []DealResponse
	for _, d := range deals {
		responses = append(responses, *ToDealResponse(&d))
	}
	return responses, nil
}

func (s *service) GetByStageID(stageID, tenantID uint) ([]DealResponse, error) {
	deals, err := s.repo.FindByStageID(stageID, tenantID)
	if err != nil {
		return nil, err
	}

	var responses []DealResponse
	for _, d := range deals {
		responses = append(responses, *ToDealResponse(&d))
	}
	return responses, nil
}

func (s *service) Update(id uint, req DealRequest, tenantID uint) (*DealResponse, error) {
	deal, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("deal tidak ditemukan")
		}
		return nil, err
	}

	deal.CustomerID = req.CustomerID
	deal.Title = req.Title
	deal.Value = req.Value
	deal.StageID = req.StageID
	deal.Probability = req.Probability
	deal.ExpectedCloseDate = req.ExpectedCloseDate
	deal.AssignedTo = req.AssignedTo

	if err := s.repo.Update(deal); err != nil {
		return nil, err
	}

	return ToDealResponse(deal), nil
}

func (s *service) Delete(id, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("deal tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
