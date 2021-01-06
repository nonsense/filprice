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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Struct10 is an auto generated low-level Go binding around an user-defined struct.
type Struct10 struct {
	MarginRatio       Struct0
	LiquidationSpread Struct0
	EarningsRate      Struct0
	MinBorrowedValue  Struct0
}

// Struct11 is an auto generated low-level Go binding around an user-defined struct.
type Struct11 struct {
	Token          common.Address
	TotalPar       Struct7
	Index          Struct2
	PriceOracle    common.Address
	InterestSetter common.Address
	MarginPremium  Struct0
	SpreadPremium  Struct0
	IsClosing      bool
}

// Struct9 is an auto generated low-level Go binding around an user-defined struct.
type Struct9 struct {
	Operator common.Address
	Trusted  bool
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Owner  common.Address
	Number *big.Int
}

// Struct5 is an auto generated low-level Go binding around an user-defined struct.
type Struct5 struct {
	Sign  bool
	Value *big.Int
}

// Struct8 is an auto generated low-level Go binding around an user-defined struct.
type Struct8 struct {
	Sign  bool
	Value *big.Int
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	Sign         bool
	Denomination uint8
	Ref          uint8
	Value        *big.Int
}

// Struct7 is an auto generated low-level Go binding around an user-defined struct.
type Struct7 struct {
	Borrow *big.Int
	Supply *big.Int
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Value *big.Int
}

// Struct6 is an auto generated low-level Go binding around an user-defined struct.
type Struct6 struct {
	MarginRatioMax       uint64
	LiquidationSpreadMax uint64
	EarningsRateMax      uint64
	MarginPremiumMax     uint64
	SpreadPremiumMax     uint64
	MinBorrowedValueMax  *big.Int
}

// Struct4 is an auto generated low-level Go binding around an user-defined struct.
type Struct4 struct {
	ActionType        uint8
	AccountId         *big.Int
	Amount            Struct3
	PrimaryMarketId   *big.Int
	SecondaryMarketId *big.Int
	OtherAddress      common.Address
	OtherAccountId    *big.Int
	Data              []byte
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Borrow     *big.Int
	Supply     *big.Int
	LastUpdate uint32
}

