package google

import (
	"context"

	google_pubsub "cloud.google.com/go/pubsub"
	"github.com/yabslabs/utils/logging"
	"github.com/yabslabs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Google struct {
	*google_pubsub.Client
	defaultSubscribers []*DefaultSubscriber
	subscriberName     string
}

func NewGoogle(ctx context.Context, config *Config) (*Google, error) {
	client, err := google_pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		logging.Log("UTILS-zXtY").WithError(err).Error("connection to google pubsub failed")
		return nil, status.Error(codes.Internal, "connection to google pubsub failed")
	}
	return &Google{
		Client:             client,
		defaultSubscribers: config.DefaultSubscribers,
		subscriberName:     config.SubscriberName,
	}, nil
}

func (g *Google) Topic(ctx context.Context, topicName string) pubsub.Topic {
	return &Topic{Topic: g.Client.Topic(topicName), client: g}
}

func (g *Google) EnsureTopic(ctx context.Context, topicName string) (pubsub.Topic, error) {
	topic := g.Client.Topic(topicName)
	exists, err := topic.Exists(ctx)

	if err != nil {
		logging.Log("UTILS-Qomi").WithError(err).Error("error calling existing topic")
		return nil, err
	}

	if !exists {
		topic, err = g.CreateTopic(ctx, topicName)
		if err != nil {
			logging.Log("UTILS-CtEI").WithError(err).Error("error calling exists topic")
			return nil, err
		}
	}

	return &Topic{
		client: g,
		Topic:  topic,
	}, nil
}

func (g *Google) EnsureTopics(ctx context.Context, topicNames ...string) (pubsub.Topics, error) {
	topics := make(pubsub.Topics, 0, len(topicNames))
	for _, topicName := range topicNames {
		topic, err := g.EnsureTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
