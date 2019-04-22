package google

import (
	"context"

	yabs_trace "github.com/yabslabs/utils/tracing"
)

type Config struct {
	ProjectID      string
	MetricPrefix   string
	Fraction       float64
	GitProjectPath string
}

func (c *Config) NewTracer(ctx context.Context) (yabs_trace.Tracer, error) {
	return NewGoogleTracing(c)
}
