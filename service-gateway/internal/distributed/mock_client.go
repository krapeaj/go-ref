// Code generated by MockGen. DO NOT EDIT.
// Source: internal/distributed/client.go

// Package distributed is a generated GoMock package.
package distributed

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConnector is a mock of Connector interface
type MockConnector struct {
	ctrl     *gomock.Controller
	recorder *MockConnectorMockRecorder
}

// MockConnectorMockRecorder is the mock recorder for MockConnector
type MockConnectorMockRecorder struct {
	mock *MockConnector
}

// NewMockConnector creates a new mock instance
func NewMockConnector(ctrl *gomock.Controller) *MockConnector {
	mock := &MockConnector{ctrl: ctrl}
	mock.recorder = &MockConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnector) EXPECT() *MockConnectorMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockConnector) Publish(message *Message) {
	m.ctrl.Call(m, "Publish", message)
}

// Publish indicates an expected call of Publish
func (mr *MockConnectorMockRecorder) Publish(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockConnector)(nil).Publish), message)
}

// Stop mocks base method
func (m *MockConnector) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockConnectorMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockConnector)(nil).Stop))
}