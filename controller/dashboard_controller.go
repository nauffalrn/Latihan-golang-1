package controller

import (
	"encoding/json"
	"net/http"
	"latihan/service"
	"latihan/response"
)

// GetSummaryHandler menangani Endpoint A: /api/v1/dashboard/summary
func GetSummaryHandler(w http.ResponseWriter, r *http.Request) {
	// Pastikan hanya menerima metode GET
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, response.BuildErrorResponse("Method Not Allowed"))
		return
	}

	// 1. Panggil Service untuk mengambil data
	// r.Context() penting untuk meneruskan pembatalan request dari client ke database
	summaryData, err := service.GetDashboardSummary(r.Context())
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, response.BuildErrorResponse("Terjadi kesalahan saat memuat rekap tugas"))
		return
	}

	// 2. Jika sukses, gunakan Global Response untuk mengirim JSON
	successResponse := response.BuildSuccessResponse("Dashboard summary retrieved", summaryData)
	sendJSONResponse(w, http.StatusOK, successResponse)
}

// GetWorkloadHandler menangani Endpoint B: /api/v1/dashboard/workload
func GetWorkloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, response.BuildErrorResponse("Method Not Allowed"))
		return
	}

	// 1. Panggil Service
	workloadData, err := service.GetTeamWorkload(r.Context())
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, response.BuildErrorResponse("Terjadi kesalahan saat memuat data workload"))
		return
	}

	// 2. Balas dengan Response sukses
	successResponse := response.BuildSuccessResponse("Team workload retrieved", workloadData)
	sendJSONResponse(w, http.StatusOK, successResponse)
}

// Helper: sendJSONResponse untuk membungkus pengiriman data ke browser
func sendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	// Beritahu client bahwa kita mengirim data berformat JSON
	w.Header().Set("Content-Type", "application/json")
	// Set kode status HTTP (misal 200 OK, atau 500 Error)
	w.WriteHeader(statusCode)

	// Ubah Struct golang menjadi JSON text
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, `{"success":false,"message":"Gagal memproses JSON response"}`, http.StatusInternalServerError)
		return
	}

	// Kirim string JSON ke client
	w.Write(jsonBytes)
}