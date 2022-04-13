package organization

import (
	"novocaine-dev/user"
	"time"
)

type Organization struct {
	ID        int
	Name      string
	Status    int8
	UserID    int
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
