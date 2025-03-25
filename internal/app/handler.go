package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go-app/internal/parser"
	"go-app/internal/response"
)

func (a *App) HandleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	query := r.URL.Query()

	if waitStr := query.Get("wait"); waitStr != "" {
		waitDuration, err := parser.ParseDuration(waitStr)
		if err != nil {
			http.Error(w, `{"error":"invalid wait"}`, http.StatusBadRequest)
			return
		}
		time.Sleep(waitDuration)
	}

	statusCode := http.StatusOK
	if statusStr := query.Get("status"); statusStr != "" {
		code, err := parser.ParseStatus(statusStr)
		if err != nil {
			http.Error(w, `{"error":"invalid status"}`, http.StatusBadRequest)
			return
		}
		statusCode = code
	}

	body := ""
	if sizeStr := query.Get("response_size"); sizeStr != "" {
		targetSize, err := parser.ParseSize(sizeStr)
		if err != nil {
			http.Error(w, `{"error":"invalid response_size"}`, http.StatusBadRequest)
			return
		}
		if targetSize > 0 {
			body = strings.Repeat("A", targetSize)
		}
	}

	duration := time.Since(start)
	resp := response.JSON{
		Status:   statusCode,
		URI:      r.RequestURI,
		Duration: duration.String(),
		Body:     body,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)

	a.logger.Log(r, statusCode, duration)
}