// SoloMarginABI is the input ABI used to generate the binding from.
const SoloMarginABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"spreadPremium\",\"type\":\"tuple\"}],\"name\":\"ownerSetSpreadPremium\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getIsGlobalOperator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"name\":\"interestSetter\",\"type\":\"address\"}],\"name\":\"ownerSetInterestSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getAccountValues\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketPriceOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketInterestSetter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketSpreadPremium\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumMarkets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ownerWithdrawUnsupportedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"minBorrowedValue\",\"type\":\"tuple\"}],\"name\":\"ownerSetMinBorrowedValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"spread\",\"type\":\"tuple\"}],\"name\":\"ownerSetLiquidationSpread\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"earningsRate\",\"type\":\"tuple\"}],\"name\":\"ownerSetEarningsRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getIsLocalOperator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"},{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getAccountPar\",\"outputs\":[{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginPremium\",\"type\":\"tuple\"}],\"name\":\"ownerSetMarginPremium\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarginRatio\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketCurrentIndex\",\"outputs\":[{\"components\":[{\"name\":\"borrow\",\"type\":\"uint96\"},{\"name\":\"supply\",\"type\":\"uint96\"},{\"name\":\"lastUpdate\",\"type\":\"uint32\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketIsClosing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRiskParams\",\"outputs\":[{\"components\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginRatio\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"liquidationSpread\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"earningsRate\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"minBorrowedValue\",\"type\":\"tuple\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getAccountBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint128\"}],\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinBorrowedValue\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"trusted\",\"type\":\"bool\"}],\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"setOperators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketPrice\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ownerWithdrawExcessTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"priceOracle\",\"type\":\"address\"},{\"name\":\"interestSetter\",\"type\":\"address\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginPremium\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"spreadPremium\",\"type\":\"tuple\"}],\"name\":\"ownerAddMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"type\":\"tuple[]\"},{\"components\":[{\"name\":\"actionType\",\"type\":\"uint8\"},{\"name\":\"accountId\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"denomination\",\"type\":\"uint8\"},{\"name\":\"ref\",\"type\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"amount\",\"type\":\"tuple\"},{\"name\":\"primaryMarketId\",\"type\":\"uint256\"},{\"name\":\"secondaryMarketId\",\"type\":\"uint256\"},{\"name\":\"otherAddress\",\"type\":\"address\"},{\"name\":\"otherAccountId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"actions\",\"type\":\"tuple[]\"}],\"name\":\"operate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketWithInfo\",\"outputs\":[{\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"name\":\"borrow\",\"type\":\"uint128\"},{\"name\":\"supply\",\"type\":\"uint128\"}],\"name\":\"totalPar\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"borrow\",\"type\":\"uint96\"},{\"name\":\"supply\",\"type\":\"uint96\"},{\"name\":\"lastUpdate\",\"type\":\"uint32\"}],\"name\":\"index\",\"type\":\"tuple\"},{\"name\":\"priceOracle\",\"type\":\"address\"},{\"name\":\"interestSetter\",\"type\":\"address\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginPremium\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"spreadPremium\",\"type\":\"tuple\"},{\"name\":\"isClosing\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"borrow\",\"type\":\"uint96\"},{\"name\":\"supply\",\"type\":\"uint96\"},{\"name\":\"lastUpdate\",\"type\":\"uint32\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ratio\",\"type\":\"tuple\"}],\"name\":\"ownerSetMarginRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLiquidationSpread\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"},{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getAccountWei\",\"outputs\":[{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketTotalPar\",\"outputs\":[{\"components\":[{\"name\":\"borrow\",\"type\":\"uint128\"},{\"name\":\"supply\",\"type\":\"uint128\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"heldMarketId\",\"type\":\"uint256\"},{\"name\":\"owedMarketId\",\"type\":\"uint256\"}],\"name\":\"getLiquidationSpreadForPair\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getNumExcessTokens\",\"outputs\":[{\"components\":[{\"name\":\"sign\",\"type\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketCachedIndex\",\"outputs\":[{\"components\":[{\"name\":\"borrow\",\"type\":\"uint96\"},{\"name\":\"supply\",\"type\":\"uint96\"},{\"name\":\"lastUpdate\",\"type\":\"uint32\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getAccountStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEarningsRate\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"ownerSetPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRiskLimits\",\"outputs\":[{\"components\":[{\"name\":\"marginRatioMax\",\"type\":\"uint64\"},{\"name\":\"liquidationSpreadMax\",\"type\":\"uint64\"},{\"name\":\"earningsRateMax\",\"type\":\"uint64\"},{\"name\":\"marginPremiumMax\",\"type\":\"uint64\"},{\"name\":\"spreadPremiumMax\",\"type\":\"uint64\"},{\"name\":\"minBorrowedValueMax\",\"type\":\"uint128\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"name\":\"borrow\",\"type\":\"uint128\"},{\"name\":\"supply\",\"type\":\"uint128\"}],\"name\":\"totalPar\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"borrow\",\"type\":\"uint96\"},{\"name\":\"supply\",\"type\":\"uint96\"},{\"name\":\"lastUpdate\",\"type\":\"uint32\"}],\"name\":\"index\",\"type\":\"tuple\"},{\"name\":\"priceOracle\",\"type\":\"address\"},{\"name\":\"interestSetter\",\"type\":\"address\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginPremium\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"spreadPremium\",\"type\":\"tuple\"},{\"name\":\"isClosing\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"},{\"name\":\"isClosing\",\"type\":\"bool\"}],\"name\":\"ownerSetIsClosing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ownerSetGlobalOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getAdjustedAccountValues\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketMarginPremium\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketInterestRate\",\"outputs\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"marginRatio\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"liquidationSpread\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"earningsRate\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"minBorrowedValue\",\"type\":\"tuple\"}],\"name\":\"riskParams\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"marginRatioMax\",\"type\":\"uint64\"},{\"name\":\"liquidationSpreadMax\",\"type\":\"uint64\"},{\"name\":\"earningsRateMax\",\"type\":\"uint64\"},{\"name\":\"marginPremiumMax\",\"type\":\"uint64\"},{\"name\":\"spreadPremiumMax\",\"type\":\"uint64\"},{\"name\":\"minBorrowedValueMax\",\"type\":\"uint128\"}],\"name\":\"riskLimits\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"trusted\",\"type\":\"bool\"}],\"name\":\"LogOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SoloMargin is an auto generated Go binding around an Ethereum contract.
type SoloMargin struct {
	SoloMarginCaller     // Read-only binding to the contract
	SoloMarginTransactor // Write-only binding to the contract
	SoloMarginFilterer   // Log filterer for contract events
}

// SoloMarginCaller is an auto generated read-only Go binding around an Ethereum contract.
type SoloMarginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoloMarginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SoloMarginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoloMarginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SoloMarginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoloMarginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoloMarginSession struct {
	Contract     *SoloMargin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoloMarginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SoloMarginCallerSession struct {
	Contract *SoloMarginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SoloMarginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SoloMarginTransactorSession struct {
	Contract     *SoloMarginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SoloMarginRaw is an auto generated low-level Go binding around an Ethereum contract.
type SoloMarginRaw struct {
	Contract *SoloMargin // Generic contract binding to access the raw methods on
}

// SoloMarginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SoloMarginCallerRaw struct {
	Contract *SoloMarginCaller // Generic read-only contract binding to access the raw methods on
}

// SoloMarginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SoloMarginTransactorRaw struct {
	Contract *SoloMarginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSoloMargin creates a new instance of SoloMargin, bound to a specific deployed contract.
func NewSoloMargin(address common.Address, backend bind.ContractBackend) (*SoloMargin, error) {
	contract, err := bindSoloMargin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SoloMargin{SoloMarginCaller: SoloMarginCaller{contract: contract}, SoloMarginTransactor: SoloMarginTransactor{contract: contract}, SoloMarginFilterer: SoloMarginFilterer{contract: contract}}, nil
}

// NewSoloMarginCaller creates a new read-only instance of SoloMargin, bound to a specific deployed contract.
func NewSoloMarginCaller(address common.Address, caller bind.ContractCaller) (*SoloMarginCaller, error) {
	contract, err := bindSoloMargin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SoloMarginCaller{contract: contract}, nil
}

// NewSoloMarginTransactor creates a new write-only instance of SoloMargin, bound to a specific deployed contract.
func NewSoloMarginTransactor(address common.Address, transactor bind.ContractTransactor) (*SoloMarginTransactor, error) {
	contract, err := bindSoloMargin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SoloMarginTransactor{contract: contract}, nil
}

// NewSoloMarginFilterer creates a new log filterer instance of SoloMargin, bound to a specific deployed contract.
func NewSoloMarginFilterer(address common.Address, filterer bind.ContractFilterer) (*SoloMarginFilterer, error) {
	contract, err := bindSoloMargin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SoloMarginFilterer{contract: contract}, nil
}

// bindSoloMargin binds a generic wrapper to an already deployed contract.
func bindSoloMargin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SoloMarginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoloMargin *SoloMarginRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SoloMargin.Contract.SoloMarginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoloMargin *SoloMarginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoloMargin.Contract.SoloMarginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoloMargin *SoloMarginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoloMargin.Contract.SoloMarginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoloMargin *SoloMarginCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SoloMargin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoloMargin *SoloMarginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoloMargin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoloMargin *SoloMarginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoloMargin.Contract.contract.Transact(opts, method, params...)
}

// GetAccountBalances is a free data retrieval call binding the contract method 0x6a8194e7.
//
// Solidity: function getAccountBalances(Struct1 account) constant returns(address[], []Struct5, []Struct8)
func (_SoloMargin *SoloMarginCaller) GetAccountBalances(opts *bind.CallOpts, account Struct1) ([]common.Address, []Struct5, []Struct8, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]Struct5)
		ret2 = new([]Struct8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SoloMargin.contract.Call(opts, out, "getAccountBalances", account)
	return *ret0, *ret1, *ret2, err
}

// GetAccountBalances is a free data retrieval call binding the contract method 0x6a8194e7.
//
// Solidity: function getAccountBalances(Struct1 account) constant returns(address[], []Struct5, []Struct8)
func (_SoloMargin *SoloMarginSession) GetAccountBalances(account Struct1) ([]common.Address, []Struct5, []Struct8, error) {
	return _SoloMargin.Contract.GetAccountBalances(&_SoloMargin.CallOpts, account)
}

// GetAccountBalances is a free data retrieval call binding the contract method 0x6a8194e7.
//
// Solidity: function getAccountBalances(Struct1 account) constant returns(address[], []Struct5, []Struct8)
func (_SoloMargin *SoloMarginCallerSession) GetAccountBalances(account Struct1) ([]common.Address, []Struct5, []Struct8, error) {
	return _SoloMargin.Contract.GetAccountBalances(&_SoloMargin.CallOpts, account)
}

// GetAccountPar is a free data retrieval call binding the contract method 0x47d1b53c.
//
// Solidity: function getAccountPar(Struct1 account, uint256 marketId) constant returns(Struct5)
func (_SoloMargin *SoloMarginCaller) GetAccountPar(opts *bind.CallOpts, account Struct1, marketId *big.Int) (Struct5, error) {
	var (
		ret0 = new(Struct5)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getAccountPar", account, marketId)
	return *ret0, err
}

// GetAccountPar is a free data retrieval call binding the contract method 0x47d1b53c.
//
// Solidity: function getAccountPar(Struct1 account, uint256 marketId) constant returns(Struct5)
func (_SoloMargin *SoloMarginSession) GetAccountPar(account Struct1, marketId *big.Int) (Struct5, error) {
	return _SoloMargin.Contract.GetAccountPar(&_SoloMargin.CallOpts, account, marketId)
}

// GetAccountPar is a free data retrieval call binding the contract method 0x47d1b53c.
//
// Solidity: function getAccountPar(Struct1 account, uint256 marketId) constant returns(Struct5)
func (_SoloMargin *SoloMarginCallerSession) GetAccountPar(account Struct1, marketId *big.Int) (Struct5, error) {
	return _SoloMargin.Contract.GetAccountPar(&_SoloMargin.CallOpts, account, marketId)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xe51bfcb4.
//
// Solidity: function getAccountStatus(Struct1 account) constant returns(uint8)
func (_SoloMargin *SoloMarginCaller) GetAccountStatus(opts *bind.CallOpts, account Struct1) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getAccountStatus", account)
	return *ret0, err
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xe51bfcb4.
//
// Solidity: function getAccountStatus(Struct1 account) constant returns(uint8)
func (_SoloMargin *SoloMarginSession) GetAccountStatus(account Struct1) (uint8, error) {
	return _SoloMargin.Contract.GetAccountStatus(&_SoloMargin.CallOpts, account)
}

// GetAccountStatus is a free data retrieval call binding the contract method 0xe51bfcb4.
//
// Solidity: function getAccountStatus(Struct1 account) constant returns(uint8)
func (_SoloMargin *SoloMarginCallerSession) GetAccountStatus(account Struct1) (uint8, error) {
	return _SoloMargin.Contract.GetAccountStatus(&_SoloMargin.CallOpts, account)
}

// GetAccountValues is a free data retrieval call binding the contract method 0x124f914c.
//
// Solidity: function getAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginCaller) GetAccountValues(opts *bind.CallOpts, account Struct1) (Struct0, Struct0, error) {
	var (
		ret0 = new(Struct0)
		ret1 = new(Struct0)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SoloMargin.contract.Call(opts, out, "getAccountValues", account)
	return *ret0, *ret1, err
}

// GetAccountValues is a free data retrieval call binding the contract method 0x124f914c.
//
// Solidity: function getAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginSession) GetAccountValues(account Struct1) (Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetAccountValues(&_SoloMargin.CallOpts, account)
}

// GetAccountValues is a free data retrieval call binding the contract method 0x124f914c.
//
// Solidity: function getAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetAccountValues(account Struct1) (Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetAccountValues(&_SoloMargin.CallOpts, account)
}

// GetAccountWei is a free data retrieval call binding the contract method 0xc190c2ec.
//
// Solidity: function getAccountWei(Struct1 account, uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginCaller) GetAccountWei(opts *bind.CallOpts, account Struct1, marketId *big.Int) (Struct8, error) {
	var (
		ret0 = new(Struct8)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getAccountWei", account, marketId)
	return *ret0, err
}

// GetAccountWei is a free data retrieval call binding the contract method 0xc190c2ec.
//
// Solidity: function getAccountWei(Struct1 account, uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginSession) GetAccountWei(account Struct1, marketId *big.Int) (Struct8, error) {
	return _SoloMargin.Contract.GetAccountWei(&_SoloMargin.CallOpts, account, marketId)
}

