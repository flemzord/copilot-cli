// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/ssm/ssm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

// Mockapi is a mock of api interface.
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi.
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance.
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// AddTagsToResource mocks base method.
func (m *Mockapi) AddTagsToResource(input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagsToResource", input)
	ret0, _ := ret[0].(*ssm.AddTagsToResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTagsToResource indicates an expected call of AddTagsToResource.
func (mr *MockapiMockRecorder) AddTagsToResource(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagsToResource", reflect.TypeOf((*Mockapi)(nil).AddTagsToResource), input)
}

// PutParameter mocks base method.
func (m *Mockapi) PutParameter(input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutParameter", input)
	ret0, _ := ret[0].(*ssm.PutParameterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutParameter indicates an expected call of PutParameter.
func (mr *MockapiMockRecorder) PutParameter(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutParameter", reflect.TypeOf((*Mockapi)(nil).PutParameter), input)
}