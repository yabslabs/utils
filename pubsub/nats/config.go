package nats

import (
	"context"
	"regexp"

	"github.com/yabslabs/utils/pubsub"
)

type Config struct {
	ClusterID          string
	URL                string
	DefaultSubscribers []*DefaultSubscriber
	SubscriberName     string
}

type DefaultSubscriber struct {
	Servicename string
	Topic       *regexp.Regexp
}

func (c *Config) NewPubsub() (pubsub.Pubsub, error) {
	return NewNats(context.Background(), c)
}
