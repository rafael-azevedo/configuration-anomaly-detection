// Code generated by MockGen. DO NOT EDIT.
// Source: chgm.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	cloudtrail "github.com/aws/aws-sdk-go/service/cloudtrail"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	v10 "github.com/openshift/hive/apis/hive/v1"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetClusterDeployment mocks base method.
func (m *MockService) GetClusterDeployment(clusterID string) (*v10.ClusterDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterDeployment", clusterID)
	ret0, _ := ret[0].(*v10.ClusterDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterDeployment indicates an expected call of GetClusterDeployment.
func (mr *MockServiceMockRecorder) GetClusterDeployment(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterDeployment", reflect.TypeOf((*MockService)(nil).GetClusterDeployment), clusterID)
}

// GetClusterInfo mocks base method.
func (m *MockService) GetClusterInfo(identifier string) (*v1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", identifier)
	ret0, _ := ret[0].(*v1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockServiceMockRecorder) GetClusterInfo(identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockService)(nil).GetClusterInfo), identifier)
}

// ListRunningInstances mocks base method.
func (m *MockService) ListRunningInstances(infraID string) ([]*ec2.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunningInstances", infraID)
	ret0, _ := ret[0].([]*ec2.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunningInstances indicates an expected call of ListRunningInstances.
func (mr *MockServiceMockRecorder) ListRunningInstances(infraID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunningInstances", reflect.TypeOf((*MockService)(nil).ListRunningInstances), infraID)
}

// ListStoppedInstances mocks base method.
func (m *MockService) ListStoppedInstances(infraID string) ([]*ec2.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoppedInstances", infraID)
	ret0, _ := ret[0].([]*ec2.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoppedInstances indicates an expected call of ListStoppedInstances.
func (mr *MockServiceMockRecorder) ListStoppedInstances(infraID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoppedInstances", reflect.TypeOf((*MockService)(nil).ListStoppedInstances), infraID)
}

// PollInstanceStopEventsFor mocks base method.
func (m *MockService) PollInstanceStopEventsFor(instances []*ec2.Instance, retryTimes int) ([]*cloudtrail.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollInstanceStopEventsFor", instances, retryTimes)
	ret0, _ := ret[0].([]*cloudtrail.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollInstanceStopEventsFor indicates an expected call of PollInstanceStopEventsFor.
func (mr *MockServiceMockRecorder) PollInstanceStopEventsFor(instances, retryTimes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollInstanceStopEventsFor", reflect.TypeOf((*MockService)(nil).PollInstanceStopEventsFor), instances, retryTimes)
}