// Code generated by MockGen. DO NOT EDIT.
// Source: find_user_by_document_repository.go

// Package mock_data_user_protocols is a generated GoMock package.
package mock_data_user_protocols

import (
	reflect "reflect"

	domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"
	gomock "github.com/golang/mock/gomock"
)

// MockFindUserByDocumentRepository is a mock of FindUserByDocumentRepository interface.
type MockFindUserByDocumentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFindUserByDocumentRepositoryMockRecorder
}

// MockFindUserByDocumentRepositoryMockRecorder is the mock recorder for MockFindUserByDocumentRepository.
type MockFindUserByDocumentRepositoryMockRecorder struct {
	mock *MockFindUserByDocumentRepository
}

// NewMockFindUserByDocumentRepository creates a new mock instance.
func NewMockFindUserByDocumentRepository(ctrl *gomock.Controller) *MockFindUserByDocumentRepository {
	mock := &MockFindUserByDocumentRepository{ctrl: ctrl}
	mock.recorder = &MockFindUserByDocumentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFindUserByDocumentRepository) EXPECT() *MockFindUserByDocumentRepositoryMockRecorder {
	return m.recorder
}

// FindByDocument mocks base method.
func (m *MockFindUserByDocumentRepository) FindByDocument(document string) (*domain_user.FindUserByDocumentUseCaseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDocument", document)
	ret0, _ := ret[0].(*domain_user.FindUserByDocumentUseCaseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDocument indicates an expected call of FindByDocument.
func (mr *MockFindUserByDocumentRepositoryMockRecorder) FindByDocument(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDocument", reflect.TypeOf((*MockFindUserByDocumentRepository)(nil).FindByDocument), document)
}