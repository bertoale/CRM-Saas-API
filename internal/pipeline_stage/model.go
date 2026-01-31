package pipeline_stage

import (
	"crm/internal/tenant"
	"time"

	"gorm.io/gorm"
)

type PipelineStage struct {
	ID         uint           `gorm:"primaryKey"`
	TenantID   uint           `gorm:"not null;index"`
	Name       string         `gorm:"not null"`
	OrderIndex int            `gorm:"not null"`
	IsWon      bool           `gorm:"default:false"`
	IsLost     bool           `gorm:"default:false"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
}

// DTOs
type PipelineStageRequest struct {
	Name       string `json:"name" binding:"required"`
	OrderIndex int    `json:"order_index" binding:"required"`
	IsWon      bool   `json:"is_won"`
	IsLost     bool   `json:"is_lost"`
}

type PipelineStageResponse struct {
	ID         uint      `json:"id"`
	TenantID   uint      `json:"tenant_id"`
	Name       string    `json:"name"`
	OrderIndex int       `json:"order_index"`
	IsWon      bool      `json:"is_won"`
	IsLost     bool      `json:"is_lost"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

