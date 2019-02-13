package nats

import (
	"context"

	"github.com/yabslabs/utils/pubsub"

	stan "github.com/nats-io/go-nats-streaming"
)

const subscriptionIDFormat = "%s-%s"

type Subscription struct {
	client *Nats
	stan.Subscription
	durableName string
	topic       *Topic
}

func (s *Subscription) Receive(ctx context.Context, handleFunc interface{}) (err error) {
	msgHandler, ok := handleFunc.(stan.MsgHandler)
	if !ok {
		return pubsub.ErrHandleFuncWrongType
	}
	s.Subscription, err = s.client.Subscribe(s.topic.name, msgHandler)
	return err
}
