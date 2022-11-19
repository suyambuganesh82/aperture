// Code generated by MockGen. DO NOT EDIT.
// Source: engine.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	checkv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/flowcontrol/check/v1"
	iface "github.com/fluxninja/aperture/pkg/policies/flowcontrol/iface"
	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// GetConcurrencyLimiter mocks base method.
func (m *MockEngine) GetConcurrencyLimiter(limiterID iface.LimiterID) iface.ConcurrencyLimiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConcurrencyLimiter", limiterID)
	ret0, _ := ret[0].(iface.ConcurrencyLimiter)
	return ret0
}

// GetConcurrencyLimiter indicates an expected call of GetConcurrencyLimiter.
func (mr *MockEngineMockRecorder) GetConcurrencyLimiter(limiterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConcurrencyLimiter", reflect.TypeOf((*MockEngine)(nil).GetConcurrencyLimiter), limiterID)
}

// GetFluxMeter mocks base method.
func (m *MockEngine) GetFluxMeter(fluxMeterName string) iface.FluxMeter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFluxMeter", fluxMeterName)
	ret0, _ := ret[0].(iface.FluxMeter)
	return ret0
}

// GetFluxMeter indicates an expected call of GetFluxMeter.
func (mr *MockEngineMockRecorder) GetFluxMeter(fluxMeterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFluxMeter", reflect.TypeOf((*MockEngine)(nil).GetFluxMeter), fluxMeterName)
}

// GetRateLimiter mocks base method.
func (m *MockEngine) GetRateLimiter(limiterID iface.LimiterID) iface.RateLimiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRateLimiter", limiterID)
	ret0, _ := ret[0].(iface.RateLimiter)
	return ret0
}

// GetRateLimiter indicates an expected call of GetRateLimiter.
func (mr *MockEngineMockRecorder) GetRateLimiter(limiterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateLimiter", reflect.TypeOf((*MockEngine)(nil).GetRateLimiter), limiterID)
}

// ProcessRequest mocks base method.
func (m *MockEngine) ProcessRequest(ctx context.Context, controlPoint string, serviceIDs []string, labels map[string]string) *checkv1.CheckResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessRequest", ctx, controlPoint, serviceIDs, labels)
	ret0, _ := ret[0].(*checkv1.CheckResponse)
	return ret0
}

// ProcessRequest indicates an expected call of ProcessRequest.
func (mr *MockEngineMockRecorder) ProcessRequest(ctx, controlPoint, serviceIDs, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessRequest", reflect.TypeOf((*MockEngine)(nil).ProcessRequest), ctx, controlPoint, serviceIDs, labels)
}

// RegisterConcurrencyLimiter mocks base method.
func (m *MockEngine) RegisterConcurrencyLimiter(sa iface.ConcurrencyLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterConcurrencyLimiter", sa)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterConcurrencyLimiter indicates an expected call of RegisterConcurrencyLimiter.
func (mr *MockEngineMockRecorder) RegisterConcurrencyLimiter(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConcurrencyLimiter", reflect.TypeOf((*MockEngine)(nil).RegisterConcurrencyLimiter), sa)
}

// RegisterFluxMeter mocks base method.
func (m *MockEngine) RegisterFluxMeter(fm iface.FluxMeter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFluxMeter", fm)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterFluxMeter indicates an expected call of RegisterFluxMeter.
func (mr *MockEngineMockRecorder) RegisterFluxMeter(fm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFluxMeter", reflect.TypeOf((*MockEngine)(nil).RegisterFluxMeter), fm)
}

// RegisterRateLimiter mocks base method.
func (m *MockEngine) RegisterRateLimiter(l iface.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRateLimiter", l)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRateLimiter indicates an expected call of RegisterRateLimiter.
func (mr *MockEngineMockRecorder) RegisterRateLimiter(l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRateLimiter", reflect.TypeOf((*MockEngine)(nil).RegisterRateLimiter), l)
}

// UnregisterConcurrencyLimiter mocks base method.
func (m *MockEngine) UnregisterConcurrencyLimiter(sa iface.ConcurrencyLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterConcurrencyLimiter", sa)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterConcurrencyLimiter indicates an expected call of UnregisterConcurrencyLimiter.
func (mr *MockEngineMockRecorder) UnregisterConcurrencyLimiter(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterConcurrencyLimiter", reflect.TypeOf((*MockEngine)(nil).UnregisterConcurrencyLimiter), sa)
}

// UnregisterFluxMeter mocks base method.
func (m *MockEngine) UnregisterFluxMeter(fm iface.FluxMeter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterFluxMeter", fm)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterFluxMeter indicates an expected call of UnregisterFluxMeter.
func (mr *MockEngineMockRecorder) UnregisterFluxMeter(fm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterFluxMeter", reflect.TypeOf((*MockEngine)(nil).UnregisterFluxMeter), fm)
}

// UnregisterRateLimiter mocks base method.
func (m *MockEngine) UnregisterRateLimiter(l iface.RateLimiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterRateLimiter", l)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterRateLimiter indicates an expected call of UnregisterRateLimiter.
func (mr *MockEngineMockRecorder) UnregisterRateLimiter(l interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterRateLimiter", reflect.TypeOf((*MockEngine)(nil).UnregisterRateLimiter), l)
}
