// Code generated by MockGen. DO NOT EDIT.
// Source: config.go
//
// Generated by this command:
//
//	mockgen -source=config.go -package=kubeletconfig -destination=capability_checker_mock.go
//
// Package kubeletconfig is a generated GoMock package.
package kubeletconfig

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCapabilityChecker is a mock of CapabilityChecker interface.
type MockCapabilityChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCapabilityCheckerMockRecorder
}

// MockCapabilityCheckerMockRecorder is the mock recorder for MockCapabilityChecker.
type MockCapabilityCheckerMockRecorder struct {
	mock *MockCapabilityChecker
}

// NewMockCapabilityChecker creates a new mock instance.
func NewMockCapabilityChecker(ctrl *gomock.Controller) *MockCapabilityChecker {
	mock := &MockCapabilityChecker{ctrl: ctrl}
	mock.recorder = &MockCapabilityCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCapabilityChecker) EXPECT() *MockCapabilityCheckerMockRecorder {
	return m.recorder
}

// IsCapabilityEnabled mocks base method.
func (m *MockCapabilityChecker) IsCapabilityEnabled(capability string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCapabilityEnabled", capability)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCapabilityEnabled indicates an expected call of IsCapabilityEnabled.
func (mr *MockCapabilityCheckerMockRecorder) IsCapabilityEnabled(capability any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCapabilityEnabled", reflect.TypeOf((*MockCapabilityChecker)(nil).IsCapabilityEnabled), capability)
}
