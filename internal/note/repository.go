package note

import "gorm.io/gorm"

type Repository interface {
	Create(note *Note) error
	FindByID(ID, tenantID uint) (*Note, error)
	FindByTenantID(tenantID uint) ([]Note, error)
	FindByEntity(entityType EntityENUM, entityID, tenantID uint) ([]Note, error)
	Update(note *Note) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(note *Note) error {
	return r.db.Create(note).Error
}

func (r *repository) FindByID(ID, tenantID uint) (*Note, error) {
	var note Note
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *repository) FindByTenantID(tenantID uint) ([]Note, error) {
	var notes []Note
	err := r.db.Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&notes).Error
	return notes, err
}

func (r *repository) FindByEntity(entityType EntityENUM, entityID, tenantID uint) ([]Note, error) {
	var notes []Note
	err := r.db.Where("entity_type = ? AND entity_id = ? AND tenant_id = ?", entityType, entityID, tenantID).
		Order("created_at DESC").Find(&notes).Error
	return notes, err
}

func (r *repository) Update(note *Note) error {
	return r.db.Save(note).Error
}

func (r *repository) Delete(ID, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&Note{}).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
