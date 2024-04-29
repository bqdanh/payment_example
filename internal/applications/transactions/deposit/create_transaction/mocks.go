// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go

// Package create_transaction is a generated GoMock package.
package create_transaction

import (
	context "context"
	reflect "reflect"
	time "time"

	account "github.com/bqdanh/money_transfer/internal/entities/account"
	transaction "github.com/bqdanh/money_transfer/internal/entities/transaction"
	gomock "github.com/golang/mock/gomock"
)

// MockdistributeLock is a mock of distributeLock interface.
type MockdistributeLock struct {
	ctrl     *gomock.Controller
	recorder *MockdistributeLockMockRecorder
}

// MockdistributeLockMockRecorder is the mock recorder for MockdistributeLock.
type MockdistributeLockMockRecorder struct {
	mock *MockdistributeLock
}

// NewMockdistributeLock creates a new mock instance.
func NewMockdistributeLock(ctrl *gomock.Controller) *MockdistributeLock {
	mock := &MockdistributeLock{ctrl: ctrl}
	mock.recorder = &MockdistributeLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdistributeLock) EXPECT() *MockdistributeLockMockRecorder {
	return m.recorder
}

// AcquireLockForCreateDepositTransaction mocks base method.
func (m *MockdistributeLock) AcquireLockForCreateDepositTransaction(ctx context.Context, requestID string, lockDuration time.Duration) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireLockForCreateDepositTransaction", ctx, requestID, lockDuration)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireLockForCreateDepositTransaction indicates an expected call of AcquireLockForCreateDepositTransaction.
func (mr *MockdistributeLockMockRecorder) AcquireLockForCreateDepositTransaction(ctx, requestID, lockDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireLockForCreateDepositTransaction", reflect.TypeOf((*MockdistributeLock)(nil).AcquireLockForCreateDepositTransaction), ctx, requestID, lockDuration)
}

// MockaccountRepository is a mock of accountRepository interface.
type MockaccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockaccountRepositoryMockRecorder
}

// MockaccountRepositoryMockRecorder is the mock recorder for MockaccountRepository.
type MockaccountRepositoryMockRecorder struct {
	mock *MockaccountRepository
}

// NewMockaccountRepository creates a new mock instance.
func NewMockaccountRepository(ctrl *gomock.Controller) *MockaccountRepository {
	mock := &MockaccountRepository{ctrl: ctrl}
	mock.recorder = &MockaccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockaccountRepository) EXPECT() *MockaccountRepositoryMockRecorder {
	return m.recorder
}

// GetAccountsByID mocks base method.
func (m *MockaccountRepository) GetAccountsByID(ctx context.Context, accountID int64) (account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByID", ctx, accountID)
	ret0, _ := ret[0].(account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByID indicates an expected call of GetAccountsByID.
func (mr *MockaccountRepositoryMockRecorder) GetAccountsByID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByID", reflect.TypeOf((*MockaccountRepository)(nil).GetAccountsByID), ctx, accountID)
}

// MocktransactionRepository is a mock of transactionRepository interface.
type MocktransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MocktransactionRepositoryMockRecorder
}

// MocktransactionRepositoryMockRecorder is the mock recorder for MocktransactionRepository.
type MocktransactionRepositoryMockRecorder struct {
	mock *MocktransactionRepository
}

// NewMocktransactionRepository creates a new mock instance.
func NewMocktransactionRepository(ctrl *gomock.Controller) *MocktransactionRepository {
	mock := &MocktransactionRepository{ctrl: ctrl}
	mock.recorder = &MocktransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktransactionRepository) EXPECT() *MocktransactionRepositoryMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MocktransactionRepository) CreateTransaction(ctx context.Context, t transaction.Transaction) (transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, t)
	ret0, _ := ret[0].(transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MocktransactionRepositoryMockRecorder) CreateTransaction(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MocktransactionRepository)(nil).CreateTransaction), ctx, t)
}

// GetTransactionByRequestID mocks base method.
func (m *MocktransactionRepository) GetTransactionByRequestID(ctx context.Context, account account.Account, requestID string) (transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByRequestID", ctx, account, requestID)
	ret0, _ := ret[0].(transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByRequestID indicates an expected call of GetTransactionByRequestID.
func (mr *MocktransactionRepositoryMockRecorder) GetTransactionByRequestID(ctx, account, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByRequestID", reflect.TypeOf((*MocktransactionRepository)(nil).GetTransactionByRequestID), ctx, account, requestID)
}
