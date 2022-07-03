// Code generated by MockGen. DO NOT EDIT.
// Source: API_Youtube/services (interfaces: TodoService)

// Package services is a generated GoMock package.
package services

import (
	dto "GO_API/dto"
	models "GO_API/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTodoService is a mock of TodoService interface.
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService.
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance.
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// TodoDelete mocks base method.
func (m *MockTodoService) TodoDelete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoDelete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TodoDelete indicates an expected call of TodoDelete.
func (mr *MockTodoServiceMockRecorder) TodoDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoDelete", reflect.TypeOf((*MockTodoService)(nil).TodoDelete), arg0)
}

// TodoGetAll mocks base method.
func (m *MockTodoService) TodoGetAll() ([]models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoGetAll")
	ret0, _ := ret[0].([]models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TodoGetAll indicates an expected call of TodoGetAll.
func (mr *MockTodoServiceMockRecorder) TodoGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoGetAll", reflect.TypeOf((*MockTodoService)(nil).TodoGetAll))
}

// TodoInsert mocks base method.
func (m *MockTodoService) TodoInsert(arg0 models.Todo) (*dto.TodoDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoInsert", arg0)
	ret0, _ := ret[0].(*dto.TodoDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TodoInsert indicates an expected call of TodoInsert.
func (mr *MockTodoServiceMockRecorder) TodoInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoInsert", reflect.TypeOf((*MockTodoService)(nil).TodoInsert), arg0)
}