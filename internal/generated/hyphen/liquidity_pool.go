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

// LiquidityPoolPermitRequest is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolPermitRequest struct {
	Nonce   *big.Int
	Expiry  *big.Int
	Allowed bool
	V       uint8
	R       [32]byte
	S       [32]byte
}

// LiquidityPoolMetaData contains all meta data concerning the LiquidityPool contract.
var LiquidityPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transferredAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"}],\"name\":\"AssetSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasFeeWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvidersAddress\",\"type\":\"address\"}],\"name\":\"LiquidityProvidersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarderAddress\",\"type\":\"address\"}],\"name\":\"TrustedForwarderChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"changePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"}],\"name\":\"checkHashStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashSendTransaction\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"depositErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gasFeeAccumulated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gasFeeAccumulatedByToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getCurrentLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLiquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutorManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTransferFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"incentivePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executorManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityProviders\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityProviders\",\"outputs\":[{\"internalType\":\"contractILiquidityProviders\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLiquidityPool.PermitRequest\",\"name\":\"permitOptions\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"permitAndDepositErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structLiquidityPool.PermitRequest\",\"name\":\"permitOptions\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"name\":\"permitEIP2612AndDepositErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"depositHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"}],\"name\":\"sendFundsToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"gas\",\"type\":\"uint128\"}],\"name\":\"setBaseGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executorManagerAddress\",\"type\":\"address\"}],\"name\":\"setExecutorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProviders\",\"type\":\"address\"}],\"name\":\"setLiquidityProviders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractITokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawErc20GasFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNativeGasFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityPoolMetaData.ABI instead.
var LiquidityPoolABI = LiquidityPoolMetaData.ABI

// LiquidityPool is an auto generated Go binding around an Ethereum contract.
type LiquidityPool struct {
	LiquidityPoolCaller     // Read-only binding to the contract
	LiquidityPoolTransactor // Write-only binding to the contract
	LiquidityPoolFilterer   // Log filterer for contract events
}

// LiquidityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityPoolSession struct {
	Contract     *LiquidityPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityPoolCallerSession struct {
	Contract *LiquidityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityPoolTransactorSession struct {
	Contract     *LiquidityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityPoolRaw struct {
	Contract *LiquidityPool // Generic contract binding to access the raw methods on
}

// LiquidityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityPoolCallerRaw struct {
	Contract *LiquidityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactorRaw struct {
	Contract *LiquidityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityPool creates a new instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPool(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	contract, err := bindLiquidityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// NewLiquidityPoolCaller creates a new read-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolCaller(address common.Address, caller bind.ContractCaller) (*LiquidityPoolCaller, error) {
	contract, err := bindLiquidityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolCaller{contract: contract}, nil
}

// NewLiquidityPoolTransactor creates a new write-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityPoolTransactor, error) {
	contract, err := bindLiquidityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTransactor{contract: contract}, nil
}

// NewLiquidityPoolFilterer creates a new log filterer instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityPoolFilterer, error) {
	contract, err := bindLiquidityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolFilterer{contract: contract}, nil
}

// bindLiquidityPool binds a generic wrapper to an already deployed contract.
func bindLiquidityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.LiquidityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transact(opts, method, params...)
}

// BaseGas is a free data retrieval call binding the contract method 0x583bbc40.
//
// Solidity: function baseGas() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) BaseGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "baseGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGas is a free data retrieval call binding the contract method 0x583bbc40.
//
// Solidity: function baseGas() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) BaseGas() (*big.Int, error) {
	return _LiquidityPool.Contract.BaseGas(&_LiquidityPool.CallOpts)
}

// BaseGas is a free data retrieval call binding the contract method 0x583bbc40.
//
// Solidity: function baseGas() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) BaseGas() (*big.Int, error) {
	return _LiquidityPool.Contract.BaseGas(&_LiquidityPool.CallOpts)
}

// CheckHashStatus is a free data retrieval call binding the contract method 0x85a25597.
//
// Solidity: function checkHashStatus(address tokenAddress, uint256 amount, address receiver, bytes depositHash) view returns(bytes32 hashSendTransaction, bool status)
func (_LiquidityPool *LiquidityPoolCaller) CheckHashStatus(opts *bind.CallOpts, tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte) (struct {
	HashSendTransaction [32]byte
	Status              bool
}, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "checkHashStatus", tokenAddress, amount, receiver, depositHash)

	outstruct := new(struct {
		HashSendTransaction [32]byte
		Status              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HashSendTransaction = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CheckHashStatus is a free data retrieval call binding the contract method 0x85a25597.
//
// Solidity: function checkHashStatus(address tokenAddress, uint256 amount, address receiver, bytes depositHash) view returns(bytes32 hashSendTransaction, bool status)
func (_LiquidityPool *LiquidityPoolSession) CheckHashStatus(tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte) (struct {
	HashSendTransaction [32]byte
	Status              bool
}, error) {
	return _LiquidityPool.Contract.CheckHashStatus(&_LiquidityPool.CallOpts, tokenAddress, amount, receiver, depositHash)
}

// CheckHashStatus is a free data retrieval call binding the contract method 0x85a25597.
//
// Solidity: function checkHashStatus(address tokenAddress, uint256 amount, address receiver, bytes depositHash) view returns(bytes32 hashSendTransaction, bool status)
func (_LiquidityPool *LiquidityPoolCallerSession) CheckHashStatus(tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte) (struct {
	HashSendTransaction [32]byte
	Status              bool
}, error) {
	return _LiquidityPool.Contract.CheckHashStatus(&_LiquidityPool.CallOpts, tokenAddress, amount, receiver, depositHash)
}

// DepositConfig is a free data retrieval call binding the contract method 0xab882567.
//
// Solidity: function depositConfig(uint256 , address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolCaller) DepositConfig(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "depositConfig", arg0, arg1)

	outstruct := new(struct {
		Min *big.Int
		Max *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Min = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositConfig is a free data retrieval call binding the contract method 0xab882567.
//
// Solidity: function depositConfig(uint256 , address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolSession) DepositConfig(arg0 *big.Int, arg1 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _LiquidityPool.Contract.DepositConfig(&_LiquidityPool.CallOpts, arg0, arg1)
}

// DepositConfig is a free data retrieval call binding the contract method 0xab882567.
//
// Solidity: function depositConfig(uint256 , address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolCallerSession) DepositConfig(arg0 *big.Int, arg1 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _LiquidityPool.Contract.DepositConfig(&_LiquidityPool.CallOpts, arg0, arg1)
}

// GasFeeAccumulated is a free data retrieval call binding the contract method 0xf16c2c25.
//
// Solidity: function gasFeeAccumulated(address , address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) GasFeeAccumulated(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "gasFeeAccumulated", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasFeeAccumulated is a free data retrieval call binding the contract method 0xf16c2c25.
//
// Solidity: function gasFeeAccumulated(address , address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) GasFeeAccumulated(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GasFeeAccumulated(&_LiquidityPool.CallOpts, arg0, arg1)
}

// GasFeeAccumulated is a free data retrieval call binding the contract method 0xf16c2c25.
//
// Solidity: function gasFeeAccumulated(address , address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) GasFeeAccumulated(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GasFeeAccumulated(&_LiquidityPool.CallOpts, arg0, arg1)
}

