package pubsub

import (
	"context"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/awssnssqs" // env var AWS_SDK_LOAD_CONFIG
	_ "gocloud.dev/pubsub/azuresb"   // env var SERVICEBUS_CONNECTION_STRING
	_ "gocloud.dev/pubsub/gcppubsub" // env var GOOGLE_APPLICATION_CREDENTIALS
	_ "gocloud.dev/pubsub/mempubsub"
	_ "gocloud.dev/pubsub/natspubsub"   // env var NATS_SERVER_URL
	_ "gocloud.dev/pubsub/rabbitpubsub" // env var RABBIT_SERVER_URL
)

func OpenTopic(ctx context.Context, url string) (*pubsub.Topic, error) {
	return pubsub.OpenTopic(ctx, url)
}

func OpenSubscription(ctx context.Context, url string) (*pubsub.Subscription, error) {
	return pubsub.OpenSubscription(ctx, url)
}