// GetAccountWei is a free data retrieval call binding the contract method 0xc190c2ec.
//
// Solidity: function getAccountWei(Struct1 account, uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginCallerSession) GetAccountWei(account Struct1, marketId *big.Int) (Struct8, error) {
	return _SoloMargin.Contract.GetAccountWei(&_SoloMargin.CallOpts, account, marketId)
}

// GetAdjustedAccountValues is a free data retrieval call binding the contract method 0xf9416052.
//
// Solidity: function getAdjustedAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginCaller) GetAdjustedAccountValues(opts *bind.CallOpts, account Struct1) (Struct0, Struct0, error) {
	var (
		ret0 = new(Struct0)
		ret1 = new(Struct0)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SoloMargin.contract.Call(opts, out, "getAdjustedAccountValues", account)
	return *ret0, *ret1, err
}

// GetAdjustedAccountValues is a free data retrieval call binding the contract method 0xf9416052.
//
// Solidity: function getAdjustedAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginSession) GetAdjustedAccountValues(account Struct1) (Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetAdjustedAccountValues(&_SoloMargin.CallOpts, account)
}

// GetAdjustedAccountValues is a free data retrieval call binding the contract method 0xf9416052.
//
// Solidity: function getAdjustedAccountValues(Struct1 account) constant returns(Struct0, Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetAdjustedAccountValues(account Struct1) (Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetAdjustedAccountValues(&_SoloMargin.CallOpts, account)
}

// GetEarningsRate is a free data retrieval call binding the contract method 0xe5520228.
//
// Solidity: function getEarningsRate() constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetEarningsRate(opts *bind.CallOpts) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getEarningsRate")
	return *ret0, err
}

// GetEarningsRate is a free data retrieval call binding the contract method 0xe5520228.
//
// Solidity: function getEarningsRate() constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetEarningsRate() (Struct0, error) {
	return _SoloMargin.Contract.GetEarningsRate(&_SoloMargin.CallOpts)
}

// GetEarningsRate is a free data retrieval call binding the contract method 0xe5520228.
//
// Solidity: function getEarningsRate() constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetEarningsRate() (Struct0, error) {
	return _SoloMargin.Contract.GetEarningsRate(&_SoloMargin.CallOpts)
}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) constant returns(bool)
func (_SoloMargin *SoloMarginCaller) GetIsGlobalOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getIsGlobalOperator", operator)
	return *ret0, err
}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) constant returns(bool)
func (_SoloMargin *SoloMarginSession) GetIsGlobalOperator(operator common.Address) (bool, error) {
	return _SoloMargin.Contract.GetIsGlobalOperator(&_SoloMargin.CallOpts, operator)
}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) constant returns(bool)
func (_SoloMargin *SoloMarginCallerSession) GetIsGlobalOperator(operator common.Address) (bool, error) {
	return _SoloMargin.Contract.GetIsGlobalOperator(&_SoloMargin.CallOpts, operator)
}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address owner, address operator) constant returns(bool)
func (_SoloMargin *SoloMarginCaller) GetIsLocalOperator(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getIsLocalOperator", owner, operator)
	return *ret0, err
}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address owner, address operator) constant returns(bool)
func (_SoloMargin *SoloMarginSession) GetIsLocalOperator(owner common.Address, operator common.Address) (bool, error) {
	return _SoloMargin.Contract.GetIsLocalOperator(&_SoloMargin.CallOpts, owner, operator)
}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address owner, address operator) constant returns(bool)
func (_SoloMargin *SoloMarginCallerSession) GetIsLocalOperator(owner common.Address, operator common.Address) (bool, error) {
	return _SoloMargin.Contract.GetIsLocalOperator(&_SoloMargin.CallOpts, owner, operator)
}

// GetLiquidationSpread is a free data retrieval call binding the contract method 0xc1460942.
//
// Solidity: function getLiquidationSpread() constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetLiquidationSpread(opts *bind.CallOpts) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getLiquidationSpread")
	return *ret0, err
}

// GetLiquidationSpread is a free data retrieval call binding the contract method 0xc1460942.
//
// Solidity: function getLiquidationSpread() constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetLiquidationSpread() (Struct0, error) {
	return _SoloMargin.Contract.GetLiquidationSpread(&_SoloMargin.CallOpts)
}

// GetLiquidationSpread is a free data retrieval call binding the contract method 0xc1460942.
//
// Solidity: function getLiquidationSpread() constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetLiquidationSpread() (Struct0, error) {
	return _SoloMargin.Contract.GetLiquidationSpread(&_SoloMargin.CallOpts)
}

// GetLiquidationSpreadForPair is a free data retrieval call binding the contract method 0xd24c48bc.
//
// Solidity: function getLiquidationSpreadForPair(uint256 heldMarketId, uint256 owedMarketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetLiquidationSpreadForPair(opts *bind.CallOpts, heldMarketId *big.Int, owedMarketId *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getLiquidationSpreadForPair", heldMarketId, owedMarketId)
	return *ret0, err
}

// GetLiquidationSpreadForPair is a free data retrieval call binding the contract method 0xd24c48bc.
//
// Solidity: function getLiquidationSpreadForPair(uint256 heldMarketId, uint256 owedMarketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetLiquidationSpreadForPair(heldMarketId *big.Int, owedMarketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetLiquidationSpreadForPair(&_SoloMargin.CallOpts, heldMarketId, owedMarketId)
}

// GetLiquidationSpreadForPair is a free data retrieval call binding the contract method 0xd24c48bc.
//
// Solidity: function getLiquidationSpreadForPair(uint256 heldMarketId, uint256 owedMarketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetLiquidationSpreadForPair(heldMarketId *big.Int, owedMarketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetLiquidationSpreadForPair(&_SoloMargin.CallOpts, heldMarketId, owedMarketId)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0x4f3c1542.
//
// Solidity: function getMarginRatio() constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarginRatio(opts *bind.CallOpts) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarginRatio")
	return *ret0, err
}