// GasFeeAccumulatedByToken is a free data retrieval call binding the contract method 0xe06179fe.
//
// Solidity: function gasFeeAccumulatedByToken(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) GasFeeAccumulatedByToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "gasFeeAccumulatedByToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasFeeAccumulatedByToken is a free data retrieval call binding the contract method 0xe06179fe.
//
// Solidity: function gasFeeAccumulatedByToken(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) GasFeeAccumulatedByToken(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GasFeeAccumulatedByToken(&_LiquidityPool.CallOpts, arg0)
}

// GasFeeAccumulatedByToken is a free data retrieval call binding the contract method 0xe06179fe.
//
// Solidity: function gasFeeAccumulatedByToken(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) GasFeeAccumulatedByToken(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GasFeeAccumulatedByToken(&_LiquidityPool.CallOpts, arg0)
}

// GetCurrentLiquidity is a free data retrieval call binding the contract method 0xca2ba943.
//
// Solidity: function getCurrentLiquidity(address tokenAddress) view returns(uint256 currentLiquidity)
func (_LiquidityPool *LiquidityPoolCaller) GetCurrentLiquidity(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "getCurrentLiquidity", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentLiquidity is a free data retrieval call binding the contract method 0xca2ba943.
//
// Solidity: function getCurrentLiquidity(address tokenAddress) view returns(uint256 currentLiquidity)
func (_LiquidityPool *LiquidityPoolSession) GetCurrentLiquidity(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GetCurrentLiquidity(&_LiquidityPool.CallOpts, tokenAddress)
}

// GetCurrentLiquidity is a free data retrieval call binding the contract method 0xca2ba943.
//
// Solidity: function getCurrentLiquidity(address tokenAddress) view returns(uint256 currentLiquidity)
func (_LiquidityPool *LiquidityPoolCallerSession) GetCurrentLiquidity(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GetCurrentLiquidity(&_LiquidityPool.CallOpts, tokenAddress)
}

// GetExecutorManager is a free data retrieval call binding the contract method 0xab1635b7.
//
// Solidity: function getExecutorManager() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) GetExecutorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "getExecutorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutorManager is a free data retrieval call binding the contract method 0xab1635b7.
//
// Solidity: function getExecutorManager() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) GetExecutorManager() (common.Address, error) {
	return _LiquidityPool.Contract.GetExecutorManager(&_LiquidityPool.CallOpts)
}

// GetExecutorManager is a free data retrieval call binding the contract method 0xab1635b7.
//
// Solidity: function getExecutorManager() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) GetExecutorManager() (common.Address, error) {
	return _LiquidityPool.Contract.GetExecutorManager(&_LiquidityPool.CallOpts)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x6e8ae0a3.
