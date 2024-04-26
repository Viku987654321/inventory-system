// Code generated by MockGen. DO NOT EDIT.
// Source: inventory-system/inventory-service/internal/domain/service (interfaces: IInventoryService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	dto "inventory-system/common/pkg/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockIInventoryService is a mock of IInventoryService interface.
type MockIInventoryService struct {
	ctrl     *gomock.Controller
	recorder *MockIInventoryServiceMockRecorder
}

// MockIInventoryServiceMockRecorder is the mock recorder for MockIInventoryService.
type MockIInventoryServiceMockRecorder struct {
	mock *MockIInventoryService
}

// NewMockIInventoryService creates a new mock instance.
func NewMockIInventoryService(ctrl *gomock.Controller) *MockIInventoryService {
	mock := &MockIInventoryService{ctrl: ctrl}
	mock.recorder = &MockIInventoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInventoryService) EXPECT() *MockIInventoryServiceMockRecorder {
	return m.recorder
}

// CreateNewInventory mocks base method.
func (m *MockIInventoryService) CreateNewInventory(arg0 context.Context, arg1 interface{}, arg2 string) *dto.ErrorResponseDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewInventory", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.ErrorResponseDto)
	return ret0
}

// CreateNewInventory indicates an expected call of CreateNewInventory.
func (mr *MockIInventoryServiceMockRecorder) CreateNewInventory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewInventory", reflect.TypeOf((*MockIInventoryService)(nil).CreateNewInventory), arg0, arg1, arg2)
}

// GetInventory mocks base method.
func (m *MockIInventoryService) GetInventory(arg0 context.Context, arg1 string, arg2 map[string][]string) (primitive.M, *dto.ErrorResponseDto) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInventory", arg0, arg1, arg2)
	ret0, _ := ret[0].(primitive.M)
	ret1, _ := ret[1].(*dto.ErrorResponseDto)
	return ret0, ret1
}

// GetInventory indicates an expected call of GetInventory.
func (mr *MockIInventoryServiceMockRecorder) GetInventory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInventory", reflect.TypeOf((*MockIInventoryService)(nil).GetInventory), arg0, arg1, arg2)
}
