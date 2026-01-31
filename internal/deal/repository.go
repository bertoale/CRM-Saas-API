package deal

import "gorm.io/gorm"

type Repository interface {
	Create(deal *Deal) error
	FindByID(ID, tenantID uint) (*Deal, error)
	FindByTenantID(tenantID uint) ([]Deal, error)
	FindByAssignedTo(userID, tenantID uint) ([]Deal, error)
	FindByStageID(stageID, tenantID uint) ([]Deal, error)
	Update(deal *Deal) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(deal *Deal) error {
	return r.db.Create(deal).Error
}

func (r *repository) FindByID(ID, tenantID uint) (*Deal, error) {
	var deal Deal
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&deal).Error
	if err != nil {
		return nil, err
	}
	return &deal, nil
}

func (r *repository) FindByTenantID(tenantID uint) ([]Deal, error) {
	var deals []Deal
	err := r.db.Where("tenant_id = ?", tenantID).Find(&deals).Error
	return deals, err
}

func (r *repository) FindByAssignedTo(userID, tenantID uint) ([]Deal, error) {
	var deals []Deal
	err := r.db.Where("assigned_to = ? AND tenant_id = ?", userID, tenantID).Find(&deals).Error
	return deals, err
}

func (r *repository) FindByStageID(stageID, tenantID uint) ([]Deal, error) {
	var deals []Deal
	err := r.db.Where("stage_id = ? AND tenant_id = ?", stageID, tenantID).Find(&deals).Error
	return deals, err
}

func (r *repository) Update(deal *Deal) error {
	return r.db.Save(deal).Error
}

func (r *repository) Delete(ID, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&Deal{}).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
