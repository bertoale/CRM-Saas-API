// filepath: d:\CODING\CRM-SAAS\server\internal\reminder\service.go
package reminder

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Create(req *ReminderRequest, userID uint, tenantID uint) (*ReminderResponse, error)
	GetByID(id uint, tenantID uint) (*ReminderResponse, error)
	GetAll(tenantID uint) ([]*ReminderResponse, error)
	GetMyReminders(userID uint, tenantID uint) ([]*ReminderResponse, error)
	GetMyPendingReminders(userID uint, tenantID uint) ([]*ReminderResponse, error)
	GetUpcoming(tenantID uint, from time.Time, to time.Time) ([]*ReminderResponse, error)
	Update(id uint, req *ReminderRequest, tenantID uint) (*ReminderResponse, error)
	UpdateStatus(id uint, status string, tenantID uint) (*ReminderResponse, error)
	Delete(id uint, tenantID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req *ReminderRequest, userID uint, tenantID uint) (*ReminderResponse, error) {
	reminder := &Reminder{
		TenantID:   tenantID,
		UserID:     userID,
		EntityType: EntityENUM(req.EntityType),
		EntityID:   req.EntityID,
		RemindAt:   req.RemindAt,
		Status:     StatusPending,
	}

	if err := s.repo.Create(reminder); err != nil {
		return nil, err
	}

	return ToReminderResponse(reminder), nil
}

func (s *service) GetByID(id uint, tenantID uint) (*ReminderResponse, error) {
	reminder, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reminder not found")
		}
		return nil, err
	}

	return ToReminderResponse(reminder), nil
}

func (s *service) GetAll(tenantID uint) ([]*ReminderResponse, error) {
	reminders, err := s.repo.FindAll(tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ReminderResponse, len(reminders))
	for i, reminder := range reminders {
		responses[i] = ToReminderResponse(&reminder)
	}

	return responses, nil
}

func (s *service) GetMyReminders(userID uint, tenantID uint) ([]*ReminderResponse, error) {
	reminders, err := s.repo.FindByUserID(userID, tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ReminderResponse, len(reminders))
	for i, reminder := range reminders {
		responses[i] = ToReminderResponse(&reminder)
	}

	return responses, nil
}

func (s *service) GetMyPendingReminders(userID uint, tenantID uint) ([]*ReminderResponse, error) {
	reminders, err := s.repo.FindPendingByUserID(userID, tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*ReminderResponse, len(reminders))
	for i, reminder := range reminders {
		responses[i] = ToReminderResponse(&reminder)
	}

	return responses, nil
}

func (s *service) GetUpcoming(tenantID uint, from time.Time, to time.Time) ([]*ReminderResponse, error) {
	reminders, err := s.repo.FindUpcoming(tenantID, from, to)
	if err != nil {
		return nil, err
	}

	responses := make([]*ReminderResponse, len(reminders))
	for i, reminder := range reminders {
		responses[i] = ToReminderResponse(&reminder)
	}

	return responses, nil
}

func (s *service) Update(id uint, req *ReminderRequest, tenantID uint) (*ReminderResponse, error) {
	reminder, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reminder not found")
		}
		return nil, err
	}

	reminder.EntityType = EntityENUM(req.EntityType)
	reminder.EntityID = req.EntityID
	reminder.RemindAt = req.RemindAt

	if err := s.repo.Update(reminder); err != nil {
		return nil, err
	}

	return ToReminderResponse(reminder), nil
}

func (s *service) UpdateStatus(id uint, status string, tenantID uint) (*ReminderResponse, error) {
	reminder, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reminder not found")
		}
		return nil, err
	}

	reminder.Status = StatusENUM(status)

	if err := s.repo.Update(reminder); err != nil {
		return nil, err
	}

	return ToReminderResponse(reminder), nil
}

func (s *service) Delete(id uint, tenantID uint) error {
	_, err := s.repo.FindByID(id, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reminder not found")
		}
		return err
	}

	return s.repo.Delete(id, tenantID)
}
