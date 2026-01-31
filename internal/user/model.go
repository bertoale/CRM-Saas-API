package user

import (
	"crm/internal/tenant"
	"time"
)

type RoleENUM = string
type StatusENUM = string

const (
	RoleOwner RoleENUM = "owner"
	RoleAdmin RoleENUM = "admin"
	RoleSales RoleENUM = "sales"
)
const (
	StatusActive    StatusENUM = "active"
	StatusInactive  StatusENUM = "inactive"
	StatusSuspended StatusENUM = "suspended"
)

type User struct {
	ID        uint       `gorm:"primaryKey"`
	TenantID  uint       `gorm:"not null;index"`
	Name      string     `gorm:"not null"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Role      RoleENUM   `gorm:"default:'sales'"`
	Status    StatusENUM `gorm:"default:'active'"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	//relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=owner admin sales"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	TenantID  uint      `json:"tenant_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type CreateTenantWithOwnerRequest struct {
	Tenant tenant.TenantRequest `json:"tenant" binding:"required"`
	User   UserRequest          `json:"user" binding:"required"`
}

type CreateTenantWithOwnerResponse struct {
	Tenant tenant.TenantResponse `json:"tenant"`
	User   UserResponse          `json:"user"`
}