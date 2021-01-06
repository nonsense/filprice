// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PriceOracleABI is the input ABI used to generate the binding from.
const PriceOracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"anchorAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSwingMantissa\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestedState\",\"type\":\"bool\"}],\"name\":\"_setPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"assets\",\"type\":\"address[]\"},{\"name\":\"requestedPriceMantissas\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAnchorAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numBlocksPerPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"readers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"assetPrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"anchors\",\"outputs\":[{\"name\":\"period\",\"type\":\"uint256\"},{\"name\":\"priceMantissa\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poster\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPendingAnchorAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAnchorAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingAnchors\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSwing\",\"outputs\":[{\"name\":\"mantissa\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAnchorAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"newScaledPrice\",\"type\":\"uint256\"}],\"name\":\"_setPendingAnchor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_poster\",\"type\":\"address\"},{\"name\":\"addr0\",\"type\":\"address\"},{\"name\":\"reader0\",\"type\":\"address\"},{\"name\":\"addr1\",\"type\":\"address\"},{\"name\":\"reader1\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"OracleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"anchorAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldScaledPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newScaledPrice\",\"type\":\"uint256\"}],\"name\":\"NewPendingAnchor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"PricePosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"anchorPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cappedPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"CappedPricePosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newState\",\"type\":\"bool\"}],\"name\":\"SetPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldPendingAnchorAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newPendingAnchorAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAnchorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAnchorAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAnchorAdmin\",\"type\":\"address\"}],\"name\":\"NewAnchorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"}]"

// PriceOracle is an auto generated Go binding around an Ethereum contract.
type PriceOracle struct {
	PriceOracleCaller     // Read-only binding to the contract
	PriceOracleTransactor // Write-only binding to the contract
	PriceOracleFilterer   // Log filterer for contract events
}

// PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleSession struct {
	Contract     *PriceOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleCallerSession struct {
	Contract *PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleTransactorSession struct {
	Contract     *PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleRaw struct {
	Contract *PriceOracle // Generic contract binding to access the raw methods on
}

// PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleCallerRaw struct {
	Contract *PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleTransactorRaw struct {
	Contract *PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracle creates a new instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracle(address common.Address, backend bind.ContractBackend) (*PriceOracle, error) {
	contract, err := bindPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracle{PriceOracleCaller: PriceOracleCaller{contract: contract}, PriceOracleTransactor: PriceOracleTransactor{contract: contract}, PriceOracleFilterer: PriceOracleFilterer{contract: contract}}, nil
}

// NewPriceOracleCaller creates a new read-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleCaller, error) {
	contract, err := bindPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleCaller{contract: contract}, nil
}

// NewPriceOracleTransactor creates a new write-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleTransactor, error) {
	contract, err := bindPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleTransactor{contract: contract}, nil
}

// NewPriceOracleFilterer creates a new log filterer instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleFilterer, error) {
	contract, err := bindPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleFilterer{contract: contract}, nil
}

// bindPriceOracle binds a generic wrapper to an already deployed contract.
func bindPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// AnchorAdmin is a free data retrieval call binding the contract method 0x08f31857.
//
// Solidity: function anchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleCaller) AnchorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "anchorAdmin")
	return *ret0, err
}

// AnchorAdmin is a free data retrieval call binding the contract method 0x08f31857.
//
// Solidity: function anchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleSession) AnchorAdmin() (common.Address, error) {
	return _PriceOracle.Contract.AnchorAdmin(&_PriceOracle.CallOpts)
}

// AnchorAdmin is a free data retrieval call binding the contract method 0x08f31857.
//
// Solidity: function anchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleCallerSession) AnchorAdmin() (common.Address, error) {
	return _PriceOracle.Contract.AnchorAdmin(&_PriceOracle.CallOpts)
}

// Anchors is a free data retrieval call binding the contract method 0x692374e3.
//
// Solidity: function anchors(address ) view returns(uint256 period, uint256 priceMantissa)
func (_PriceOracle *PriceOracleCaller) Anchors(opts *bind.CallOpts, arg0 common.Address) (struct {
	Period        *big.Int
	PriceMantissa *big.Int
}, error) {
	ret := new(struct {
		Period        *big.Int
		PriceMantissa *big.Int
	})
	out := ret
	err := _PriceOracle.contract.Call(opts, out, "anchors", arg0)
	return *ret, err
}

