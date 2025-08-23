package models

import (
	"gorm.io/gorm"
)

type LinkPrecedence string

const (
	Primary   LinkPrecedence = "primary"
	Secondary LinkPrecedence = "secondary"
)

type Contact struct {
	gorm.Model
	ID             uint           `gorm:"primaryKey" json:"id"`
	PhoneNumber    *string        `json:"phoneNumber,omitempty"`
	Email          *string        `json:"email,omitempty"`
	LinkedID       *uint          `json:"linkedId,omitempty"`
	LinkPrecedence LinkPrecedence `gorm:"type:varchar(10);default:'primary'" json:"linkPrecedence,omitempty"`
}
