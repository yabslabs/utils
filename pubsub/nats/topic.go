package nats

import (
	"context"
	"fmt"
	"sync"

	"github.com/yabslabs/utils/logging"

	stan "github.com/nats-io/go-nats-streaming"

	"github.com/pkg/errors"

	"github.com/yabslabs/utils/pubsub"
)

type Topic struct {
	client *Nats
	name   string
}

func (t *Topic) Publish(ctx context.Context, data []byte, attsKV ...string) (_ string, err error) {
	var wg sync.WaitGroup
	ackHandler := func(ackedNuid string, ackErr error) {
		wg.Add(1)
		err = ackErr
	}
	id, publishErr := t.client.PublishAsync(t.name, data, ackHandler)
	if publishErr != nil {
		return "", publishErr
	}
	wg.Wait()
	if err != nil {
		return "", err
	}
	return id, nil
}
func (t *Topic) PublishAsync(ctx context.Context, data []byte, attsKV ...string) {
	ackHandler := func(ackedNuid string, err error) {
		logging.Log("NATS-JQkG").OnError(err).Warn("message not published")
	}
	t.client.PublishAsync(t.name, data, ackHandler)
}
func (t *Topic) EnsureSubscription(ctx context.Context) (pubsub.Subscription, error) {
	return t.EnsureSubscriptionForSubscriber(ctx, t.client.subscriberName)
}

func (t *Topic) EnsureSubscriptionForSubscriber(ctx context.Context, subscriberName string) (pubsub.Subscription, error) {
	durableName := fmt.Sprintf(subscriptionIDFormat, t.name, subscriberName)
	opts := []stan.SubscriptionOption{
		stan.DurableName(durableName),
		stan.StartWithLastReceived(),
		stan.SetManualAckMode(),
	}
	dummy := func(msg *stan.Msg) {}
	dummySubscription, err := t.client.Subscribe(t.name, dummy, opts...)
	if err != nil || !dummySubscription.IsValid() {
		return nil, errors.Wrap(err, "subscribe failed")
	}

	err = dummySubscription.Unsubscribe()
	logging.Log("UTILS-lvJN").OnError(err).Warn("unsubscribe of dummy subscription failed")

	return &Subscription{client: t.client, durableName: durableName, topic: t}, nil
}
