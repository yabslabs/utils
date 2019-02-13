package pubsub

import (
	"context"
	"errors"
)

var ErrHandleFuncWrongType = errors.New("handle function has wrong type")

type Pubsub interface {
	EnsureTopic(ctx context.Context, topicName string) (Topic, error)
	EnsureTopics(ctx context.Context, topicNames ...string) (Topics, error)
	Topic(ctx context.Context, topicName string) Topic
}

type Topic interface {
	Publish(ctx context.Context, data []byte, attsKV ...string) (string, error)
	PublishAsync(ctx context.Context, data []byte, attsKV ...string)
	EnsureSubscription(ctx context.Context) (Subscription, error)
	EnsureSubscriptionForSubscriber(ctx context.Context, subscriberName string) (Subscription, error)
}

type Topics []Topic

type Subscription interface {
	Receive(ctx context.Context, handleFunc interface{}) error
}

type Config interface {
	NewPubsub() (Pubsub, error)
}
