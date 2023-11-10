// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/aws/services (interfaces: Tagging)

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	resourcegroupstaggingapi "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	gomock "github.com/golang/mock/gomock"
)

// MockTagging is a mock of Tagging interface.
type MockTagging struct {
	ctrl     *gomock.Controller
	recorder *MockTaggingMockRecorder
}

// MockTaggingMockRecorder is the mock recorder for MockTagging.
type MockTaggingMockRecorder struct {
	mock *MockTagging
}

// NewMockTagging creates a new mock instance.
func NewMockTagging(ctrl *gomock.Controller) *MockTagging {
	mock := &MockTagging{ctrl: ctrl}
	mock.recorder = &MockTaggingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagging) EXPECT() *MockTaggingMockRecorder {
	return m.recorder
}

// DescribeReportCreation mocks base method.
func (m *MockTagging) DescribeReportCreation(arg0 *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeReportCreation", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportCreation indicates an expected call of DescribeReportCreation.
func (mr *MockTaggingMockRecorder) DescribeReportCreation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportCreation", reflect.TypeOf((*MockTagging)(nil).DescribeReportCreation), arg0)
}

// DescribeReportCreationRequest mocks base method.
func (m *MockTagging) DescribeReportCreationRequest(arg0 *resourcegroupstaggingapi.DescribeReportCreationInput) (*request.Request, *resourcegroupstaggingapi.DescribeReportCreationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeReportCreationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	return ret0, ret1
}

// DescribeReportCreationRequest indicates an expected call of DescribeReportCreationRequest.
func (mr *MockTaggingMockRecorder) DescribeReportCreationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportCreationRequest", reflect.TypeOf((*MockTagging)(nil).DescribeReportCreationRequest), arg0)
}

// DescribeReportCreationWithContext mocks base method.
func (m *MockTagging) DescribeReportCreationWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.DescribeReportCreationInput, arg2 ...request.Option) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReportCreationWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.DescribeReportCreationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportCreationWithContext indicates an expected call of DescribeReportCreationWithContext.
func (mr *MockTaggingMockRecorder) DescribeReportCreationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportCreationWithContext", reflect.TypeOf((*MockTagging)(nil).DescribeReportCreationWithContext), varargs...)
}

// FindResourceWithTags mocks base method.
func (m *MockTagging) FindResourceWithTags(arg0 context.Context, arg1 string, arg2 map[string]*string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindResourceWithTags", arg0, arg1, arg2)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResourceWithTags indicates an expected call of FindResourceWithTags.
func (mr *MockTaggingMockRecorder) FindResourceWithTags(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResourceWithTags", reflect.TypeOf((*MockTagging)(nil).FindResourceWithTags), arg0, arg1, arg2)
}

// GetComplianceSummary mocks base method.
func (m *MockTagging) GetComplianceSummary(arg0 *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceSummary", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplianceSummary indicates an expected call of GetComplianceSummary.
func (mr *MockTaggingMockRecorder) GetComplianceSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceSummary", reflect.TypeOf((*MockTagging)(nil).GetComplianceSummary), arg0)
}

// GetComplianceSummaryPages mocks base method.
func (m *MockTagging) GetComplianceSummaryPages(arg0 *resourcegroupstaggingapi.GetComplianceSummaryInput, arg1 func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceSummaryPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetComplianceSummaryPages indicates an expected call of GetComplianceSummaryPages.
func (mr *MockTaggingMockRecorder) GetComplianceSummaryPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceSummaryPages", reflect.TypeOf((*MockTagging)(nil).GetComplianceSummaryPages), arg0, arg1)
}

// GetComplianceSummaryPagesWithContext mocks base method.
func (m *MockTagging) GetComplianceSummaryPagesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetComplianceSummaryInput, arg2 func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetComplianceSummaryPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetComplianceSummaryPagesWithContext indicates an expected call of GetComplianceSummaryPagesWithContext.
func (mr *MockTaggingMockRecorder) GetComplianceSummaryPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceSummaryPagesWithContext", reflect.TypeOf((*MockTagging)(nil).GetComplianceSummaryPagesWithContext), varargs...)
}

