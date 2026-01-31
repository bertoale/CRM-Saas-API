package pipeline_stage

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(req PipelineStageRequest, tenantID uint) (*PipelineStageResponse, error)
	GetAll(tenantID uint) ([]PipelineStageResponse, error)
	GetByID(id, tenantID uint) (*PipelineStageResponse, error)
	Update(id uint, req PipelineStageRequest, tenantID uint) (*PipelineStageResponse, error)
	Delete(id, tenantID uint) error
}

type service struct {
	repo Repository
}

func (s *service) Create(req PipelineStageRequest, tenantID uint) (*PipelineStageResponse, error) {
	stage := &PipelineStage{
		TenantID:   tenantID,
		Name:       req.Name,
		OrderIndex: req.OrderIndex,
		IsWon:      req.IsWon,
		IsLost:     req.IsLost,
	}

	if err := s.repo.Create(stage); err != nil {
		return nil, err
	}

	return ToPipelineStageResponse(stage), nil
}

func (s *service) GetAll(tenantID uint) ([]PipelineStageResponse, error) {
	stages, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []PipelineStageResponse
	for _, st := range stages {
		responses = append(responses, *ToPipelineStageResponse(&st))
	}
	return responses, nil
}

func (s *service) GetByID(id, tenantID uint) (*PipelineStageResponse, error) {
	stage, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("pipeline stage tidak ditemukan")
		}
		return nil, err
	}
	return ToPipelineStageResponse(stage), nil
}

func (s *service) Update(id uint, req PipelineStageRequest, tenantID uint) (*PipelineStageResponse, error) {
	stage, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("pipeline stage tidak ditemukan")
		}
		return nil, err
	}

	stage.Name = req.Name
	stage.OrderIndex = req.OrderIndex
	stage.IsWon = req.IsWon
	stage.IsLost = req.IsLost

	if err := s.repo.Update(stage); err != nil {
		return nil, err
	}

	return ToPipelineStageResponse(stage), nil
}

func (s *service) Delete(id, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("pipeline stage tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
