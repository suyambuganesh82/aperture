// Code generated by MockGen. DO NOT EDIT.
// Source: limiter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	checkv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/flowcontrol/check/v1"
	languagev1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	labels "github.com/fluxninja/aperture/v2/pkg/labels"
	iface "github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/iface"
	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

// MockLimiter is a mock of Limiter interface.
type MockLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockLimiterMockRecorder
}

// MockLimiterMockRecorder is the mock recorder for MockLimiter.
type MockLimiterMockRecorder struct {
	mock *MockLimiter
}

// NewMockLimiter creates a new mock instance.
func NewMockLimiter(ctrl *gomock.Controller) *MockLimiter {
	mock := &MockLimiter{ctrl: ctrl}
	mock.recorder = &MockLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLimiter) EXPECT() *MockLimiterMockRecorder {
	return m.recorder
}

// Decide mocks base method.
func (m *MockLimiter) Decide(arg0 context.Context, arg1 labels.Labels) *checkv1.LimiterDecision {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decide", arg0, arg1)
	ret0, _ := ret[0].(*checkv1.LimiterDecision)
	return ret0
}

// Decide indicates an expected call of Decide.
func (mr *MockLimiterMockRecorder) Decide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decide", reflect.TypeOf((*MockLimiter)(nil).Decide), arg0, arg1)
}

// GetLimiterID mocks base method.
func (m *MockLimiter) GetLimiterID() iface.LimiterID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimiterID")
	ret0, _ := ret[0].(iface.LimiterID)
	return ret0
}

// GetLimiterID indicates an expected call of GetLimiterID.
func (mr *MockLimiterMockRecorder) GetLimiterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimiterID", reflect.TypeOf((*MockLimiter)(nil).GetLimiterID))
}

// GetPolicyName mocks base method.
func (m *MockLimiter) GetPolicyName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPolicyName indicates an expected call of GetPolicyName.
func (mr *MockLimiterMockRecorder) GetPolicyName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyName", reflect.TypeOf((*MockLimiter)(nil).GetPolicyName))
}

// GetRequestCounter mocks base method.
func (m *MockLimiter) GetRequestCounter(labels map[string]string) prometheus.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestCounter", labels)
	ret0, _ := ret[0].(prometheus.Counter)
	return ret0
}

// GetRequestCounter indicates an expected call of GetRequestCounter.
func (mr *MockLimiterMockRecorder) GetRequestCounter(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestCounter", reflect.TypeOf((*MockLimiter)(nil).GetRequestCounter), labels)
}

// GetSelectors mocks base method.
func (m *MockLimiter) GetSelectors() []*languagev1.Selector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectors")
	ret0, _ := ret[0].([]*languagev1.Selector)
	return ret0
}

// GetSelectors indicates an expected call of GetSelectors.
func (mr *MockLimiterMockRecorder) GetSelectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectors", reflect.TypeOf((*MockLimiter)(nil).GetSelectors))
}

// Revert mocks base method.
func (m *MockLimiter) Revert(arg0 context.Context, arg1 labels.Labels, arg2 *checkv1.LimiterDecision) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Revert", arg0, arg1, arg2)
}

// Revert indicates an expected call of Revert.
func (mr *MockLimiterMockRecorder) Revert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockLimiter)(nil).Revert), arg0, arg1, arg2)
}

// MockRateLimiter is a mock of RateLimiter interface.
type MockRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimiterMockRecorder
}

// MockRateLimiterMockRecorder is the mock recorder for MockRateLimiter.
type MockRateLimiterMockRecorder struct {
	mock *MockRateLimiter
}

// NewMockRateLimiter creates a new mock instance.
func NewMockRateLimiter(ctrl *gomock.Controller) *MockRateLimiter {
	mock := &MockRateLimiter{ctrl: ctrl}
	mock.recorder = &MockRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateLimiter) EXPECT() *MockRateLimiterMockRecorder {
	return m.recorder
}

// Decide mocks base method.
func (m *MockRateLimiter) Decide(arg0 context.Context, arg1 labels.Labels) *checkv1.LimiterDecision {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decide", arg0, arg1)
	ret0, _ := ret[0].(*checkv1.LimiterDecision)
	return ret0
}

// Decide indicates an expected call of Decide.
func (mr *MockRateLimiterMockRecorder) Decide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decide", reflect.TypeOf((*MockRateLimiter)(nil).Decide), arg0, arg1)
}

// GetLimiterID mocks base method.
func (m *MockRateLimiter) GetLimiterID() iface.LimiterID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimiterID")
	ret0, _ := ret[0].(iface.LimiterID)
	return ret0
}

// GetLimiterID indicates an expected call of GetLimiterID.
func (mr *MockRateLimiterMockRecorder) GetLimiterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimiterID", reflect.TypeOf((*MockRateLimiter)(nil).GetLimiterID))
}

// GetPolicyName mocks base method.
func (m *MockRateLimiter) GetPolicyName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPolicyName indicates an expected call of GetPolicyName.
func (mr *MockRateLimiterMockRecorder) GetPolicyName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyName", reflect.TypeOf((*MockRateLimiter)(nil).GetPolicyName))
}

// GetRequestCounter mocks base method.
func (m *MockRateLimiter) GetRequestCounter(labels map[string]string) prometheus.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestCounter", labels)
	ret0, _ := ret[0].(prometheus.Counter)
	return ret0
}

