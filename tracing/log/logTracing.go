package log

import (
	"context"
	"net/http"

	"github.com/yabslabs/utils/logging"
	yabs_trace "github.com/yabslabs/utils/tracing"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"

	"go.opencensus.io/trace"

	"go.opencensus.io/stats/view"
)

type Tracer struct {
	config *Config
}

func NewLogTracing(config *Config) (*Tracer, error) {
	if config.Fraction < 1 {
		config.Fraction = 1
	}
	trace.RegisterExporter(&exporter.PrintExporter{})

	err := view.Register(ocgrpc.DefaultServerViews...)
	logging.Log("LOG-jiZkh").OnError(err).Error("register failed")

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(config.Fraction)})

	return &Tracer{config: config}, nil
}

func (t *Tracer) NewServerInterceptorSpan(ctx context.Context, name string) (context.Context, *yabs_trace.Span) {
	return t.newSpanFromName(ctx, name, trace.WithSpanKind(trace.SpanKindServer))
}

func (t *Tracer) NewServerSpan(ctx context.Context, pkg, method string) (context.Context, *yabs_trace.Span) {
	return t.newSpan(ctx, pkg, method, trace.WithSpanKind(trace.SpanKindServer))
}

func (t *Tracer) NewClientInterceptorSpan(ctx context.Context, name string) (context.Context, *yabs_trace.Span) {
	return t.newSpanFromName(ctx, name, trace.WithSpanKind(trace.SpanKindClient))
}

func (t *Tracer) NewClientSpan(ctx context.Context, pkg, method string) (context.Context, *yabs_trace.Span) {
	return t.newSpan(ctx, pkg, method, trace.WithSpanKind(trace.SpanKindClient))
}

func (t *Tracer) NewSpan(ctx context.Context, pkg, method string) (context.Context, *yabs_trace.Span) {
	return t.newSpan(ctx, pkg, method)
}

func (t *Tracer) newSpan(ctx context.Context, pkg, method string, options ...trace.StartOption) (context.Context, *yabs_trace.Span) {
	name := yabs_trace.CreateSpanName(t.config.GitProjectPath, pkg, method)
	return t.newSpanFromName(ctx, name, options...)
}

func (t *Tracer) newSpanFromName(ctx context.Context, name string, options ...trace.StartOption) (context.Context, *yabs_trace.Span) {
	ctx, span := trace.StartSpan(ctx, name, options...)
	return ctx, yabs_trace.NewSpan(span)
}

func (t *Tracer) NewSpanHTTP(r *http.Request, pkg, method string) (*http.Request, *yabs_trace.Span) {
	ctx, span := t.NewSpan(r.Context(), pkg, method)
	r = r.WithContext(ctx)
	return r, span
}