// GetMarginRatio is a free data retrieval call binding the contract method 0x4f3c1542.
//
// Solidity: function getMarginRatio() constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMarginRatio() (Struct0, error) {
	return _SoloMargin.Contract.GetMarginRatio(&_SoloMargin.CallOpts)
}

// GetMarginRatio is a free data retrieval call binding the contract method 0x4f3c1542.
//
// Solidity: function getMarginRatio() constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarginRatio() (Struct0, error) {
	return _SoloMargin.Contract.GetMarginRatio(&_SoloMargin.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) constant returns(Struct11)
func (_SoloMargin *SoloMarginCaller) GetMarket(opts *bind.CallOpts, marketId *big.Int) (Struct11, error) {
	var (
		ret0 = new(Struct11)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarket", marketId)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) constant returns(Struct11)
func (_SoloMargin *SoloMarginSession) GetMarket(marketId *big.Int) (Struct11, error) {
	return _SoloMargin.Contract.GetMarket(&_SoloMargin.CallOpts, marketId)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 marketId) constant returns(Struct11)
func (_SoloMargin *SoloMarginCallerSession) GetMarket(marketId *big.Int) (Struct11, error) {
	return _SoloMargin.Contract.GetMarket(&_SoloMargin.CallOpts, marketId)
}

// GetMarketCachedIndex is a free data retrieval call binding the contract method 0xdeec053d.
//
// Solidity: function getMarketCachedIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginCaller) GetMarketCachedIndex(opts *bind.CallOpts, marketId *big.Int) (Struct2, error) {
	var (
		ret0 = new(Struct2)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketCachedIndex", marketId)
	return *ret0, err
}

// GetMarketCachedIndex is a free data retrieval call binding the contract method 0xdeec053d.
//
// Solidity: function getMarketCachedIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginSession) GetMarketCachedIndex(marketId *big.Int) (Struct2, error) {
	return _SoloMargin.Contract.GetMarketCachedIndex(&_SoloMargin.CallOpts, marketId)
}

// GetMarketCachedIndex is a free data retrieval call binding the contract method 0xdeec053d.
//
// Solidity: function getMarketCachedIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginCallerSession) GetMarketCachedIndex(marketId *big.Int) (Struct2, error) {
	return _SoloMargin.Contract.GetMarketCachedIndex(&_SoloMargin.CallOpts, marketId)
}

// GetMarketCurrentIndex is a free data retrieval call binding the contract method 0x56ea84b2.
//
// Solidity: function getMarketCurrentIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginCaller) GetMarketCurrentIndex(opts *bind.CallOpts, marketId *big.Int) (Struct2, error) {
	var (
		ret0 = new(Struct2)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketCurrentIndex", marketId)
	return *ret0, err
}

// GetMarketCurrentIndex is a free data retrieval call binding the contract method 0x56ea84b2.
//
// Solidity: function getMarketCurrentIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginSession) GetMarketCurrentIndex(marketId *big.Int) (Struct2, error) {
	return _SoloMargin.Contract.GetMarketCurrentIndex(&_SoloMargin.CallOpts, marketId)
}

// GetMarketCurrentIndex is a free data retrieval call binding the contract method 0x56ea84b2.
//
// Solidity: function getMarketCurrentIndex(uint256 marketId) constant returns(Struct2)
func (_SoloMargin *SoloMarginCallerSession) GetMarketCurrentIndex(marketId *big.Int) (Struct2, error) {
	return _SoloMargin.Contract.GetMarketCurrentIndex(&_SoloMargin.CallOpts, marketId)
}

// GetMarketInterestRate is a free data retrieval call binding the contract method 0xfd47eda6.
//
// Solidity: function getMarketInterestRate(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarketInterestRate(opts *bind.CallOpts, marketId *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketInterestRate", marketId)
	return *ret0, err
}

// GetMarketInterestRate is a free data retrieval call binding the contract method 0xfd47eda6.
//
// Solidity: function getMarketInterestRate(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMarketInterestRate(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketInterestRate(&_SoloMargin.CallOpts, marketId)
}

// GetMarketInterestRate is a free data retrieval call binding the contract method 0xfd47eda6.
//
// Solidity: function getMarketInterestRate(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarketInterestRate(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketInterestRate(&_SoloMargin.CallOpts, marketId)
}

// GetMarketInterestSetter is a free data retrieval call binding the contract method 0x197f0f05.
//
// Solidity: function getMarketInterestSetter(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCaller) GetMarketInterestSetter(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketInterestSetter", marketId)
	return *ret0, err
}

// GetMarketInterestSetter is a free data retrieval call binding the contract method 0x197f0f05.
//
// Solidity: function getMarketInterestSetter(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginSession) GetMarketInterestSetter(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketInterestSetter(&_SoloMargin.CallOpts, marketId)
}

// GetMarketInterestSetter is a free data retrieval call binding the contract method 0x197f0f05.
//
// Solidity: function getMarketInterestSetter(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCallerSession) GetMarketInterestSetter(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketInterestSetter(&_SoloMargin.CallOpts, marketId)
}

// GetMarketIsClosing is a free data retrieval call binding the contract method 0x5ac7d17c.
//
// Solidity: function getMarketIsClosing(uint256 marketId) constant returns(bool)
func (_SoloMargin *SoloMarginCaller) GetMarketIsClosing(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketIsClosing", marketId)
	return *ret0, err
}

// GetMarketIsClosing is a free data retrieval call binding the contract method 0x5ac7d17c.
//
// Solidity: function getMarketIsClosing(uint256 marketId) constant returns(bool)
func (_SoloMargin *SoloMarginSession) GetMarketIsClosing(marketId *big.Int) (bool, error) {
	return _SoloMargin.Contract.GetMarketIsClosing(&_SoloMargin.CallOpts, marketId)
}

// GetMarketIsClosing is a free data retrieval call binding the contract method 0x5ac7d17c.
//
// Solidity: function getMarketIsClosing(uint256 marketId) constant returns(bool)
func (_SoloMargin *SoloMarginCallerSession) GetMarketIsClosing(marketId *big.Int) (bool, error) {
	return _SoloMargin.Contract.GetMarketIsClosing(&_SoloMargin.CallOpts, marketId)
}

// GetMarketMarginPremium is a free data retrieval call binding the contract method 0xfd04b606.
//
// Solidity: function getMarketMarginPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarketMarginPremium(opts *bind.CallOpts, marketId *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketMarginPremium", marketId)
	return *ret0, err
}

// GetMarketMarginPremium is a free data retrieval call binding the contract method 0xfd04b606.
//
// Solidity: function getMarketMarginPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMarketMarginPremium(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketMarginPremium(&_SoloMargin.CallOpts, marketId)
}

// GetMarketMarginPremium is a free data retrieval call binding the contract method 0xfd04b606.
//
// Solidity: function getMarketMarginPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarketMarginPremium(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketMarginPremium(&_SoloMargin.CallOpts, marketId)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x8928378e.
//
// Solidity: function getMarketPrice(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarketPrice(opts *bind.CallOpts, marketId *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketPrice", marketId)
	return *ret0, err
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x8928378e.
//
// Solidity: function getMarketPrice(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMarketPrice(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketPrice(&_SoloMargin.CallOpts, marketId)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x8928378e.
//
// Solidity: function getMarketPrice(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarketPrice(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketPrice(&_SoloMargin.CallOpts, marketId)
}

// GetMarketPriceOracle is a free data retrieval call binding the contract method 0x13368364.
//
// Solidity: function getMarketPriceOracle(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCaller) GetMarketPriceOracle(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketPriceOracle", marketId)
	return *ret0, err
}

