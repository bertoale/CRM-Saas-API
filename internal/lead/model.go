package lead

import (
	"crm/internal/tenant"
	"crm/internal/user"
	"time"

	"gorm.io/gorm"
)

type StatusEnum = string

const (
	StatusNew         StatusEnum = "new"
	StatusContacted   StatusEnum = "contacted"
	StatusQualified   StatusEnum = "qualified"
	StatusUnqualified StatusEnum = "unqualified"
)

type Lead struct {
	ID         uint           `gorm:"primaryKey"`
	TenantID   uint           `gorm:"not null;index"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"not null"`
	Phone      string         `gorm:"not null"`
	Source     string         `gorm:"not null"`
	Status     StatusEnum     `gorm:"default:'new'"`
	AssignedTo uint           `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
	User   user.User     `gorm:"foreignKey:AssignedTo"`
}

// DTOs
type LeadRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
	Source     string `json:"source" binding:"required"`
	Status     string `json:"status" binding:"omitempty,oneof=new contacted qualified unqualified"`
	AssignedTo uint   `json:"assigned_to" binding:"required"`
}

type LeadResponse struct {
	ID         uint      `json:"id"`
	TenantID   uint      `json:"tenant_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Source     string    `json:"source"`
	Status     string    `json:"status"`
	AssignedTo uint      `json:"assigned_to"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
