package google

import (
	"context"
	"fmt"

	google_pubsub "cloud.google.com/go/pubsub"
	"github.com/yabslabs/utils/logging"
	"github.com/yabslabs/utils/pairs"
	"github.com/yabslabs/utils/pubsub"
)

type Topic struct {
	client *Google
	*google_pubsub.Topic
}

type Topics []Topic

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
	logging.Log("UTILS-ySLW").OnError(err).Debug("error checking subscription")
	if exists {
		return &Subscription{Subscription: subscription, client: t.client}, nil
	}
	cfg := new(google_pubsub.SubscriptionConfig)
	subscription, err = t.client.CreateSubscription(ctx, id, *cfg)
	if err != nil {
		return nil, err
	}
	return &Subscription{client: t.client, Subscription: subscription}, nil
}
