package google

import (
	"context"

	"github.com/yabslabs/utils/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/pubsub"
)

const subscriptionIDFormat = "%s-%s"

type Subscription struct {
	client *Google
	*pubsub.Subscription
}

func (s *Subscription) Receive(ctx context.Context, handleFunc interface{}) error {
	typedHandleFunc, ok := handleFunc.(func(context.Context, *pubsub.Message))
	if !ok {
		return status.Error(codes.InvalidArgument, "handle function has wrong type")
	}
	go func() {
		err := s.Subscription.Receive(ctx, typedHandleFunc)
		logging.Log("UTILS-9t61").OnError(err).Warn("receive failed")
	}()
	return nil
}