// Anchors is a free data retrieval call binding the contract method 0x692374e3.
//
// Solidity: function anchors(address ) view returns(uint256 period, uint256 priceMantissa)
func (_PriceOracle *PriceOracleSession) Anchors(arg0 common.Address) (struct {
	Period        *big.Int
	PriceMantissa *big.Int
}, error) {
	return _PriceOracle.Contract.Anchors(&_PriceOracle.CallOpts, arg0)
}

// Anchors is a free data retrieval call binding the contract method 0x692374e3.
//
// Solidity: function anchors(address ) view returns(uint256 period, uint256 priceMantissa)
func (_PriceOracle *PriceOracleCallerSession) Anchors(arg0 common.Address) (struct {
	Period        *big.Int
	PriceMantissa *big.Int
}, error) {
	return _PriceOracle.Contract.Anchors(&_PriceOracle.CallOpts, arg0)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) AssetPrices(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "assetPrices", asset)
	return *ret0, err
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.AssetPrices(&_PriceOracle.CallOpts, asset)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.AssetPrices(&_PriceOracle.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "getPrice", asset)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetPrice(&_PriceOracle.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetPrice(&_PriceOracle.CallOpts, asset)
}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256 mantissa)
func (_PriceOracle *PriceOracleCaller) MaxSwing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "maxSwing")
	return *ret0, err
}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256 mantissa)
func (_PriceOracle *PriceOracleSession) MaxSwing() (*big.Int, error) {
	return _PriceOracle.Contract.MaxSwing(&_PriceOracle.CallOpts)
}

// MaxSwing is a free data retrieval call binding the contract method 0xc5faf1d5.
//
// Solidity: function maxSwing() view returns(uint256 mantissa)
func (_PriceOracle *PriceOracleCallerSession) MaxSwing() (*big.Int, error) {
	return _PriceOracle.Contract.MaxSwing(&_PriceOracle.CallOpts)
}

// MaxSwingMantissa is a free data retrieval call binding the contract method 0x0c9c6301.
//
// Solidity: function maxSwingMantissa() view returns(uint256)
func (_PriceOracle *PriceOracleCaller) MaxSwingMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "maxSwingMantissa")
	return *ret0, err
}

// MaxSwingMantissa is a free data retrieval call binding the contract method 0x0c9c6301.
//
// Solidity: function maxSwingMantissa() view returns(uint256)
func (_PriceOracle *PriceOracleSession) MaxSwingMantissa() (*big.Int, error) {
	return _PriceOracle.Contract.MaxSwingMantissa(&_PriceOracle.CallOpts)
}

// MaxSwingMantissa is a free data retrieval call binding the contract method 0x0c9c6301.
//
// Solidity: function maxSwingMantissa() view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) MaxSwingMantissa() (*big.Int, error) {
	return _PriceOracle.Contract.MaxSwingMantissa(&_PriceOracle.CallOpts)
}

// NumBlocksPerPeriod is a free data retrieval call binding the contract method 0x485feabe.
//
// Solidity: function numBlocksPerPeriod() view returns(uint256)
func (_PriceOracle *PriceOracleCaller) NumBlocksPerPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "numBlocksPerPeriod")
	return *ret0, err
}

// NumBlocksPerPeriod is a free data retrieval call binding the contract method 0x485feabe.
//
// Solidity: function numBlocksPerPeriod() view returns(uint256)
func (_PriceOracle *PriceOracleSession) NumBlocksPerPeriod() (*big.Int, error) {
	return _PriceOracle.Contract.NumBlocksPerPeriod(&_PriceOracle.CallOpts)
}

// NumBlocksPerPeriod is a free data retrieval call binding the contract method 0x485feabe.
//
// Solidity: function numBlocksPerPeriod() view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) NumBlocksPerPeriod() (*big.Int, error) {
	return _PriceOracle.Contract.NumBlocksPerPeriod(&_PriceOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleSession) Paused() (bool, error) {
	return _PriceOracle.Contract.Paused(&_PriceOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOracle *PriceOracleCallerSession) Paused() (bool, error) {
	return _PriceOracle.Contract.Paused(&_PriceOracle.CallOpts)
}

// PendingAnchorAdmin is a free data retrieval call binding the contract method 0x451b1e3a.
//
// Solidity: function pendingAnchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleCaller) PendingAnchorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "pendingAnchorAdmin")
	return *ret0, err
}

