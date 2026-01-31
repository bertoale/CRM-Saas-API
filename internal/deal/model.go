package deal

import (
	"crm/internal/customer"
	"crm/internal/pipeline_stage"
	"crm/internal/tenant"
	"crm/internal/user"
	"time"

	"gorm.io/gorm"
)

type Deal struct {
	ID                uint           `gorm:"primaryKey"`
	TenantID          uint           `gorm:"not null;index"`
	CustomerID        uint           `gorm:"not null"`
	Title             string         `gorm:"not null"`
	Value             float64        `gorm:"not null"`
	StageID           uint           `gorm:"not null"`
	Probability       int            `gorm:"not null"`
	ExpectedCloseDate time.Time      `gorm:"not null"`
	AssignedTo        uint           `gorm:"not null"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	// Relations
	Tenant        tenant.Tenant                `gorm:"foreignKey:TenantID"`
	Customer      customer.Customer            `gorm:"foreignKey:CustomerID"`
	PipelineStage pipeline_stage.PipelineStage `gorm:"foreignKey:StageID"`
	User          user.User                    `gorm:"foreignKey:AssignedTo"`
}

// DTOs
type DealRequest struct {
	CustomerID        uint      `json:"customer_id" binding:"required"`
	Title             string    `json:"title" binding:"required"`
	Value             float64   `json:"value" binding:"required"`
	StageID           uint      `json:"stage_id" binding:"required"`
	Probability       int       `json:"probability" binding:"required,min=0,max=100"`
	ExpectedCloseDate time.Time `json:"expected_close_date" binding:"required"`
	AssignedTo        uint      `json:"assigned_to" binding:"required"`
}

type DealResponse struct {
	ID                uint      `json:"id"`
	TenantID          uint      `json:"tenant_id"`
	CustomerID        uint      `json:"customer_id"`
	Title             string    `json:"title"`
	Value             float64   `json:"value"`
	StageID           uint      `json:"stage_id"`
	Probability       int       `json:"probability"`
	ExpectedCloseDate time.Time `json:"expected_close_date"`
	AssignedTo        uint      `json:"assigned_to"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