// GetComplianceSummaryRequest mocks base method.
func (m *MockTagging) GetComplianceSummaryRequest(arg0 *resourcegroupstaggingapi.GetComplianceSummaryInput) (*request.Request, *resourcegroupstaggingapi.GetComplianceSummaryOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplianceSummaryRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	return ret0, ret1
}

// GetComplianceSummaryRequest indicates an expected call of GetComplianceSummaryRequest.
func (mr *MockTaggingMockRecorder) GetComplianceSummaryRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceSummaryRequest", reflect.TypeOf((*MockTagging)(nil).GetComplianceSummaryRequest), arg0)
}

// GetComplianceSummaryWithContext mocks base method.
func (m *MockTagging) GetComplianceSummaryWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetComplianceSummaryInput, arg2 ...request.Option) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetComplianceSummaryWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetComplianceSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplianceSummaryWithContext indicates an expected call of GetComplianceSummaryWithContext.
func (mr *MockTaggingMockRecorder) GetComplianceSummaryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplianceSummaryWithContext", reflect.TypeOf((*MockTagging)(nil).GetComplianceSummaryWithContext), varargs...)
}

// GetResources mocks base method.
func (m *MockTagging) GetResources(arg0 *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResources", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResources indicates an expected call of GetResources.
func (mr *MockTaggingMockRecorder) GetResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResources", reflect.TypeOf((*MockTagging)(nil).GetResources), arg0)
}

// GetResourcesPages mocks base method.
func (m *MockTagging) GetResourcesPages(arg0 *resourcegroupstaggingapi.GetResourcesInput, arg1 func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourcesPages indicates an expected call of GetResourcesPages.
func (mr *MockTaggingMockRecorder) GetResourcesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesPages", reflect.TypeOf((*MockTagging)(nil).GetResourcesPages), arg0, arg1)
}

// GetResourcesPagesWithContext mocks base method.
func (m *MockTagging) GetResourcesPagesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetResourcesInput, arg2 func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourcesPagesWithContext indicates an expected call of GetResourcesPagesWithContext.
func (mr *MockTaggingMockRecorder) GetResourcesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesPagesWithContext", reflect.TypeOf((*MockTagging)(nil).GetResourcesPagesWithContext), varargs...)
}

// GetResourcesRequest mocks base method.
func (m *MockTagging) GetResourcesRequest(arg0 *resourcegroupstaggingapi.GetResourcesInput) (*request.Request, *resourcegroupstaggingapi.GetResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.GetResourcesOutput)
	return ret0, ret1
}

// GetResourcesRequest indicates an expected call of GetResourcesRequest.
func (mr *MockTaggingMockRecorder) GetResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesRequest", reflect.TypeOf((*MockTagging)(nil).GetResourcesRequest), arg0)
}

// GetResourcesWithContext mocks base method.
func (m *MockTagging) GetResourcesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetResourcesInput, arg2 ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesWithContext indicates an expected call of GetResourcesWithContext.
func (mr *MockTaggingMockRecorder) GetResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesWithContext", reflect.TypeOf((*MockTagging)(nil).GetResourcesWithContext), varargs...)
}

// GetTagKeys mocks base method.
func (m *MockTagging) GetTagKeys(arg0 *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagKeys", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetTagKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagKeys indicates an expected call of GetTagKeys.
func (mr *MockTaggingMockRecorder) GetTagKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagKeys", reflect.TypeOf((*MockTagging)(nil).GetTagKeys), arg0)
}

// GetTagKeysPages mocks base method.
func (m *MockTagging) GetTagKeysPages(arg0 *resourcegroupstaggingapi.GetTagKeysInput, arg1 func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagKeysPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTagKeysPages indicates an expected call of GetTagKeysPages.
func (mr *MockTaggingMockRecorder) GetTagKeysPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagKeysPages", reflect.TypeOf((*MockTagging)(nil).GetTagKeysPages), arg0, arg1)
}