// GetMarketPriceOracle is a free data retrieval call binding the contract method 0x13368364.
//
// Solidity: function getMarketPriceOracle(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginSession) GetMarketPriceOracle(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketPriceOracle(&_SoloMargin.CallOpts, marketId)
}

// GetMarketPriceOracle is a free data retrieval call binding the contract method 0x13368364.
//
// Solidity: function getMarketPriceOracle(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCallerSession) GetMarketPriceOracle(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketPriceOracle(&_SoloMargin.CallOpts, marketId)
}

// GetMarketSpreadPremium is a free data retrieval call binding the contract method 0x1a7777bb.
//
// Solidity: function getMarketSpreadPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarketSpreadPremium(opts *bind.CallOpts, marketId *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketSpreadPremium", marketId)
	return *ret0, err
}

// GetMarketSpreadPremium is a free data retrieval call binding the contract method 0x1a7777bb.
//
// Solidity: function getMarketSpreadPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMarketSpreadPremium(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketSpreadPremium(&_SoloMargin.CallOpts, marketId)
}

// GetMarketSpreadPremium is a free data retrieval call binding the contract method 0x1a7777bb.
//
// Solidity: function getMarketSpreadPremium(uint256 marketId) constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarketSpreadPremium(marketId *big.Int) (Struct0, error) {
	return _SoloMargin.Contract.GetMarketSpreadPremium(&_SoloMargin.CallOpts, marketId)
}

// GetMarketTokenAddress is a free data retrieval call binding the contract method 0x062bd3e9.
//
// Solidity: function getMarketTokenAddress(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCaller) GetMarketTokenAddress(opts *bind.CallOpts, marketId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketTokenAddress", marketId)
	return *ret0, err
}

// GetMarketTokenAddress is a free data retrieval call binding the contract method 0x062bd3e9.
//
// Solidity: function getMarketTokenAddress(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginSession) GetMarketTokenAddress(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketTokenAddress(&_SoloMargin.CallOpts, marketId)
}

// GetMarketTokenAddress is a free data retrieval call binding the contract method 0x062bd3e9.
//
// Solidity: function getMarketTokenAddress(uint256 marketId) constant returns(address)
func (_SoloMargin *SoloMarginCallerSession) GetMarketTokenAddress(marketId *big.Int) (common.Address, error) {
	return _SoloMargin.Contract.GetMarketTokenAddress(&_SoloMargin.CallOpts, marketId)
}

// GetMarketTotalPar is a free data retrieval call binding the contract method 0xcb04a34c.
//
// Solidity: function getMarketTotalPar(uint256 marketId) constant returns(Struct7)
func (_SoloMargin *SoloMarginCaller) GetMarketTotalPar(opts *bind.CallOpts, marketId *big.Int) (Struct7, error) {
	var (
		ret0 = new(Struct7)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMarketTotalPar", marketId)
	return *ret0, err
}

// GetMarketTotalPar is a free data retrieval call binding the contract method 0xcb04a34c.
//
// Solidity: function getMarketTotalPar(uint256 marketId) constant returns(Struct7)
func (_SoloMargin *SoloMarginSession) GetMarketTotalPar(marketId *big.Int) (Struct7, error) {
	return _SoloMargin.Contract.GetMarketTotalPar(&_SoloMargin.CallOpts, marketId)
}

// GetMarketTotalPar is a free data retrieval call binding the contract method 0xcb04a34c.
//
// Solidity: function getMarketTotalPar(uint256 marketId) constant returns(Struct7)
func (_SoloMargin *SoloMarginCallerSession) GetMarketTotalPar(marketId *big.Int) (Struct7, error) {
	return _SoloMargin.Contract.GetMarketTotalPar(&_SoloMargin.CallOpts, marketId)
}

// GetMarketWithInfo is a free data retrieval call binding the contract method 0xb548b892.
//
// Solidity: function getMarketWithInfo(uint256 marketId) constant returns(Struct11, Struct2, Struct0, Struct0)
func (_SoloMargin *SoloMarginCaller) GetMarketWithInfo(opts *bind.CallOpts, marketId *big.Int) (Struct11, Struct2, Struct0, Struct0, error) {
	var (
		ret0 = new(Struct11)
		ret1 = new(Struct2)
		ret2 = new(Struct0)
		ret3 = new(Struct0)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _SoloMargin.contract.Call(opts, out, "getMarketWithInfo", marketId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetMarketWithInfo is a free data retrieval call binding the contract method 0xb548b892.
//
// Solidity: function getMarketWithInfo(uint256 marketId) constant returns(Struct11, Struct2, Struct0, Struct0)
func (_SoloMargin *SoloMarginSession) GetMarketWithInfo(marketId *big.Int) (Struct11, Struct2, Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetMarketWithInfo(&_SoloMargin.CallOpts, marketId)
}

// GetMarketWithInfo is a free data retrieval call binding the contract method 0xb548b892.
//
// Solidity: function getMarketWithInfo(uint256 marketId) constant returns(Struct11, Struct2, Struct0, Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMarketWithInfo(marketId *big.Int) (Struct11, Struct2, Struct0, Struct0, error) {
	return _SoloMargin.Contract.GetMarketWithInfo(&_SoloMargin.CallOpts, marketId)
}

// GetMinBorrowedValue is a free data retrieval call binding the contract method 0x7e9eaf41.
//
// Solidity: function getMinBorrowedValue() constant returns(Struct0)
func (_SoloMargin *SoloMarginCaller) GetMinBorrowedValue(opts *bind.CallOpts) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getMinBorrowedValue")
	return *ret0, err
}

// GetMinBorrowedValue is a free data retrieval call binding the contract method 0x7e9eaf41.
//
// Solidity: function getMinBorrowedValue() constant returns(Struct0)
func (_SoloMargin *SoloMarginSession) GetMinBorrowedValue() (Struct0, error) {
	return _SoloMargin.Contract.GetMinBorrowedValue(&_SoloMargin.CallOpts)
}

// GetMinBorrowedValue is a free data retrieval call binding the contract method 0x7e9eaf41.
//
// Solidity: function getMinBorrowedValue() constant returns(Struct0)
func (_SoloMargin *SoloMarginCallerSession) GetMinBorrowedValue() (Struct0, error) {
	return _SoloMargin.Contract.GetMinBorrowedValue(&_SoloMargin.CallOpts)
}

// GetNumExcessTokens is a free data retrieval call binding the contract method 0xd5ecf7c5.
//
// Solidity: function getNumExcessTokens(uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginCaller) GetNumExcessTokens(opts *bind.CallOpts, marketId *big.Int) (Struct8, error) {
	var (
		ret0 = new(Struct8)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getNumExcessTokens", marketId)
	return *ret0, err
}

// GetNumExcessTokens is a free data retrieval call binding the contract method 0xd5ecf7c5.
//
// Solidity: function getNumExcessTokens(uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginSession) GetNumExcessTokens(marketId *big.Int) (Struct8, error) {
	return _SoloMargin.Contract.GetNumExcessTokens(&_SoloMargin.CallOpts, marketId)
}

// GetNumExcessTokens is a free data retrieval call binding the contract method 0xd5ecf7c5.
//
// Solidity: function getNumExcessTokens(uint256 marketId) constant returns(Struct8)
func (_SoloMargin *SoloMarginCallerSession) GetNumExcessTokens(marketId *big.Int) (Struct8, error) {
	return _SoloMargin.Contract.GetNumExcessTokens(&_SoloMargin.CallOpts, marketId)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_SoloMargin *SoloMarginCaller) GetNumMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getNumMarkets")
	return *ret0, err
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_SoloMargin *SoloMarginSession) GetNumMarkets() (*big.Int, error) {
	return _SoloMargin.Contract.GetNumMarkets(&_SoloMargin.CallOpts)
}

