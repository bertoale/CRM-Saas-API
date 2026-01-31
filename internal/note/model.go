package note

import (
	"crm/internal/tenant"
	"crm/internal/user"
	"time"

	"gorm.io/gorm"
)

type EntityENUM = string

const (
	EntityLead     EntityENUM = "lead"
	EntityCustomer EntityENUM = "customer"
	EntityDeal     EntityENUM = "deal"
)

type Note struct {
	ID         uint           `gorm:"primaryKey"`
	TenantID   uint           `gorm:"not null;index"`
	EntityType EntityENUM     `gorm:"not null"`
	EntityID   uint           `gorm:"not null"`
	Content    string         `gorm:"not null"`
	CreatedBy  uint           `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
	User   user.User     `gorm:"foreignKey:CreatedBy"`
}

// DTOs
type NoteRequest struct {
	EntityType string `json:"entity_type" binding:"required,oneof=lead customer deal"`
	EntityID   uint   `json:"entity_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type NoteResponse struct {
	ID         uint      `json:"id"`
	TenantID   uint      `json:"tenant_id"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	Content    string    `json:"content"`
	CreatedBy  uint      `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
