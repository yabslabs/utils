package config

import (
	"encoding/json"

	"github.com/yabslabs/utils/pubsub"
	"github.com/yabslabs/utils/pubsub/google"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PubsubConfig struct {
	ID     string
	Type   string
	Config pubsub.Config
}

var caches = map[string]func() pubsub.Config{
	"google": func() pubsub.Config { return &google.Config{} },
}

func (c *PubsubConfig) UnmarshalJSON(data []byte) error {
	var rc struct {
		ID     string
		Type   string
		Config json.RawMessage
	}

	if err := json.Unmarshal(data, &rc); err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "PUBSUB-vmjS", err)
	}

	c.Type = rc.Type
	c.ID = rc.ID

	var err error
	c.Config, err = newPubsubConfig(c.Type, rc.Config)
	if err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "PUBSUB-Ws9E", err)
	}

	return nil
}

func newPubsubConfig(cacheType string, configData []byte) (pubsub.Config, error) {
	t, ok := caches[cacheType]
	if !ok {
		return nil, status.Errorf(codes.Internal, "%v No config: %v", "PUBSUB-HMEJ", cacheType)
	}

	PubsubConfig := t()
	if len(configData) == 0 {
		return PubsubConfig, nil
	}

	if err := json.Unmarshal(configData, PubsubConfig); err != nil {
		return nil, status.Errorf(codes.Internal, "%v Could not read conifg: %v", "PUBSUB-1tSS", err)
	}

	return PubsubConfig, nil
}
