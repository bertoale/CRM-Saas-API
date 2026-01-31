package reminder

import (
	"crm/internal/tenant"
	"crm/internal/user"
	"time"
)

type EntityENUM = string
type StatusENUM = string

const (
	StatusPending   StatusENUM = "pending"
	StatusDone      StatusENUM = "done"
	StatusCancelled StatusENUM = "cancelled"
)

const (
	EntityLead     EntityENUM = "lead"
	EntityCustomer EntityENUM = "customer"
	EntityDeal     EntityENUM = "deal"
)

type Reminder struct {
	ID         uint       `gorm:"primaryKey"`
	TenantID   uint       `gorm:"not null;index"`
	UserID     uint       `gorm:"not null"`
	EntityType EntityENUM `gorm:"not null"`
	EntityID   uint       `gorm:"not null"`
	RemindAt   time.Time  `gorm:"not null"`
	Status     StatusENUM `gorm:"default:'pending'"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	// Relations
	Tenant tenant.Tenant `gorm:"foreignKey:TenantID"`
	User   user.User     `gorm:"foreignKey:UserID"`
}

// DTOs
type ReminderRequest struct {
	EntityType string    `json:"entity_type" binding:"required,oneof=lead customer deal"`
	EntityID   uint      `json:"entity_id" binding:"required"`
	RemindAt   time.Time `json:"remind_at" binding:"required"`
}

type ReminderResponse struct {
	ID         uint      `json:"id"`
	TenantID   uint      `json:"tenant_id"`
	UserID     uint      `json:"user_id"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	RemindAt   time.Time `json:"remind_at"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateReminderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending done cancelled"`
}