//
// Solidity: function getRewardAmount(uint256 amount, address tokenAddress) view returns(uint256 rewardAmount)
func (_LiquidityPool *LiquidityPoolCaller) GetRewardAmount(opts *bind.CallOpts, amount *big.Int, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "getRewardAmount", amount, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardAmount is a free data retrieval call binding the contract method 0x6e8ae0a3.
//
// Solidity: function getRewardAmount(uint256 amount, address tokenAddress) view returns(uint256 rewardAmount)
func (_LiquidityPool *LiquidityPoolSession) GetRewardAmount(amount *big.Int, tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GetRewardAmount(&_LiquidityPool.CallOpts, amount, tokenAddress)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x6e8ae0a3.
//
// Solidity: function getRewardAmount(uint256 amount, address tokenAddress) view returns(uint256 rewardAmount)
func (_LiquidityPool *LiquidityPoolCallerSession) GetRewardAmount(amount *big.Int, tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.GetRewardAmount(&_LiquidityPool.CallOpts, amount, tokenAddress)
}

// GetTransferFee is a free data retrieval call binding the contract method 0xfc42b58f.
//
// Solidity: function getTransferFee(address tokenAddress, uint256 amount) view returns(uint256 fee)
func (_LiquidityPool *LiquidityPoolCaller) GetTransferFee(opts *bind.CallOpts, tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "getTransferFee", tokenAddress, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferFee is a free data retrieval call binding the contract method 0xfc42b58f.
//
// Solidity: function getTransferFee(address tokenAddress, uint256 amount) view returns(uint256 fee)
func (_LiquidityPool *LiquidityPoolSession) GetTransferFee(tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.GetTransferFee(&_LiquidityPool.CallOpts, tokenAddress, amount)
}

// GetTransferFee is a free data retrieval call binding the contract method 0xfc42b58f.
//
// Solidity: function getTransferFee(address tokenAddress, uint256 amount) view returns(uint256 fee)
func (_LiquidityPool *LiquidityPoolCallerSession) GetTransferFee(tokenAddress common.Address, amount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.GetTransferFee(&_LiquidityPool.CallOpts, tokenAddress, amount)
}

// IncentivePool is a free data retrieval call binding the contract method 0xa4479b48.
//
// Solidity: function incentivePool(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) IncentivePool(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "incentivePool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncentivePool is a free data retrieval call binding the contract method 0xa4479b48.
//
// Solidity: function incentivePool(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) IncentivePool(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.IncentivePool(&_LiquidityPool.CallOpts, arg0)
}

// IncentivePool is a free data retrieval call binding the contract method 0xa4479b48.
//
// Solidity: function incentivePool(address ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) IncentivePool(arg0 common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.IncentivePool(&_LiquidityPool.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityPool *LiquidityPoolCaller) IsPauser(opts *bind.CallOpts, pauser common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "isPauser", pauser)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityPool *LiquidityPoolSession) IsPauser(pauser common.Address) (bool, error) {
	return _LiquidityPool.Contract.IsPauser(&_LiquidityPool.CallOpts, pauser)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityPool *LiquidityPoolCallerSession) IsPauser(pauser common.Address) (bool, error) {
	return _LiquidityPool.Contract.IsPauser(&_LiquidityPool.CallOpts, pauser)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityPool *LiquidityPoolCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityPool *LiquidityPoolSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _LiquidityPool.Contract.IsTrustedForwarder(&_LiquidityPool.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityPool *LiquidityPoolCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _LiquidityPool.Contract.IsTrustedForwarder(&_LiquidityPool.CallOpts, forwarder)
}

// LiquidityProviders is a free data retrieval call binding the contract method 0xa2419a6b.
//
// Solidity: function liquidityProviders() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) LiquidityProviders(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "liquidityProviders")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityProviders is a free data retrieval call binding the contract method 0xa2419a6b.
//
// Solidity: function liquidityProviders() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) LiquidityProviders() (common.Address, error) {
	return _LiquidityPool.Contract.LiquidityProviders(&_LiquidityPool.CallOpts)
}

// LiquidityProviders is a free data retrieval call binding the contract method 0xa2419a6b.
//
// Solidity: function liquidityProviders() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) LiquidityProviders() (common.Address, error) {
	return _LiquidityPool.Contract.LiquidityProviders(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityPool *LiquidityPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityPool *LiquidityPoolSession) Paused() (bool, error) {
	return _LiquidityPool.Contract.Paused(&_LiquidityPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityPool *LiquidityPoolCallerSession) Paused() (bool, error) {
	return _LiquidityPool.Contract.Paused(&_LiquidityPool.CallOpts)
}

// ProcessedHash is a free data retrieval call binding the contract method 0x760d098a.
//
// Solidity: function processedHash(bytes32 ) view returns(bool)
func (_LiquidityPool *LiquidityPoolCaller) ProcessedHash(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "processedHash", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedHash is a free data retrieval call binding the contract method 0x760d098a.
//
// Solidity: function processedHash(bytes32 ) view returns(bool)
func (_LiquidityPool *LiquidityPoolSession) ProcessedHash(arg0 [32]byte) (bool, error) {
	return _LiquidityPool.Contract.ProcessedHash(&_LiquidityPool.CallOpts, arg0)
}

// ProcessedHash is a free data retrieval call binding the contract method 0x760d098a.
//
// Solidity: function processedHash(bytes32 ) view returns(bool)
func (_LiquidityPool *LiquidityPoolCallerSession) ProcessedHash(arg0 [32]byte) (bool, error) {
	return _LiquidityPool.Contract.ProcessedHash(&_LiquidityPool.CallOpts, arg0)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) TokenManager() (common.Address, error) {
	return _LiquidityPool.Contract.TokenManager(&_LiquidityPool.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) TokenManager() (common.Address, error) {
	return _LiquidityPool.Contract.TokenManager(&_LiquidityPool.CallOpts)
}

// TransferConfig is a free data retrieval call binding the contract method 0xa53a225f.
//
// Solidity: function transferConfig(address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolCaller) TransferConfig(opts *bind.CallOpts, arg0 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "transferConfig", arg0)

	outstruct := new(struct {
		Min *big.Int
		Max *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Min = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferConfig is a free data retrieval call binding the contract method 0xa53a225f.
//
// Solidity: function transferConfig(address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolSession) TransferConfig(arg0 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _LiquidityPool.Contract.TransferConfig(&_LiquidityPool.CallOpts, arg0)
}

// TransferConfig is a free data retrieval call binding the contract method 0xa53a225f.
//
// Solidity: function transferConfig(address ) view returns(uint256 min, uint256 max)
func (_LiquidityPool *LiquidityPoolCallerSession) TransferConfig(arg0 common.Address) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _LiquidityPool.Contract.TransferConfig(&_LiquidityPool.CallOpts, arg0)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityPool *LiquidityPoolTransactor) ChangePauser(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "changePauser", newPauser)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityPool *LiquidityPoolSession) ChangePauser(newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ChangePauser(&_LiquidityPool.TransactOpts, newPauser)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) ChangePauser(newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ChangePauser(&_LiquidityPool.TransactOpts, newPauser)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0x55d73595.
//
// Solidity: function depositErc20(uint256 toChainId, address tokenAddress, address receiver, uint256 amount, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactor) DepositErc20(opts *bind.TransactOpts, toChainId *big.Int, tokenAddress common.Address, receiver common.Address, amount *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "depositErc20", toChainId, tokenAddress, receiver, amount, tag)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0x55d73595.
//
// Solidity: function depositErc20(uint256 toChainId, address tokenAddress, address receiver, uint256 amount, string tag) returns()
func (_LiquidityPool *LiquidityPoolSession) DepositErc20(toChainId *big.Int, tokenAddress common.Address, receiver common.Address, amount *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DepositErc20(&_LiquidityPool.TransactOpts, toChainId, tokenAddress, receiver, amount, tag)
}

// DepositErc20 is a paid mutator transaction binding the contract method 0x55d73595.
//
// Solidity: function depositErc20(uint256 toChainId, address tokenAddress, address receiver, uint256 amount, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) DepositErc20(toChainId *big.Int, tokenAddress common.Address, receiver common.Address, amount *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DepositErc20(&_LiquidityPool.TransactOpts, toChainId, tokenAddress, receiver, amount, tag)
}

// DepositNative is a paid mutator transaction binding the contract method 0xea368421.
//
// Solidity: function depositNative(address receiver, uint256 toChainId, string tag) payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) DepositNative(opts *bind.TransactOpts, receiver common.Address, toChainId *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "depositNative", receiver, toChainId, tag)
}

// DepositNative is a paid mutator transaction binding the contract method 0xea368421.
//
// Solidity: function depositNative(address receiver, uint256 toChainId, string tag) payable returns()
func (_LiquidityPool *LiquidityPoolSession) DepositNative(receiver common.Address, toChainId *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DepositNative(&_LiquidityPool.TransactOpts, receiver, toChainId, tag)
}

// DepositNative is a paid mutator transaction binding the contract method 0xea368421.
//
// Solidity: function depositNative(address receiver, uint256 toChainId, string tag) payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) DepositNative(receiver common.Address, toChainId *big.Int, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DepositNative(&_LiquidityPool.TransactOpts, receiver, toChainId, tag)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _executorManagerAddress, address _pauser, address _trustedForwarder, address _tokenManager, address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Initialize(opts *bind.TransactOpts, _executorManagerAddress common.Address, _pauser common.Address, _trustedForwarder common.Address, _tokenManager common.Address, _liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "initialize", _executorManagerAddress, _pauser, _trustedForwarder, _tokenManager, _liquidityProviders)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _executorManagerAddress, address _pauser, address _trustedForwarder, address _tokenManager, address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolSession) Initialize(_executorManagerAddress common.Address, _pauser common.Address, _trustedForwarder common.Address, _tokenManager common.Address, _liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Initialize(&_LiquidityPool.TransactOpts, _executorManagerAddress, _pauser, _trustedForwarder, _tokenManager, _liquidityProviders)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _executorManagerAddress, address _pauser, address _trustedForwarder, address _tokenManager, address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Initialize(_executorManagerAddress common.Address, _pauser common.Address, _trustedForwarder common.Address, _tokenManager common.Address, _liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Initialize(&_LiquidityPool.TransactOpts, _executorManagerAddress, _pauser, _trustedForwarder, _tokenManager, _liquidityProviders)
}

// PermitAndDepositErc20 is a paid mutator transaction binding the contract method 0x67924a56.
//
// Solidity: function permitAndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactor) PermitAndDepositErc20(opts *bind.TransactOpts, tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "permitAndDepositErc20", tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// PermitAndDepositErc20 is a paid mutator transaction binding the contract method 0x67924a56.
//
// Solidity: function permitAndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolSession) PermitAndDepositErc20(tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PermitAndDepositErc20(&_LiquidityPool.TransactOpts, tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// PermitAndDepositErc20 is a paid mutator transaction binding the contract method 0x67924a56.
//
// Solidity: function permitAndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) PermitAndDepositErc20(tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PermitAndDepositErc20(&_LiquidityPool.TransactOpts, tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// PermitEIP2612AndDepositErc20 is a paid mutator transaction binding the contract method 0xbf33e812.
//
// Solidity: function permitEIP2612AndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactor) PermitEIP2612AndDepositErc20(opts *bind.TransactOpts, tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "permitEIP2612AndDepositErc20", tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// PermitEIP2612AndDepositErc20 is a paid mutator transaction binding the contract method 0xbf33e812.
//
// Solidity: function permitEIP2612AndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolSession) PermitEIP2612AndDepositErc20(tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PermitEIP2612AndDepositErc20(&_LiquidityPool.TransactOpts, tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// PermitEIP2612AndDepositErc20 is a paid mutator transaction binding the contract method 0xbf33e812.
//
// Solidity: function permitEIP2612AndDepositErc20(address tokenAddress, address receiver, uint256 amount, uint256 toChainId, (uint256,uint256,bool,uint8,bytes32,bytes32) permitOptions, string tag) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) PermitEIP2612AndDepositErc20(tokenAddress common.Address, receiver common.Address, amount *big.Int, toChainId *big.Int, permitOptions LiquidityPoolPermitRequest, tag string) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PermitEIP2612AndDepositErc20(&_LiquidityPool.TransactOpts, tokenAddress, receiver, amount, toChainId, permitOptions, tag)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityPool *LiquidityPoolTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityPool *LiquidityPoolSession) RenouncePauser() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenouncePauser(&_LiquidityPool.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenouncePauser(&_LiquidityPool.TransactOpts)
}

