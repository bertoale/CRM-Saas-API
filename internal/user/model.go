package user

import "crm/internal/tenant"

type RoleType = string
type StatusType = string

const (
	RoleOwner RoleType = "owner"
	RoleAdmin RoleType = "admin"
	RoleSales RoleType = "sales"
)
const (
	StatusActive    StatusType = "active"
	StatusInactive  StatusType = "inactive"
	StatusSuspended StatusType = "suspended"
)

type User struct {
	ID       uint          `gorm:"primaryKey"`
	TenantID uint          `gorm:"not null;index"`
	Name     string        `gorm:"not null"`
	Email    string        `gorm:"unique;not null"`
	Password string        `gorm:"not null"`
	Role     RoleType      `gorm:"default:'sales'"`
	Status   StatusType    `gorm:"default:'active'"`
	//relations
	Tenant   tenant.Tenant `gorm:"foreignKey:TenantID"`
}