// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/tools/providers/http_client/http_client_provider.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHTTPClientProvider is a mock of IHTTPClientProvider interface.
type MockIHTTPClientProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIHTTPClientProviderMockRecorder
}

// MockIHTTPClientProviderMockRecorder is the mock recorder for MockIHTTPClientProvider.
type MockIHTTPClientProviderMockRecorder struct {
	mock *MockIHTTPClientProvider
}

// NewMockIHTTPClientProvider creates a new mock instance.
func NewMockIHTTPClientProvider(ctrl *gomock.Controller) *MockIHTTPClientProvider {
	mock := &MockIHTTPClientProvider{ctrl: ctrl}
	mock.recorder = &MockIHTTPClientProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHTTPClientProvider) EXPECT() *MockIHTTPClientProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIHTTPClientProvider) Get(url string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIHTTPClientProviderMockRecorder) Get(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIHTTPClientProvider)(nil).Get), url)
}