// SendFundsToUser is a paid mutator transaction binding the contract method 0x1ef6626b.
//
// Solidity: function sendFundsToUser(address tokenAddress, uint256 amount, address receiver, bytes depositHash, uint256 tokenGasPrice, uint256 fromChainId) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SendFundsToUser(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte, tokenGasPrice *big.Int, fromChainId *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "sendFundsToUser", tokenAddress, amount, receiver, depositHash, tokenGasPrice, fromChainId)
}

// SendFundsToUser is a paid mutator transaction binding the contract method 0x1ef6626b.
//
// Solidity: function sendFundsToUser(address tokenAddress, uint256 amount, address receiver, bytes depositHash, uint256 tokenGasPrice, uint256 fromChainId) returns()
func (_LiquidityPool *LiquidityPoolSession) SendFundsToUser(tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte, tokenGasPrice *big.Int, fromChainId *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SendFundsToUser(&_LiquidityPool.TransactOpts, tokenAddress, amount, receiver, depositHash, tokenGasPrice, fromChainId)
}

// SendFundsToUser is a paid mutator transaction binding the contract method 0x1ef6626b.
//
// Solidity: function sendFundsToUser(address tokenAddress, uint256 amount, address receiver, bytes depositHash, uint256 tokenGasPrice, uint256 fromChainId) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SendFundsToUser(tokenAddress common.Address, amount *big.Int, receiver common.Address, depositHash []byte, tokenGasPrice *big.Int, fromChainId *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SendFundsToUser(&_LiquidityPool.TransactOpts, tokenAddress, amount, receiver, depositHash, tokenGasPrice, fromChainId)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xccb844e1.
//
// Solidity: function setBaseGas(uint128 gas) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetBaseGas(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setBaseGas", gas)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xccb844e1.
//
// Solidity: function setBaseGas(uint128 gas) returns()
func (_LiquidityPool *LiquidityPoolSession) SetBaseGas(gas *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetBaseGas(&_LiquidityPool.TransactOpts, gas)
}

