package respond

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type H map[string]any

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func JSON(w http.ResponseWriter, status int, data any) {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success":false,"error":"encode failed"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func Success(w http.ResponseWriter, status int, data any) {
	response := H{
		"success": true,
	}

	if data != nil {
		response["data"] = data
	}

	JSON(w, status, response)
}

func Error(w http.ResponseWriter, status int, message string, data any) {
	response := H{
		"success": false,
		"error":   message,
	}

	if data != nil {
		response["data"] = data
	}

	JSON(w, status, response)
}