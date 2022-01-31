// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hyphen

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ITokenManagerTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type ITokenManagerTokenInfo struct {
	TransferOverhead *big.Int
	SupportedToken   bool
	MinCap           *big.Int
	MaxCap           *big.Int
	EquilibriumFee   *big.Int
	MaxFee           *big.Int
}

// TokenManagerMetaData contains all meta data concerning the TokenManager contract.
var TokenManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"equilibriumFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"FeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minCapLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCapLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"equilibriumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_equilibriumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getEquilibriumFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getMaxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokensInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"transferOverhead\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"supportedToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"equilibriumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"internalType\":\"structITokenManager.TokenInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"}],\"name\":\"setTokenTransferOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferOverhead\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"supportedToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"equilibriumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minCapLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCapLimit\",\"type\":\"uint256\"}],\"name\":\"updateTokenCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenManagerMetaData.ABI instead.
var TokenManagerABI = TokenManagerMetaData.ABI

// TokenManager is an auto generated Go binding around an Ethereum contract.
type TokenManager struct {
	TokenManagerCaller     // Read-only binding to the contract
	TokenManagerTransactor // Write-only binding to the contract
	TokenManagerFilterer   // Log filterer for contract events
}

// TokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenManagerSession struct {
	Contract     *TokenManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenManagerCallerSession struct {
	Contract *TokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenManagerTransactorSession struct {
	Contract     *TokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenManagerRaw struct {
	Contract *TokenManager // Generic contract binding to access the raw methods on
}

// TokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenManagerCallerRaw struct {
	Contract *TokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenManagerTransactorRaw struct {
	Contract *TokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenManager creates a new instance of TokenManager, bound to a specific deployed contract.
func NewTokenManager(address common.Address, backend bind.ContractBackend) (*TokenManager, error) {
	contract, err := bindTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenManager{TokenManagerCaller: TokenManagerCaller{contract: contract}, TokenManagerTransactor: TokenManagerTransactor{contract: contract}, TokenManagerFilterer: TokenManagerFilterer{contract: contract}}, nil
}

// NewTokenManagerCaller creates a new read-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*TokenManagerCaller, error) {
	contract, err := bindTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerCaller{contract: contract}, nil
}

// NewTokenManagerTransactor creates a new write-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenManagerTransactor, error) {
	contract, err := bindTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerTransactor{contract: contract}, nil
}

// NewTokenManagerFilterer creates a new log filterer instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenManagerFilterer, error) {
	contract, err := bindTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenManagerFilterer{contract: contract}, nil
}

// bindTokenManager binds a generic wrapper to an already deployed contract.
func bindTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.TokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transact(opts, method, params...)
}

// GetEquilibriumFee is a free data retrieval call binding the contract method 0x85dc3013.
//
// Solidity: function getEquilibriumFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerCaller) GetEquilibriumFee(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "getEquilibriumFee", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEquilibriumFee is a free data retrieval call binding the contract method 0x85dc3013.
//
// Solidity: function getEquilibriumFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerSession) GetEquilibriumFee(tokenAddress common.Address) (*big.Int, error) {
	return _TokenManager.Contract.GetEquilibriumFee(&_TokenManager.CallOpts, tokenAddress)
}

// GetEquilibriumFee is a free data retrieval call binding the contract method 0x85dc3013.
//
// Solidity: function getEquilibriumFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) GetEquilibriumFee(tokenAddress common.Address) (*big.Int, error) {
	return _TokenManager.Contract.GetEquilibriumFee(&_TokenManager.CallOpts, tokenAddress)
}

// GetMaxFee is a free data retrieval call binding the contract method 0x60f5dfda.
//
// Solidity: function getMaxFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerCaller) GetMaxFee(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "getMaxFee", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFee is a free data retrieval call binding the contract method 0x60f5dfda.
//
// Solidity: function getMaxFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerSession) GetMaxFee(tokenAddress common.Address) (*big.Int, error) {
	return _TokenManager.Contract.GetMaxFee(&_TokenManager.CallOpts, tokenAddress)
}

// GetMaxFee is a free data retrieval call binding the contract method 0x60f5dfda.
//
// Solidity: function getMaxFee(address tokenAddress) view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) GetMaxFee(tokenAddress common.Address) (*big.Int, error) {
	return _TokenManager.Contract.GetMaxFee(&_TokenManager.CallOpts, tokenAddress)
}

// GetTokensInfo is a free data retrieval call binding the contract method 0xfa5e2abc.
//
// Solidity: function getTokensInfo(address tokenAddress) view returns((uint256,bool,uint256,uint256,uint256,uint256))
func (_TokenManager *TokenManagerCaller) GetTokensInfo(opts *bind.CallOpts, tokenAddress common.Address) (ITokenManagerTokenInfo, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "getTokensInfo", tokenAddress)

	if err != nil {
		return *new(ITokenManagerTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITokenManagerTokenInfo)).(*ITokenManagerTokenInfo)

	return out0, err

}

