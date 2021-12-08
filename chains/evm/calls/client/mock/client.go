// Code generated by MockGen. DO NOT EDIT.
// Source: chains/evm/calls/client/client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmclient"
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockClientContractChecker is a mock of ClientContractChecker interface.
type MockClientContractChecker struct {
	ctrl     *gomock.Controller
	recorder *MockClientContractCheckerMockRecorder
}

// MockClientContractCheckerMockRecorder is the mock recorder for MockClientContractChecker.
type MockClientContractCheckerMockRecorder struct {
	mock *MockClientContractChecker
}

// NewMockClientContractChecker creates a new mock instance.
func NewMockClientContractChecker(ctrl *gomock.Controller) *MockClientContractChecker {
	mock := &MockClientContractChecker{ctrl: ctrl}
	mock.recorder = &MockClientContractCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientContractChecker) EXPECT() *MockClientContractCheckerMockRecorder {
	return m.recorder
}

// CodeAt mocks base method.
func (m *MockClientContractChecker) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockClientContractCheckerMockRecorder) CodeAt(ctx, contract, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockClientContractChecker)(nil).CodeAt), ctx, contract, blockNumber)
}

// MockContractCallerClient is a mock of ContractCallerClient interface.
type MockContractCallerClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractCallerClientMockRecorder
}

// MockContractCallerClientMockRecorder is the mock recorder for MockContractCallerClient.
type MockContractCallerClientMockRecorder struct {
	mock *MockContractCallerClient
}

// NewMockContractCallerClient creates a new mock instance.
func NewMockContractCallerClient(ctrl *gomock.Controller) *MockContractCallerClient {
	mock := &MockContractCallerClient{ctrl: ctrl}
	mock.recorder = &MockContractCallerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCallerClient) EXPECT() *MockContractCallerClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCallerClient) CallContract(ctx context.Context, callArgs map[string]interface{}, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCallerClientMockRecorder) CallContract(ctx, callArgs, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCallerClient)(nil).CallContract), ctx, callArgs, blockNumber)
}

// MockContractCheckerCallerClient is a mock of ContractCheckerCallerClient interface.
type MockContractCheckerCallerClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractCheckerCallerClientMockRecorder
}

// MockContractCheckerCallerClientMockRecorder is the mock recorder for MockContractCheckerCallerClient.
type MockContractCheckerCallerClientMockRecorder struct {
	mock *MockContractCheckerCallerClient
}

// NewMockContractCheckerCallerClient creates a new mock instance.
func NewMockContractCheckerCallerClient(ctrl *gomock.Controller) *MockContractCheckerCallerClient {
	mock := &MockContractCheckerCallerClient{ctrl: ctrl}
	mock.recorder = &MockContractCheckerCallerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCheckerCallerClient) EXPECT() *MockContractCheckerCallerClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCheckerCallerClient) CallContract(ctx context.Context, callArgs map[string]interface{}, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCheckerCallerClientMockRecorder) CallContract(ctx, callArgs, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCheckerCallerClient)(nil).CallContract), ctx, callArgs, blockNumber)
}

// CodeAt mocks base method.
func (m *MockContractCheckerCallerClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockContractCheckerCallerClientMockRecorder) CodeAt(ctx, contract, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockContractCheckerCallerClient)(nil).CodeAt), ctx, contract, blockNumber)
}

// MockClientDeployer is a mock of ClientDeployer interface.
type MockClientDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockClientDeployerMockRecorder
}

// MockClientDeployerMockRecorder is the mock recorder for MockClientDeployer.
type MockClientDeployerMockRecorder struct {
	mock *MockClientDeployer
}

// NewMockClientDeployer creates a new mock instance.
func NewMockClientDeployer(ctrl *gomock.Controller) *MockClientDeployer {
	mock := &MockClientDeployer{ctrl: ctrl}
	mock.recorder = &MockClientDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientDeployer) EXPECT() *MockClientDeployerMockRecorder {
	return m.recorder
}

