// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/domains/user/repositories/user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/martinsd3v/planets/core/domains/user/entities"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockIUserRepository) All() (entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockIUserRepositoryMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIUserRepository)(nil).All))
}

// Create mocks base method.
func (m *MockIUserRepository) Create(arg0 entities.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), arg0)
}

// Destroy mocks base method.
func (m *MockIUserRepository) Destroy(uuid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockIUserRepositoryMockRecorder) Destroy(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockIUserRepository)(nil).Destroy), uuid)
}

// FindByEmail mocks base method.
func (m *MockIUserRepository) FindByEmail(email string) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockIUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockIUserRepository)(nil).FindByEmail), email)
}

// FindByUUID mocks base method.
func (m *MockIUserRepository) FindByUUID(arg0 string) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUUID", arg0)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUUID indicates an expected call of FindByUUID.
func (mr *MockIUserRepositoryMockRecorder) FindByUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUUID", reflect.TypeOf((*MockIUserRepository)(nil).FindByUUID), arg0)
}

// Save mocks base method.
func (m *MockIUserRepository) Save(arg0 entities.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIUserRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUserRepository)(nil).Save), arg0)
}