// GetTokensInfo is a free data retrieval call binding the contract method 0xfa5e2abc.
//
// Solidity: function getTokensInfo(address tokenAddress) view returns((uint256,bool,uint256,uint256,uint256,uint256))
func (_TokenManager *TokenManagerSession) GetTokensInfo(tokenAddress common.Address) (ITokenManagerTokenInfo, error) {
	return _TokenManager.Contract.GetTokensInfo(&_TokenManager.CallOpts, tokenAddress)
}

// GetTokensInfo is a free data retrieval call binding the contract method 0xfa5e2abc.
//
// Solidity: function getTokensInfo(address tokenAddress) view returns((uint256,bool,uint256,uint256,uint256,uint256))
func (_TokenManager *TokenManagerCallerSession) GetTokensInfo(tokenAddress common.Address) (ITokenManagerTokenInfo, error) {
	return _TokenManager.Contract.GetTokensInfo(&_TokenManager.CallOpts, tokenAddress)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenManager *TokenManagerCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenManager *TokenManagerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenManager.Contract.IsTrustedForwarder(&_TokenManager.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenManager.Contract.IsTrustedForwarder(&_TokenManager.CallOpts, forwarder)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCallerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenManager *TokenManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenManager *TokenManagerSession) Paused() (bool, error) {
	return _TokenManager.Contract.Paused(&_TokenManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenManager *TokenManagerCallerSession) Paused() (bool, error) {
	return _TokenManager.Contract.Paused(&_TokenManager.CallOpts)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 transferOverhead, bool supportedToken, uint256 minCap, uint256 maxCap, uint256 equilibriumFee, uint256 maxFee)
func (_TokenManager *TokenManagerCaller) TokensInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	TransferOverhead *big.Int
	SupportedToken   bool
	MinCap           *big.Int
	MaxCap           *big.Int
	EquilibriumFee   *big.Int
	MaxFee           *big.Int
}, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "tokensInfo", arg0)

	outstruct := new(struct {
		TransferOverhead *big.Int
		SupportedToken   bool
		MinCap           *big.Int
		MaxCap           *big.Int
		EquilibriumFee   *big.Int
		MaxFee           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransferOverhead = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SupportedToken = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.MinCap = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxCap = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EquilibriumFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 transferOverhead, bool supportedToken, uint256 minCap, uint256 maxCap, uint256 equilibriumFee, uint256 maxFee)
func (_TokenManager *TokenManagerSession) TokensInfo(arg0 common.Address) (struct {
	TransferOverhead *big.Int
	SupportedToken   bool
	MinCap           *big.Int
	MaxCap           *big.Int
	EquilibriumFee   *big.Int
	MaxFee           *big.Int
}, error) {
	return _TokenManager.Contract.TokensInfo(&_TokenManager.CallOpts, arg0)
}

// TokensInfo is a free data retrieval call binding the contract method 0xba8dbea2.
//
// Solidity: function tokensInfo(address ) view returns(uint256 transferOverhead, bool supportedToken, uint256 minCap, uint256 maxCap, uint256 equilibriumFee, uint256 maxFee)
func (_TokenManager *TokenManagerCallerSession) TokensInfo(arg0 common.Address) (struct {
	TransferOverhead *big.Int
	SupportedToken   bool
	MinCap           *big.Int
	MaxCap           *big.Int
	EquilibriumFee   *big.Int
	MaxFee           *big.Int
}, error) {
	return _TokenManager.Contract.TokensInfo(&_TokenManager.CallOpts, arg0)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x3d3ebe53.
//
// Solidity: function addSupportedToken(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit, uint256 equilibriumFee, uint256 maxFee) returns()
func (_TokenManager *TokenManagerTransactor) AddSupportedToken(opts *bind.TransactOpts, tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int, equilibriumFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "addSupportedToken", tokenAddress, minCapLimit, maxCapLimit, equilibriumFee, maxFee)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x3d3ebe53.
//
// Solidity: function addSupportedToken(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit, uint256 equilibriumFee, uint256 maxFee) returns()
func (_TokenManager *TokenManagerSession) AddSupportedToken(tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int, equilibriumFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.AddSupportedToken(&_TokenManager.TransactOpts, tokenAddress, minCapLimit, maxCapLimit, equilibriumFee, maxFee)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x3d3ebe53.
//
// Solidity: function addSupportedToken(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit, uint256 equilibriumFee, uint256 maxFee) returns()
func (_TokenManager *TokenManagerTransactorSession) AddSupportedToken(tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int, equilibriumFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.AddSupportedToken(&_TokenManager.TransactOpts, tokenAddress, minCapLimit, maxCapLimit, equilibriumFee, maxFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xb5ec9999.
//
// Solidity: function changeFee(address tokenAddress, uint256 _equilibriumFee, uint256 _maxFee) returns()
func (_TokenManager *TokenManagerTransactor) ChangeFee(opts *bind.TransactOpts, tokenAddress common.Address, _equilibriumFee *big.Int, _maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "changeFee", tokenAddress, _equilibriumFee, _maxFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xb5ec9999.
//
// Solidity: function changeFee(address tokenAddress, uint256 _equilibriumFee, uint256 _maxFee) returns()
func (_TokenManager *TokenManagerSession) ChangeFee(tokenAddress common.Address, _equilibriumFee *big.Int, _maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ChangeFee(&_TokenManager.TransactOpts, tokenAddress, _equilibriumFee, _maxFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xb5ec9999.
//
// Solidity: function changeFee(address tokenAddress, uint256 _equilibriumFee, uint256 _maxFee) returns()
func (_TokenManager *TokenManagerTransactorSession) ChangeFee(tokenAddress common.Address, _equilibriumFee *big.Int, _maxFee *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ChangeFee(&_TokenManager.TransactOpts, tokenAddress, _equilibriumFee, _maxFee)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address tokenAddress) returns()
func (_TokenManager *TokenManagerTransactor) RemoveSupportedToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "removeSupportedToken", tokenAddress)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address tokenAddress) returns()
func (_TokenManager *TokenManagerSession) RemoveSupportedToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.RemoveSupportedToken(&_TokenManager.TransactOpts, tokenAddress)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address tokenAddress) returns()
func (_TokenManager *TokenManagerTransactorSession) RemoveSupportedToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.RemoveSupportedToken(&_TokenManager.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// SetTokenTransferOverhead is a paid mutator transaction binding the contract method 0x8b56d0b7.
//
// Solidity: function setTokenTransferOverhead(address tokenAddress, uint256 gasOverhead) returns()
func (_TokenManager *TokenManagerTransactor) SetTokenTransferOverhead(opts *bind.TransactOpts, tokenAddress common.Address, gasOverhead *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setTokenTransferOverhead", tokenAddress, gasOverhead)
}

// SetTokenTransferOverhead is a paid mutator transaction binding the contract method 0x8b56d0b7.
//
// Solidity: function setTokenTransferOverhead(address tokenAddress, uint256 gasOverhead) returns()
func (_TokenManager *TokenManagerSession) SetTokenTransferOverhead(tokenAddress common.Address, gasOverhead *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetTokenTransferOverhead(&_TokenManager.TransactOpts, tokenAddress, gasOverhead)
}

// SetTokenTransferOverhead is a paid mutator transaction binding the contract method 0x8b56d0b7.
//
// Solidity: function setTokenTransferOverhead(address tokenAddress, uint256 gasOverhead) returns()
func (_TokenManager *TokenManagerTransactorSession) SetTokenTransferOverhead(tokenAddress common.Address, gasOverhead *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetTokenTransferOverhead(&_TokenManager.TransactOpts, tokenAddress, gasOverhead)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// UpdateTokenCap is a paid mutator transaction binding the contract method 0x56e3866d.
//
// Solidity: function updateTokenCap(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit) returns()
func (_TokenManager *TokenManagerTransactor) UpdateTokenCap(opts *bind.TransactOpts, tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "updateTokenCap", tokenAddress, minCapLimit, maxCapLimit)
}

// UpdateTokenCap is a paid mutator transaction binding the contract method 0x56e3866d.
//
// Solidity: function updateTokenCap(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit) returns()
func (_TokenManager *TokenManagerSession) UpdateTokenCap(tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.UpdateTokenCap(&_TokenManager.TransactOpts, tokenAddress, minCapLimit, maxCapLimit)
}

// UpdateTokenCap is a paid mutator transaction binding the contract method 0x56e3866d.
//
// Solidity: function updateTokenCap(address tokenAddress, uint256 minCapLimit, uint256 maxCapLimit) returns()
func (_TokenManager *TokenManagerTransactorSession) UpdateTokenCap(tokenAddress common.Address, minCapLimit *big.Int, maxCapLimit *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.UpdateTokenCap(&_TokenManager.TransactOpts, tokenAddress, minCapLimit, maxCapLimit)
}

// TokenManagerFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the TokenManager contract.
type TokenManagerFeeChangedIterator struct {
	Event *TokenManagerFeeChanged // Event containing the contract specifics and raw log

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
func (it *TokenManagerFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerFeeChanged)
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
		it.Event = new(TokenManagerFeeChanged)
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
func (it *TokenManagerFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerFeeChanged represents a FeeChanged event raised by the TokenManager contract.
type TokenManagerFeeChanged struct {
	TokenAddress   common.Address
	EquilibriumFee *big.Int
	MaxFee         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0xf98c81ad0a5eb3551c3561de8dc9d1512e8680fb77425ea138ebfe9a9c0065ff.
//
// Solidity: event FeeChanged(address indexed tokenAddress, uint256 indexed equilibriumFee, uint256 indexed maxFee)
func (_TokenManager *TokenManagerFilterer) FilterFeeChanged(opts *bind.FilterOpts, tokenAddress []common.Address, equilibriumFee []*big.Int, maxFee []*big.Int) (*TokenManagerFeeChangedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var equilibriumFeeRule []interface{}
	for _, equilibriumFeeItem := range equilibriumFee {
		equilibriumFeeRule = append(equilibriumFeeRule, equilibriumFeeItem)
	}
	var maxFeeRule []interface{}
	for _, maxFeeItem := range maxFee {
		maxFeeRule = append(maxFeeRule, maxFeeItem)
	}

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "FeeChanged", tokenAddressRule, equilibriumFeeRule, maxFeeRule)
	if err != nil {
		return nil, err
	}
	return &TokenManagerFeeChangedIterator{contract: _TokenManager.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0xf98c81ad0a5eb3551c3561de8dc9d1512e8680fb77425ea138ebfe9a9c0065ff.
//
// Solidity: event FeeChanged(address indexed tokenAddress, uint256 indexed equilibriumFee, uint256 indexed maxFee)
func (_TokenManager *TokenManagerFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *TokenManagerFeeChanged, tokenAddress []common.Address, equilibriumFee []*big.Int, maxFee []*big.Int) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var equilibriumFeeRule []interface{}
	for _, equilibriumFeeItem := range equilibriumFee {
		equilibriumFeeRule = append(equilibriumFeeRule, equilibriumFeeItem)
	}
	var maxFeeRule []interface{}
	for _, maxFeeItem := range maxFee {
		maxFeeRule = append(maxFeeRule, maxFeeItem)
	}

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "FeeChanged", tokenAddressRule, equilibriumFeeRule, maxFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerFeeChanged)
				if err := _TokenManager.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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

// ParseFeeChanged is a log parse operation binding the contract event 0xf98c81ad0a5eb3551c3561de8dc9d1512e8680fb77425ea138ebfe9a9c0065ff.
//
// Solidity: event FeeChanged(address indexed tokenAddress, uint256 indexed equilibriumFee, uint256 indexed maxFee)
func (_TokenManager *TokenManagerFilterer) ParseFeeChanged(log types.Log) (*TokenManagerFeeChanged, error) {
	event := new(TokenManagerFeeChanged)
	if err := _TokenManager.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenManager contract.
type TokenManagerOwnershipTransferredIterator struct {
	Event *TokenManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerOwnershipTransferred)
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
		it.Event = new(TokenManagerOwnershipTransferred)
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
func (it *TokenManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TokenManager contract.
type TokenManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenManagerOwnershipTransferredIterator{contract: _TokenManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerOwnershipTransferred)
				if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenManager *TokenManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TokenManagerOwnershipTransferred, error) {
	event := new(TokenManagerOwnershipTransferred)
	if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TokenManager contract.
type TokenManagerPausedIterator struct {
	Event *TokenManagerPaused // Event containing the contract specifics and raw log

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
func (it *TokenManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerPaused)
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
		it.Event = new(TokenManagerPaused)
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
func (it *TokenManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerPaused represents a Paused event raised by the TokenManager contract.
type TokenManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenManager *TokenManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenManagerPausedIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenManagerPausedIterator{contract: _TokenManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenManager *TokenManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenManagerPaused) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerPaused)
				if err := _TokenManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenManager *TokenManagerFilterer) ParsePaused(log types.Log) (*TokenManagerPaused, error) {
	event := new(TokenManagerPaused)
	if err := _TokenManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TokenManager contract.
type TokenManagerUnpausedIterator struct {
	Event *TokenManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerUnpaused)
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
		it.Event = new(TokenManagerUnpaused)
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
func (it *TokenManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerUnpaused represents a Unpaused event raised by the TokenManager contract.
type TokenManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenManager *TokenManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenManagerUnpausedIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenManagerUnpausedIterator{contract: _TokenManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenManager *TokenManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerUnpaused)
				if err := _TokenManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenManager *TokenManagerFilterer) ParseUnpaused(log types.Log) (*TokenManagerUnpaused, error) {
	event := new(TokenManagerUnpaused)
	if err := _TokenManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
