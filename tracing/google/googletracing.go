package google

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	yabs_trace "github.com/yabslabs/utils/tracing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opencensus.io/plugin/ocgrpc"

	"github.com/yabslabs/utils/logging"

	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"
)

type Tracer struct {
	Exporter *stackdriver.Exporter
	config   *Config
}

func NewGoogleTracing(config *Config) (*Tracer, error) {
	if !envIsSet() {
		return nil, status.Error(codes.InvalidArgument, "env not properly set, GOOGLE_APPLICATION_CREDENTIALS is misconfigured or missing")
	}
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:    config.ProjectID,
		MetricPrefix: config.MetricPrefix,
	})
	logging.Log("GOOG-VWv4F").OnError(err).Error("creating new trace exporter")

	// Register the views to collect server request count.
	err = view.Register(ocgrpc.DefaultServerViews...)
	logging.Log("GOOG-UjfKZ").OnError(err).Error("register failed")

	view.RegisterExporter(exporter)
	view.SetReportingPeriod(60 * time.Second)

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(config.Fraction)})

	return &Tracer{config: config}, err
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

func envIsSet() bool {
	gAuthCred := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	return strings.Contains(gAuthCred, ".json")
}

func (t *Tracer) SetErrStatus(span *trace.Span, code int32, err error, obj ...string) {
	span.SetStatus(trace.Status{Code: code, Message: err.Error() + strings.Join(obj, ", ")})
}
