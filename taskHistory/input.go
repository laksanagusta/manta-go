package taskHistory

type CreateTaskHistoryInput struct {
	Content      string `json:"content" binding:"required"`
	LatestStatus string `json:"latestStatus" binding:"required"`
	UserId       int    `json:"userId" binding:"required"`
	TaskId       int    `json:"taskId" binding:"required"`
}
