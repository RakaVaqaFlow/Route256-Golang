// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	repository "homework/internal/pkg/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsersRepo is a mock of UsersRepo interface.
type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoMockRecorder
}

// MockUsersRepoMockRecorder is the mock recorder for MockUsersRepo.
type MockUsersRepoMockRecorder struct {
	mock *MockUsersRepo
}

// NewMockUsersRepo creates a new mock instance.
func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepo) EXPECT() *MockUsersRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockUsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockUsersRepoMockRecorder) Add(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsersRepo)(nil).Add), ctx, user)
}

// Delete mocks base method.
func (m *MockUsersRepo) Delete(ctx context.Context, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsersRepo)(nil).Delete), ctx, id)
}

// GetById mocks base method.
func (m *MockUsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUsersRepoMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUsersRepo)(nil).GetById), ctx, id)
}

// List mocks base method.
func (m *MockUsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersRepoMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsersRepo)(nil).List), ctx)
}

// Update mocks base method.
func (m *MockUsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersRepoMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersRepo)(nil).Update), ctx, user)
}

// MockTasksRepo is a mock of TasksRepo interface.
type MockTasksRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTasksRepoMockRecorder
}

// MockTasksRepoMockRecorder is the mock recorder for MockTasksRepo.
type MockTasksRepoMockRecorder struct {
	mock *MockTasksRepo
}

// NewMockTasksRepo creates a new mock instance.
func NewMockTasksRepo(ctrl *gomock.Controller) *MockTasksRepo {
	mock := &MockTasksRepo{ctrl: ctrl}
	mock.recorder = &MockTasksRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTasksRepo) EXPECT() *MockTasksRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTasksRepo) Add(ctx context.Context, task *repository.Task) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, task)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTasksRepoMockRecorder) Add(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTasksRepo)(nil).Add), ctx, task)
}

// Delete mocks base method.
func (m *MockTasksRepo) Delete(ctx context.Context, id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTasksRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTasksRepo)(nil).Delete), ctx, id)
}

// GetById mocks base method.
func (m *MockTasksRepo) GetById(ctx context.Context, id int64) (*repository.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*repository.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTasksRepoMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTasksRepo)(nil).GetById), ctx, id)
}

// List mocks base method.
func (m *MockTasksRepo) List(ctx context.Context) ([]*repository.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*repository.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTasksRepoMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTasksRepo)(nil).List), ctx)
}

// Update mocks base method.
func (m *MockTasksRepo) Update(ctx context.Context, task *repository.Task) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, task)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTasksRepoMockRecorder) Update(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTasksRepo)(nil).Update), ctx, task)
}