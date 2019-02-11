package google

import (
	"context"
	"fmt"

	google_pubsub "cloud.google.com/go/pubsub"
	"github.com/pkg/errors"

	"github.com/yabslabs/utils/pairs"
	"github.com/yabslabs/utils/pubsub"
)

type Topic struct {
	client *Google
	*google_pubsub.Topic
}

func (t *Topic) Publish(ctx context.Context, data []byte, attsKV ...string) (id string, err error) {
	msg := &google_pubsub.Message{Data: data, Attributes: pairs.PairsString(attsKV...)}
	return t.Topic.Publish(ctx, msg).Get(ctx)
}
func (t *Topic) PublishAsync(ctx context.Context, data []byte, attsKV ...string) {
	msg := &google_pubsub.Message{Data: data, Attributes: pairs.PairsString(attsKV...)}
	t.Topic.Publish(ctx, msg)
	return
}
func (t *Topic) EnsureSubscription(ctx context.Context) (pubsub.Subscription, error) {
	return t.EnsureSubscriptionForSubscriber(ctx, t.client.subscriberName)
}

func (t *Topic) EnsureSubscriptionForSubscriber(ctx context.Context, subscriberName string) (pubsub.Subscription, error) {
	id := fmt.Sprintf(subscriptionIDFormat, t.ID(), subscriberName)
	subscription := t.client.Subscription(id)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "subscription: check failed")
	}
	if exists {
		return &Subscription{Subscription: subscription, client: t.client}, nil
	}
	cfg := google_pubsub.SubscriptionConfig{Topic: t.Topic}
	subscription, err = t.client.CreateSubscription(ctx, id, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "subscription: create failed")
	}
	return &Subscription{client: t.client, Subscription: subscription}, nil
}
