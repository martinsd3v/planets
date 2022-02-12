// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/tools/providers/cache/cache_provider.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	cache "github.com/martinsd3v/planets/core/tools/providers/cache"
)

// MockICacheProvider is a mock of ICacheProvider interface.
type MockICacheProvider struct {
	ctrl     *gomock.Controller
	recorder *MockICacheProviderMockRecorder
}

// MockICacheProviderMockRecorder is the mock recorder for MockICacheProvider.
type MockICacheProviderMockRecorder struct {
	mock *MockICacheProvider
}

// NewMockICacheProvider creates a new mock instance.
func NewMockICacheProvider(ctrl *gomock.Controller) *MockICacheProvider {
	mock := &MockICacheProvider{ctrl: ctrl}
	mock.recorder = &MockICacheProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheProvider) EXPECT() *MockICacheProviderMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockICacheProvider) Clear() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockICacheProviderMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockICacheProvider)(nil).Clear))
}

// Delete mocks base method.
func (m *MockICacheProvider) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICacheProviderMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICacheProvider)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockICacheProvider) Get(key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockICacheProviderMockRecorder) Get(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICacheProvider)(nil).Get), key, value)
}

// Set mocks base method.
func (m *MockICacheProvider) Set(key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockICacheProviderMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockICacheProvider)(nil).Set), key, value)
}

// WithExpiration mocks base method.
func (m *MockICacheProvider) WithExpiration(arg0 time.Duration) cache.ICacheProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithExpiration", arg0)
	ret0, _ := ret[0].(cache.ICacheProvider)
	return ret0
}

// WithExpiration indicates an expected call of WithExpiration.
func (mr *MockICacheProviderMockRecorder) WithExpiration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithExpiration", reflect.TypeOf((*MockICacheProvider)(nil).WithExpiration), arg0)
}
