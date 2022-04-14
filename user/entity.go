package user

import (
	"time"
)

type User struct {
	ID             int
	Username       string
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	Role           string
	OrganizationId int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
