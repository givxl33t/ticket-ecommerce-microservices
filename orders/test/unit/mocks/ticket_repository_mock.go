// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/ticket_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/ticket_repository.go -destination=test/unit/mocks/ticket_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	domain "ticketing/orders/internal/domain"

	gomock "go.uber.org/mock/gomock"
)

// MockTicketRepository is a mock of TicketRepository interface.
type MockTicketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepositoryMockRecorder
	isgomock struct{}
}

// MockTicketRepositoryMockRecorder is the mock recorder for MockTicketRepository.
type MockTicketRepositoryMockRecorder struct {
	mock *MockTicketRepository
}

// NewMockTicketRepository creates a new mock instance.
func NewMockTicketRepository(ctrl *gomock.Controller) *MockTicketRepository {
	mock := &MockTicketRepository{ctrl: ctrl}
	mock.recorder = &MockTicketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRepository) EXPECT() *MockTicketRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTicketRepository) Create(ctx context.Context, order *domain.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTicketRepositoryMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTicketRepository)(nil).Create), ctx, order)
}

// FindById mocks base method.
func (m *MockTicketRepository) FindById(ctx context.Context, id int32) (*domain.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*domain.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockTicketRepositoryMockRecorder) FindById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTicketRepository)(nil).FindById), ctx, id)
}

// Update mocks base method.
func (m *MockTicketRepository) Update(ctx context.Context, order *domain.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTicketRepositoryMockRecorder) Update(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTicketRepository)(nil).Update), ctx, order)
}
