package response

// BaseResponse adalah format standar API untuk seluruh endpoint di proyek ini.
// Kita menggunakan interface{} untuk `Data` agar bisa menerima tipe data apa pun 
// (bisa struct TaskSummary, atau slice/array Workload).
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"` // omitempty: jika kosong, field ini tidak akan dikirim di JSON
	Data    interface{} `json:"data,omitempty"`
}

// BuildSuccessResponse adalah helper function untuk mempermudah pembuatan respons sukses.
// Fungsi ini dipanggil dari Controller.
func BuildSuccessResponse(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// BuildErrorResponse adalah helper function untuk mempercepat pembuatan respons error.
func BuildErrorResponse(message string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
	}
}