package note

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(req NoteRequest, tenantID, userID uint) (*NoteResponse, error)
	GetAll(tenantID uint) ([]NoteResponse, error)
	GetByID(id, tenantID uint) (*NoteResponse, error)
	GetByEntity(entityType string, entityID, tenantID uint) ([]NoteResponse, error)
	Update(id uint, req NoteRequest, tenantID uint) (*NoteResponse, error)
	Delete(id, tenantID uint) error
}

type service struct {
	repo Repository
}

func (s *service) Create(req NoteRequest, tenantID, userID uint) (*NoteResponse, error) {
	note := &Note{
		TenantID:   tenantID,
		EntityType: EntityENUM(req.EntityType),
		EntityID:   req.EntityID,
		Content:    req.Content,
		CreatedBy:  userID,
	}

	if err := s.repo.Create(note); err != nil {
		return nil, err
	}

	return ToNoteResponse(note), nil
}

func (s *service) GetAll(tenantID uint) ([]NoteResponse, error) {
	notes, err := s.repo.FindByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	var responses []NoteResponse
	for _, n := range notes {
		responses = append(responses, *ToNoteResponse(&n))
	}
	return responses, nil
}

func (s *service) GetByID(id, tenantID uint) (*NoteResponse, error) {
	note, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("note tidak ditemukan")
		}
		return nil, err
	}
	return ToNoteResponse(note), nil
}

func (s *service) GetByEntity(entityType string, entityID, tenantID uint) ([]NoteResponse, error) {
	notes, err := s.repo.FindByEntity(EntityENUM(entityType), entityID, tenantID)
	if err != nil {
		return nil, err
	}

	var responses []NoteResponse
	for _, n := range notes {
		responses = append(responses, *ToNoteResponse(&n))
	}
	return responses, nil
}

func (s *service) Update(id uint, req NoteRequest, tenantID uint) (*NoteResponse, error) {
	note, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("note tidak ditemukan")
		}
		return nil, err
	}

	note.Content = req.Content

	if err := s.repo.Update(note); err != nil {
		return nil, err
	}

	return ToNoteResponse(note), nil
}

func (s *service) Delete(id, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("note tidak ditemukan")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
