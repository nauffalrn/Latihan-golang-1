package service

import (
	"context"
	"log"
	"latihan/model"
	"latihan/repository"
)

// GetDashboardSummary memanggil repository untuk mengambil rekap status tugas
func GetDashboardSummary(ctx context.Context) (model.TaskSummary, error) {
	summary, err := repository.GetTaskSummary(ctx)
	if err != nil {
		// Log error di service agar kita tahu lokasi gagalnya
		log.Printf("[DashboardService.GetSummary] Gagal mengambil data summary: %v", err)
		return model.TaskSummary{}, err
	}

	return summary, nil
}

// GetTeamWorkload memanggil repository untuk mendapatkan top workload
func GetTeamWorkload(ctx context.Context) ([]model.Workload, error) {
	workloads, err := repository.GetTopWorkload(ctx)
	if err != nil {
		log.Printf("[DashboardService.GetWorkload] Gagal mengambil data workload: %v", err)
		return nil, err
	}

	// Opsional: Jika hasil query dari database bernilai `nil` (tidak ada data aktif), 
	// kita pastikan service mengembalikan slice kosong [] agar Frontend menerima array kosong
	// alih-alih `null` dalam bentuk JSON.
	if workloads == nil {
		workloads = []model.Workload{}
	}

	return workloads, nil
}