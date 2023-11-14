// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/deploy/lattice (interfaces: ServiceNetworkManager)

// Package lattice is a generated GoMock package.
package lattice

import (
	context "context"
	reflect "reflect"

	lattice "github.com/aws/aws-application-networking-k8s/pkg/model/lattice"
	gomock "github.com/golang/mock/gomock"
)

// MockServiceNetworkManager is a mock of ServiceNetworkManager interface.
type MockServiceNetworkManager struct {
	ctrl     *gomock.Controller
	recorder *MockServiceNetworkManagerMockRecorder
}

// MockServiceNetworkManagerMockRecorder is the mock recorder for MockServiceNetworkManager.
type MockServiceNetworkManagerMockRecorder struct {
	mock *MockServiceNetworkManager
}

// NewMockServiceNetworkManager creates a new mock instance.
func NewMockServiceNetworkManager(ctrl *gomock.Controller) *MockServiceNetworkManager {
	mock := &MockServiceNetworkManager{ctrl: ctrl}
	mock.recorder = &MockServiceNetworkManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceNetworkManager) EXPECT() *MockServiceNetworkManagerMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockServiceNetworkManager) CreateOrUpdate(arg0 context.Context, arg1 *lattice.ServiceNetwork) (lattice.ServiceNetworkStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1)
	ret0, _ := ret[0].(lattice.ServiceNetworkStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockServiceNetworkManagerMockRecorder) CreateOrUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockServiceNetworkManager)(nil).CreateOrUpdate), arg0, arg1)
}

// DeleteVpcAssociation mocks base method.
func (m *MockServiceNetworkManager) DeleteVpcAssociation(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVpcAssociation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVpcAssociation indicates an expected call of DeleteVpcAssociation.
func (mr *MockServiceNetworkManagerMockRecorder) DeleteVpcAssociation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVpcAssociation", reflect.TypeOf((*MockServiceNetworkManager)(nil).DeleteVpcAssociation), arg0, arg1)
}

// UpsertVpcAssociation mocks base method.
func (m *MockServiceNetworkManager) UpsertVpcAssociation(arg0 context.Context, arg1 string, arg2 []*string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertVpcAssociation", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertVpcAssociation indicates an expected call of UpsertVpcAssociation.
func (mr *MockServiceNetworkManagerMockRecorder) UpsertVpcAssociation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertVpcAssociation", reflect.TypeOf((*MockServiceNetworkManager)(nil).UpsertVpcAssociation), arg0, arg1, arg2)
}
