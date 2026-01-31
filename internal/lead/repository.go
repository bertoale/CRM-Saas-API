package lead

import "gorm.io/gorm"

type Repository interface {
	Create(lead *Lead) error
	FindByID(ID, tenantID uint) (*Lead, error)
	FindByTenantID(tenantID uint) ([]Lead, error)
	FindByAssignedTo(userID, tenantID uint) ([]Lead, error)
	Update(lead *Lead) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(lead *Lead) error {
	return r.db.Create(lead).Error
}

func (r *repository) FindByID(ID, tenantID uint) (*Lead, error) {
	var lead Lead
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&lead).Error
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

func (r *repository) FindByTenantID(tenantID uint) ([]Lead, error) {
	var leads []Lead
	err := r.db.Where("tenant_id = ?", tenantID).Find(&leads).Error
	return leads, err
}

func (r *repository) FindByAssignedTo(userID, tenantID uint) ([]Lead, error) {
	var leads []Lead
	err := r.db.Where("assigned_to = ? AND tenant_id = ?", userID, tenantID).Find(&leads).Error
	return leads, err
}

func (r *repository) Update(lead *Lead) error {
	return r.db.Save(lead).Error
}

func (r *repository) Delete(ID, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&Lead{}).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}