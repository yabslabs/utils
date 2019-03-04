package tracing

import (
	"context"
	"net/http"
)

type Tracer interface {
	NewSpan(ctx context.Context, pkg, method string) (context.Context, *Span)
	NewClientSpan(ctx context.Context, pkg, method string) (context.Context, *Span)
	NewServerSpan(ctx context.Context, pkg, method string) (context.Context, *Span)
	NewClientInterceptorSpan(ctx context.Context, name string) (context.Context, *Span)
	NewServerInterceptorSpan(ctx context.Context, name string) (context.Context, *Span)
	NewSpanHTTP(r *http.Request, pkg, method string) (*http.Request, *Span)
}

type Config interface {
	NewTracer(ctx context.Context) (Tracer, error)
}