// CodeAt mocks base method.
func (m *MockClientDeployer) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockClientDeployerMockRecorder) CodeAt(ctx, contract, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockClientDeployer)(nil).CodeAt), ctx, contract, blockNumber)
}

// From mocks base method.
func (m *MockClientDeployer) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockClientDeployerMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockClientDeployer)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockClientDeployer) GetTransactionByHash(h common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", h)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockClientDeployerMockRecorder) GetTransactionByHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockClientDeployer)(nil).GetTransactionByHash), h)
}

// LockNonce mocks base method.
func (m *MockClientDeployer) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockClientDeployerMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockClientDeployer)(nil).LockNonce))
}

// SignAndSendTransaction mocks base method.
func (m *MockClientDeployer) SignAndSendTransaction(ctx context.Context, tx evmclient.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", ctx, tx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockClientDeployerMockRecorder) SignAndSendTransaction(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockClientDeployer)(nil).SignAndSendTransaction), ctx, tx)
}

// UnlockNonce mocks base method.
func (m *MockClientDeployer) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockClientDeployerMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockClientDeployer)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockClientDeployer) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockClientDeployerMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockClientDeployer)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockClientDeployer) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockClientDeployerMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockClientDeployer)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockClientDeployer) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockClientDeployerMockRecorder) WaitAndReturnTxReceipt(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockClientDeployer)(nil).WaitAndReturnTxReceipt), h)
}

// MockContractCallerDispatcherClient is a mock of ContractCallerDispatcherClient interface.
type MockContractCallerDispatcherClient struct {
	ctrl     *gomock.Controller
	recorder *MockContractCallerDispatcherClientMockRecorder
}

// MockContractCallerDispatcherClientMockRecorder is the mock recorder for MockContractCallerDispatcherClient.
type MockContractCallerDispatcherClientMockRecorder struct {
	mock *MockContractCallerDispatcherClient
}

// NewMockContractCallerDispatcherClient creates a new mock instance.
func NewMockContractCallerDispatcherClient(ctrl *gomock.Controller) *MockContractCallerDispatcherClient {
	mock := &MockContractCallerDispatcherClient{ctrl: ctrl}
	mock.recorder = &MockContractCallerDispatcherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCallerDispatcherClient) EXPECT() *MockContractCallerDispatcherClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCallerDispatcherClient) CallContract(ctx context.Context, callArgs map[string]interface{}, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCallerDispatcherClientMockRecorder) CallContract(ctx, callArgs, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).CallContract), ctx, callArgs, blockNumber)
}

// CodeAt mocks base method.
func (m *MockContractCallerDispatcherClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockContractCallerDispatcherClientMockRecorder) CodeAt(ctx, contract, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).CodeAt), ctx, contract, blockNumber)
}

// From mocks base method.
func (m *MockContractCallerDispatcherClient) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockContractCallerDispatcherClientMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockContractCallerDispatcherClient) GetTransactionByHash(h common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", h)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockContractCallerDispatcherClientMockRecorder) GetTransactionByHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).GetTransactionByHash), h)
}

// LockNonce mocks base method.
func (m *MockContractCallerDispatcherClient) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockContractCallerDispatcherClientMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).LockNonce))
}

// SignAndSendTransaction mocks base method.
func (m *MockContractCallerDispatcherClient) SignAndSendTransaction(ctx context.Context, tx evmclient.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", ctx, tx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockContractCallerDispatcherClientMockRecorder) SignAndSendTransaction(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).SignAndSendTransaction), ctx, tx)
}

// UnlockNonce mocks base method.
func (m *MockContractCallerDispatcherClient) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockContractCallerDispatcherClientMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockContractCallerDispatcherClient) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockContractCallerDispatcherClientMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockContractCallerDispatcherClient) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockContractCallerDispatcherClientMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockContractCallerDispatcherClient) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockContractCallerDispatcherClientMockRecorder) WaitAndReturnTxReceipt(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockContractCallerDispatcherClient)(nil).WaitAndReturnTxReceipt), h)
}

