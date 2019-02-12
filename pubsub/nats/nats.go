package nats

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yabslabs/utils/logging"
	"github.com/yabslabs/utils/pubsub"

	stan "github.com/nats-io/go-nats-streaming"
)

type Nats struct {
	stan.Conn
	subscriberName string
}

func NewNats(ctx context.Context, config *Config) (*Nats, error) {
	conn, err := stan.Connect(config.ClusterID, config.SubscriberName, stan.NatsURL(config.URL))
	if err != nil {
		logging.Log("NATS-zXtY").WithError(err).Error("connection to gats pubsub failed")
		return nil, status.Error(codes.Internal, "connection to gats pubsub failed")
	}
	return &Nats{
		Conn:           conn,
		subscriberName: config.SubscriberName,
	}, nil
}

func (n *Nats) Topic(ctx context.Context, topicName string) pubsub.Topic {
	return &Topic{client: n, name: topicName}
}

func (n *Nats) EnsureTopic(ctx context.Context, topicName string) (pubsub.Topic, error) {
	topic := n.Topic(ctx, topicName)
	return topic, nil
}

func (n *Nats) EnsureTopics(ctx context.Context, topicNames ...string) (pubsub.Topics, error) {
	topics := make(pubsub.Topics, 0, len(topicNames))
	for _, topicName := range topicNames {
		topic, err := n.EnsureTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
