// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/repository/transaction/transaction.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	sql "database/sql"
	entity "dating-service/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockTransactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockTransactionRepositoryMockRecorder) BeginTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockTransactionRepository)(nil).BeginTx), ctx)
}

// CheckPaymentMethod mocks base method.
func (m *MockTransactionRepository) CheckPaymentMethod(ctx context.Context, id int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPaymentMethod", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPaymentMethod indicates an expected call of CheckPaymentMethod.
func (mr *MockTransactionRepositoryMockRecorder) CheckPaymentMethod(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPaymentMethod", reflect.TypeOf((*MockTransactionRepository)(nil).CheckPaymentMethod), ctx, id)
}

// CommitTx mocks base method.
func (m *MockTransactionRepository) CommitTx(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTx", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTx indicates an expected call of CommitTx.
func (mr *MockTransactionRepositoryMockRecorder) CommitTx(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTx", reflect.TypeOf((*MockTransactionRepository)(nil).CommitTx), ctx, tx)
}

// CountUserSwipped mocks base method.
func (m *MockTransactionRepository) CountUserSwipped(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserSwipped", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserSwipped indicates an expected call of CountUserSwipped.
func (mr *MockTransactionRepositoryMockRecorder) CountUserSwipped(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserSwipped", reflect.TypeOf((*MockTransactionRepository)(nil).CountUserSwipped), ctx)
}

// CreateSubscription mocks base method.
func (m *MockTransactionRepository) CreateSubscription(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest, trxId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, tx, request, trxId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockTransactionRepositoryMockRecorder) CreateSubscription(ctx, tx, request, trxId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockTransactionRepository)(nil).CreateSubscription), ctx, tx, request, trxId)
}

// CreateTransaction mocks base method.
func (m *MockTransactionRepository) CreateTransaction(ctx context.Context, tx *sql.Tx, request entity.TransactionRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, tx, request)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionRepositoryMockRecorder) CreateTransaction(ctx, tx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).CreateTransaction), ctx, tx, request)
}

// GetPackageById mocks base method.
func (m *MockTransactionRepository) GetPackageById(ctx context.Context, id int) (entity.PackageType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackageById", ctx, id)
	ret0, _ := ret[0].(entity.PackageType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackageById indicates an expected call of GetPackageById.
func (mr *MockTransactionRepositoryMockRecorder) GetPackageById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackageById", reflect.TypeOf((*MockTransactionRepository)(nil).GetPackageById), ctx, id)
}

// GetPackages mocks base method.
func (m *MockTransactionRepository) GetPackages(ctx context.Context) ([]*entity.PackageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackages", ctx)
	ret0, _ := ret[0].([]*entity.PackageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackages indicates an expected call of GetPackages.
func (mr *MockTransactionRepositoryMockRecorder) GetPackages(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackages", reflect.TypeOf((*MockTransactionRepository)(nil).GetPackages), ctx)
}

// GetPaymentMethods mocks base method.
func (m *MockTransactionRepository) GetPaymentMethods(ctx context.Context) ([]*entity.PaymentMethodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentMethods", ctx)
	ret0, _ := ret[0].([]*entity.PaymentMethodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentMethods indicates an expected call of GetPaymentMethods.
func (mr *MockTransactionRepositoryMockRecorder) GetPaymentMethods(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentMethods", reflect.TypeOf((*MockTransactionRepository)(nil).GetPaymentMethods), ctx)
}

// RollbackTx mocks base method.
func (m *MockTransactionRepository) RollbackTx(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTx", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTx indicates an expected call of RollbackTx.
func (mr *MockTransactionRepositoryMockRecorder) RollbackTx(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTx", reflect.TypeOf((*MockTransactionRepository)(nil).RollbackTx), ctx, tx)
}

// UpdateUserIsPremium mocks base method.
func (m *MockTransactionRepository) UpdateUserIsPremium(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserIsPremium", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserIsPremium indicates an expected call of UpdateUserIsPremium.
func (mr *MockTransactionRepositoryMockRecorder) UpdateUserIsPremium(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserIsPremium", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateUserIsPremium), ctx, tx)
}

// UpdateUserIsVerified mocks base method.
func (m *MockTransactionRepository) UpdateUserIsVerified(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserIsVerified", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserIsVerified indicates an expected call of UpdateUserIsVerified.
func (mr *MockTransactionRepositoryMockRecorder) UpdateUserIsVerified(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserIsVerified", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateUserIsVerified), ctx, tx)
}