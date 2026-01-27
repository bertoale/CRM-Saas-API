package tenant

import (
	"time"
)

type StatusType = string

const (
	StatusActive   StatusType = "active"
	StatusInactive StatusType = "inactive"
	StatusSuspended StatusType = "suspended"
)

type Tenant struct {
	ID				uint 		 `gorm:"primaryKey"`
	Name			string		 `gorm:"unique;not null"`
	Domain			string		 `gorm:"unique;not null"`
	Status			StatusType		 `gorm:"default:'active'"`
	CreatedAt		time.Time	`gorm:"autoCreateTime"`
	UpdatedAt		time.Time 	`gorm:"autoUpdateTime"`
}


