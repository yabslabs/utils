package mock

import (
	context "context"
	"testing"

	"github.com/golang/mock/gomock"

	"git.workshop21.ch/go/abraxas/tracing"
)

func NewSimpleMockTracer(t *testing.T) *MockTracing {
	return NewMockTracing(gomock.NewController(t))
}

func ExpectServerSpan(ctx context.Context, mock interface{}) {
	m := mock.(*MockTracing)
	any := gomock.Any()
	m.EXPECT().NewServerSpan(any, any, any).AnyTimes().Return(ctx, &tracing.Span{})
}
