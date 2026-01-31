package tenant

import (
	"time"
)

type StatusENUM = string

const (
	StatusActive    StatusENUM = "active"
	StatusInactive  StatusENUM = "inactive"
	StatusSuspended StatusENUM = "suspended"
)

type Tenant struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"unique;not null"`
	Domain    string     `gorm:"unique;not null"`
	Status    StatusENUM `gorm:"default:'active'"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type TenantRequest struct {
	Name   string `json:"name" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

type TenantResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}