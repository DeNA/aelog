package log

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.opencensus.io/trace"
)

const (
	LevelDebug    = "DEBUG"
	LevelInfo     = "INFO"
	LevelWarning  = "WARNING"
	LevelError    = "ERROR"
	LevelCritical = "CRITICAL"
)

// Debugf is output of debug level log
func Debugf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, LevelDebug, format, a)
}

// Infof is output of infomation level log
func Infof(ctx context.Context, format string, a ...interface{}) {
	out(ctx, LevelInfo, format, a)
}

// Warningf is output of warning level log
func Warningf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, LevelWarning, format, a)
}

// Errorf is output of error level log
func Errorf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, LevelError, format, a)
}

// Criticalf is output of critical level log
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, LevelCritical, format, a)
}

type jsonPayload struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`
	SpanID   string `json:"logging.googleapis.com/spanId,omitempty"`
}

func out(ctx context.Context, severity, format string, a ...interface{}) {
	sc := trace.FromContext(ctx).SpanContext()
	payload := &jsonPayload{
		Severity: severity,
		Message:  fmt.Sprintf(format, a),
		Trace:    sc.TraceID.String(),
		SpanID:   sc.SpanID.String(),
	}

	json.NewEncoder(os.Stdout).Encode(payload)
}
