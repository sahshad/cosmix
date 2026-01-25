package events

import "time"

type UserCreatedEvent struct {
	AuthUserID uint      `json:"auth_user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CreatedAt  time.Time `json:"created_at"`
}
