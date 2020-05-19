// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/describe/lb_web_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	ecs "github.com/aws/amazon-ecs-cli-v2/internal/pkg/aws/ecs"
	config "github.com/aws/amazon-ecs-cli-v2/internal/pkg/config"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockstoreSvc is a mock of storeSvc interface
type MockstoreSvc struct {
	ctrl     *gomock.Controller
	recorder *MockstoreSvcMockRecorder
}

// MockstoreSvcMockRecorder is the mock recorder for MockstoreSvc
type MockstoreSvcMockRecorder struct {
	mock *MockstoreSvc
}

// NewMockstoreSvc creates a new mock instance
func NewMockstoreSvc(ctrl *gomock.Controller) *MockstoreSvc {
	mock := &MockstoreSvc{ctrl: ctrl}
	mock.recorder = &MockstoreSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstoreSvc) EXPECT() *MockstoreSvcMockRecorder {
	return m.recorder
}

// GetEnvironment mocks base method
func (m *MockstoreSvc) GetEnvironment(appName, environmentName string) (*config.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", appName, environmentName)
	ret0, _ := ret[0].(*config.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockstoreSvcMockRecorder) GetEnvironment(appName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockstoreSvc)(nil).GetEnvironment), appName, environmentName)
}

// ListEnvironments mocks base method
func (m *MockstoreSvc) ListEnvironments(appName string) ([]*config.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnvironments", appName)
	ret0, _ := ret[0].([]*config.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments
func (mr *MockstoreSvcMockRecorder) ListEnvironments(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockstoreSvc)(nil).ListEnvironments), appName)
}

// MocksvcDescriber is a mock of svcDescriber interface
type MocksvcDescriber struct {
	ctrl     *gomock.Controller
	recorder *MocksvcDescriberMockRecorder
}

// MocksvcDescriberMockRecorder is the mock recorder for MocksvcDescriber
type MocksvcDescriberMockRecorder struct {
	mock *MocksvcDescriber
}

// NewMocksvcDescriber creates a new mock instance
func NewMocksvcDescriber(ctrl *gomock.Controller) *MocksvcDescriber {
	mock := &MocksvcDescriber{ctrl: ctrl}
	mock.recorder = &MocksvcDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocksvcDescriber) EXPECT() *MocksvcDescriberMockRecorder {
	return m.recorder
}

// Params mocks base method
func (m *MocksvcDescriber) Params() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Params indicates an expected call of Params
func (mr *MocksvcDescriberMockRecorder) Params() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MocksvcDescriber)(nil).Params))
}

// EnvOutputs mocks base method
func (m *MocksvcDescriber) EnvOutputs() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvOutputs")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvOutputs indicates an expected call of EnvOutputs
func (mr *MocksvcDescriberMockRecorder) EnvOutputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvOutputs", reflect.TypeOf((*MocksvcDescriber)(nil).EnvOutputs))
}

// EnvVars mocks base method
func (m *MocksvcDescriber) EnvVars() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvVars")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvVars indicates an expected call of EnvVars
func (mr *MocksvcDescriberMockRecorder) EnvVars() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvVars", reflect.TypeOf((*MocksvcDescriber)(nil).EnvVars))
}

// GetServiceArn mocks base method
func (m *MocksvcDescriber) GetServiceArn() (*ecs.ServiceArn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceArn")
	ret0, _ := ret[0].(*ecs.ServiceArn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceArn indicates an expected call of GetServiceArn
func (mr *MocksvcDescriberMockRecorder) GetServiceArn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceArn", reflect.TypeOf((*MocksvcDescriber)(nil).GetServiceArn))
}

// ServiceStackResources mocks base method
func (m *MocksvcDescriber) ServiceStackResources() ([]*cloudformation.StackResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceStackResources")
	ret0, _ := ret[0].([]*cloudformation.StackResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceStackResources indicates an expected call of ServiceStackResources
func (mr *MocksvcDescriberMockRecorder) ServiceStackResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceStackResources", reflect.TypeOf((*MocksvcDescriber)(nil).ServiceStackResources))
}

// MockHumanJSONStringer is a mock of HumanJSONStringer interface
type MockHumanJSONStringer struct {
	ctrl     *gomock.Controller
	recorder *MockHumanJSONStringerMockRecorder
}

// MockHumanJSONStringerMockRecorder is the mock recorder for MockHumanJSONStringer
type MockHumanJSONStringerMockRecorder struct {
	mock *MockHumanJSONStringer
}

// NewMockHumanJSONStringer creates a new mock instance
func NewMockHumanJSONStringer(ctrl *gomock.Controller) *MockHumanJSONStringer {
	mock := &MockHumanJSONStringer{ctrl: ctrl}
	mock.recorder = &MockHumanJSONStringerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHumanJSONStringer) EXPECT() *MockHumanJSONStringerMockRecorder {
	return m.recorder
}

// HumanString mocks base method
func (m *MockHumanJSONStringer) HumanString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HumanString")
	ret0, _ := ret[0].(string)
	return ret0
}

// HumanString indicates an expected call of HumanString
func (mr *MockHumanJSONStringerMockRecorder) HumanString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HumanString", reflect.TypeOf((*MockHumanJSONStringer)(nil).HumanString))
}

// JSONString mocks base method
func (m *MockHumanJSONStringer) JSONString() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSONString indicates an expected call of JSONString
func (mr *MockHumanJSONStringerMockRecorder) JSONString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONString", reflect.TypeOf((*MockHumanJSONStringer)(nil).JSONString))
}