// GetNumMarkets is a free data retrieval call binding the contract method 0x295c39a5.
//
// Solidity: function getNumMarkets() constant returns(uint256)
func (_SoloMargin *SoloMarginCallerSession) GetNumMarkets() (*big.Int, error) {
	return _SoloMargin.Contract.GetNumMarkets(&_SoloMargin.CallOpts)
}

// GetRiskLimits is a free data retrieval call binding the contract method 0xeb1c6e6b.
//
// Solidity: function getRiskLimits() constant returns(Struct6)
func (_SoloMargin *SoloMarginCaller) GetRiskLimits(opts *bind.CallOpts) (Struct6, error) {
	var (
		ret0 = new(Struct6)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getRiskLimits")
	return *ret0, err
}

// GetRiskLimits is a free data retrieval call binding the contract method 0xeb1c6e6b.
//
// Solidity: function getRiskLimits() constant returns(Struct6)
func (_SoloMargin *SoloMarginSession) GetRiskLimits() (Struct6, error) {
	return _SoloMargin.Contract.GetRiskLimits(&_SoloMargin.CallOpts)
}

// GetRiskLimits is a free data retrieval call binding the contract method 0xeb1c6e6b.
//
// Solidity: function getRiskLimits() constant returns(Struct6)
func (_SoloMargin *SoloMarginCallerSession) GetRiskLimits() (Struct6, error) {
	return _SoloMargin.Contract.GetRiskLimits(&_SoloMargin.CallOpts)
}

// GetRiskParams is a free data retrieval call binding the contract method 0x69794795.
//
// Solidity: function getRiskParams() constant returns(Struct10)
func (_SoloMargin *SoloMarginCaller) GetRiskParams(opts *bind.CallOpts) (Struct10, error) {
	var (
		ret0 = new(Struct10)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "getRiskParams")
	return *ret0, err
}

// GetRiskParams is a free data retrieval call binding the contract method 0x69794795.
//
// Solidity: function getRiskParams() constant returns(Struct10)
func (_SoloMargin *SoloMarginSession) GetRiskParams() (Struct10, error) {
	return _SoloMargin.Contract.GetRiskParams(&_SoloMargin.CallOpts)
}

// GetRiskParams is a free data retrieval call binding the contract method 0x69794795.
//
// Solidity: function getRiskParams() constant returns(Struct10)
func (_SoloMargin *SoloMarginCallerSession) GetRiskParams() (Struct10, error) {
	return _SoloMargin.Contract.GetRiskParams(&_SoloMargin.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SoloMargin *SoloMarginCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SoloMargin *SoloMarginSession) IsOwner() (bool, error) {
	return _SoloMargin.Contract.IsOwner(&_SoloMargin.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SoloMargin *SoloMarginCallerSession) IsOwner() (bool, error) {
	return _SoloMargin.Contract.IsOwner(&_SoloMargin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SoloMargin *SoloMarginCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SoloMargin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SoloMargin *SoloMarginSession) Owner() (common.Address, error) {
	return _SoloMargin.Contract.Owner(&_SoloMargin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SoloMargin *SoloMarginCallerSession) Owner() (common.Address, error) {
	return _SoloMargin.Contract.Owner(&_SoloMargin.CallOpts)
}

// Operate is a paid mutator transaction binding the contract method 0xa67a6a45.
//
// Solidity: function operate([]Struct1 accounts, []Struct4 actions) returns()
func (_SoloMargin *SoloMarginTransactor) Operate(opts *bind.TransactOpts, accounts []Struct1, actions []Struct4) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "operate", accounts, actions)
}

// Operate is a paid mutator transaction binding the contract method 0xa67a6a45.
//
// Solidity: function operate([]Struct1 accounts, []Struct4 actions) returns()
func (_SoloMargin *SoloMarginSession) Operate(accounts []Struct1, actions []Struct4) (*types.Transaction, error) {
	return _SoloMargin.Contract.Operate(&_SoloMargin.TransactOpts, accounts, actions)
}

// Operate is a paid mutator transaction binding the contract method 0xa67a6a45.
//
// Solidity: function operate([]Struct1 accounts, []Struct4 actions) returns()
func (_SoloMargin *SoloMarginTransactorSession) Operate(accounts []Struct1, actions []Struct4) (*types.Transaction, error) {
	return _SoloMargin.Contract.Operate(&_SoloMargin.TransactOpts, accounts, actions)
}

// OwnerAddMarket is a paid mutator transaction binding the contract method 0x982f323c.
//
// Solidity: function ownerAddMarket(address token, address priceOracle, address interestSetter, Struct0 marginPremium, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerAddMarket(opts *bind.TransactOpts, token common.Address, priceOracle common.Address, interestSetter common.Address, marginPremium Struct0, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerAddMarket", token, priceOracle, interestSetter, marginPremium, spreadPremium)
}

// OwnerAddMarket is a paid mutator transaction binding the contract method 0x982f323c.
//
// Solidity: function ownerAddMarket(address token, address priceOracle, address interestSetter, Struct0 marginPremium, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginSession) OwnerAddMarket(token common.Address, priceOracle common.Address, interestSetter common.Address, marginPremium Struct0, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerAddMarket(&_SoloMargin.TransactOpts, token, priceOracle, interestSetter, marginPremium, spreadPremium)
}

// OwnerAddMarket is a paid mutator transaction binding the contract method 0x982f323c.
//
// Solidity: function ownerAddMarket(address token, address priceOracle, address interestSetter, Struct0 marginPremium, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerAddMarket(token common.Address, priceOracle common.Address, interestSetter common.Address, marginPremium Struct0, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerAddMarket(&_SoloMargin.TransactOpts, token, priceOracle, interestSetter, marginPremium, spreadPremium)
}

// OwnerSetEarningsRate is a paid mutator transaction binding the contract method 0x387a498a.
//
// Solidity: function ownerSetEarningsRate(Struct0 earningsRate) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetEarningsRate(opts *bind.TransactOpts, earningsRate Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetEarningsRate", earningsRate)
}

// OwnerSetEarningsRate is a paid mutator transaction binding the contract method 0x387a498a.
//
// Solidity: function ownerSetEarningsRate(Struct0 earningsRate) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetEarningsRate(earningsRate Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetEarningsRate(&_SoloMargin.TransactOpts, earningsRate)
}

// OwnerSetEarningsRate is a paid mutator transaction binding the contract method 0x387a498a.
//
// Solidity: function ownerSetEarningsRate(Struct0 earningsRate) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetEarningsRate(earningsRate Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetEarningsRate(&_SoloMargin.TransactOpts, earningsRate)
}

// OwnerSetGlobalOperator is a paid mutator transaction binding the contract method 0xf2901ae2.
//
// Solidity: function ownerSetGlobalOperator(address operator, bool approved) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetGlobalOperator(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetGlobalOperator", operator, approved)
}

// OwnerSetGlobalOperator is a paid mutator transaction binding the contract method 0xf2901ae2.
//
// Solidity: function ownerSetGlobalOperator(address operator, bool approved) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetGlobalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetGlobalOperator(&_SoloMargin.TransactOpts, operator, approved)
}

// OwnerSetGlobalOperator is a paid mutator transaction binding the contract method 0xf2901ae2.
//
// Solidity: function ownerSetGlobalOperator(address operator, bool approved) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetGlobalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetGlobalOperator(&_SoloMargin.TransactOpts, operator, approved)
}

