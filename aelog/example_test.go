package aelog

import (
	"context"

	"go.opencensus.io/trace"
)

func ExampleDebugf() {
	ctx := trace.NewContext(context.Background(), nil)
	Debugf(ctx, "Debugf: %s", "Sample Text")

	// Output:
	// {"severity":"DEBUG","message":"Debugf: Sample Text","logging.googleapis.com/trace":"projects//traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000"}
}

func ExampleInfof() {
	ctx := trace.NewContext(context.Background(), nil)
	Debugf(ctx, "Infof: %s", "Sample Text")

	// Output:
	// {"severity":"INFO","message":"Infof: Sample Text","logging.googleapis.com/trace":"projects//traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000"}
}

func ExampleWarningf() {
	ctx := trace.NewContext(context.Background(), nil)
	Warningf(ctx, "Warningf: %s", "Sample Text")

	// Output:
	// {"severity":"WARNING","message":"Warningf: Sample Text","logging.googleapis.com/trace":"projects//traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000"}
}

func ExampleErrorf() {
	ctx := trace.NewContext(context.Background(), nil)
	Errorf(ctx, "Errorf: %s", "Sample Text")

	// Output:
	// {"severity":"ERROR","message":"Errorf: Sample Text","logging.googleapis.com/trace":"projects//traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000"}
}

func ExampleCriticalf() {
	ctx := trace.NewContext(context.Background(), nil)
	Criticalf(ctx, "Criticalf: %s", "Sample Text")

	// Output:
	// {"severity":"CRITICAL","message":"Criticalf: Sample Text","logging.googleapis.com/trace":"projects//traces/00000000000000000000000000000000","logging.googleapis.com/spanId":"0000000000000000"}
}
