// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	dtos "RedisLesson/dtos"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockService) CreateBook(book dtos.Book) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", book)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockServiceMockRecorder) CreateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockService)(nil).CreateBook), book)
}

// DeleteBook mocks base method.
func (m *MockService) DeleteBook(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockServiceMockRecorder) DeleteBook(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockService)(nil).DeleteBook), id)
}

// GetAllBooks mocks base method.
func (m *MockService) GetAllBooks(pagination dtos.PaginationParams) ([]dtos.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks", pagination)
	ret0, _ := ret[0].([]dtos.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockServiceMockRecorder) GetAllBooks(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockService)(nil).GetAllBooks), pagination)
}

// GetBookById mocks base method.
func (m *MockService) GetBookById(id int) (dtos.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", id)
	ret0, _ := ret[0].(dtos.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockServiceMockRecorder) GetBookById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockService)(nil).GetBookById), id)
}

// GetBooksByFilter mocks base method.
func (m *MockService) GetBooksByFilter(filter dtos.BookFilter, pagination dtos.PaginationParams) ([]dtos.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByFilter", filter, pagination)
	ret0, _ := ret[0].([]dtos.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooksByFilter indicates an expected call of GetBooksByFilter.
func (mr *MockServiceMockRecorder) GetBooksByFilter(filter, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByFilter", reflect.TypeOf((*MockService)(nil).GetBooksByFilter), filter, pagination)
}

// UpdateBook mocks base method.
func (m *MockService) UpdateBook(id int, input dtos.BookUpdateInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockServiceMockRecorder) UpdateBook(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockService)(nil).UpdateBook), id, input)
}
