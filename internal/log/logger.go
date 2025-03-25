package log

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"
)

type JSONLogger struct {
	out *os.File
}

func NewJSONLogger(out *os.File) *JSONLogger {
	return &JSONLogger{out: out}
}

func (l *JSONLogger) Log(r *http.Request, status int, duration time.Duration) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	entry := map[string]interface{}{
		"remote_ip":    ip,
		"method":       r.Method,
		"uri":          r.RequestURI,
		"status":       status,
		"duration_ms":  duration.Milliseconds(),
		"user_agent":   r.UserAgent(),
		"timestamp":    time.Now().Format(time.RFC3339),
		"proto":        r.Proto,
		"content_type": r.Header.Get("Content-Type"),
	}

	_ = json.NewEncoder(l.out).Encode(entry)
}
