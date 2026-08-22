package route

import (
	"net/http"
	"latihan/controller"
)

// SetupRouter mendaftarkan semua endpoint URL ke handler yang sesuai
func SetupRouter() *http.ServeMux {
	// ServeMux adalah router standar bawaan Golang
	mux := http.NewServeMux()

	// Daftarkan Endpoint A: Dashboard Summary
	mux.HandleFunc("/api/v1/dashboard/summary", controller.GetSummaryHandler)

	// Daftarkan Endpoint B: Team Workload
	mux.HandleFunc("/api/v1/dashboard/workload", controller.GetWorkloadHandler)

	return mux
}