// GetRequestCounter indicates an expected call of GetRequestCounter.
func (mr *MockRateLimiterMockRecorder) GetRequestCounter(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestCounter", reflect.TypeOf((*MockRateLimiter)(nil).GetRequestCounter), labels)
}

// GetSelectors mocks base method.
func (m *MockRateLimiter) GetSelectors() []*languagev1.Selector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectors")
	ret0, _ := ret[0].([]*languagev1.Selector)
	return ret0
}

// GetSelectors indicates an expected call of GetSelectors.
func (mr *MockRateLimiterMockRecorder) GetSelectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectors", reflect.TypeOf((*MockRateLimiter)(nil).GetSelectors))
}

// Revert mocks base method.
func (m *MockRateLimiter) Revert(arg0 context.Context, arg1 labels.Labels, arg2 *checkv1.LimiterDecision) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Revert", arg0, arg1, arg2)
}

// Revert indicates an expected call of Revert.
func (mr *MockRateLimiterMockRecorder) Revert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockRateLimiter)(nil).Revert), arg0, arg1, arg2)
}

// TakeIfAvailable mocks base method.
func (m *MockRateLimiter) TakeIfAvailable(ctx context.Context, labels labels.Labels, count float64) (string, bool, time.Duration, float64, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeIfAvailable", ctx, labels, count)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(time.Duration)
	ret3, _ := ret[3].(float64)
	ret4, _ := ret[4].(float64)
	return ret0, ret1, ret2, ret3, ret4
}

// TakeIfAvailable indicates an expected call of TakeIfAvailable.
func (mr *MockRateLimiterMockRecorder) TakeIfAvailable(ctx, labels, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeIfAvailable", reflect.TypeOf((*MockRateLimiter)(nil).TakeIfAvailable), ctx, labels, count)
}

// MockScheduler is a mock of Scheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// Decide mocks base method.
func (m *MockScheduler) Decide(arg0 context.Context, arg1 labels.Labels) *checkv1.LimiterDecision {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decide", arg0, arg1)
	ret0, _ := ret[0].(*checkv1.LimiterDecision)
	return ret0
}

// Decide indicates an expected call of Decide.
func (mr *MockSchedulerMockRecorder) Decide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decide", reflect.TypeOf((*MockScheduler)(nil).Decide), arg0, arg1)
}

// GetFlowDurationSummary mocks base method.
func (m *MockScheduler) GetFlowDurationSummary(labels map[string]string) prometheus.Observer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowDurationSummary", labels)
	ret0, _ := ret[0].(prometheus.Observer)
	return ret0
}

// GetFlowDurationSummary indicates an expected call of GetFlowDurationSummary.
func (mr *MockSchedulerMockRecorder) GetFlowDurationSummary(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowDurationSummary", reflect.TypeOf((*MockScheduler)(nil).GetFlowDurationSummary), labels)
}

// GetLatencyObserver mocks base method.
func (m *MockScheduler) GetLatencyObserver(labels map[string]string) prometheus.Observer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatencyObserver", labels)
	ret0, _ := ret[0].(prometheus.Observer)
	return ret0
}

// GetLatencyObserver indicates an expected call of GetLatencyObserver.
func (mr *MockSchedulerMockRecorder) GetLatencyObserver(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatencyObserver", reflect.TypeOf((*MockScheduler)(nil).GetLatencyObserver), labels)
}

// GetLimiterID mocks base method.
func (m *MockScheduler) GetLimiterID() iface.LimiterID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLimiterID")
	ret0, _ := ret[0].(iface.LimiterID)
	return ret0
}

// GetLimiterID indicates an expected call of GetLimiterID.
func (mr *MockSchedulerMockRecorder) GetLimiterID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLimiterID", reflect.TypeOf((*MockScheduler)(nil).GetLimiterID))
}

// GetPolicyName mocks base method.
func (m *MockScheduler) GetPolicyName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPolicyName indicates an expected call of GetPolicyName.
func (mr *MockSchedulerMockRecorder) GetPolicyName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyName", reflect.TypeOf((*MockScheduler)(nil).GetPolicyName))
}

// GetRequestCounter mocks base method.
func (m *MockScheduler) GetRequestCounter(labels map[string]string) prometheus.Counter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestCounter", labels)
	ret0, _ := ret[0].(prometheus.Counter)
	return ret0
}

// GetRequestCounter indicates an expected call of GetRequestCounter.
func (mr *MockSchedulerMockRecorder) GetRequestCounter(labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestCounter", reflect.TypeOf((*MockScheduler)(nil).GetRequestCounter), labels)
}

// GetSelectors mocks base method.
func (m *MockScheduler) GetSelectors() []*languagev1.Selector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectors")
	ret0, _ := ret[0].([]*languagev1.Selector)
	return ret0
}

// GetSelectors indicates an expected call of GetSelectors.
func (mr *MockSchedulerMockRecorder) GetSelectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectors", reflect.TypeOf((*MockScheduler)(nil).GetSelectors))
}

// Revert mocks base method.
func (m *MockScheduler) Revert(arg0 context.Context, arg1 labels.Labels, arg2 *checkv1.LimiterDecision) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Revert", arg0, arg1, arg2)
}

// Revert indicates an expected call of Revert.
func (mr *MockSchedulerMockRecorder) Revert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockScheduler)(nil).Revert), arg0, arg1, arg2)
}
