package model

// TaskSummary merepresentasikan data agregasi status tugas (Endpoint A)
type TaskSummary struct {
	TotalTasks     int `json:"total_tasks"`
	StatusBreakdown struct {
		Todo       int `json:"todo"`
		InProgress int `json:"in_progress"`
		Review     int `json:"review"`
		Done       int `json:"done"`
	} `json:"status_breakdown"`
	OverdueTasks   int `json:"overdue_tasks"`
}

// Workload merepresentasikan beban kerja tiap karyawan (Endpoint B)
type Workload struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	ActiveTasks int    `json:"active_tasks"`
}