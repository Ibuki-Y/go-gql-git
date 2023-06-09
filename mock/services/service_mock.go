// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	model "github.com/Ibuki-Y/go-gql-git/graph/model"
	gomock "github.com/golang/mock/gomock"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// GetIssueByID mocks base method.
func (m *MockServices) GetIssueByID(ctx context.Context, id string) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueByID", ctx, id)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueByID indicates an expected call of GetIssueByID.
func (mr *MockServicesMockRecorder) GetIssueByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueByID", reflect.TypeOf((*MockServices)(nil).GetIssueByID), ctx, id)
}

// GetIssueByRepoAndNumber mocks base method.
func (m *MockServices) GetIssueByRepoAndNumber(ctx context.Context, repID string, number int) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueByRepoAndNumber", ctx, repID, number)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueByRepoAndNumber indicates an expected call of GetIssueByRepoAndNumber.
func (mr *MockServicesMockRecorder) GetIssueByRepoAndNumber(ctx, repID, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueByRepoAndNumber", reflect.TypeOf((*MockServices)(nil).GetIssueByRepoAndNumber), ctx, repID, number)
}

// GetRepositoryByFullName mocks base method.
func (m *MockServices) GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByFullName", ctx, owner, name)
	ret0, _ := ret[0].(*model.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByFullName indicates an expected call of GetRepositoryByFullName.
func (mr *MockServicesMockRecorder) GetRepositoryByFullName(ctx, owner, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByFullName", reflect.TypeOf((*MockServices)(nil).GetRepositoryByFullName), ctx, owner, name)
}

// GetRepositoryByID mocks base method.
func (m *MockServices) GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByID", ctx, id)
	ret0, _ := ret[0].(*model.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByID indicates an expected call of GetRepositoryByID.
func (mr *MockServicesMockRecorder) GetRepositoryByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByID", reflect.TypeOf((*MockServices)(nil).GetRepositoryByID), ctx, id)
}

// GetUserByID mocks base method.
func (m *MockServices) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockServicesMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockServices)(nil).GetUserByID), ctx, id)
}

// GetUserByName mocks base method.
func (m *MockServices) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", ctx, name)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockServicesMockRecorder) GetUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockServices)(nil).GetUserByName), ctx, name)
}

// ListUsersByID mocks base method.
func (m *MockServices) ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersByID", ctx, IDs)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersByID indicates an expected call of ListUsersByID.
func (mr *MockServicesMockRecorder) ListUsersByID(ctx, IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersByID", reflect.TypeOf((*MockServices)(nil).ListUsersByID), ctx, IDs)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), ctx, id)
}

// GetUserByName mocks base method.
func (m *MockUserService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", ctx, name)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserServiceMockRecorder) GetUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserService)(nil).GetUserByName), ctx, name)
}

// ListUsersByID mocks base method.
func (m *MockUserService) ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersByID", ctx, IDs)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsersByID indicates an expected call of ListUsersByID.
func (mr *MockUserServiceMockRecorder) ListUsersByID(ctx, IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersByID", reflect.TypeOf((*MockUserService)(nil).ListUsersByID), ctx, IDs)
}

// MockRepositoryService is a mock of RepositoryService interface.
type MockRepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryServiceMockRecorder
}

// MockRepositoryServiceMockRecorder is the mock recorder for MockRepositoryService.
type MockRepositoryServiceMockRecorder struct {
	mock *MockRepositoryService
}

// NewMockRepositoryService creates a new mock instance.
func NewMockRepositoryService(ctrl *gomock.Controller) *MockRepositoryService {
	mock := &MockRepositoryService{ctrl: ctrl}
	mock.recorder = &MockRepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryService) EXPECT() *MockRepositoryServiceMockRecorder {
	return m.recorder
}

// GetRepositoryByFullName mocks base method.
func (m *MockRepositoryService) GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByFullName", ctx, owner, name)
	ret0, _ := ret[0].(*model.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByFullName indicates an expected call of GetRepositoryByFullName.
func (mr *MockRepositoryServiceMockRecorder) GetRepositoryByFullName(ctx, owner, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByFullName", reflect.TypeOf((*MockRepositoryService)(nil).GetRepositoryByFullName), ctx, owner, name)
}

// GetRepositoryByID mocks base method.
func (m *MockRepositoryService) GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryByID", ctx, id)
	ret0, _ := ret[0].(*model.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryByID indicates an expected call of GetRepositoryByID.
func (mr *MockRepositoryServiceMockRecorder) GetRepositoryByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryByID", reflect.TypeOf((*MockRepositoryService)(nil).GetRepositoryByID), ctx, id)
}

// MockIssueService is a mock of IssueService interface.
type MockIssueService struct {
	ctrl     *gomock.Controller
	recorder *MockIssueServiceMockRecorder
}

// MockIssueServiceMockRecorder is the mock recorder for MockIssueService.
type MockIssueServiceMockRecorder struct {
	mock *MockIssueService
}

// NewMockIssueService creates a new mock instance.
func NewMockIssueService(ctrl *gomock.Controller) *MockIssueService {
	mock := &MockIssueService{ctrl: ctrl}
	mock.recorder = &MockIssueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueService) EXPECT() *MockIssueServiceMockRecorder {
	return m.recorder
}

// GetIssueByID mocks base method.
func (m *MockIssueService) GetIssueByID(ctx context.Context, id string) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueByID", ctx, id)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueByID indicates an expected call of GetIssueByID.
func (mr *MockIssueServiceMockRecorder) GetIssueByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueByID", reflect.TypeOf((*MockIssueService)(nil).GetIssueByID), ctx, id)
}

// GetIssueByRepoAndNumber mocks base method.
func (m *MockIssueService) GetIssueByRepoAndNumber(ctx context.Context, repID string, number int) (*model.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueByRepoAndNumber", ctx, repID, number)
	ret0, _ := ret[0].(*model.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueByRepoAndNumber indicates an expected call of GetIssueByRepoAndNumber.
func (mr *MockIssueServiceMockRecorder) GetIssueByRepoAndNumber(ctx, repID, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueByRepoAndNumber", reflect.TypeOf((*MockIssueService)(nil).GetIssueByRepoAndNumber), ctx, repID, number)
}
