package taskHistory

type CreateTaskHistoryInput struct {
	Content      string `json:"contet" binding:"required"`
	LatestStatus string `json:"latestStatus" binding:"required"`
	UserId       string `json:"userId" binding:"required"`
	TaskId       string `json:"taskId" binding:"required"`
}
