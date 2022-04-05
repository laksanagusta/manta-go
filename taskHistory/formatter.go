package taskHistory

type TaskHistoryFormatter struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	LatestStatus string `json:"latestStatus"`
}

func FormatTaskHistory(taskHistory TaskHistory) TaskHistoryFormatter {
	taskHistoryFormatter := TaskHistoryFormatter{}
	taskHistoryFormatter.ID = taskHistory.ID
	taskHistoryFormatter.Content = taskHistory.Content
	taskHistoryFormatter.LatestStatus = taskHistory.LatestStatus

	return taskHistoryFormatter
}

func FormatTaskHistories(taskHistories []TaskHistory) []TaskHistoryFormatter {
	tasksFormatter := []TaskHistoryFormatter{}

	for _, task := range taskHistories {
		taskHistoryFormatter := FormatTaskHistory(task)
		tasksFormatter = append(tasksFormatter, taskHistoryFormatter)
	}

	return tasksFormatter
}
