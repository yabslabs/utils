package google

import (
	"context"

	ggl_pubsub "cloud.google.com/go/pubsub"
	"github.com/yabslabs/utils/logging"
	"github.com/yabslabs/utils/pubsub"
)

const subscriptionIDFormat = "%s-%s"

type Subscription struct {
	client *Google
	*ggl_pubsub.Subscription
}

func (s *Subscription) Receive(ctx context.Context, handleFunc interface{}) error {
	typedHandleFunc, ok := handleFunc.(func(context.Context, *ggl_pubsub.Message))
	if !ok {
		return pubsub.ErrHandleFuncWrongType
	}
	go func() {
		err := s.Subscription.Receive(ctx, typedHandleFunc)
		logging.Log("UTILS-9t61").OnError(err).Warn("receive failed")
	}()
	return nil
}
