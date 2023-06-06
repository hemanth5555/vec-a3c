// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler-node/internal/store (interfaces: WorkflowStore,QueueStore,VENStore)

// Package store_mock is a generated GoMock package.
package store_mock

import (
	context "context"
	reflect "reflect"
	store "scheduler-node/internal/store"

	gomock "github.com/golang/mock/gomock"
)

// MockWorkflowStore is a mock of WorkflowStore interface.
type MockWorkflowStore struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowStoreMockRecorder
}

// MockWorkflowStoreMockRecorder is the mock recorder for MockWorkflowStore.
type MockWorkflowStoreMockRecorder struct {
	mock *MockWorkflowStore
}

// NewMockWorkflowStore creates a new mock instance.
func NewMockWorkflowStore(ctrl *gomock.Controller) *MockWorkflowStore {
	mock := &MockWorkflowStore{ctrl: ctrl}
	mock.recorder = &MockWorkflowStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowStore) EXPECT() *MockWorkflowStoreMockRecorder {
	return m.recorder
}

// AssignWorkflow mocks base method.
func (m *MockWorkflowStore) AssignWorkflow(arg0 context.Context, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignWorkflow indicates an expected call of AssignWorkflow.
func (mr *MockWorkflowStoreMockRecorder) AssignWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignWorkflow", reflect.TypeOf((*MockWorkflowStore)(nil).AssignWorkflow), arg0, arg1, arg2)
}

// CompleteWorkflow mocks base method.
func (m *MockWorkflowStore) CompleteWorkflow(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteWorkflow indicates an expected call of CompleteWorkflow.
func (mr *MockWorkflowStoreMockRecorder) CompleteWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteWorkflow", reflect.TypeOf((*MockWorkflowStore)(nil).CompleteWorkflow), arg0, arg1)
}

// GetWorkflowByID mocks base method.
func (m *MockWorkflowStore) GetWorkflowByID(arg0 context.Context, arg1 int) (*store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowByID", arg0, arg1)
	ret0, _ := ret[0].(*store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByID indicates an expected call of GetWorkflowByID.
func (mr *MockWorkflowStoreMockRecorder) GetWorkflowByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowByID", reflect.TypeOf((*MockWorkflowStore)(nil).GetWorkflowByID), arg0, arg1)
}

// GetWorkflows mocks base method.
func (m *MockWorkflowStore) GetWorkflows(arg0 context.Context) ([]store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflows", arg0)
	ret0, _ := ret[0].([]store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflows indicates an expected call of GetWorkflows.
func (mr *MockWorkflowStoreMockRecorder) GetWorkflows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflows", reflect.TypeOf((*MockWorkflowStore)(nil).GetWorkflows), arg0)
}

// SaveWorkflow mocks base method.
func (m *MockWorkflowStore) SaveWorkflow(arg0 context.Context, arg1 *store.WorkflowInfo) (*store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveWorkflow indicates an expected call of SaveWorkflow.
func (mr *MockWorkflowStoreMockRecorder) SaveWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWorkflow", reflect.TypeOf((*MockWorkflowStore)(nil).SaveWorkflow), arg0, arg1)
}

// StartWorkflow mocks base method.
func (m *MockWorkflowStore) StartWorkflow(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWorkflow indicates an expected call of StartWorkflow.
func (mr *MockWorkflowStoreMockRecorder) StartWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflow", reflect.TypeOf((*MockWorkflowStore)(nil).StartWorkflow), arg0, arg1)
}

// UpdateWorkflow mocks base method.
func (m *MockWorkflowStore) UpdateWorkflow(arg0 context.Context, arg1 *store.WorkflowInfo) (*store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflow indicates an expected call of UpdateWorkflow.
func (mr *MockWorkflowStoreMockRecorder) UpdateWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflow", reflect.TypeOf((*MockWorkflowStore)(nil).UpdateWorkflow), arg0, arg1)
}

// MockQueueStore is a mock of QueueStore interface.
type MockQueueStore struct {
	ctrl     *gomock.Controller
	recorder *MockQueueStoreMockRecorder
}

// MockQueueStoreMockRecorder is the mock recorder for MockQueueStore.
type MockQueueStoreMockRecorder struct {
	mock *MockQueueStore
}

// NewMockQueueStore creates a new mock instance.
func NewMockQueueStore(ctrl *gomock.Controller) *MockQueueStore {
	mock := &MockQueueStore{ctrl: ctrl}
	mock.recorder = &MockQueueStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueStore) EXPECT() *MockQueueStoreMockRecorder {
	return m.recorder
}

// AssignWorkflow mocks base method.
func (m *MockQueueStore) AssignWorkflow(arg0 context.Context, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignWorkflow indicates an expected call of AssignWorkflow.
func (mr *MockQueueStoreMockRecorder) AssignWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignWorkflow", reflect.TypeOf((*MockQueueStore)(nil).AssignWorkflow), arg0, arg1, arg2)
}

// CompleteWorkflow mocks base method.
func (m *MockQueueStore) CompleteWorkflow(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteWorkflow indicates an expected call of CompleteWorkflow.
func (mr *MockQueueStoreMockRecorder) CompleteWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteWorkflow", reflect.TypeOf((*MockQueueStore)(nil).CompleteWorkflow), arg0, arg1)
}

// GetQueue mocks base method.
func (m *MockQueueStore) GetQueue(arg0 context.Context) ([]store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueue", arg0)
	ret0, _ := ret[0].([]store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueue indicates an expected call of GetQueue.
func (mr *MockQueueStoreMockRecorder) GetQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueue", reflect.TypeOf((*MockQueueStore)(nil).GetQueue), arg0)
}

// Peek mocks base method.
func (m *MockQueueStore) Peek(arg0 context.Context) (*store.WorkflowInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek", arg0)
	ret0, _ := ret[0].(*store.WorkflowInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peek indicates an expected call of Peek.
func (mr *MockQueueStoreMockRecorder) Peek(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockQueueStore)(nil).Peek), arg0)
}

// StartWorkflow mocks base method.
func (m *MockQueueStore) StartWorkflow(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWorkflow indicates an expected call of StartWorkflow.
func (mr *MockQueueStoreMockRecorder) StartWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflow", reflect.TypeOf((*MockQueueStore)(nil).StartWorkflow), arg0, arg1)
}

// MockVENStore is a mock of VENStore interface.
type MockVENStore struct {
	ctrl     *gomock.Controller
	recorder *MockVENStoreMockRecorder
}

// MockVENStoreMockRecorder is the mock recorder for MockVENStore.
type MockVENStoreMockRecorder struct {
	mock *MockVENStore
}

// NewMockVENStore creates a new mock instance.
func NewMockVENStore(ctrl *gomock.Controller) *MockVENStore {
	mock := &MockVENStore{ctrl: ctrl}
	mock.recorder = &MockVENStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVENStore) EXPECT() *MockVENStoreMockRecorder {
	return m.recorder
}

// GetVENInfos mocks base method.
func (m *MockVENStore) GetVENInfos() ([]store.VENInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVENInfos")
	ret0, _ := ret[0].([]store.VENInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVENInfos indicates an expected call of GetVENInfos.
func (mr *MockVENStoreMockRecorder) GetVENInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVENInfos", reflect.TypeOf((*MockVENStore)(nil).GetVENInfos))
}