// OwnerSetInterestSetter is a paid mutator transaction binding the contract method 0x121fb72f.
//
// Solidity: function ownerSetInterestSetter(uint256 marketId, address interestSetter) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetInterestSetter(opts *bind.TransactOpts, marketId *big.Int, interestSetter common.Address) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetInterestSetter", marketId, interestSetter)
}

// OwnerSetInterestSetter is a paid mutator transaction binding the contract method 0x121fb72f.
//
// Solidity: function ownerSetInterestSetter(uint256 marketId, address interestSetter) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetInterestSetter(marketId *big.Int, interestSetter common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetInterestSetter(&_SoloMargin.TransactOpts, marketId, interestSetter)
}

// OwnerSetInterestSetter is a paid mutator transaction binding the contract method 0x121fb72f.
//
// Solidity: function ownerSetInterestSetter(uint256 marketId, address interestSetter) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetInterestSetter(marketId *big.Int, interestSetter common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetInterestSetter(&_SoloMargin.TransactOpts, marketId, interestSetter)
}

// OwnerSetIsClosing is a paid mutator transaction binding the contract method 0xef6957d0.
//
// Solidity: function ownerSetIsClosing(uint256 marketId, bool isClosing) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetIsClosing(opts *bind.TransactOpts, marketId *big.Int, isClosing bool) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetIsClosing", marketId, isClosing)
}

// OwnerSetIsClosing is a paid mutator transaction binding the contract method 0xef6957d0.
//
// Solidity: function ownerSetIsClosing(uint256 marketId, bool isClosing) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetIsClosing(marketId *big.Int, isClosing bool) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetIsClosing(&_SoloMargin.TransactOpts, marketId, isClosing)
}

// OwnerSetIsClosing is a paid mutator transaction binding the contract method 0xef6957d0.
//
// Solidity: function ownerSetIsClosing(uint256 marketId, bool isClosing) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetIsClosing(marketId *big.Int, isClosing bool) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetIsClosing(&_SoloMargin.TransactOpts, marketId, isClosing)
}

// OwnerSetLiquidationSpread is a paid mutator transaction binding the contract method 0x3063bce2.
//
// Solidity: function ownerSetLiquidationSpread(Struct0 spread) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetLiquidationSpread(opts *bind.TransactOpts, spread Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetLiquidationSpread", spread)
}

// OwnerSetLiquidationSpread is a paid mutator transaction binding the contract method 0x3063bce2.
//
// Solidity: function ownerSetLiquidationSpread(Struct0 spread) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetLiquidationSpread(spread Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetLiquidationSpread(&_SoloMargin.TransactOpts, spread)
}

// OwnerSetLiquidationSpread is a paid mutator transaction binding the contract method 0x3063bce2.
//
// Solidity: function ownerSetLiquidationSpread(Struct0 spread) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetLiquidationSpread(spread Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetLiquidationSpread(&_SoloMargin.TransactOpts, spread)
}

// OwnerSetMarginPremium is a paid mutator transaction binding the contract method 0x4be87414.
//
// Solidity: function ownerSetMarginPremium(uint256 marketId, Struct0 marginPremium) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetMarginPremium(opts *bind.TransactOpts, marketId *big.Int, marginPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetMarginPremium", marketId, marginPremium)
}

// OwnerSetMarginPremium is a paid mutator transaction binding the contract method 0x4be87414.
//
// Solidity: function ownerSetMarginPremium(uint256 marketId, Struct0 marginPremium) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetMarginPremium(marketId *big.Int, marginPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMarginPremium(&_SoloMargin.TransactOpts, marketId, marginPremium)
}

// OwnerSetMarginPremium is a paid mutator transaction binding the contract method 0x4be87414.
//
// Solidity: function ownerSetMarginPremium(uint256 marketId, Struct0 marginPremium) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetMarginPremium(marketId *big.Int, marginPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMarginPremium(&_SoloMargin.TransactOpts, marketId, marginPremium)
}

// OwnerSetMarginRatio is a paid mutator transaction binding the contract method 0xc0bb72b7.
//
// Solidity: function ownerSetMarginRatio(Struct0 ratio) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetMarginRatio(opts *bind.TransactOpts, ratio Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetMarginRatio", ratio)
}

// OwnerSetMarginRatio is a paid mutator transaction binding the contract method 0xc0bb72b7.
//
// Solidity: function ownerSetMarginRatio(Struct0 ratio) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetMarginRatio(ratio Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMarginRatio(&_SoloMargin.TransactOpts, ratio)
}

// OwnerSetMarginRatio is a paid mutator transaction binding the contract method 0xc0bb72b7.
//
// Solidity: function ownerSetMarginRatio(Struct0 ratio) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetMarginRatio(ratio Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMarginRatio(&_SoloMargin.TransactOpts, ratio)
}

// OwnerSetMinBorrowedValue is a paid mutator transaction binding the contract method 0x2e822af3.
//
// Solidity: function ownerSetMinBorrowedValue(Struct0 minBorrowedValue) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetMinBorrowedValue(opts *bind.TransactOpts, minBorrowedValue Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetMinBorrowedValue", minBorrowedValue)
}

// OwnerSetMinBorrowedValue is a paid mutator transaction binding the contract method 0x2e822af3.
//
// Solidity: function ownerSetMinBorrowedValue(Struct0 minBorrowedValue) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetMinBorrowedValue(minBorrowedValue Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMinBorrowedValue(&_SoloMargin.TransactOpts, minBorrowedValue)
}

// OwnerSetMinBorrowedValue is a paid mutator transaction binding the contract method 0x2e822af3.
//
// Solidity: function ownerSetMinBorrowedValue(Struct0 minBorrowedValue) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetMinBorrowedValue(minBorrowedValue Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetMinBorrowedValue(&_SoloMargin.TransactOpts, minBorrowedValue)
}

// OwnerSetPriceOracle is a paid mutator transaction binding the contract method 0xe8e72f75.
//
// Solidity: function ownerSetPriceOracle(uint256 marketId, address priceOracle) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetPriceOracle(opts *bind.TransactOpts, marketId *big.Int, priceOracle common.Address) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetPriceOracle", marketId, priceOracle)
}

// OwnerSetPriceOracle is a paid mutator transaction binding the contract method 0xe8e72f75.
//
// Solidity: function ownerSetPriceOracle(uint256 marketId, address priceOracle) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetPriceOracle(marketId *big.Int, priceOracle common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetPriceOracle(&_SoloMargin.TransactOpts, marketId, priceOracle)
}

// OwnerSetPriceOracle is a paid mutator transaction binding the contract method 0xe8e72f75.
//
// Solidity: function ownerSetPriceOracle(uint256 marketId, address priceOracle) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetPriceOracle(marketId *big.Int, priceOracle common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetPriceOracle(&_SoloMargin.TransactOpts, marketId, priceOracle)
}

// OwnerSetSpreadPremium is a paid mutator transaction binding the contract method 0x042069d6.
//
// Solidity: function ownerSetSpreadPremium(uint256 marketId, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginTransactor) OwnerSetSpreadPremium(opts *bind.TransactOpts, marketId *big.Int, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerSetSpreadPremium", marketId, spreadPremium)
}

// OwnerSetSpreadPremium is a paid mutator transaction binding the contract method 0x042069d6.
//
// Solidity: function ownerSetSpreadPremium(uint256 marketId, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginSession) OwnerSetSpreadPremium(marketId *big.Int, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetSpreadPremium(&_SoloMargin.TransactOpts, marketId, spreadPremium)
}

