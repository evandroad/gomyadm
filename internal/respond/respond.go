package respond

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type H map[string]any

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"success":false,"error":"encode failed"}`))
		return
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	JSON(w, status, response)
}

func Error(w http.ResponseWriter, status int, message string, data any) {
	response := Response{
		Success: false,
		Error:   message,
		Data:    data,
	}

	JSON(w, status, response)
}