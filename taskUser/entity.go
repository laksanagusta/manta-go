package taskUser

import (
	"time"
)

type TaskUser struct {
	ID                int
	Task_id 		  int
	User_id			  int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
