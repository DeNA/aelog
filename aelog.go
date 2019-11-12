/*
Package aelog provides the structured log of an application's logs
from within an App Engine application.

Example:
	ctx := r.Context() // r is *http.Request
	aelog.Infof(ctx, "info log. requestURL: %s", r.URL.String())
*/
package aelog

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/emahiro/aelog/internal/spancontext"
)

type logger struct {
	mu  sync.Mutex
	out func(context.Context, string, string, ...interface{})
}

func (l *logger) output(ctx context.Context, severity, format string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out(ctx, severity, format, a...)
}

// SetOutput allows to set the log output format. The default log setting is json format.
func SetOutput(f func(context.Context, string, string, ...interface{})) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.out = f
}

var std = &logger{out: OutputJSON}

// Debugf formats its arguments according to the format, analogous to fmt.Printf,
// and records the text as log message at debug level.
// The log message will be associated with the platform request linked with the context.
func Debugf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "DEBUG", format, a...)
}

// Infof is like Debugf, but the severity is info level.
func Infof(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "INFO", format, a...)
}

// Warningf is like Debugf, but the severity is warning level.
func Warningf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "WARNING", format, a...)
}

// Errorf is like Debugf, but the severity is error level.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "ERROR", format, a...)
}

// Criticalf is like Debugf, but the severity is critical level.
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "CRITICAL", format, a...)
}

// OutputJSON fills the fields required with the structured log described in
// https://cloud.google.com/logging/docs/agent/configuration#special-fields
// and outputs the log in json format.
func OutputJSON(ctx context.Context, severity, format string, a ...interface{}) {
	sc := spancontext.Get(ctx)
	payload := &struct {
		Severity string `json:"severity"`
		Message  string `json:"message"`
		Trace    string `json:"logging.googleapis.com/trace"`
		SpanID   string `json:"logging.googleapis.com/spanId"`
	}{
		Severity: severity,
		Message:  fmt.Sprintf(format, a...),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", os.Getenv("GOOGLE_CLOUD_PROJECT"), sc.TraceID),
		SpanID:   sc.SpanID,
	}

	json.NewEncoder(os.Stdout).Encode(payload)
}

// OutputText outputs log using standard log package.
func OutputText(ctx context.Context, severity, format string, a ...interface{}) {
	log.Printf(severity+": "+format, a...)
}
