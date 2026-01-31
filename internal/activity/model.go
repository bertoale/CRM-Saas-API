package activity

import (
	"crm/internal/tenant"
	"crm/internal/user"
	"time"

	"gorm.io/datatypes"
)

type EntityENUM = string

const (
	EntityLead     EntityENUM = "lead"
	EntityCustomer EntityENUM = "customer"
	EntityDeal     EntityENUM = "deal"
)

type Activity struct {
	ID         uint            `gorm:"primaryKey"`
	TenantID   uint            `gorm:"not null;index"`
	UserID     uint            `gorm:"not null"`
	EntityType EntityENUM      `gorm:"not null"`
	EntityID   uint            `gorm:"not null"`
	Action     string          `gorm:"not null"`
	Meta       datatypes.JSON  `gorm:"type:json"`
	CreatedAt  time.Time       `gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
	User   user.User     `gorm:"foreignKey:UserID"`
}

// DTOs
type ActivityRequest struct {
	EntityType string                 `json:"entity_type" binding:"required,oneof=lead customer deal"`
	EntityID   uint                   `json:"entity_id" binding:"required"`
	Action     string                 `json:"action" binding:"required"`
	Meta       map[string]interface{} `json:"meta"`
}

type ActivityResponse struct {
	ID         uint                   `json:"id"`
	TenantID   uint                   `json:"tenant_id"`
	UserID     uint                   `json:"user_id"`
	EntityType string                 `json:"entity_type"`
	EntityID   uint                   `json:"entity_id"`
	Action     string                 `json:"action"`
	Meta       map[string]interface{} `json:"meta"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}
