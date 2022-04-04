package taskHistory

type TaskHistoryFormatter struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	LatestStatus string `json:"latestStatus"`
}

func FormatTask(taskHistory TaskHistory) TaskHistoryFormatter {
	taskHistoryFormatter := TaskHistoryFormatter{}
	taskHistoryFormatter.ID = taskHistory.ID
	taskHistoryFormatter.Content = taskHistory.Content
	taskHistoryFormatter.LatestStatus = taskHistory.LatestStatus

	return taskHistoryFormatter
}

func FormatTasks(taskHistories []TaskHistory) []TaskHistoryFormatter {
	tasksFormatter := []TaskHistoryFormatter{}

	for _, task := range taskHistories {
		taskHistoryFormatter := FormatTask(task)
		tasksFormatter = append(tasksFormatter, taskHistoryFormatter)
	}

	return tasksFormatter
}
