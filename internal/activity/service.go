// filepath: d:\CODING\CRM-SAAS\server\internal\activity\service.go
package activity

import (
	"encoding/json"
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Service interface {
	Create(req *ActivityRequest, userID uint, tenantID uint) (*ActivityResponse, error)
	GetByID(id uint, tenantID uint) (*ActivityResponse, error)
	GetAll(tenantID uint) ([]*ActivityResponse, error)
	GetByEntity(entityType string, entityID uint, tenantID uint) ([]*ActivityResponse, error)
	GetByUserID(userID uint, tenantID uint) ([]*ActivityResponse, error)
	Update(id uint, req *ActivityRequest, tenantID uint) (*ActivityResponse, error)
	Delete(id uint, tenantID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req *ActivityRequest, userID uint, tenantID uint) (*ActivityResponse, error) {
	// Marshal meta to JSON
	var metaJSON datatypes.JSON
	if req.Meta != nil {
		metaBytes, err := json.Marshal(req.Meta)
		if err != nil {
			return nil, errors.New("invalid meta data")
		}
		metaJSON = metaBytes
	}

	activity := &Activity{
		TenantID:   tenantID,
		UserID:     userID,
		EntityType: EntityENUM(req.EntityType),
		EntityID:   req.EntityID,
		Action:     req.Action,
		Meta:       metaJSON,
	}

	if err := s.repo.Create(activity); err != nil {
		return nil, err
	}

	return ToActivityResponse(activity), nil
}

func (s *service) GetByID(id uint, tenantID uint) (*ActivityResponse, error) {
	activity, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("activity not found")
		}
		return nil, err
	}

	return ToActivityResponse(activity), nil
}

func (s *service) GetAll(tenantID uint) ([]*ActivityResponse, error) {
	activities, err := s.repo.FindAll(tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ActivityResponse, len(activities))
	for i, activity := range activities {
		responses[i] = ToActivityResponse(&activity)
	}

	return responses, nil
}

func (s *service) GetByEntity(entityType string, entityID uint, tenantID uint) ([]*ActivityResponse, error) {
	activities, err := s.repo.FindByEntity(entityType, entityID, tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ActivityResponse, len(activities))
	for i, activity := range activities {
		responses[i] = ToActivityResponse(&activity)
	}

	return responses, nil
}

func (s *service) GetByUserID(userID uint, tenantID uint) ([]*ActivityResponse, error) {
	activities, err := s.repo.FindByUserID(userID, tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ActivityResponse, len(activities))
	for i, activity := range activities {
		responses[i] = ToActivityResponse(&activity)
	}

	return responses, nil
}

func (s *service) Update(id uint, req *ActivityRequest, tenantID uint) (*ActivityResponse, error) {
	activity, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("activity not found")
		}
		return nil, err
	}

	// Marshal meta to JSON
	var metaJSON datatypes.JSON
	if req.Meta != nil {
		metaBytes, err := json.Marshal(req.Meta)
		if err != nil {
			return nil, errors.New("invalid meta data")
		}
		metaJSON = metaBytes
	}

	activity.EntityType = EntityENUM(req.EntityType)
	activity.EntityID = req.EntityID
	activity.Action = req.Action
	activity.Meta = metaJSON

	if err := s.repo.Update(activity); err != nil {
		return nil, err
	}

	return ToActivityResponse(activity), nil
}

func (s *service) Delete(id uint, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("activity not found")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}
