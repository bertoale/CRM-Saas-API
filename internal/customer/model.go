package customer

import (
	"crm/internal/lead"
	"crm/internal/tenant"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint           `gorm:"primaryKey"`
	TenantID    uint           `gorm:"index;not null"`
	Name        string         `gorm:"not null"`
	Email       string         `gorm:"not null"`
	Phone       string         `gorm:"not null"`
	CompanyName string         `gorm:"not null"`
	LeadID      uint           `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
	Lead   lead.Lead     `gorm:"foreignKey:LeadID"`
}

// DTOs
type CustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	LeadID      uint   `json:"lead_id" binding:"required"`
}

type CustomerResponse struct {
	ID          uint      `json:"id"`
	TenantID    uint      `json:"tenant_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CompanyName string    `json:"company_name"`
	LeadID      uint      `json:"lead_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
