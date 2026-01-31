package customer

import "gorm.io/gorm"

type Repository interface {
	Create(customer *Customer) error
	FindByID(ID, tenantID uint) (*Customer, error)
	FindByTenantID(tenantID uint) ([]Customer, error)
	Update(customer *Customer) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *repository) FindByID(ID, tenantID uint) (*Customer, error) {
	var customer Customer
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *repository) FindByTenantID(tenantID uint) ([]Customer, error) {
	var customers []Customer
	err := r.db.Where("tenant_id = ?", tenantID).Find(&customers).Error
	return customers, err
}

func (r *repository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

func (r *repository) Delete(ID, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&Customer{}).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