// SetBaseGas is a paid mutator transaction binding the contract method 0xccb844e1.
//
// Solidity: function setBaseGas(uint128 gas) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetBaseGas(gas *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetBaseGas(&_LiquidityPool.TransactOpts, gas)
}

// SetExecutorManager is a paid mutator transaction binding the contract method 0xef0d4572.
//
// Solidity: function setExecutorManager(address _executorManagerAddress) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetExecutorManager(opts *bind.TransactOpts, _executorManagerAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setExecutorManager", _executorManagerAddress)
}

// SetExecutorManager is a paid mutator transaction binding the contract method 0xef0d4572.
//
// Solidity: function setExecutorManager(address _executorManagerAddress) returns()
func (_LiquidityPool *LiquidityPoolSession) SetExecutorManager(_executorManagerAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetExecutorManager(&_LiquidityPool.TransactOpts, _executorManagerAddress)
}

// SetExecutorManager is a paid mutator transaction binding the contract method 0xef0d4572.
//
// Solidity: function setExecutorManager(address _executorManagerAddress) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetExecutorManager(_executorManagerAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetExecutorManager(&_LiquidityPool.TransactOpts, _executorManagerAddress)
}

// SetLiquidityProviders is a paid mutator transaction binding the contract method 0x59ca6c01.
//
// Solidity: function setLiquidityProviders(address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetLiquidityProviders(opts *bind.TransactOpts, _liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setLiquidityProviders", _liquidityProviders)
}

// SetLiquidityProviders is a paid mutator transaction binding the contract method 0x59ca6c01.
//
// Solidity: function setLiquidityProviders(address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolSession) SetLiquidityProviders(_liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetLiquidityProviders(&_LiquidityPool.TransactOpts, _liquidityProviders)
}

