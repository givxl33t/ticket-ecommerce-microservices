// Code generated by MockGen. DO NOT EDIT.
// Source: internal/publisher/payment_publisher.go
//
// Generated by this command:
//
//	mockgen -source=internal/publisher/payment_publisher.go -destination=test/unit/mocks/payment_publisher_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "ticketing/payments/internal/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockPaymentPublisher is a mock of PaymentPublisher interface.
type MockPaymentPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentPublisherMockRecorder
	isgomock struct{}
}

// MockPaymentPublisherMockRecorder is the mock recorder for MockPaymentPublisher.
type MockPaymentPublisherMockRecorder struct {
	mock *MockPaymentPublisher
}

// NewMockPaymentPublisher creates a new mock instance.
func NewMockPaymentPublisher(ctrl *gomock.Controller) *MockPaymentPublisher {
	mock := &MockPaymentPublisher{ctrl: ctrl}
	mock.recorder = &MockPaymentPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentPublisher) EXPECT() *MockPaymentPublisherMockRecorder {
	return m.recorder
}

// Created mocks base method.
func (m *MockPaymentPublisher) Created(order *domain.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Created", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Created indicates an expected call of Created.
func (mr *MockPaymentPublisherMockRecorder) Created(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Created", reflect.TypeOf((*MockPaymentPublisher)(nil).Created), order)
}
