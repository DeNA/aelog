package log

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/emahiro/ae-plain-logger/spancontext"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// Debugf formats its arguments according to the format, analogous to fmt.Printf,
// and records the text as log message at debug level.
// The log message will be associated with the platform request linked with the context.
func Debugf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "DEBUG", format, a...)
}

// Infof is like Debugf, but the severity is info level.
func Infof(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "INFO", format, a...)
}

// Warningf is like Debugf, but the severity is warning level.
func Warningf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "WARNING", format, a...)
}

// Errorf is like Debugf, but the severity is error level.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "ERROR", format, a...)
}

// Criticalf is like Debugf, but the severity is critical level.
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	out(ctx, "CRITICAL", format, a...)
}

type jsonPayload struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Trace    string `json:"logging.googleapis.com/trace"`
	SpanID   string `json:"logging.googleapis.com/spanId"`
}

func out(ctx context.Context, severity, format string, a ...interface{}) {
	sc := spancontext.Get(ctx)
	payload := &jsonPayload{
		Severity: severity,
		Message:  fmt.Sprintf(format, a...),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", projectID, sc.TraceID),
		SpanID:   sc.SpanID,
	}

	json.NewEncoder(os.Stdout).Encode(payload)
}
