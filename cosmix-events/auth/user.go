package auth

import "time"

type UserCreated struct {
	EventVersion string    `json:"event_version"`
	AuthUserID   uint      `json:"auth_user_id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserUpdated struct {
	EventVersion string    `json:"event_version"`
	AuthUserID   uint      `json:"auth_user_id"`
	Email        string    `json:"email"`
	UpdatedAt    time.Time `json:"updated_at"`
}