package spancontext

import (
	"context"
	"net/http"

	"go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/trace"
)

type AELogSpanContext struct {
	SpanID  string
	TraceID string
}

func Get(ctx context.Context) AELogSpanContext {
	sc := trace.FromContext(ctx).SpanContext()
	return AELogSpanContext{
		SpanID:  sc.SpanID.String(),
		TraceID: sc.TraceID.String(),
	}
}

func Set(r *http.Request, label string) (context.Context, func()) {
	ctx := r.Context()
	span := new(trace.Span)
	httpFormat := propagation.HTTPFormat{}
	if sc, ok := httpFormat.SpanContextFromRequest(r); ok {
		ctx, span = trace.StartSpanWithRemoteParent(ctx, label, sc)
	} else {
		ctx, span = trace.StartSpan(ctx, label)
	}

	return ctx, span.End
}