// OwnerSetSpreadPremium is a paid mutator transaction binding the contract method 0x042069d6.
//
// Solidity: function ownerSetSpreadPremium(uint256 marketId, Struct0 spreadPremium) returns()
func (_SoloMargin *SoloMarginTransactorSession) OwnerSetSpreadPremium(marketId *big.Int, spreadPremium Struct0) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerSetSpreadPremium(&_SoloMargin.TransactOpts, marketId, spreadPremium)
}

// OwnerWithdrawExcessTokens is a paid mutator transaction binding the contract method 0x8f6bc659.
//
// Solidity: function ownerWithdrawExcessTokens(uint256 marketId, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginTransactor) OwnerWithdrawExcessTokens(opts *bind.TransactOpts, marketId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerWithdrawExcessTokens", marketId, recipient)
}

// OwnerWithdrawExcessTokens is a paid mutator transaction binding the contract method 0x8f6bc659.
//
// Solidity: function ownerWithdrawExcessTokens(uint256 marketId, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginSession) OwnerWithdrawExcessTokens(marketId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerWithdrawExcessTokens(&_SoloMargin.TransactOpts, marketId, recipient)
}

// OwnerWithdrawExcessTokens is a paid mutator transaction binding the contract method 0x8f6bc659.
//
// Solidity: function ownerWithdrawExcessTokens(uint256 marketId, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginTransactorSession) OwnerWithdrawExcessTokens(marketId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerWithdrawExcessTokens(&_SoloMargin.TransactOpts, marketId, recipient)
}

// OwnerWithdrawUnsupportedTokens is a paid mutator transaction binding the contract method 0x2a560845.
//
// Solidity: function ownerWithdrawUnsupportedTokens(address token, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginTransactor) OwnerWithdrawUnsupportedTokens(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "ownerWithdrawUnsupportedTokens", token, recipient)
}

// OwnerWithdrawUnsupportedTokens is a paid mutator transaction binding the contract method 0x2a560845.
//
// Solidity: function ownerWithdrawUnsupportedTokens(address token, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginSession) OwnerWithdrawUnsupportedTokens(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerWithdrawUnsupportedTokens(&_SoloMargin.TransactOpts, token, recipient)
}

// OwnerWithdrawUnsupportedTokens is a paid mutator transaction binding the contract method 0x2a560845.
//
// Solidity: function ownerWithdrawUnsupportedTokens(address token, address recipient) returns(uint256)
func (_SoloMargin *SoloMarginTransactorSession) OwnerWithdrawUnsupportedTokens(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.OwnerWithdrawUnsupportedTokens(&_SoloMargin.TransactOpts, token, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoloMargin *SoloMarginTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoloMargin *SoloMarginSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoloMargin.Contract.RenounceOwnership(&_SoloMargin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoloMargin *SoloMarginTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoloMargin.Contract.RenounceOwnership(&_SoloMargin.TransactOpts)
}

// SetOperators is a paid mutator transaction binding the contract method 0x85b53fc8.
//
// Solidity: function setOperators([]Struct9 args) returns()
func (_SoloMargin *SoloMarginTransactor) SetOperators(opts *bind.TransactOpts, args []Struct9) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "setOperators", args)
}

// SetOperators is a paid mutator transaction binding the contract method 0x85b53fc8.
//
// Solidity: function setOperators([]Struct9 args) returns()
func (_SoloMargin *SoloMarginSession) SetOperators(args []Struct9) (*types.Transaction, error) {
	return _SoloMargin.Contract.SetOperators(&_SoloMargin.TransactOpts, args)
}

// SetOperators is a paid mutator transaction binding the contract method 0x85b53fc8.
//
// Solidity: function setOperators([]Struct9 args) returns()
func (_SoloMargin *SoloMarginTransactorSession) SetOperators(args []Struct9) (*types.Transaction, error) {
	return _SoloMargin.Contract.SetOperators(&_SoloMargin.TransactOpts, args)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoloMargin *SoloMarginTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SoloMargin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoloMargin *SoloMarginSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.TransferOwnership(&_SoloMargin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoloMargin *SoloMarginTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoloMargin.Contract.TransferOwnership(&_SoloMargin.TransactOpts, newOwner)
}

// SoloMarginLogOperatorSetIterator is returned from FilterLogOperatorSet and is used to iterate over the raw logs and unpacked data for LogOperatorSet events raised by the SoloMargin contract.
type SoloMarginLogOperatorSetIterator struct {
	Event *SoloMarginLogOperatorSet // Event containing the contract specifics and raw log

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
func (it *SoloMarginLogOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoloMarginLogOperatorSet)
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
		it.Event = new(SoloMarginLogOperatorSet)
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
func (it *SoloMarginLogOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoloMarginLogOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoloMarginLogOperatorSet represents a LogOperatorSet event raised by the SoloMargin contract.
type SoloMarginLogOperatorSet struct {
	Owner    common.Address
	Operator common.Address
	Trusted  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorSet is a free log retrieval operation binding the contract event 0x4d7f317d2088d039c2a95a09fcbf9cc9191fad5905f883c937cc3d317c4a6327.
//
// Solidity: event LogOperatorSet(address indexed owner, address operator, bool trusted)
func (_SoloMargin *SoloMarginFilterer) FilterLogOperatorSet(opts *bind.FilterOpts, owner []common.Address) (*SoloMarginLogOperatorSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoloMargin.contract.FilterLogs(opts, "LogOperatorSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SoloMarginLogOperatorSetIterator{contract: _SoloMargin.contract, event: "LogOperatorSet", logs: logs, sub: sub}, nil
}

// WatchLogOperatorSet is a free log subscription operation binding the contract event 0x4d7f317d2088d039c2a95a09fcbf9cc9191fad5905f883c937cc3d317c4a6327.
//
// Solidity: event LogOperatorSet(address indexed owner, address operator, bool trusted)
func (_SoloMargin *SoloMarginFilterer) WatchLogOperatorSet(opts *bind.WatchOpts, sink chan<- *SoloMarginLogOperatorSet, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SoloMargin.contract.WatchLogs(opts, "LogOperatorSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoloMarginLogOperatorSet)
				if err := _SoloMargin.contract.UnpackLog(event, "LogOperatorSet", log); err != nil {
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

// ParseLogOperatorSet is a log parse operation binding the contract event 0x4d7f317d2088d039c2a95a09fcbf9cc9191fad5905f883c937cc3d317c4a6327.
//
// Solidity: event LogOperatorSet(address indexed owner, address operator, bool trusted)
func (_SoloMargin *SoloMarginFilterer) ParseLogOperatorSet(log types.Log) (*SoloMarginLogOperatorSet, error) {
	event := new(SoloMarginLogOperatorSet)
	if err := _SoloMargin.contract.UnpackLog(event, "LogOperatorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SoloMarginOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SoloMargin contract.
type SoloMarginOwnershipTransferredIterator struct {
	Event *SoloMarginOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SoloMarginOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoloMarginOwnershipTransferred)
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
		it.Event = new(SoloMarginOwnershipTransferred)
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
func (it *SoloMarginOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoloMarginOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoloMarginOwnershipTransferred represents a OwnershipTransferred event raised by the SoloMargin contract.
type SoloMarginOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoloMargin *SoloMarginFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SoloMarginOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoloMargin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SoloMarginOwnershipTransferredIterator{contract: _SoloMargin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoloMargin *SoloMarginFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SoloMarginOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoloMargin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoloMarginOwnershipTransferred)
				if err := _SoloMargin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoloMargin *SoloMarginFilterer) ParseOwnershipTransferred(log types.Log) (*SoloMarginOwnershipTransferred, error) {
	event := new(SoloMarginOwnershipTransferred)
	if err := _SoloMargin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
