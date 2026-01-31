package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindByID(ID, tenantID uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByTenantID(tenantID uint) ([]User, error)
	Update(user *User, tenantID uint) error
	Delete(ID, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

// Create implements Repository.
func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

// Delete implements Repository.
func (r *repository) Delete(ID uint, tenantID uint) error {
	return r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).Delete(&User{}).Error
}

// FindByEmail implements Repository.
func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}	
	return &user, nil
}

// FindByID implements Repository.
func (r *repository) FindByID(ID uint, tenantID uint) (*User, error) {
	var user User
	err := r.db.Where("id = ? AND tenant_id = ?", ID, tenantID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByTenantID implements Repository.
func (r *repository) FindByTenantID(tenantID uint) ([]User, error) {
	var users []User
	err := r.db.Where("tenant_id = ?", tenantID).Find(&users).Error
	return users, err
}

// Update implements Repository.
func (r *repository) Update(user *User, tenantID uint) error {
	return r.db.Model(&User{}).Where("id = ? AND tenant_id = ?", user.ID, tenantID).Updates(user).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
