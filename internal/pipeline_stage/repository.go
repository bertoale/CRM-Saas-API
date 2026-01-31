package pipeline_stage

import "gorm.io/gorm"

type Repository interface {
	Create(stage *PipelineStage) error
	FindByID(ID, tenantID uint) (*PipelineStage, error)
	FindByTenantID(tenantID uint) ([]PipelineStage, error)
	Update(stage *PipelineStage) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(stage *PipelineStage) error {
	return r.db.Create(stage).Error
}

func (r *repository) FindByID(ID, tenantID uint) (*PipelineStage, error) {
	var stage PipelineStage
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&stage).Error
	if err != nil {
		return nil, err
	}
	return &stage, nil
}

func (r *repository) FindByTenantID(tenantID uint) ([]PipelineStage, error) {
	var stages []PipelineStage
	err := r.db.Where("tenant_id = ?", tenantID).Order("order_index ASC").Find(&stages).Error
	return stages, err
}

func (r *repository) Update(stage *PipelineStage) error {
	return r.db.Save(stage).Error
}

func (r *repository) Delete(ID, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&PipelineStage{}).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