// PendingAnchorAdmin is a free data retrieval call binding the contract method 0x451b1e3a.
//
// Solidity: function pendingAnchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleSession) PendingAnchorAdmin() (common.Address, error) {
	return _PriceOracle.Contract.PendingAnchorAdmin(&_PriceOracle.CallOpts)
}

// PendingAnchorAdmin is a free data retrieval call binding the contract method 0x451b1e3a.
//
// Solidity: function pendingAnchorAdmin() view returns(address)
func (_PriceOracle *PriceOracleCallerSession) PendingAnchorAdmin() (common.Address, error) {
	return _PriceOracle.Contract.PendingAnchorAdmin(&_PriceOracle.CallOpts)
}

// PendingAnchors is a free data retrieval call binding the contract method 0x9e8c4d95.
//
// Solidity: function pendingAnchors(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) PendingAnchors(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "pendingAnchors", arg0)
	return *ret0, err
}

// PendingAnchors is a free data retrieval call binding the contract method 0x9e8c4d95.
//
// Solidity: function pendingAnchors(address ) view returns(uint256)
func (_PriceOracle *PriceOracleSession) PendingAnchors(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.PendingAnchors(&_PriceOracle.CallOpts, arg0)
}

// PendingAnchors is a free data retrieval call binding the contract method 0x9e8c4d95.
//
// Solidity: function pendingAnchors(address ) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) PendingAnchors(arg0 common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.PendingAnchors(&_PriceOracle.CallOpts, arg0)
}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_PriceOracle *PriceOracleCaller) Poster(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "poster")
	return *ret0, err
}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_PriceOracle *PriceOracleSession) Poster() (common.Address, error) {
	return _PriceOracle.Contract.Poster(&_PriceOracle.CallOpts)
}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_PriceOracle *PriceOracleCallerSession) Poster() (common.Address, error) {
	return _PriceOracle.Contract.Poster(&_PriceOracle.CallOpts)
}

// Readers is a free data retrieval call binding the contract method 0x51e59ffb.
//
// Solidity: function readers(address ) view returns(address)
func (_PriceOracle *PriceOracleCaller) Readers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceOracle.contract.Call(opts, out, "readers", arg0)
	return *ret0, err
}

// Readers is a free data retrieval call binding the contract method 0x51e59ffb.
//
// Solidity: function readers(address ) view returns(address)
func (_PriceOracle *PriceOracleSession) Readers(arg0 common.Address) (common.Address, error) {
	return _PriceOracle.Contract.Readers(&_PriceOracle.CallOpts, arg0)
}

// Readers is a free data retrieval call binding the contract method 0x51e59ffb.
//
// Solidity: function readers(address ) view returns(address)
func (_PriceOracle *PriceOracleCallerSession) Readers(arg0 common.Address) (common.Address, error) {
	return _PriceOracle.Contract.Readers(&_PriceOracle.CallOpts, arg0)
}

// AcceptAnchorAdmin is a paid mutator transaction binding the contract method 0xccb13cbd.
//
// Solidity: function _acceptAnchorAdmin() returns(uint256)
func (_PriceOracle *PriceOracleTransactor) AcceptAnchorAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "_acceptAnchorAdmin")
}

// AcceptAnchorAdmin is a paid mutator transaction binding the contract method 0xccb13cbd.
//
// Solidity: function _acceptAnchorAdmin() returns(uint256)
func (_PriceOracle *PriceOracleSession) AcceptAnchorAdmin() (*types.Transaction, error) {
	return _PriceOracle.Contract.AcceptAnchorAdmin(&_PriceOracle.TransactOpts)
}

// AcceptAnchorAdmin is a paid mutator transaction binding the contract method 0xccb13cbd.
//
// Solidity: function _acceptAnchorAdmin() returns(uint256)
func (_PriceOracle *PriceOracleTransactorSession) AcceptAnchorAdmin() (*types.Transaction, error) {
	return _PriceOracle.Contract.AcceptAnchorAdmin(&_PriceOracle.TransactOpts)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool requestedState) returns(uint256)
