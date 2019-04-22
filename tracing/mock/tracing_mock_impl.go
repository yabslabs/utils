package mock

import (
	context "context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/yabslabs/utils/tracing"
)

func NewSimpleMockTracer(t *testing.T) *MockTracer {
	return NewMockTracer(gomock.NewController(t))
}

func ExpectServerSpan(ctx context.Context, mock interface{}) {
	m := mock.(*MockTracer)
	any := gomock.Any()
	m.EXPECT().NewServerSpan(any, any, any).AnyTimes().Return(ctx, &tracing.Span{})
}
