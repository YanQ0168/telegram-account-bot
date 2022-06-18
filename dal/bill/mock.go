// Code generated by MockGen. DO NOT EDIT.
// Source: dal/bill/repository.go

// Package bill is a generated GoMock package.
package bill

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/orenoid/account-bot/models"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateBillAndUpdateUserBalance mocks base method.
func (m *MockRepository) CreateBillAndUpdateUserBalance(userID uint, amount float64, category string, opts ...CreateBillOptions) (*models.Bill, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{userID, amount, category}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBillAndUpdateUserBalance", varargs...)
	ret0, _ := ret[0].(*models.Bill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBillAndUpdateUserBalance indicates an expected call of CreateBillAndUpdateUserBalance.
func (mr *MockRepositoryMockRecorder) CreateBillAndUpdateUserBalance(userID, amount, category interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{userID, amount, category}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillAndUpdateUserBalance", reflect.TypeOf((*MockRepository)(nil).CreateBillAndUpdateUserBalance), varargs...)
}
