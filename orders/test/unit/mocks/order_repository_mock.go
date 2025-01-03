// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/order_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/order_repository.go -destination=test/unit/mocks/order_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	domain "ticketing/orders/internal/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
	isgomock struct{}
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderRepository) Create(ctx context.Context, order *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepository)(nil).Create), ctx, order)
}

// FindAll mocks base method.
func (m *MockOrderRepository) FindAll(ctx context.Context, userId string) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, userId)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockOrderRepositoryMockRecorder) FindAll(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockOrderRepository)(nil).FindAll), ctx, userId)
}

// FindById mocks base method.
func (m *MockOrderRepository) FindById(ctx context.Context, id int32) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockOrderRepositoryMockRecorder) FindById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockOrderRepository)(nil).FindById), ctx, id)
}

// IsTicketReserved mocks base method.
func (m *MockOrderRepository) IsTicketReserved(ctx context.Context, ticketId int32) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTicketReserved", ctx, ticketId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsTicketReserved indicates an expected call of IsTicketReserved.
func (mr *MockOrderRepositoryMockRecorder) IsTicketReserved(ctx, ticketId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTicketReserved", reflect.TypeOf((*MockOrderRepository)(nil).IsTicketReserved), ctx, ticketId)
}

// Update mocks base method.
func (m *MockOrderRepository) Update(ctx context.Context, order *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderRepositoryMockRecorder) Update(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderRepository)(nil).Update), ctx, order)
}