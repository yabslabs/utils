// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yabslabs/utils/tracing (interfaces: Tracer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	tracing "github.com/yabslabs/utils/tracing"
	http "net/http"
	reflect "reflect"
)

// MockTracer is a mock of Tracer interface
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// NewClientInterceptorSpan mocks base method
func (m *MockTracer) NewClientInterceptorSpan(arg0 context.Context, arg1 string) (context.Context, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewClientInterceptorSpan", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewClientInterceptorSpan indicates an expected call of NewClientInterceptorSpan
func (mr *MockTracerMockRecorder) NewClientInterceptorSpan(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClientInterceptorSpan", reflect.TypeOf((*MockTracer)(nil).NewClientInterceptorSpan), arg0, arg1)
}

// NewClientSpan mocks base method
func (m *MockTracer) NewClientSpan(arg0 context.Context, arg1, arg2 string) (context.Context, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewClientSpan", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewClientSpan indicates an expected call of NewClientSpan
func (mr *MockTracerMockRecorder) NewClientSpan(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClientSpan", reflect.TypeOf((*MockTracer)(nil).NewClientSpan), arg0, arg1, arg2)
}

// NewServerInterceptorSpan mocks base method
func (m *MockTracer) NewServerInterceptorSpan(arg0 context.Context, arg1 string) (context.Context, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewServerInterceptorSpan", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewServerInterceptorSpan indicates an expected call of NewServerInterceptorSpan
func (mr *MockTracerMockRecorder) NewServerInterceptorSpan(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewServerInterceptorSpan", reflect.TypeOf((*MockTracer)(nil).NewServerInterceptorSpan), arg0, arg1)
}

// NewServerSpan mocks base method
func (m *MockTracer) NewServerSpan(arg0 context.Context, arg1, arg2 string) (context.Context, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewServerSpan", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewServerSpan indicates an expected call of NewServerSpan
func (mr *MockTracerMockRecorder) NewServerSpan(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewServerSpan", reflect.TypeOf((*MockTracer)(nil).NewServerSpan), arg0, arg1, arg2)
}

// NewSpan mocks base method
func (m *MockTracer) NewSpan(arg0 context.Context, arg1, arg2 string) (context.Context, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewSpan", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewSpan indicates an expected call of NewSpan
func (mr *MockTracerMockRecorder) NewSpan(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSpan", reflect.TypeOf((*MockTracer)(nil).NewSpan), arg0, arg1, arg2)
}

// NewSpanHTTP mocks base method
func (m *MockTracer) NewSpanHTTP(arg0 *http.Request, arg1, arg2 string) (*http.Request, *tracing.Span) {
	ret := m.ctrl.Call(m, "NewSpanHTTP", arg0, arg1, arg2)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(*tracing.Span)
	return ret0, ret1
}

// NewSpanHTTP indicates an expected call of NewSpanHTTP
func (mr *MockTracerMockRecorder) NewSpanHTTP(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSpanHTTP", reflect.TypeOf((*MockTracer)(nil).NewSpanHTTP), arg0, arg1, arg2)
}