// GetTagKeysPagesWithContext mocks base method.
func (m *MockTagging) GetTagKeysPagesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetTagKeysInput, arg2 func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagKeysPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTagKeysPagesWithContext indicates an expected call of GetTagKeysPagesWithContext.
func (mr *MockTaggingMockRecorder) GetTagKeysPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagKeysPagesWithContext", reflect.TypeOf((*MockTagging)(nil).GetTagKeysPagesWithContext), varargs...)
}

// GetTagKeysRequest mocks base method.
func (m *MockTagging) GetTagKeysRequest(arg0 *resourcegroupstaggingapi.GetTagKeysInput) (*request.Request, *resourcegroupstaggingapi.GetTagKeysOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagKeysRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.GetTagKeysOutput)
	return ret0, ret1
}

// GetTagKeysRequest indicates an expected call of GetTagKeysRequest.
func (mr *MockTaggingMockRecorder) GetTagKeysRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagKeysRequest", reflect.TypeOf((*MockTagging)(nil).GetTagKeysRequest), arg0)
}

// GetTagKeysWithContext mocks base method.
func (m *MockTagging) GetTagKeysWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetTagKeysInput, arg2 ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagKeysWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetTagKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagKeysWithContext indicates an expected call of GetTagKeysWithContext.
func (mr *MockTaggingMockRecorder) GetTagKeysWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagKeysWithContext", reflect.TypeOf((*MockTagging)(nil).GetTagKeysWithContext), varargs...)
}

// GetTagValues mocks base method.
func (m *MockTagging) GetTagValues(arg0 *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagValues", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetTagValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagValues indicates an expected call of GetTagValues.
func (mr *MockTaggingMockRecorder) GetTagValues(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagValues", reflect.TypeOf((*MockTagging)(nil).GetTagValues), arg0)
}

// GetTagValuesPages mocks base method.
func (m *MockTagging) GetTagValuesPages(arg0 *resourcegroupstaggingapi.GetTagValuesInput, arg1 func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagValuesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTagValuesPages indicates an expected call of GetTagValuesPages.
func (mr *MockTaggingMockRecorder) GetTagValuesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagValuesPages", reflect.TypeOf((*MockTagging)(nil).GetTagValuesPages), arg0, arg1)
}

// GetTagValuesPagesWithContext mocks base method.
func (m *MockTagging) GetTagValuesPagesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetTagValuesInput, arg2 func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagValuesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTagValuesPagesWithContext indicates an expected call of GetTagValuesPagesWithContext.
func (mr *MockTaggingMockRecorder) GetTagValuesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagValuesPagesWithContext", reflect.TypeOf((*MockTagging)(nil).GetTagValuesPagesWithContext), varargs...)
}

// GetTagValuesRequest mocks base method.
func (m *MockTagging) GetTagValuesRequest(arg0 *resourcegroupstaggingapi.GetTagValuesInput) (*request.Request, *resourcegroupstaggingapi.GetTagValuesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagValuesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.GetTagValuesOutput)
	return ret0, ret1
}

// GetTagValuesRequest indicates an expected call of GetTagValuesRequest.
func (mr *MockTaggingMockRecorder) GetTagValuesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagValuesRequest", reflect.TypeOf((*MockTagging)(nil).GetTagValuesRequest), arg0)
}

// GetTagValuesWithContext mocks base method.
func (m *MockTagging) GetTagValuesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.GetTagValuesInput, arg2 ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagValuesWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.GetTagValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagValuesWithContext indicates an expected call of GetTagValuesWithContext.
func (mr *MockTaggingMockRecorder) GetTagValuesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagValuesWithContext", reflect.TypeOf((*MockTagging)(nil).GetTagValuesWithContext), varargs...)
}

// GetTagsFromArns mocks base method.
func (m *MockTagging) GetTagsFromArns(arg0 context.Context, arg1 []*string) (map[string]map[string]*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsFromArns", arg0, arg1)
	ret0, _ := ret[0].(map[string]map[string]*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsFromArns indicates an expected call of GetTagsFromArns.
func (mr *MockTaggingMockRecorder) GetTagsFromArns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsFromArns", reflect.TypeOf((*MockTagging)(nil).GetTagsFromArns), arg0, arg1)
}

// StartReportCreation mocks base method.
func (m *MockTagging) StartReportCreation(arg0 *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartReportCreation", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.StartReportCreationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartReportCreation indicates an expected call of StartReportCreation.
func (mr *MockTaggingMockRecorder) StartReportCreation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartReportCreation", reflect.TypeOf((*MockTagging)(nil).StartReportCreation), arg0)
}

// StartReportCreationRequest mocks base method.
func (m *MockTagging) StartReportCreationRequest(arg0 *resourcegroupstaggingapi.StartReportCreationInput) (*request.Request, *resourcegroupstaggingapi.StartReportCreationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartReportCreationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.StartReportCreationOutput)
	return ret0, ret1
}

