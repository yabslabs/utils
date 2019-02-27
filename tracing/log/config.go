package log

import (
	"context"

	yabs_trace "github.com/yabslabs/utils/tracing"
)

type Config struct {
	Fraction       float64
	GitProjectPath string
}

func (c *Config) NewTracer(ctx context.Context) (yabs_trace.Tracing, error) {
	return NewLogTracing(c)
}
