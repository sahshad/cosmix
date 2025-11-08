package models

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	FirstName    string     `gorm:"size:100" json:"firstName"`
	LastName     string     `gorm:"size:100" json:"lastName"`
	Username     string     `gorm:"size:50;uniqueIndex" json:"username"`
	Email        string     `gorm:"size:150;uniqueIndex" json:"email"`
	PasswordHash string     `gorm:"size:255" json:"-"`
	DateOfBirth  *time.Time `json:"dateOfBirth,omitempty"`
	Role         Role       `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
