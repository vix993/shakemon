// Code generated by MockGen. DO NOT EDIT.
// Source: services.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/vix993/shakemon/domain/entities"
	models "github.com/vix993/shakemon/internal/api/models"
)

// MockTranslateService is a mock of TranslateService interface.
type MockTranslateService struct {
	ctrl     *gomock.Controller
	recorder *MockTranslateServiceMockRecorder
}

// MockTranslateServiceMockRecorder is the mock recorder for MockTranslateService.
type MockTranslateServiceMockRecorder struct {
	mock *MockTranslateService
}

// NewMockTranslateService creates a new mock instance.
func NewMockTranslateService(ctrl *gomock.Controller) *MockTranslateService {
	mock := &MockTranslateService{ctrl: ctrl}
	mock.recorder = &MockTranslateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslateService) EXPECT() *MockTranslateServiceMockRecorder {
	return m.recorder
}

// DoGetPokemonThenTranslate mocks base method.
func (m *MockTranslateService) DoGetPokemonThenTranslate(name string) (*entities.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoGetPokemonThenTranslate", name)
	ret0, _ := ret[0].(*entities.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoGetPokemonThenTranslate indicates an expected call of DoGetPokemonThenTranslate.
func (mr *MockTranslateServiceMockRecorder) DoGetPokemonThenTranslate(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoGetPokemonThenTranslate", reflect.TypeOf((*MockTranslateService)(nil).DoGetPokemonThenTranslate), name)
}

// GetPokemon mocks base method.
func (m *MockTranslateService) GetPokemon(name string) (*entities.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemon", name)
	ret0, _ := ret[0].(*entities.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPokemon indicates an expected call of GetPokemon.
func (mr *MockTranslateServiceMockRecorder) GetPokemon(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemon", reflect.TypeOf((*MockTranslateService)(nil).GetPokemon), name)
}

// GetShakespeareTranslation mocks base method.
func (m *MockTranslateService) GetShakespeareTranslation(request *models.ShakespeareTranslationRequest) (*models.ShakespeareTranslationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShakespeareTranslation", request)
	ret0, _ := ret[0].(*models.ShakespeareTranslationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShakespeareTranslation indicates an expected call of GetShakespeareTranslation.
func (mr *MockTranslateServiceMockRecorder) GetShakespeareTranslation(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShakespeareTranslation", reflect.TypeOf((*MockTranslateService)(nil).GetShakespeareTranslation), request)
}
