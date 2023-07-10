// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "pearshop_backend/app/domain/entity"
	dto "pearshop_backend/app/usecase/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductFind is a mock of ProductFind interface.
type MockProductFind struct {
	ctrl     *gomock.Controller
	recorder *MockProductFindMockRecorder
}

// MockProductFindMockRecorder is the mock recorder for MockProductFind.
type MockProductFindMockRecorder struct {
	mock *MockProductFind
}

// NewMockProductFind creates a new mock instance.
func NewMockProductFind(ctrl *gomock.Controller) *MockProductFind {
	mock := &MockProductFind{ctrl: ctrl}
	mock.recorder = &MockProductFindMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductFind) EXPECT() *MockProductFindMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockProductFind) Execute(ctx context.Context, req dto.ProductFindRequest, paging entity.Paging) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, req, paging)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockProductFindMockRecorder) Execute(ctx, req, paging interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockProductFind)(nil).Execute), ctx, req, paging)
}

// MockProductUpdate is a mock of ProductUpdate interface.
type MockProductUpdate struct {
	ctrl     *gomock.Controller
	recorder *MockProductUpdateMockRecorder
}

// MockProductUpdateMockRecorder is the mock recorder for MockProductUpdate.
type MockProductUpdateMockRecorder struct {
	mock *MockProductUpdate
}

// NewMockProductUpdate creates a new mock instance.
func NewMockProductUpdate(ctrl *gomock.Controller) *MockProductUpdate {
	mock := &MockProductUpdate{ctrl: ctrl}
	mock.recorder = &MockProductUpdateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUpdate) EXPECT() *MockProductUpdateMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockProductUpdate) Execute(ctx context.Context, userID, productID int, req dto.ProductSaveRequest) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, userID, productID, req)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockProductUpdateMockRecorder) Execute(ctx, userID, productID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockProductUpdate)(nil).Execute), ctx, userID, productID, req)
}

// MockProductCreate is a mock of ProductCreate interface.
type MockProductCreate struct {
	ctrl     *gomock.Controller
	recorder *MockProductCreateMockRecorder
}

// MockProductCreateMockRecorder is the mock recorder for MockProductCreate.
type MockProductCreateMockRecorder struct {
	mock *MockProductCreate
}

// NewMockProductCreate creates a new mock instance.
func NewMockProductCreate(ctrl *gomock.Controller) *MockProductCreate {
	mock := &MockProductCreate{ctrl: ctrl}
	mock.recorder = &MockProductCreateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductCreate) EXPECT() *MockProductCreateMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockProductCreate) Execute(ctx context.Context, userID, req dto.ProductSaveRequest) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, userID, req)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockProductCreateMockRecorder) Execute(ctx, userID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockProductCreate)(nil).Execute), ctx, userID, req)
}
