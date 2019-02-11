package google

import (
	"context"
	"regexp"

	"github.com/yabslabs/utils/pubsub"
)

type Config struct {
	ProjectID          string
	DefaultSubscribers []*DefaultSubscriber
	SubscriberName     string
}

type DefaultSubscriber struct {
	Servicename string
	Topic       *regexp.Regexp
}

func (c *Config) NewPubsub() (pubsub.Pubsub, error) {
	return NewGoogle(context.Background(), c)
}
