package tenant

import "gorm.io/gorm"

type Repository interface {
	Create(tenant *Tenant) error
	FindByID(ID uint) (*Tenant, error)
	FindByDomain(domain string) (*Tenant, error)
	FindAll() ([]Tenant, error)
	Update(tenant *Tenant) error
	Delete(ID uint) error
}

type repository struct {
	db *gorm.DB
}

// Create implements Repository.
func (r *repository) Create(tenant *Tenant) error {
	return r.db.Create(tenant).Error
}

// Delete implements Repository.
func (r *repository) Delete(ID uint) error {
	return r.db.Delete(&Tenant{}, ID).Error
}

// FindAll implements Repository.
func (r *repository) FindAll() ([]Tenant, error) {
	var tenants []Tenant
	err := r.db.Find(&tenants).Error
	return tenants, err
}

// FindByDomain implements Repository.
func (r *repository) FindByDomain(domain string) (*Tenant, error) {
	var tenant Tenant
	err := r.db.Where("domain = ?", domain).First(&tenant).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// FindByID implements Repository.
func (r *repository) FindByID(ID uint) (*Tenant, error) {
	var tenant Tenant
	err := r.db.First(&tenant, ID).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// Update implements Repository.
func (r *repository) Update(tenant *Tenant) error {
	return r.db.Save(tenant).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
