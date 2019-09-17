// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import big "math/big"
import common "github.com/ethereum/go-ethereum/common"
import heimdalltypes "github.com/maticnetwork/heimdall/types"

import mock "github.com/stretchr/testify/mock"
import types "github.com/ethereum/go-ethereum/core/types"

// IContractCaller is an autogenerated mock type for the IContractCaller type
type IContractCaller struct {
	mock.Mock
}

// CurrentChildBlock provides a mock function with given fields:
func (_m *IContractCaller) CurrentChildBlock() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: address
func (_m *IContractCaller) GetBalance(address common.Address) (*big.Int, error) {
	ret := _m.Called(address)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockNumberFromTxHash provides a mock function with given fields: tx
func (_m *IContractCaller) GetBlockNumberFromTxHash(tx common.Hash) (big.Int, error) {
	ret := _m.Called(tx)

	var r0 big.Int
	if rf, ok := ret.Get(0).(func(common.Hash) big.Int); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(big.Int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderInfo provides a mock function with given fields: headerID
func (_m *IContractCaller) GetHeaderInfo(headerID uint64) (common.Hash, uint64, uint64, uint64, error) {
	ret := _m.Called(headerID)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(uint64) common.Hash); ok {
		r0 = rf(headerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(uint64) uint64); ok {
		r1 = rf(headerID)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func(uint64) uint64); ok {
		r2 = rf(headerID)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 uint64
	if rf, ok := ret.Get(3).(func(uint64) uint64); ok {
		r3 = rf(headerID)
	} else {
		r3 = ret.Get(3).(uint64)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(uint64) error); ok {
		r4 = rf(headerID)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// GetMainChainBlock provides a mock function with given fields: blockNum
func (_m *IContractCaller) GetMainChainBlock(blockNum *big.Int) (*types.Header, error) {
	ret := _m.Called(blockNum)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(blockNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticChainBlock provides a mock function with given fields: blockNum
func (_m *IContractCaller) GetMaticChainBlock(blockNum *big.Int) (*types.Header, error) {
	ret := _m.Called(blockNum)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(blockNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatorInfo provides a mock function with given fields: valID
func (_m *IContractCaller) GetValidatorInfo(valID heimdalltypes.ValidatorID) (heimdalltypes.Validator, error) {
	ret := _m.Called(valID)

	var r0 heimdalltypes.Validator
	if rf, ok := ret.Get(0).(func(heimdalltypes.ValidatorID) heimdalltypes.Validator); ok {
		r0 = rf(valID)
	} else {
		r0 = ret.Get(0).(heimdalltypes.Validator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(heimdalltypes.ValidatorID) error); ok {
		r1 = rf(valID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTxConfirmed provides a mock function with given fields: tx
func (_m *IContractCaller) IsTxConfirmed(tx common.Hash) bool {
	ret := _m.Called(tx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Hash) bool); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SendCheckpoint provides a mock function with given fields: voteSignBytes, sigs, txData
func (_m *IContractCaller) SendCheckpoint(voteSignBytes []byte, sigs []byte, txData []byte) {
	_m.Called(voteSignBytes, sigs, txData)
}

// SigUpdateEvent provides a mock function with given fields: tx
func (_m *IContractCaller) SigUpdateEvent(tx common.Hash) (uint64, common.Address, common.Address, error) {
	ret := _m.Called(tx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(common.Hash) uint64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 common.Address
	if rf, ok := ret.Get(1).(func(common.Hash) common.Address); ok {
		r1 = rf(tx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(common.Address)
		}
	}

	var r2 common.Address
	if rf, ok := ret.Get(2).(func(common.Hash) common.Address); ok {
		r2 = rf(tx)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(common.Address)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(common.Hash) error); ok {
		r3 = rf(tx)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}