// SetLiquidityProviders is a paid mutator transaction binding the contract method 0x59ca6c01.
//
// Solidity: function setLiquidityProviders(address _liquidityProviders) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetLiquidityProviders(_liquidityProviders common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetLiquidityProviders(&_LiquidityPool.TransactOpts, _liquidityProviders)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address trustedForwarder) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetTrustedForwarder(opts *bind.TransactOpts, trustedForwarder common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setTrustedForwarder", trustedForwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address trustedForwarder) returns()
func (_LiquidityPool *LiquidityPoolSession) SetTrustedForwarder(trustedForwarder common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetTrustedForwarder(&_LiquidityPool.TransactOpts, trustedForwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address trustedForwarder) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetTrustedForwarder(trustedForwarder common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetTrustedForwarder(&_LiquidityPool.TransactOpts, trustedForwarder)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _tokenAddress, address receiver, uint256 _tokenAmount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Transfer(opts *bind.TransactOpts, _tokenAddress common.Address, receiver common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "transfer", _tokenAddress, receiver, _tokenAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _tokenAddress, address receiver, uint256 _tokenAmount) returns()
func (_LiquidityPool *LiquidityPoolSession) Transfer(_tokenAddress common.Address, receiver common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Transfer(&_LiquidityPool.TransactOpts, _tokenAddress, receiver, _tokenAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _tokenAddress, address receiver, uint256 _tokenAmount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Transfer(_tokenAddress common.Address, receiver common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Transfer(&_LiquidityPool.TransactOpts, _tokenAddress, receiver, _tokenAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// WithdrawErc20GasFee is a paid mutator transaction binding the contract method 0xc87e4e25.
//
// Solidity: function withdrawErc20GasFee(address tokenAddress) returns()
func (_LiquidityPool *LiquidityPoolTransactor) WithdrawErc20GasFee(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdrawErc20GasFee", tokenAddress)
}

// WithdrawErc20GasFee is a paid mutator transaction binding the contract method 0xc87e4e25.
//
// Solidity: function withdrawErc20GasFee(address tokenAddress) returns()
func (_LiquidityPool *LiquidityPoolSession) WithdrawErc20GasFee(tokenAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawErc20GasFee(&_LiquidityPool.TransactOpts, tokenAddress)
}

// WithdrawErc20GasFee is a paid mutator transaction binding the contract method 0xc87e4e25.
//
// Solidity: function withdrawErc20GasFee(address tokenAddress) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) WithdrawErc20GasFee(tokenAddress common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawErc20GasFee(&_LiquidityPool.TransactOpts, tokenAddress)
}

// WithdrawNativeGasFee is a paid mutator transaction binding the contract method 0xda6f6794.
//
// Solidity: function withdrawNativeGasFee() returns()
func (_LiquidityPool *LiquidityPoolTransactor) WithdrawNativeGasFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdrawNativeGasFee")
}

// WithdrawNativeGasFee is a paid mutator transaction binding the contract method 0xda6f6794.
//
// Solidity: function withdrawNativeGasFee() returns()
func (_LiquidityPool *LiquidityPoolSession) WithdrawNativeGasFee() (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawNativeGasFee(&_LiquidityPool.TransactOpts)
}

// WithdrawNativeGasFee is a paid mutator transaction binding the contract method 0xda6f6794.
//
// Solidity: function withdrawNativeGasFee() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) WithdrawNativeGasFee() (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawNativeGasFee(&_LiquidityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// LiquidityPoolAssetSentIterator is returned from FilterAssetSent and is used to iterate over the raw logs and unpacked data for AssetSent events raised by the LiquidityPool contract.
type LiquidityPoolAssetSentIterator struct {
	Event *LiquidityPoolAssetSent // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolAssetSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolAssetSent)
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
		it.Event = new(LiquidityPoolAssetSent)
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
func (it *LiquidityPoolAssetSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolAssetSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolAssetSent represents a AssetSent event raised by the LiquidityPool contract.
type LiquidityPoolAssetSent struct {
	Asset             common.Address
	Amount            *big.Int
	TransferredAmount *big.Int
	Target            common.Address
	DepositHash       []byte
	FromChainId       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAssetSent is a free log retrieval operation binding the contract event 0xfa67019f292323b49b589fc709d66c232c7b0ce022f3f32a39af2f91028bbf2c.
//
// Solidity: event AssetSent(address indexed asset, uint256 indexed amount, uint256 indexed transferredAmount, address target, bytes depositHash, uint256 fromChainId)
func (_LiquidityPool *LiquidityPoolFilterer) FilterAssetSent(opts *bind.FilterOpts, asset []common.Address, amount []*big.Int, transferredAmount []*big.Int) (*LiquidityPoolAssetSentIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var transferredAmountRule []interface{}
	for _, transferredAmountItem := range transferredAmount {
		transferredAmountRule = append(transferredAmountRule, transferredAmountItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "AssetSent", assetRule, amountRule, transferredAmountRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolAssetSentIterator{contract: _LiquidityPool.contract, event: "AssetSent", logs: logs, sub: sub}, nil
}

// WatchAssetSent is a free log subscription operation binding the contract event 0xfa67019f292323b49b589fc709d66c232c7b0ce022f3f32a39af2f91028bbf2c.
//
// Solidity: event AssetSent(address indexed asset, uint256 indexed amount, uint256 indexed transferredAmount, address target, bytes depositHash, uint256 fromChainId)
func (_LiquidityPool *LiquidityPoolFilterer) WatchAssetSent(opts *bind.WatchOpts, sink chan<- *LiquidityPoolAssetSent, asset []common.Address, amount []*big.Int, transferredAmount []*big.Int) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var transferredAmountRule []interface{}
	for _, transferredAmountItem := range transferredAmount {
		transferredAmountRule = append(transferredAmountRule, transferredAmountItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "AssetSent", assetRule, amountRule, transferredAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolAssetSent)
				if err := _LiquidityPool.contract.UnpackLog(event, "AssetSent", log); err != nil {
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

// ParseAssetSent is a log parse operation binding the contract event 0xfa67019f292323b49b589fc709d66c232c7b0ce022f3f32a39af2f91028bbf2c.
//
// Solidity: event AssetSent(address indexed asset, uint256 indexed amount, uint256 indexed transferredAmount, address target, bytes depositHash, uint256 fromChainId)
func (_LiquidityPool *LiquidityPoolFilterer) ParseAssetSent(log types.Log) (*LiquidityPoolAssetSent, error) {
	event := new(LiquidityPoolAssetSent)
	if err := _LiquidityPool.contract.UnpackLog(event, "AssetSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LiquidityPool contract.
type LiquidityPoolDepositIterator struct {
	Event *LiquidityPoolDeposit // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolDeposit)
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
		it.Event = new(LiquidityPoolDeposit)
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
func (it *LiquidityPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolDeposit represents a Deposit event raised by the LiquidityPool contract.
type LiquidityPoolDeposit struct {
	From         common.Address
	TokenAddress common.Address
	Receiver     common.Address
	ToChainId    *big.Int
	Amount       *big.Int
	Reward       *big.Int
	Tag          string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x522e11fa05593b306c8df10d2b0b8e01eec48f9d0a9427a7a93f21ff90d66fb1.
//
// Solidity: event Deposit(address indexed from, address indexed tokenAddress, address indexed receiver, uint256 toChainId, uint256 amount, uint256 reward, string tag)
func (_LiquidityPool *LiquidityPoolFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, tokenAddress []common.Address, receiver []common.Address) (*LiquidityPoolDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Deposit", fromRule, tokenAddressRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolDepositIterator{contract: _LiquidityPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x522e11fa05593b306c8df10d2b0b8e01eec48f9d0a9427a7a93f21ff90d66fb1.
//
// Solidity: event Deposit(address indexed from, address indexed tokenAddress, address indexed receiver, uint256 toChainId, uint256 amount, uint256 reward, string tag)
func (_LiquidityPool *LiquidityPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LiquidityPoolDeposit, from []common.Address, tokenAddress []common.Address, receiver []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Deposit", fromRule, tokenAddressRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolDeposit)
				if err := _LiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x522e11fa05593b306c8df10d2b0b8e01eec48f9d0a9427a7a93f21ff90d66fb1.
//
// Solidity: event Deposit(address indexed from, address indexed tokenAddress, address indexed receiver, uint256 toChainId, uint256 amount, uint256 reward, string tag)
func (_LiquidityPool *LiquidityPoolFilterer) ParseDeposit(log types.Log) (*LiquidityPoolDeposit, error) {
	event := new(LiquidityPoolDeposit)
	if err := _LiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the LiquidityPool contract.
type LiquidityPoolEthReceivedIterator struct {
	Event *LiquidityPoolEthReceived // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolEthReceived)
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
		it.Event = new(LiquidityPoolEthReceived)
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
func (it *LiquidityPoolEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolEthReceived represents a EthReceived event raised by the LiquidityPool contract.
type LiquidityPoolEthReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address arg0, uint256 arg1)
func (_LiquidityPool *LiquidityPoolFilterer) FilterEthReceived(opts *bind.FilterOpts) (*LiquidityPoolEthReceivedIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "EthReceived")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolEthReceivedIterator{contract: _LiquidityPool.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address arg0, uint256 arg1)
func (_LiquidityPool *LiquidityPoolFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *LiquidityPoolEthReceived) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "EthReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolEthReceived)
				if err := _LiquidityPool.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address arg0, uint256 arg1)
func (_LiquidityPool *LiquidityPoolFilterer) ParseEthReceived(log types.Log) (*LiquidityPoolEthReceived, error) {
	event := new(LiquidityPoolEthReceived)
	if err := _LiquidityPool.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolGasFeeWithdrawIterator is returned from FilterGasFeeWithdraw and is used to iterate over the raw logs and unpacked data for GasFeeWithdraw events raised by the LiquidityPool contract.
type LiquidityPoolGasFeeWithdrawIterator struct {
	Event *LiquidityPoolGasFeeWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolGasFeeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolGasFeeWithdraw)
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
		it.Event = new(LiquidityPoolGasFeeWithdraw)
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
func (it *LiquidityPoolGasFeeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolGasFeeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolGasFeeWithdraw represents a GasFeeWithdraw event raised by the LiquidityPool contract.
type LiquidityPoolGasFeeWithdraw struct {
	TokenAddress common.Address
	Owner        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasFeeWithdraw is a free log retrieval operation binding the contract event 0xc129cf8730ce23dfea8953a499f9b276ad5777855ff41fafe198ebc2a0d23243.
//
// Solidity: event GasFeeWithdraw(address indexed tokenAddress, address indexed owner, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterGasFeeWithdraw(opts *bind.FilterOpts, tokenAddress []common.Address, owner []common.Address, amount []*big.Int) (*LiquidityPoolGasFeeWithdrawIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "GasFeeWithdraw", tokenAddressRule, ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolGasFeeWithdrawIterator{contract: _LiquidityPool.contract, event: "GasFeeWithdraw", logs: logs, sub: sub}, nil
}

// WatchGasFeeWithdraw is a free log subscription operation binding the contract event 0xc129cf8730ce23dfea8953a499f9b276ad5777855ff41fafe198ebc2a0d23243.
//
// Solidity: event GasFeeWithdraw(address indexed tokenAddress, address indexed owner, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchGasFeeWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidityPoolGasFeeWithdraw, tokenAddress []common.Address, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "GasFeeWithdraw", tokenAddressRule, ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolGasFeeWithdraw)
				if err := _LiquidityPool.contract.UnpackLog(event, "GasFeeWithdraw", log); err != nil {
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

// ParseGasFeeWithdraw is a log parse operation binding the contract event 0xc129cf8730ce23dfea8953a499f9b276ad5777855ff41fafe198ebc2a0d23243.
//
// Solidity: event GasFeeWithdraw(address indexed tokenAddress, address indexed owner, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) ParseGasFeeWithdraw(log types.Log) (*LiquidityPoolGasFeeWithdraw, error) {
	event := new(LiquidityPoolGasFeeWithdraw)
	if err := _LiquidityPool.contract.UnpackLog(event, "GasFeeWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolLiquidityProvidersChangedIterator is returned from FilterLiquidityProvidersChanged and is used to iterate over the raw logs and unpacked data for LiquidityProvidersChanged events raised by the LiquidityPool contract.
type LiquidityPoolLiquidityProvidersChangedIterator struct {
	Event *LiquidityPoolLiquidityProvidersChanged // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolLiquidityProvidersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolLiquidityProvidersChanged)
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
		it.Event = new(LiquidityPoolLiquidityProvidersChanged)
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
func (it *LiquidityPoolLiquidityProvidersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolLiquidityProvidersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolLiquidityProvidersChanged represents a LiquidityProvidersChanged event raised by the LiquidityPool contract.
type LiquidityPoolLiquidityProvidersChanged struct {
	LiquidityProvidersAddress common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLiquidityProvidersChanged is a free log retrieval operation binding the contract event 0x0988d9952b4295aedfb755253e8d1dcebd105e8a5cf7507e419b544da7647066.
//
// Solidity: event LiquidityProvidersChanged(address indexed liquidityProvidersAddress)
func (_LiquidityPool *LiquidityPoolFilterer) FilterLiquidityProvidersChanged(opts *bind.FilterOpts, liquidityProvidersAddress []common.Address) (*LiquidityPoolLiquidityProvidersChangedIterator, error) {

	var liquidityProvidersAddressRule []interface{}
	for _, liquidityProvidersAddressItem := range liquidityProvidersAddress {
		liquidityProvidersAddressRule = append(liquidityProvidersAddressRule, liquidityProvidersAddressItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "LiquidityProvidersChanged", liquidityProvidersAddressRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolLiquidityProvidersChangedIterator{contract: _LiquidityPool.contract, event: "LiquidityProvidersChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityProvidersChanged is a free log subscription operation binding the contract event 0x0988d9952b4295aedfb755253e8d1dcebd105e8a5cf7507e419b544da7647066.
//
// Solidity: event LiquidityProvidersChanged(address indexed liquidityProvidersAddress)
func (_LiquidityPool *LiquidityPoolFilterer) WatchLiquidityProvidersChanged(opts *bind.WatchOpts, sink chan<- *LiquidityPoolLiquidityProvidersChanged, liquidityProvidersAddress []common.Address) (event.Subscription, error) {

	var liquidityProvidersAddressRule []interface{}
	for _, liquidityProvidersAddressItem := range liquidityProvidersAddress {
		liquidityProvidersAddressRule = append(liquidityProvidersAddressRule, liquidityProvidersAddressItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "LiquidityProvidersChanged", liquidityProvidersAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolLiquidityProvidersChanged)
				if err := _LiquidityPool.contract.UnpackLog(event, "LiquidityProvidersChanged", log); err != nil {
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

// ParseLiquidityProvidersChanged is a log parse operation binding the contract event 0x0988d9952b4295aedfb755253e8d1dcebd105e8a5cf7507e419b544da7647066.
//
// Solidity: event LiquidityProvidersChanged(address indexed liquidityProvidersAddress)
func (_LiquidityPool *LiquidityPoolFilterer) ParseLiquidityProvidersChanged(log types.Log) (*LiquidityPoolLiquidityProvidersChanged, error) {
	event := new(LiquidityPoolLiquidityProvidersChanged)
	if err := _LiquidityPool.contract.UnpackLog(event, "LiquidityProvidersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferredIterator struct {
	Event *LiquidityPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolOwnershipTransferred)
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
		it.Event = new(LiquidityPoolOwnershipTransferred)
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
func (it *LiquidityPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolOwnershipTransferredIterator{contract: _LiquidityPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolOwnershipTransferred)
				if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquidityPool *LiquidityPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityPoolOwnershipTransferred, error) {
	event := new(LiquidityPoolOwnershipTransferred)
	if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LiquidityPool contract.
type LiquidityPoolPausedIterator struct {
	Event *LiquidityPoolPaused // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolPaused)
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
		it.Event = new(LiquidityPoolPaused)
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
func (it *LiquidityPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolPaused represents a Paused event raised by the LiquidityPool contract.
type LiquidityPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LiquidityPool *LiquidityPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LiquidityPoolPausedIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolPausedIterator{contract: _LiquidityPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LiquidityPool *LiquidityPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LiquidityPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolPaused)
				if err := _LiquidityPool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LiquidityPool *LiquidityPoolFilterer) ParsePaused(log types.Log) (*LiquidityPoolPaused, error) {
	event := new(LiquidityPoolPaused)
	if err := _LiquidityPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the LiquidityPool contract.
type LiquidityPoolPauserChangedIterator struct {
	Event *LiquidityPoolPauserChanged // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolPauserChanged)
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
		it.Event = new(LiquidityPoolPauserChanged)
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
func (it *LiquidityPoolPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolPauserChanged represents a PauserChanged event raised by the LiquidityPool contract.
type LiquidityPoolPauserChanged struct {
	PreviousPauser common.Address
	NewPauser      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed previousPauser, address indexed newPauser)
func (_LiquidityPool *LiquidityPoolFilterer) FilterPauserChanged(opts *bind.FilterOpts, previousPauser []common.Address, newPauser []common.Address) (*LiquidityPoolPauserChangedIterator, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "PauserChanged", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolPauserChangedIterator{contract: _LiquidityPool.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed previousPauser, address indexed newPauser)
func (_LiquidityPool *LiquidityPoolFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *LiquidityPoolPauserChanged, previousPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "PauserChanged", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolPauserChanged)
				if err := _LiquidityPool.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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

// ParsePauserChanged is a log parse operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed previousPauser, address indexed newPauser)
func (_LiquidityPool *LiquidityPoolFilterer) ParsePauserChanged(log types.Log) (*LiquidityPoolPauserChanged, error) {
	event := new(LiquidityPoolPauserChanged)
	if err := _LiquidityPool.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the LiquidityPool contract.
type LiquidityPoolReceivedIterator struct {
	Event *LiquidityPoolReceived // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolReceived)
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
		it.Event = new(LiquidityPoolReceived)
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
func (it *LiquidityPoolReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolReceived represents a Received event raised by the LiquidityPool contract.
type LiquidityPoolReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed from, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterReceived(opts *bind.FilterOpts, from []common.Address, amount []*big.Int) (*LiquidityPoolReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Received", fromRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolReceivedIterator{contract: _LiquidityPool.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed from, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *LiquidityPoolReceived, from []common.Address, amount []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Received", fromRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolReceived)
				if err := _LiquidityPool.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed from, uint256 indexed amount)
func (_LiquidityPool *LiquidityPoolFilterer) ParseReceived(log types.Log) (*LiquidityPoolReceived, error) {
	event := new(LiquidityPoolReceived)
	if err := _LiquidityPool.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolTrustedForwarderChangedIterator is returned from FilterTrustedForwarderChanged and is used to iterate over the raw logs and unpacked data for TrustedForwarderChanged events raised by the LiquidityPool contract.
type LiquidityPoolTrustedForwarderChangedIterator struct {
	Event *LiquidityPoolTrustedForwarderChanged // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolTrustedForwarderChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolTrustedForwarderChanged)
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
		it.Event = new(LiquidityPoolTrustedForwarderChanged)
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
func (it *LiquidityPoolTrustedForwarderChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolTrustedForwarderChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolTrustedForwarderChanged represents a TrustedForwarderChanged event raised by the LiquidityPool contract.
type LiquidityPoolTrustedForwarderChanged struct {
	ForwarderAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTrustedForwarderChanged is a free log retrieval operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarderAddress)
func (_LiquidityPool *LiquidityPoolFilterer) FilterTrustedForwarderChanged(opts *bind.FilterOpts, forwarderAddress []common.Address) (*LiquidityPoolTrustedForwarderChangedIterator, error) {

	var forwarderAddressRule []interface{}
	for _, forwarderAddressItem := range forwarderAddress {
		forwarderAddressRule = append(forwarderAddressRule, forwarderAddressItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "TrustedForwarderChanged", forwarderAddressRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTrustedForwarderChangedIterator{contract: _LiquidityPool.contract, event: "TrustedForwarderChanged", logs: logs, sub: sub}, nil
}

// WatchTrustedForwarderChanged is a free log subscription operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarderAddress)
func (_LiquidityPool *LiquidityPoolFilterer) WatchTrustedForwarderChanged(opts *bind.WatchOpts, sink chan<- *LiquidityPoolTrustedForwarderChanged, forwarderAddress []common.Address) (event.Subscription, error) {

	var forwarderAddressRule []interface{}
	for _, forwarderAddressItem := range forwarderAddress {
		forwarderAddressRule = append(forwarderAddressRule, forwarderAddressItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "TrustedForwarderChanged", forwarderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolTrustedForwarderChanged)
				if err := _LiquidityPool.contract.UnpackLog(event, "TrustedForwarderChanged", log); err != nil {
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

// ParseTrustedForwarderChanged is a log parse operation binding the contract event 0x871264f4293af7d2865ae7eae628b228f4991c57cb45b39c99f0b774ebe29018.
//
// Solidity: event TrustedForwarderChanged(address indexed forwarderAddress)
func (_LiquidityPool *LiquidityPoolFilterer) ParseTrustedForwarderChanged(log types.Log) (*LiquidityPoolTrustedForwarderChanged, error) {
	event := new(LiquidityPoolTrustedForwarderChanged)
	if err := _LiquidityPool.contract.UnpackLog(event, "TrustedForwarderChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LiquidityPool contract.
type LiquidityPoolUnpausedIterator struct {
	Event *LiquidityPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolUnpaused)
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
		it.Event = new(LiquidityPoolUnpaused)
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
func (it *LiquidityPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolUnpaused represents a Unpaused event raised by the LiquidityPool contract.
type LiquidityPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LiquidityPool *LiquidityPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LiquidityPoolUnpausedIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolUnpausedIterator{contract: _LiquidityPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LiquidityPool *LiquidityPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LiquidityPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolUnpaused)
				if err := _LiquidityPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LiquidityPool *LiquidityPoolFilterer) ParseUnpaused(log types.Log) (*LiquidityPoolUnpaused, error) {
	event := new(LiquidityPoolUnpaused)
	if err := _LiquidityPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
