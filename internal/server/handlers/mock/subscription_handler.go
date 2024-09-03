// Code generated by MockGen. DO NOT EDIT.
// Source: subscription_handler.go

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	models "rate/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISubscriptionService is a mock of ISubscriptionService interface.
type MockISubscriptionService struct {
	ctrl     *gomock.Controller
	recorder *MockISubscriptionServiceMockRecorder
}

// MockISubscriptionServiceMockRecorder is the mock recorder for MockISubscriptionService.
type MockISubscriptionServiceMockRecorder struct {
	mock *MockISubscriptionService
}

// NewMockISubscriptionService creates a new mock instance.
func NewMockISubscriptionService(ctrl *gomock.Controller) *MockISubscriptionService {
	mock := &MockISubscriptionService{ctrl: ctrl}
	mock.recorder = &MockISubscriptionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISubscriptionService) EXPECT() *MockISubscriptionServiceMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockISubscriptionService) Subscribe(email models.Email) (*models.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", email)
	ret0, _ := ret[0].(*models.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockISubscriptionServiceMockRecorder) Subscribe(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockISubscriptionService)(nil).Subscribe), email)
}