package pubsub_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yabslabs/utils/pubsub"
)

func TestNatsTopic(t *testing.T) {
	ctx := context.Background()
	os.Setenv("NATS_SERVER_URL", "nats://localhost:4222")
	topic, err := pubsub.OpenTopic(ctx, "nats://testTopic")
	assert.NoError(t, err)
	assert.NotNil(t, topic)
}