// MockGasPricer is a mock of GasPricer interface.
type MockGasPricer struct {
	ctrl     *gomock.Controller
	recorder *MockGasPricerMockRecorder
}

// MockGasPricerMockRecorder is the mock recorder for MockGasPricer.
type MockGasPricerMockRecorder struct {
	mock *MockGasPricer
}

// NewMockGasPricer creates a new mock instance.
func NewMockGasPricer(ctrl *gomock.Controller) *MockGasPricer {
	mock := &MockGasPricer{ctrl: ctrl}
	mock.recorder = &MockGasPricerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasPricer) EXPECT() *MockGasPricerMockRecorder {
	return m.recorder
}

// GasPrice mocks base method.
func (m *MockGasPricer) GasPrice() ([]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice")
	ret0, _ := ret[0].([]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice.
func (mr *MockGasPricerMockRecorder) GasPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockGasPricer)(nil).GasPrice))
}

// MockClientDispatcher is a mock of ClientDispatcher interface.
type MockClientDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockClientDispatcherMockRecorder
}

// MockClientDispatcherMockRecorder is the mock recorder for MockClientDispatcher.
type MockClientDispatcherMockRecorder struct {
	mock *MockClientDispatcher
}

// NewMockClientDispatcher creates a new mock instance.
func NewMockClientDispatcher(ctrl *gomock.Controller) *MockClientDispatcher {
	mock := &MockClientDispatcher{ctrl: ctrl}
	mock.recorder = &MockClientDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientDispatcher) EXPECT() *MockClientDispatcherMockRecorder {
	return m.recorder
}

// From mocks base method.
func (m *MockClientDispatcher) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockClientDispatcherMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockClientDispatcher)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockClientDispatcher) GetTransactionByHash(h common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", h)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockClientDispatcherMockRecorder) GetTransactionByHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockClientDispatcher)(nil).GetTransactionByHash), h)
}

// LockNonce mocks base method.
func (m *MockClientDispatcher) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockClientDispatcherMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockClientDispatcher)(nil).LockNonce))
}

// SignAndSendTransaction mocks base method.
func (m *MockClientDispatcher) SignAndSendTransaction(ctx context.Context, tx evmclient.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", ctx, tx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockClientDispatcherMockRecorder) SignAndSendTransaction(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockClientDispatcher)(nil).SignAndSendTransaction), ctx, tx)
}

// UnlockNonce mocks base method.
func (m *MockClientDispatcher) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockClientDispatcherMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockClientDispatcher)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockClientDispatcher) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockClientDispatcherMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockClientDispatcher)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockClientDispatcher) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockClientDispatcherMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockClientDispatcher)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockClientDispatcher) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockClientDispatcherMockRecorder) WaitAndReturnTxReceipt(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockClientDispatcher)(nil).WaitAndReturnTxReceipt), h)
}

// MockSimulateCallerClient is a mock of SimulateCallerClient interface.
type MockSimulateCallerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSimulateCallerClientMockRecorder
}

// MockSimulateCallerClientMockRecorder is the mock recorder for MockSimulateCallerClient.
type MockSimulateCallerClientMockRecorder struct {
	mock *MockSimulateCallerClient
}

// NewMockSimulateCallerClient creates a new mock instance.
func NewMockSimulateCallerClient(ctrl *gomock.Controller) *MockSimulateCallerClient {
	mock := &MockSimulateCallerClient{ctrl: ctrl}
	mock.recorder = &MockSimulateCallerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimulateCallerClient) EXPECT() *MockSimulateCallerClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockSimulateCallerClient) CallContract(ctx context.Context, callArgs map[string]interface{}, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockSimulateCallerClientMockRecorder) CallContract(ctx, callArgs, blockNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockSimulateCallerClient)(nil).CallContract), ctx, callArgs, blockNumber)
}

// TransactionByHash mocks base method.
func (m *MockSimulateCallerClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockSimulateCallerClientMockRecorder) TransactionByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockSimulateCallerClient)(nil).TransactionByHash), ctx, hash)
}
