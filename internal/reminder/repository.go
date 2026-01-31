package reminder

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(reminder *Reminder) error
	FindByID(id uint, tenantID uint) (*Reminder, error)
	FindAll(tenantID uint) ([]Reminder, error)
	FindByUserID(userID uint, tenantID uint) ([]Reminder, error)
	FindPendingByUserID(userID uint, tenantID uint) ([]Reminder, error)
	FindUpcoming(tenantID uint, from time.Time, to time.Time) ([]Reminder, error)
	Update(reminder *Reminder) error
	Delete(id uint, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(reminder *Reminder) error {
	return r.db.Create(reminder).Error
}

func (r *repository) FindByID(id uint, tenantID uint) (*Reminder, error) {
	var reminder Reminder
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&reminder).Error
	return &reminder, err
}

func (r *repository) FindAll(tenantID uint) ([]Reminder, error) {
	var reminders []Reminder
	err := r.db.Where("tenant_id = ?", tenantID).Order("remind_at ASC").Find(&reminders).Error
	return reminders, err
}

func (r *repository) FindByUserID(userID uint, tenantID uint) ([]Reminder, error) {
	var reminders []Reminder
	err := r.db.Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Order("remind_at ASC").
		Find(&reminders).Error
	return reminders, err
}

func (r *repository) FindPendingByUserID(userID uint, tenantID uint) ([]Reminder, error) {
	var reminders []Reminder
	err := r.db.Where("user_id = ? AND tenant_id = ? AND status = ?", userID, tenantID, StatusPending).
		Order("remind_at ASC").
		Find(&reminders).Error
	return reminders, err
}

func (r *repository) FindUpcoming(tenantID uint, from time.Time, to time.Time) ([]Reminder, error) {
	var reminders []Reminder
	err := r.db.Where("tenant_id = ? AND status = ? AND remind_at BETWEEN ? AND ?", 
		tenantID, StatusPending, from, to).
		Order("remind_at ASC").
		Find(&reminders).Error
	return reminders, err
}

func (r *repository) Update(reminder *Reminder) error {
	return r.db.Save(reminder).Error
}

func (r *repository) Delete(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Reminder{}).Error
}
