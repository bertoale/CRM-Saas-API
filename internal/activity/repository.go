package activity

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(activity *Activity) error
	FindByID(id uint, tenantID uint) (*Activity, error)
	FindAll(tenantID uint) ([]Activity, error)
	FindByEntity(entityType string, entityID uint, tenantID uint) ([]Activity, error)
	FindByUserID(userID uint, tenantID uint) ([]Activity, error)
	Update(activity *Activity) error
	Delete(id uint, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(activity *Activity) error {
	return r.db.Create(activity).Error
}

func (r *repository) FindByID(id uint, tenantID uint) (*Activity, error) {
	var activity Activity
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&activity).Error
	return &activity, err
}

func (r *repository) FindAll(tenantID uint) ([]Activity, error) {
	var activities []Activity
	err := r.db.Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&activities).Error
	return activities, err
}

func (r *repository) FindByEntity(entityType string, entityID uint, tenantID uint) ([]Activity, error) {
	var activities []Activity
	err := r.db.Where("entity_type = ? AND entity_id = ? AND tenant_id = ?", entityType, entityID, tenantID).
		Order("created_at DESC").
		Find(&activities).Error
	return activities, err
}

func (r *repository) FindByUserID(userID uint, tenantID uint) ([]Activity, error) {
	var activities []Activity
	err := r.db.Where("user_id = ? AND tenant_id = ?", userID, tenantID).
		Order("created_at DESC").
		Find(&activities).Error
	return activities, err
}

func (r *repository) Update(activity *Activity) error {
	return r.db.Save(activity).Error
}

func (r *repository) Delete(id uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&Activity{}).Error
}