func (_PriceOracle *PriceOracleTransactor) SetPaused(opts *bind.TransactOpts, requestedState bool) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "_setPaused", requestedState)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool requestedState) returns(uint256)
func (_PriceOracle *PriceOracleSession) SetPaused(requestedState bool) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPaused(&_PriceOracle.TransactOpts, requestedState)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool requestedState) returns(uint256)
func (_PriceOracle *PriceOracleTransactorSession) SetPaused(requestedState bool) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPaused(&_PriceOracle.TransactOpts, requestedState)
}

// SetPendingAnchor is a paid mutator transaction binding the contract method 0xde9d0e85.
//
// Solidity: function _setPendingAnchor(address asset, uint256 newScaledPrice) returns(uint256)
func (_PriceOracle *PriceOracleTransactor) SetPendingAnchor(opts *bind.TransactOpts, asset common.Address, newScaledPrice *big.Int) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "_setPendingAnchor", asset, newScaledPrice)
}

// SetPendingAnchor is a paid mutator transaction binding the contract method 0xde9d0e85.
//
// Solidity: function _setPendingAnchor(address asset, uint256 newScaledPrice) returns(uint256)
func (_PriceOracle *PriceOracleSession) SetPendingAnchor(asset common.Address, newScaledPrice *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPendingAnchor(&_PriceOracle.TransactOpts, asset, newScaledPrice)
}

// SetPendingAnchor is a paid mutator transaction binding the contract method 0xde9d0e85.
//
// Solidity: function _setPendingAnchor(address asset, uint256 newScaledPrice) returns(uint256)
func (_PriceOracle *PriceOracleTransactorSession) SetPendingAnchor(asset common.Address, newScaledPrice *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPendingAnchor(&_PriceOracle.TransactOpts, asset, newScaledPrice)
}

// SetPendingAnchorAdmin is a paid mutator transaction binding the contract method 0x9964622c.
//
// Solidity: function _setPendingAnchorAdmin(address newPendingAnchorAdmin) returns(uint256)
func (_PriceOracle *PriceOracleTransactor) SetPendingAnchorAdmin(opts *bind.TransactOpts, newPendingAnchorAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "_setPendingAnchorAdmin", newPendingAnchorAdmin)
}

// SetPendingAnchorAdmin is a paid mutator transaction binding the contract method 0x9964622c.
//
// Solidity: function _setPendingAnchorAdmin(address newPendingAnchorAdmin) returns(uint256)
func (_PriceOracle *PriceOracleSession) SetPendingAnchorAdmin(newPendingAnchorAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPendingAnchorAdmin(&_PriceOracle.TransactOpts, newPendingAnchorAdmin)
}

// SetPendingAnchorAdmin is a paid mutator transaction binding the contract method 0x9964622c.
//
// Solidity: function _setPendingAnchorAdmin(address newPendingAnchorAdmin) returns(uint256)
func (_PriceOracle *PriceOracleTransactorSession) SetPendingAnchorAdmin(newPendingAnchorAdmin common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPendingAnchorAdmin(&_PriceOracle.TransactOpts, newPendingAnchorAdmin)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 requestedPriceMantissa) returns(uint256)
func (_PriceOracle *PriceOracleTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, requestedPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setPrice", asset, requestedPriceMantissa)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 requestedPriceMantissa) returns(uint256)
func (_PriceOracle *PriceOracleSession) SetPrice(asset common.Address, requestedPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPrice(&_PriceOracle.TransactOpts, asset, requestedPriceMantissa)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 requestedPriceMantissa) returns(uint256)
func (_PriceOracle *PriceOracleTransactorSession) SetPrice(asset common.Address, requestedPriceMantissa *big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPrice(&_PriceOracle.TransactOpts, asset, requestedPriceMantissa)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] assets, uint256[] requestedPriceMantissas) returns(uint256[])
func (_PriceOracle *PriceOracleTransactor) SetPrices(opts *bind.TransactOpts, assets []common.Address, requestedPriceMantissas []*big.Int) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setPrices", assets, requestedPriceMantissas)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] assets, uint256[] requestedPriceMantissas) returns(uint256[])
func (_PriceOracle *PriceOracleSession) SetPrices(assets []common.Address, requestedPriceMantissas []*big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPrices(&_PriceOracle.TransactOpts, assets, requestedPriceMantissas)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] assets, uint256[] requestedPriceMantissas) returns(uint256[])
func (_PriceOracle *PriceOracleTransactorSession) SetPrices(assets []common.Address, requestedPriceMantissas []*big.Int) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetPrices(&_PriceOracle.TransactOpts, assets, requestedPriceMantissas)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
//func (_PriceOracle *PriceOracleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
//return _PriceOracle.contract.RawTransact(opts, calldata)
//}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
//func (_PriceOracle *PriceOracleSession) Fallback(calldata []byte) (*types.Transaction, error) {
//return _PriceOracle.Contract.Fallback(&_PriceOracle.TransactOpts, calldata)
//}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
//func (_PriceOracle *PriceOracleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
//return _PriceOracle.Contract.Fallback(&_PriceOracle.TransactOpts, calldata)
//}

