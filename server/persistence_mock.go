// Code generated by MockGen. DO NOT EDIT.
// Source: server/persistence.go

// Package server is a generated GoMock package.
package server

import (
	config "github.com/DrmagicE/gmqtt/config"
	queue "github.com/DrmagicE/gmqtt/persistence/queue"
	session "github.com/DrmagicE/gmqtt/persistence/session"
	subscription "github.com/DrmagicE/gmqtt/persistence/subscription"
	unack "github.com/DrmagicE/gmqtt/persistence/unack"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPersistence is a mock of Persistence interface
type MockPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceMockRecorder
}

// MockPersistenceMockRecorder is the mock recorder for MockPersistence
type MockPersistenceMockRecorder struct {
	mock *MockPersistence
}

// NewMockPersistence creates a new mock instance
func NewMockPersistence(ctrl *gomock.Controller) *MockPersistence {
	mock := &MockPersistence{ctrl: ctrl}
	mock.recorder = &MockPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersistence) EXPECT() *MockPersistenceMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockPersistence) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockPersistenceMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockPersistence)(nil).Open))
}

// NewQueueStore mocks base method
func (m *MockPersistence) NewQueueStore(config config.Config, defaultNotifier queue.Notifier, clientID string) (queue.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewQueueStore", config, defaultNotifier, clientID)
	ret0, _ := ret[0].(queue.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewQueueStore indicates an expected call of NewQueueStore
func (mr *MockPersistenceMockRecorder) NewQueueStore(config, defaultNotifier, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQueueStore", reflect.TypeOf((*MockPersistence)(nil).NewQueueStore), config, defaultNotifier, clientID)
}

// NewSubscriptionStore mocks base method
func (m *MockPersistence) NewSubscriptionStore(config config.Config) (subscription.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSubscriptionStore", config)
	ret0, _ := ret[0].(subscription.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSubscriptionStore indicates an expected call of NewSubscriptionStore
func (mr *MockPersistenceMockRecorder) NewSubscriptionStore(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSubscriptionStore", reflect.TypeOf((*MockPersistence)(nil).NewSubscriptionStore), config)
}

// NewSessionStore mocks base method
func (m *MockPersistence) NewSessionStore(config config.Config) (session.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSessionStore", config)
	ret0, _ := ret[0].(session.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSessionStore indicates an expected call of NewSessionStore
func (mr *MockPersistenceMockRecorder) NewSessionStore(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSessionStore", reflect.TypeOf((*MockPersistence)(nil).NewSessionStore), config)
}

// NewUnackStore mocks base method
func (m *MockPersistence) NewUnackStore(config config.Config, clientID string) (unack.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUnackStore", config, clientID)
	ret0, _ := ret[0].(unack.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUnackStore indicates an expected call of NewUnackStore
func (mr *MockPersistenceMockRecorder) NewUnackStore(config, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUnackStore", reflect.TypeOf((*MockPersistence)(nil).NewUnackStore), config, clientID)
}

// Close mocks base method
func (m *MockPersistence) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPersistenceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPersistence)(nil).Close))
}
