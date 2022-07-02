// Code generated by MockGen. DO NOT EDIT.
// Source: dal/telegram/repository.go

// Package telegram is a generated GoMock package.
package telegram

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/orenoid/telegram-account-bot/models"
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

// CreateOrUpdateTelegramUser mocks base method.
func (m *MockRepository) CreateOrUpdateTelegramUser(userID int64, userName string, chatID int64) (*models.TelegramUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateTelegramUser", userID, userName, chatID)
	ret0, _ := ret[0].(*models.TelegramUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateTelegramUser indicates an expected call of CreateOrUpdateTelegramUser.
func (mr *MockRepositoryMockRecorder) CreateOrUpdateTelegramUser(userID, userName, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateTelegramUser", reflect.TypeOf((*MockRepository)(nil).CreateOrUpdateTelegramUser), userID, userName, chatID)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(teleUserID int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", teleUserID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(teleUserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), teleUserID)
}