// PriceOracleCappedPricePostedIterator is returned from FilterCappedPricePosted and is used to iterate over the raw logs and unpacked data for CappedPricePosted events raised by the PriceOracle contract.
type PriceOracleCappedPricePostedIterator struct {
	Event *PriceOracleCappedPricePosted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleCappedPricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleCappedPricePosted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleCappedPricePosted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleCappedPricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleCappedPricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleCappedPricePosted represents a CappedPricePosted event raised by the PriceOracle contract.
type PriceOracleCappedPricePosted struct {
	Asset                  common.Address
	RequestedPriceMantissa *big.Int
	AnchorPriceMantissa    *big.Int
	CappedPriceMantissa    *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCappedPricePosted is a free log retrieval operation binding the contract event 0x7221f7a2708437039cc63319145b6b873a40594b9782a3bee45b975e2f3b0f68.
//
// Solidity: event CappedPricePosted(address asset, uint256 requestedPriceMantissa, uint256 anchorPriceMantissa, uint256 cappedPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) FilterCappedPricePosted(opts *bind.FilterOpts) (*PriceOracleCappedPricePostedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "CappedPricePosted")
	if err != nil {
		return nil, err
	}
	return &PriceOracleCappedPricePostedIterator{contract: _PriceOracle.contract, event: "CappedPricePosted", logs: logs, sub: sub}, nil
}

// WatchCappedPricePosted is a free log subscription operation binding the contract event 0x7221f7a2708437039cc63319145b6b873a40594b9782a3bee45b975e2f3b0f68.
//
// Solidity: event CappedPricePosted(address asset, uint256 requestedPriceMantissa, uint256 anchorPriceMantissa, uint256 cappedPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) WatchCappedPricePosted(opts *bind.WatchOpts, sink chan<- *PriceOracleCappedPricePosted) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "CappedPricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleCappedPricePosted)
				if err := _PriceOracle.contract.UnpackLog(event, "CappedPricePosted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCappedPricePosted is a log parse operation binding the contract event 0x7221f7a2708437039cc63319145b6b873a40594b9782a3bee45b975e2f3b0f68.
//
// Solidity: event CappedPricePosted(address asset, uint256 requestedPriceMantissa, uint256 anchorPriceMantissa, uint256 cappedPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) ParseCappedPricePosted(log types.Log) (*PriceOracleCappedPricePosted, error) {
	event := new(PriceOracleCappedPricePosted)
	if err := _PriceOracle.contract.UnpackLog(event, "CappedPricePosted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the PriceOracle contract.
type PriceOracleFailureIterator struct {
	Event *PriceOracleFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleFailure represents a Failure event raised by the PriceOracle contract.
type PriceOracleFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) FilterFailure(opts *bind.FilterOpts) (*PriceOracleFailureIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &PriceOracleFailureIterator{contract: _PriceOracle.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *PriceOracleFailure) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleFailure)
				if err := _PriceOracle.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) ParseFailure(log types.Log) (*PriceOracleFailure, error) {
	event := new(PriceOracleFailure)
	if err := _PriceOracle.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleNewAnchorAdminIterator is returned from FilterNewAnchorAdmin and is used to iterate over the raw logs and unpacked data for NewAnchorAdmin events raised by the PriceOracle contract.
type PriceOracleNewAnchorAdminIterator struct {
	Event *PriceOracleNewAnchorAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleNewAnchorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleNewAnchorAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleNewAnchorAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleNewAnchorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleNewAnchorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleNewAnchorAdmin represents a NewAnchorAdmin event raised by the PriceOracle contract.
type PriceOracleNewAnchorAdmin struct {
	OldAnchorAdmin common.Address
	NewAnchorAdmin common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewAnchorAdmin is a free log retrieval operation binding the contract event 0xbef9248fe57ae972dd47833f68c43f0b3b2d14217612dfbc804a520a23730d46.
//
// Solidity: event NewAnchorAdmin(address oldAnchorAdmin, address newAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) FilterNewAnchorAdmin(opts *bind.FilterOpts) (*PriceOracleNewAnchorAdminIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "NewAnchorAdmin")
	if err != nil {
		return nil, err
	}
	return &PriceOracleNewAnchorAdminIterator{contract: _PriceOracle.contract, event: "NewAnchorAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAnchorAdmin is a free log subscription operation binding the contract event 0xbef9248fe57ae972dd47833f68c43f0b3b2d14217612dfbc804a520a23730d46.
//
// Solidity: event NewAnchorAdmin(address oldAnchorAdmin, address newAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) WatchNewAnchorAdmin(opts *bind.WatchOpts, sink chan<- *PriceOracleNewAnchorAdmin) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "NewAnchorAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleNewAnchorAdmin)
				if err := _PriceOracle.contract.UnpackLog(event, "NewAnchorAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAnchorAdmin is a log parse operation binding the contract event 0xbef9248fe57ae972dd47833f68c43f0b3b2d14217612dfbc804a520a23730d46.
//
// Solidity: event NewAnchorAdmin(address oldAnchorAdmin, address newAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) ParseNewAnchorAdmin(log types.Log) (*PriceOracleNewAnchorAdmin, error) {
	event := new(PriceOracleNewAnchorAdmin)
	if err := _PriceOracle.contract.UnpackLog(event, "NewAnchorAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleNewPendingAnchorIterator is returned from FilterNewPendingAnchor and is used to iterate over the raw logs and unpacked data for NewPendingAnchor events raised by the PriceOracle contract.
type PriceOracleNewPendingAnchorIterator struct {
	Event *PriceOracleNewPendingAnchor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleNewPendingAnchorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleNewPendingAnchor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleNewPendingAnchor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleNewPendingAnchorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleNewPendingAnchorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleNewPendingAnchor represents a NewPendingAnchor event raised by the PriceOracle contract.
type PriceOracleNewPendingAnchor struct {
	AnchorAdmin    common.Address
	Asset          common.Address
	OldScaledPrice *big.Int
	NewScaledPrice *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAnchor is a free log retrieval operation binding the contract event 0xf0aa20c29c1f8e751bfe0a78bc49a520ed14f2dab274087d90d1341d8b76af5c.
//
// Solidity: event NewPendingAnchor(address anchorAdmin, address asset, uint256 oldScaledPrice, uint256 newScaledPrice)
func (_PriceOracle *PriceOracleFilterer) FilterNewPendingAnchor(opts *bind.FilterOpts) (*PriceOracleNewPendingAnchorIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "NewPendingAnchor")
	if err != nil {
		return nil, err
	}
	return &PriceOracleNewPendingAnchorIterator{contract: _PriceOracle.contract, event: "NewPendingAnchor", logs: logs, sub: sub}, nil
}

// WatchNewPendingAnchor is a free log subscription operation binding the contract event 0xf0aa20c29c1f8e751bfe0a78bc49a520ed14f2dab274087d90d1341d8b76af5c.
//
// Solidity: event NewPendingAnchor(address anchorAdmin, address asset, uint256 oldScaledPrice, uint256 newScaledPrice)
func (_PriceOracle *PriceOracleFilterer) WatchNewPendingAnchor(opts *bind.WatchOpts, sink chan<- *PriceOracleNewPendingAnchor) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "NewPendingAnchor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleNewPendingAnchor)
				if err := _PriceOracle.contract.UnpackLog(event, "NewPendingAnchor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPendingAnchor is a log parse operation binding the contract event 0xf0aa20c29c1f8e751bfe0a78bc49a520ed14f2dab274087d90d1341d8b76af5c.
//
// Solidity: event NewPendingAnchor(address anchorAdmin, address asset, uint256 oldScaledPrice, uint256 newScaledPrice)
func (_PriceOracle *PriceOracleFilterer) ParseNewPendingAnchor(log types.Log) (*PriceOracleNewPendingAnchor, error) {
	event := new(PriceOracleNewPendingAnchor)
	if err := _PriceOracle.contract.UnpackLog(event, "NewPendingAnchor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleNewPendingAnchorAdminIterator is returned from FilterNewPendingAnchorAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAnchorAdmin events raised by the PriceOracle contract.
type PriceOracleNewPendingAnchorAdminIterator struct {
	Event *PriceOracleNewPendingAnchorAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleNewPendingAnchorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleNewPendingAnchorAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleNewPendingAnchorAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleNewPendingAnchorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleNewPendingAnchorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleNewPendingAnchorAdmin represents a NewPendingAnchorAdmin event raised by the PriceOracle contract.
type PriceOracleNewPendingAnchorAdmin struct {
	OldPendingAnchorAdmin common.Address
	NewPendingAnchorAdmin common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAnchorAdmin is a free log retrieval operation binding the contract event 0x6c773973d5bcf264b509f4194ceb99e891251f6aabb325523a863c282958b13e.
//
// Solidity: event NewPendingAnchorAdmin(address oldPendingAnchorAdmin, address newPendingAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) FilterNewPendingAnchorAdmin(opts *bind.FilterOpts) (*PriceOracleNewPendingAnchorAdminIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "NewPendingAnchorAdmin")
	if err != nil {
		return nil, err
	}
	return &PriceOracleNewPendingAnchorAdminIterator{contract: _PriceOracle.contract, event: "NewPendingAnchorAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAnchorAdmin is a free log subscription operation binding the contract event 0x6c773973d5bcf264b509f4194ceb99e891251f6aabb325523a863c282958b13e.
//
// Solidity: event NewPendingAnchorAdmin(address oldPendingAnchorAdmin, address newPendingAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) WatchNewPendingAnchorAdmin(opts *bind.WatchOpts, sink chan<- *PriceOracleNewPendingAnchorAdmin) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "NewPendingAnchorAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleNewPendingAnchorAdmin)
				if err := _PriceOracle.contract.UnpackLog(event, "NewPendingAnchorAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPendingAnchorAdmin is a log parse operation binding the contract event 0x6c773973d5bcf264b509f4194ceb99e891251f6aabb325523a863c282958b13e.
//
// Solidity: event NewPendingAnchorAdmin(address oldPendingAnchorAdmin, address newPendingAnchorAdmin)
func (_PriceOracle *PriceOracleFilterer) ParseNewPendingAnchorAdmin(log types.Log) (*PriceOracleNewPendingAnchorAdmin, error) {
	event := new(PriceOracleNewPendingAnchorAdmin)
	if err := _PriceOracle.contract.UnpackLog(event, "NewPendingAnchorAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleOracleFailureIterator is returned from FilterOracleFailure and is used to iterate over the raw logs and unpacked data for OracleFailure events raised by the PriceOracle contract.
type PriceOracleOracleFailureIterator struct {
	Event *PriceOracleOracleFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleOracleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleOracleFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleOracleFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleOracleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleOracleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleOracleFailure represents a OracleFailure event raised by the PriceOracle contract.
type PriceOracleOracleFailure struct {
	MsgSender common.Address
	Asset     common.Address
	Error     *big.Int
	Info      *big.Int
	Detail    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleFailure is a free log retrieval operation binding the contract event 0x96f29b65cebbd6816352fb242b6af7180b49e8a09e19e589225d35bc8444f0b7.
//
// Solidity: event OracleFailure(address msgSender, address asset, uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) FilterOracleFailure(opts *bind.FilterOpts) (*PriceOracleOracleFailureIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "OracleFailure")
	if err != nil {
		return nil, err
	}
	return &PriceOracleOracleFailureIterator{contract: _PriceOracle.contract, event: "OracleFailure", logs: logs, sub: sub}, nil
}

// WatchOracleFailure is a free log subscription operation binding the contract event 0x96f29b65cebbd6816352fb242b6af7180b49e8a09e19e589225d35bc8444f0b7.
//
// Solidity: event OracleFailure(address msgSender, address asset, uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) WatchOracleFailure(opts *bind.WatchOpts, sink chan<- *PriceOracleOracleFailure) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "OracleFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleOracleFailure)
				if err := _PriceOracle.contract.UnpackLog(event, "OracleFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleFailure is a log parse operation binding the contract event 0x96f29b65cebbd6816352fb242b6af7180b49e8a09e19e589225d35bc8444f0b7.
//
// Solidity: event OracleFailure(address msgSender, address asset, uint256 error, uint256 info, uint256 detail)
func (_PriceOracle *PriceOracleFilterer) ParseOracleFailure(log types.Log) (*PriceOracleOracleFailure, error) {
	event := new(PriceOracleOracleFailure)
	if err := _PriceOracle.contract.UnpackLog(event, "OracleFailure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOraclePricePostedIterator is returned from FilterPricePosted and is used to iterate over the raw logs and unpacked data for PricePosted events raised by the PriceOracle contract.
type PriceOraclePricePostedIterator struct {
	Event *PriceOraclePricePosted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOraclePricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclePricePosted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOraclePricePosted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOraclePricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclePricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclePricePosted represents a PricePosted event raised by the PriceOracle contract.
type PriceOraclePricePosted struct {
	Asset                  common.Address
	PreviousPriceMantissa  *big.Int
	RequestedPriceMantissa *big.Int
	NewPriceMantissa       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPricePosted is a free log retrieval operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) FilterPricePosted(opts *bind.FilterOpts) (*PriceOraclePricePostedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return &PriceOraclePricePostedIterator{contract: _PriceOracle.contract, event: "PricePosted", logs: logs, sub: sub}, nil
}

// WatchPricePosted is a free log subscription operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) WatchPricePosted(opts *bind.WatchOpts, sink chan<- *PriceOraclePricePosted) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclePricePosted)
				if err := _PriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePricePosted is a log parse operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_PriceOracle *PriceOracleFilterer) ParsePricePosted(log types.Log) (*PriceOraclePricePosted, error) {
	event := new(PriceOraclePricePosted)
	if err := _PriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceOracleSetPausedIterator is returned from FilterSetPaused and is used to iterate over the raw logs and unpacked data for SetPaused events raised by the PriceOracle contract.
type PriceOracleSetPausedIterator struct {
	Event *PriceOracleSetPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceOracleSetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleSetPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceOracleSetPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceOracleSetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleSetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleSetPaused represents a SetPaused event raised by the PriceOracle contract.
type PriceOracleSetPaused struct {
	NewState bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPaused is a free log retrieval operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_PriceOracle *PriceOracleFilterer) FilterSetPaused(opts *bind.FilterOpts) (*PriceOracleSetPausedIterator, error) {

	logs, sub, err := _PriceOracle.contract.FilterLogs(opts, "SetPaused")
	if err != nil {
		return nil, err
	}
	return &PriceOracleSetPausedIterator{contract: _PriceOracle.contract, event: "SetPaused", logs: logs, sub: sub}, nil
}

// WatchSetPaused is a free log subscription operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_PriceOracle *PriceOracleFilterer) WatchSetPaused(opts *bind.WatchOpts, sink chan<- *PriceOracleSetPaused) (event.Subscription, error) {

	logs, sub, err := _PriceOracle.contract.WatchLogs(opts, "SetPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleSetPaused)
				if err := _PriceOracle.contract.UnpackLog(event, "SetPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPaused is a log parse operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_PriceOracle *PriceOracleFilterer) ParseSetPaused(log types.Log) (*PriceOracleSetPaused, error) {
	event := new(PriceOracleSetPaused)
	if err := _PriceOracle.contract.UnpackLog(event, "SetPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
