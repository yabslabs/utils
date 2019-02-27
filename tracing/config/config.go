package config

import (
	"encoding/json"

	cmn_trace "git.workshop21.ch/go/abraxas/tracing"
	cmn_trace_g "git.workshop21.ch/go/abraxas/tracing/google"
	cmn_trace_log "git.workshop21.ch/go/abraxas/tracing/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TracingConfig struct {
	Type   string
	Config cmn_trace.Config
}

var tracer = map[string]func() cmn_trace.Config{
	"google": func() cmn_trace.Config { return &cmn_trace_g.Config{} },
	"log":    func() cmn_trace.Config { return &cmn_trace_log.Config{} },
}

func (c *TracingConfig) UnmarshalJSON(data []byte) error {
	var rc struct {
		Type   string
		Config json.RawMessage
	}

	if err := json.Unmarshal(data, &rc); err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "TRACE-vmjS", err)
	}

	c.Type = rc.Type

	var err error
	c.Config, err = newTracingConfig(c.Type, rc.Config)
	if err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "TRACE-Ws9E", err)
	}

	return nil
}

func newTracingConfig(tracerType string, configData []byte) (cmn_trace.Config, error) {
	t, ok := tracer[tracerType]
	if !ok {
		return nil, status.Errorf(codes.Internal, "%v No config: %v", "TRACE-HMEJ", tracerType)
	}

	tracingConfig := t()
	if len(configData) == 0 {
		return tracingConfig, nil
	}

	if err := json.Unmarshal(configData, tracingConfig); err != nil {
		return nil, status.Errorf(codes.Internal, "%v Could not read conifg: %v", "TRACE-1tSS", err)
	}

	return tracingConfig, nil
}