// StartReportCreationRequest indicates an expected call of StartReportCreationRequest.
func (mr *MockTaggingMockRecorder) StartReportCreationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartReportCreationRequest", reflect.TypeOf((*MockTagging)(nil).StartReportCreationRequest), arg0)
}

// StartReportCreationWithContext mocks base method.
func (m *MockTagging) StartReportCreationWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.StartReportCreationInput, arg2 ...request.Option) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartReportCreationWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.StartReportCreationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartReportCreationWithContext indicates an expected call of StartReportCreationWithContext.
func (mr *MockTaggingMockRecorder) StartReportCreationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartReportCreationWithContext", reflect.TypeOf((*MockTagging)(nil).StartReportCreationWithContext), varargs...)
}

// TagResources mocks base method.
func (m *MockTagging) TagResources(arg0 *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResources", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.TagResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResources indicates an expected call of TagResources.
func (mr *MockTaggingMockRecorder) TagResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResources", reflect.TypeOf((*MockTagging)(nil).TagResources), arg0)
}

// TagResourcesRequest mocks base method.
func (m *MockTagging) TagResourcesRequest(arg0 *resourcegroupstaggingapi.TagResourcesInput) (*request.Request, *resourcegroupstaggingapi.TagResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.TagResourcesOutput)
	return ret0, ret1
}

// TagResourcesRequest indicates an expected call of TagResourcesRequest.
func (mr *MockTaggingMockRecorder) TagResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourcesRequest", reflect.TypeOf((*MockTagging)(nil).TagResourcesRequest), arg0)
}

// TagResourcesWithContext mocks base method.
func (m *MockTagging) TagResourcesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.TagResourcesInput, arg2 ...request.Option) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.TagResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagResourcesWithContext indicates an expected call of TagResourcesWithContext.
func (mr *MockTaggingMockRecorder) TagResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagResourcesWithContext", reflect.TypeOf((*MockTagging)(nil).TagResourcesWithContext), varargs...)
}

// UntagResources mocks base method.
func (m *MockTagging) UntagResources(arg0 *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResources", arg0)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.UntagResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResources indicates an expected call of UntagResources.
func (mr *MockTaggingMockRecorder) UntagResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResources", reflect.TypeOf((*MockTagging)(nil).UntagResources), arg0)
}

// UntagResourcesRequest mocks base method.
func (m *MockTagging) UntagResourcesRequest(arg0 *resourcegroupstaggingapi.UntagResourcesInput) (*request.Request, *resourcegroupstaggingapi.UntagResourcesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagResourcesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*resourcegroupstaggingapi.UntagResourcesOutput)
	return ret0, ret1
}

// UntagResourcesRequest indicates an expected call of UntagResourcesRequest.
func (mr *MockTaggingMockRecorder) UntagResourcesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourcesRequest", reflect.TypeOf((*MockTagging)(nil).UntagResourcesRequest), arg0)
}

// UntagResourcesWithContext mocks base method.
func (m *MockTagging) UntagResourcesWithContext(arg0 context.Context, arg1 *resourcegroupstaggingapi.UntagResourcesInput, arg2 ...request.Option) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagResourcesWithContext", varargs...)
	ret0, _ := ret[0].(*resourcegroupstaggingapi.UntagResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagResourcesWithContext indicates an expected call of UntagResourcesWithContext.
func (mr *MockTaggingMockRecorder) UntagResourcesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagResourcesWithContext", reflect.TypeOf((*MockTagging)(nil).UntagResourcesWithContext), varargs...)
}
