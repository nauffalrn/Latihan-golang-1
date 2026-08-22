package repository

import (
	"context"
	"latihan/config"
	"latihan/model"
)

// GetTaskSummary mengambil rekapitulasi status tugas
// Menggunakan Context agar query bisa dibatalkan jika request dari client terputus
func GetTaskSummary(ctx context.Context) (model.TaskSummary, error) {
	var summary model.TaskSummary

	// Query 1: Menghitung total dan per status (Menggunakan Conditional Aggregation / CASE WHEN)
	// Ini JAUH lebih cepat daripada melakukan 4 query terpisah
	queryStatus := `
		SELECT 
			COUNT(id) as total_tasks,
			SUM(CASE WHEN status = 'TODO' THEN 1 ELSE 0 END) as todo,
			SUM(CASE WHEN status = 'IN_PROGRESS' THEN 1 ELSE 0 END) as in_progress,
			SUM(CASE WHEN status = 'REVIEW' THEN 1 ELSE 0 END) as review,
			SUM(CASE WHEN status = 'DONE' THEN 1 ELSE 0 END) as done
		FROM tasks
	`
	// Menjalankan query (menggunakan connection pool dari config)
	err := config.DB.QueryRowContext(ctx, queryStatus).Scan(
		&summary.TotalTasks,
		&summary.StatusBreakdown.Todo,
		&summary.StatusBreakdown.InProgress,
		&summary.StatusBreakdown.Review,
		&summary.StatusBreakdown.Done,
	)
	if err != nil {
		return summary, err
	}

	// Query 2: Menghitung Overdue Tasks (Tugas yang belum selesai dan melewati tanggal hari ini)
	// Karena ini tahun 2026, kita hardcode CURDATE() MySQL, atau jika perlu, gunakan parameter waktu dari Go
	queryOverdue := `
		SELECT COUNT(id) 
		FROM tasks 
		WHERE status != 'DONE' AND due_date < CURDATE()
	`
	err = config.DB.QueryRowContext(ctx, queryOverdue).Scan(&summary.OverdueTasks)
	if err != nil {
		return summary, err
	}

	return summary, nil
}

// GetTopWorkload mengambil 5 pengguna dengan tugas terbanyak yang masih berjalan
func GetTopWorkload(ctx context.Context) ([]model.Workload, error) {
	// Query: Join tabel users dan tasks, kelompokkan berdasarkan user, hitung tugas yang belum selesai
	query := `
		SELECT 
			u.id, 
			u.name, 
			COUNT(t.id) as active_tasks
		FROM users u
		LEFT JOIN tasks t ON u.id = t.assignee_id AND t.status IN ('TODO', 'IN_PROGRESS', 'REVIEW')
		GROUP BY u.id, u.name
		HAVING active_tasks > 0
		ORDER BY active_tasks DESC
		LIMIT 5
	`

	rows, err := config.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Pastikan selalu menutup rows!

	var workloads []model.Workload

	// Looping data yang dikembalikan MySQL
	for rows.Next() {
		var w model.Workload
		if err := rows.Scan(&w.UserID, &w.Name, &w.ActiveTasks); err != nil {
			return nil, err
		}
		workloads = append(workloads, w)
	}

	return workloads, nil
}