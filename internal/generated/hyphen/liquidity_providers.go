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

// LiquidityProvidersMetaData contains all meta data concerning the LiquidityProviders contract.
var LiquidityProvidersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesBurnt\",\"type\":\"uint256\"}],\"name\":\"FeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLPFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addTokenLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"changePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"getFeeAccumulatedOnNft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"getSuppliedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getSuppliedLiquidityByToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"getTokenPriceInLPShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTotalLPFeeByToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTotalReserveByToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"increaseNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseTokenLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityPool\",\"type\":\"address\"}],\"name\":\"setLiquidityPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"setLpToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whiteListPeriodManager\",\"type\":\"address\"}],\"name\":\"setWhiteListPeriodManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"sharesToTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalLPFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalSharesMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidityProvidersABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityProvidersMetaData.ABI instead.
var LiquidityProvidersABI = LiquidityProvidersMetaData.ABI

// LiquidityProviders is an auto generated Go binding around an Ethereum contract.
type LiquidityProviders struct {
	LiquidityProvidersCaller     // Read-only binding to the contract
	LiquidityProvidersTransactor // Write-only binding to the contract
	LiquidityProvidersFilterer   // Log filterer for contract events
}

// LiquidityProvidersCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityProvidersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityProvidersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityProvidersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityProvidersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityProvidersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityProvidersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityProvidersSession struct {
	Contract     *LiquidityProviders // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LiquidityProvidersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityProvidersCallerSession struct {
	Contract *LiquidityProvidersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LiquidityProvidersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityProvidersTransactorSession struct {
	Contract     *LiquidityProvidersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LiquidityProvidersRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityProvidersRaw struct {
	Contract *LiquidityProviders // Generic contract binding to access the raw methods on
}

// LiquidityProvidersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityProvidersCallerRaw struct {
	Contract *LiquidityProvidersCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityProvidersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityProvidersTransactorRaw struct {
	Contract *LiquidityProvidersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityProviders creates a new instance of LiquidityProviders, bound to a specific deployed contract.
func NewLiquidityProviders(address common.Address, backend bind.ContractBackend) (*LiquidityProviders, error) {
	contract, err := bindLiquidityProviders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityProviders{LiquidityProvidersCaller: LiquidityProvidersCaller{contract: contract}, LiquidityProvidersTransactor: LiquidityProvidersTransactor{contract: contract}, LiquidityProvidersFilterer: LiquidityProvidersFilterer{contract: contract}}, nil
}

// NewLiquidityProvidersCaller creates a new read-only instance of LiquidityProviders, bound to a specific deployed contract.
func NewLiquidityProvidersCaller(address common.Address, caller bind.ContractCaller) (*LiquidityProvidersCaller, error) {
	contract, err := bindLiquidityProviders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersCaller{contract: contract}, nil
}

// NewLiquidityProvidersTransactor creates a new write-only instance of LiquidityProviders, bound to a specific deployed contract.
func NewLiquidityProvidersTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityProvidersTransactor, error) {
	contract, err := bindLiquidityProviders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersTransactor{contract: contract}, nil
}

// NewLiquidityProvidersFilterer creates a new log filterer instance of LiquidityProviders, bound to a specific deployed contract.
func NewLiquidityProvidersFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityProvidersFilterer, error) {
	contract, err := bindLiquidityProviders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersFilterer{contract: contract}, nil
}

// bindLiquidityProviders binds a generic wrapper to an already deployed contract.
func bindLiquidityProviders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityProvidersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityProviders *LiquidityProvidersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityProviders.Contract.LiquidityProvidersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityProviders *LiquidityProvidersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.LiquidityProvidersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityProviders *LiquidityProvidersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.LiquidityProvidersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityProviders *LiquidityProvidersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityProviders.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityProviders *LiquidityProvidersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityProviders *LiquidityProvidersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.contract.Transact(opts, method, params...)
}

// GetFeeAccumulatedOnNft is a free data retrieval call binding the contract method 0xc47296bd.
//
// Solidity: function getFeeAccumulatedOnNft(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetFeeAccumulatedOnNft(opts *bind.CallOpts, _nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getFeeAccumulatedOnNft", _nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeAccumulatedOnNft is a free data retrieval call binding the contract method 0xc47296bd.
//
// Solidity: function getFeeAccumulatedOnNft(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetFeeAccumulatedOnNft(_nftId *big.Int) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetFeeAccumulatedOnNft(&_LiquidityProviders.CallOpts, _nftId)
}

// GetFeeAccumulatedOnNft is a free data retrieval call binding the contract method 0xc47296bd.
//
// Solidity: function getFeeAccumulatedOnNft(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetFeeAccumulatedOnNft(_nftId *big.Int) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetFeeAccumulatedOnNft(&_LiquidityProviders.CallOpts, _nftId)
}

// GetSuppliedLiquidity is a free data retrieval call binding the contract method 0xd7ac683c.
//
// Solidity: function getSuppliedLiquidity(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetSuppliedLiquidity(opts *bind.CallOpts, _nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getSuppliedLiquidity", _nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuppliedLiquidity is a free data retrieval call binding the contract method 0xd7ac683c.
//
// Solidity: function getSuppliedLiquidity(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetSuppliedLiquidity(_nftId *big.Int) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetSuppliedLiquidity(&_LiquidityProviders.CallOpts, _nftId)
}

// GetSuppliedLiquidity is a free data retrieval call binding the contract method 0xd7ac683c.
//
// Solidity: function getSuppliedLiquidity(uint256 _nftId) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetSuppliedLiquidity(_nftId *big.Int) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetSuppliedLiquidity(&_LiquidityProviders.CallOpts, _nftId)
}

// GetSuppliedLiquidityByToken is a free data retrieval call binding the contract method 0xb3524e30.
//
// Solidity: function getSuppliedLiquidityByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetSuppliedLiquidityByToken(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getSuppliedLiquidityByToken", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSuppliedLiquidityByToken is a free data retrieval call binding the contract method 0xb3524e30.
//
// Solidity: function getSuppliedLiquidityByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetSuppliedLiquidityByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetSuppliedLiquidityByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// GetSuppliedLiquidityByToken is a free data retrieval call binding the contract method 0xb3524e30.
//
// Solidity: function getSuppliedLiquidityByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetSuppliedLiquidityByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetSuppliedLiquidityByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// GetTokenPriceInLPShares is a free data retrieval call binding the contract method 0x52a30127.
//
// Solidity: function getTokenPriceInLPShares(address _baseToken) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetTokenPriceInLPShares(opts *bind.CallOpts, _baseToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getTokenPriceInLPShares", _baseToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPriceInLPShares is a free data retrieval call binding the contract method 0x52a30127.
//
// Solidity: function getTokenPriceInLPShares(address _baseToken) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetTokenPriceInLPShares(_baseToken common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTokenPriceInLPShares(&_LiquidityProviders.CallOpts, _baseToken)
}

// GetTokenPriceInLPShares is a free data retrieval call binding the contract method 0x52a30127.
//
// Solidity: function getTokenPriceInLPShares(address _baseToken) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetTokenPriceInLPShares(_baseToken common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTokenPriceInLPShares(&_LiquidityProviders.CallOpts, _baseToken)
}

// GetTotalLPFeeByToken is a free data retrieval call binding the contract method 0xf0ff372e.
//
// Solidity: function getTotalLPFeeByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetTotalLPFeeByToken(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getTotalLPFeeByToken", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalLPFeeByToken is a free data retrieval call binding the contract method 0xf0ff372e.
//
// Solidity: function getTotalLPFeeByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetTotalLPFeeByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTotalLPFeeByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// GetTotalLPFeeByToken is a free data retrieval call binding the contract method 0xf0ff372e.
//
// Solidity: function getTotalLPFeeByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetTotalLPFeeByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTotalLPFeeByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// GetTotalReserveByToken is a free data retrieval call binding the contract method 0x0102c154.
//
// Solidity: function getTotalReserveByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) GetTotalReserveByToken(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "getTotalReserveByToken", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReserveByToken is a free data retrieval call binding the contract method 0x0102c154.
//
// Solidity: function getTotalReserveByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) GetTotalReserveByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTotalReserveByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// GetTotalReserveByToken is a free data retrieval call binding the contract method 0x0102c154.
//
// Solidity: function getTotalReserveByToken(address tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) GetTotalReserveByToken(tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.GetTotalReserveByToken(&_LiquidityProviders.CallOpts, tokenAddress)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCaller) IsPauser(opts *bind.CallOpts, pauser common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "isPauser", pauser)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersSession) IsPauser(pauser common.Address) (bool, error) {
	return _LiquidityProviders.Contract.IsPauser(&_LiquidityProviders.CallOpts, pauser)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address pauser) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCallerSession) IsPauser(pauser common.Address) (bool, error) {
	return _LiquidityProviders.Contract.IsPauser(&_LiquidityProviders.CallOpts, pauser)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _LiquidityProviders.Contract.IsTrustedForwarder(&_LiquidityProviders.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _LiquidityProviders.Contract.IsTrustedForwarder(&_LiquidityProviders.CallOpts, forwarder)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityProviders *LiquidityProvidersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityProviders *LiquidityProvidersSession) Owner() (common.Address, error) {
	return _LiquidityProviders.Contract.Owner(&_LiquidityProviders.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityProviders *LiquidityProvidersCallerSession) Owner() (common.Address, error) {
	return _LiquidityProviders.Contract.Owner(&_LiquidityProviders.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityProviders *LiquidityProvidersSession) Paused() (bool, error) {
	return _LiquidityProviders.Contract.Paused(&_LiquidityProviders.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LiquidityProviders *LiquidityProvidersCallerSession) Paused() (bool, error) {
	return _LiquidityProviders.Contract.Paused(&_LiquidityProviders.CallOpts)
}

// SharesToTokenAmount is a free data retrieval call binding the contract method 0x4f20f840.
//
// Solidity: function sharesToTokenAmount(uint256 _shares, address _tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) SharesToTokenAmount(opts *bind.CallOpts, _shares *big.Int, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "sharesToTokenAmount", _shares, _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToTokenAmount is a free data retrieval call binding the contract method 0x4f20f840.
//
// Solidity: function sharesToTokenAmount(uint256 _shares, address _tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) SharesToTokenAmount(_shares *big.Int, _tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.SharesToTokenAmount(&_LiquidityProviders.CallOpts, _shares, _tokenAddress)
}

// SharesToTokenAmount is a free data retrieval call binding the contract method 0x4f20f840.
//
// Solidity: function sharesToTokenAmount(uint256 _shares, address _tokenAddress) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) SharesToTokenAmount(_shares *big.Int, _tokenAddress common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.SharesToTokenAmount(&_LiquidityProviders.CallOpts, _shares, _tokenAddress)
}

// TotalLPFees is a free data retrieval call binding the contract method 0xd1b4f192.
//
// Solidity: function totalLPFees(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) TotalLPFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "totalLPFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLPFees is a free data retrieval call binding the contract method 0xd1b4f192.
//
// Solidity: function totalLPFees(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) TotalLPFees(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalLPFees(&_LiquidityProviders.CallOpts, arg0)
}

// TotalLPFees is a free data retrieval call binding the contract method 0xd1b4f192.
//
// Solidity: function totalLPFees(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) TotalLPFees(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalLPFees(&_LiquidityProviders.CallOpts, arg0)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0xc56326de.
//
// Solidity: function totalLiquidity(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) TotalLiquidity(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "totalLiquidity", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidity is a free data retrieval call binding the contract method 0xc56326de.
//
// Solidity: function totalLiquidity(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) TotalLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalLiquidity(&_LiquidityProviders.CallOpts, arg0)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0xc56326de.
//
// Solidity: function totalLiquidity(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) TotalLiquidity(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalLiquidity(&_LiquidityProviders.CallOpts, arg0)
}

// TotalReserve is a free data retrieval call binding the contract method 0x4e7e36e2.
//
// Solidity: function totalReserve(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) TotalReserve(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "totalReserve", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserve is a free data retrieval call binding the contract method 0x4e7e36e2.
//
// Solidity: function totalReserve(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) TotalReserve(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalReserve(&_LiquidityProviders.CallOpts, arg0)
}

// TotalReserve is a free data retrieval call binding the contract method 0x4e7e36e2.
//
// Solidity: function totalReserve(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) TotalReserve(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalReserve(&_LiquidityProviders.CallOpts, arg0)
}

// TotalSharesMinted is a free data retrieval call binding the contract method 0xb0261e53.
//
// Solidity: function totalSharesMinted(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCaller) TotalSharesMinted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityProviders.contract.Call(opts, &out, "totalSharesMinted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSharesMinted is a free data retrieval call binding the contract method 0xb0261e53.
//
// Solidity: function totalSharesMinted(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersSession) TotalSharesMinted(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalSharesMinted(&_LiquidityProviders.CallOpts, arg0)
}

// TotalSharesMinted is a free data retrieval call binding the contract method 0xb0261e53.
//
// Solidity: function totalSharesMinted(address ) view returns(uint256)
func (_LiquidityProviders *LiquidityProvidersCallerSession) TotalSharesMinted(arg0 common.Address) (*big.Int, error) {
	return _LiquidityProviders.Contract.TotalSharesMinted(&_LiquidityProviders.CallOpts, arg0)
}

// AddLPFee is a paid mutator transaction binding the contract method 0x9326b08a.
//
// Solidity: function addLPFee(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) AddLPFee(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "addLPFee", _token, _amount)
}

// AddLPFee is a paid mutator transaction binding the contract method 0x9326b08a.
//
// Solidity: function addLPFee(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersSession) AddLPFee(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddLPFee(&_LiquidityProviders.TransactOpts, _token, _amount)
}

// AddLPFee is a paid mutator transaction binding the contract method 0x9326b08a.
//
// Solidity: function addLPFee(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) AddLPFee(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddLPFee(&_LiquidityProviders.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) AddNativeLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "addNativeLiquidity")
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityProviders *LiquidityProvidersSession) AddNativeLiquidity() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddNativeLiquidity(&_LiquidityProviders.TransactOpts)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0xb238b533.
//
// Solidity: function addNativeLiquidity() payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) AddNativeLiquidity() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddNativeLiquidity(&_LiquidityProviders.TransactOpts)
}

// AddTokenLiquidity is a paid mutator transaction binding the contract method 0x14fe72aa.
//
// Solidity: function addTokenLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) AddTokenLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "addTokenLiquidity", _token, _amount)
}

// AddTokenLiquidity is a paid mutator transaction binding the contract method 0x14fe72aa.
//
// Solidity: function addTokenLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersSession) AddTokenLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddTokenLiquidity(&_LiquidityProviders.TransactOpts, _token, _amount)
}

// AddTokenLiquidity is a paid mutator transaction binding the contract method 0x14fe72aa.
//
// Solidity: function addTokenLiquidity(address _token, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) AddTokenLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.AddTokenLiquidity(&_LiquidityProviders.TransactOpts, _token, _amount)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) ChangePauser(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "changePauser", newPauser)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityProviders *LiquidityProvidersSession) ChangePauser(newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.ChangePauser(&_LiquidityProviders.TransactOpts, newPauser)
}

// ChangePauser is a paid mutator transaction binding the contract method 0x2cd271e7.
//
// Solidity: function changePauser(address newPauser) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) ChangePauser(newPauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.ChangePauser(&_LiquidityProviders.TransactOpts, newPauser)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xf667526a.
//
// Solidity: function claimFee(uint256 _nftId) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) ClaimFee(opts *bind.TransactOpts, _nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "claimFee", _nftId)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xf667526a.
//
// Solidity: function claimFee(uint256 _nftId) returns()
func (_LiquidityProviders *LiquidityProvidersSession) ClaimFee(_nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.ClaimFee(&_LiquidityProviders.TransactOpts, _nftId)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xf667526a.
//
// Solidity: function claimFee(uint256 _nftId) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) ClaimFee(_nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.ClaimFee(&_LiquidityProviders.TransactOpts, _nftId)
}

// IncreaseNativeLiquidity is a paid mutator transaction binding the contract method 0x9500fefc.
//
// Solidity: function increaseNativeLiquidity(uint256 _nftId) payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) IncreaseNativeLiquidity(opts *bind.TransactOpts, _nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "increaseNativeLiquidity", _nftId)
}

// IncreaseNativeLiquidity is a paid mutator transaction binding the contract method 0x9500fefc.
//
// Solidity: function increaseNativeLiquidity(uint256 _nftId) payable returns()
func (_LiquidityProviders *LiquidityProvidersSession) IncreaseNativeLiquidity(_nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.IncreaseNativeLiquidity(&_LiquidityProviders.TransactOpts, _nftId)
}

// IncreaseNativeLiquidity is a paid mutator transaction binding the contract method 0x9500fefc.
//
// Solidity: function increaseNativeLiquidity(uint256 _nftId) payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) IncreaseNativeLiquidity(_nftId *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.IncreaseNativeLiquidity(&_LiquidityProviders.TransactOpts, _nftId)
}

// IncreaseTokenLiquidity is a paid mutator transaction binding the contract method 0x39db7d0e.
//
// Solidity: function increaseTokenLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) IncreaseTokenLiquidity(opts *bind.TransactOpts, _nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "increaseTokenLiquidity", _nftId, _amount)
}

// IncreaseTokenLiquidity is a paid mutator transaction binding the contract method 0x39db7d0e.
//
// Solidity: function increaseTokenLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersSession) IncreaseTokenLiquidity(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.IncreaseTokenLiquidity(&_LiquidityProviders.TransactOpts, _nftId, _amount)
}

// IncreaseTokenLiquidity is a paid mutator transaction binding the contract method 0x39db7d0e.
//
// Solidity: function increaseTokenLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) IncreaseTokenLiquidity(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.IncreaseTokenLiquidity(&_LiquidityProviders.TransactOpts, _nftId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _trustedForwarder, address _lpToken, address _tokenManager, address _pauser) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) Initialize(opts *bind.TransactOpts, _trustedForwarder common.Address, _lpToken common.Address, _tokenManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "initialize", _trustedForwarder, _lpToken, _tokenManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _trustedForwarder, address _lpToken, address _tokenManager, address _pauser) returns()
func (_LiquidityProviders *LiquidityProvidersSession) Initialize(_trustedForwarder common.Address, _lpToken common.Address, _tokenManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.Initialize(&_LiquidityProviders.TransactOpts, _trustedForwarder, _lpToken, _tokenManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _trustedForwarder, address _lpToken, address _tokenManager, address _pauser) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) Initialize(_trustedForwarder common.Address, _lpToken common.Address, _tokenManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.Initialize(&_LiquidityProviders.TransactOpts, _trustedForwarder, _lpToken, _tokenManager, _pauser)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9d7de6b3.
//
// Solidity: function removeLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) RemoveLiquidity(opts *bind.TransactOpts, _nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "removeLiquidity", _nftId, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9d7de6b3.
//
// Solidity: function removeLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersSession) RemoveLiquidity(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RemoveLiquidity(&_LiquidityProviders.TransactOpts, _nftId, _amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9d7de6b3.
//
// Solidity: function removeLiquidity(uint256 _nftId, uint256 _amount) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) RemoveLiquidity(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RemoveLiquidity(&_LiquidityProviders.TransactOpts, _nftId, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityProviders *LiquidityProvidersSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RenounceOwnership(&_LiquidityProviders.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RenounceOwnership(&_LiquidityProviders.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityProviders *LiquidityProvidersSession) RenouncePauser() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RenouncePauser(&_LiquidityProviders.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.RenouncePauser(&_LiquidityProviders.TransactOpts)
}

// SetLiquidityPool is a paid mutator transaction binding the contract method 0x01877020.
//
// Solidity: function setLiquidityPool(address _liquidityPool) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) SetLiquidityPool(opts *bind.TransactOpts, _liquidityPool common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "setLiquidityPool", _liquidityPool)
}

// SetLiquidityPool is a paid mutator transaction binding the contract method 0x01877020.
//
// Solidity: function setLiquidityPool(address _liquidityPool) returns()
func (_LiquidityProviders *LiquidityProvidersSession) SetLiquidityPool(_liquidityPool common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetLiquidityPool(&_LiquidityProviders.TransactOpts, _liquidityPool)
}

// SetLiquidityPool is a paid mutator transaction binding the contract method 0x01877020.
//
// Solidity: function setLiquidityPool(address _liquidityPool) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) SetLiquidityPool(_liquidityPool common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetLiquidityPool(&_LiquidityProviders.TransactOpts, _liquidityPool)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) SetLpToken(opts *bind.TransactOpts, _lpToken common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "setLpToken", _lpToken)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_LiquidityProviders *LiquidityProvidersSession) SetLpToken(_lpToken common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetLpToken(&_LiquidityProviders.TransactOpts, _lpToken)
}

// SetLpToken is a paid mutator transaction binding the contract method 0x9ee933b5.
//
// Solidity: function setLpToken(address _lpToken) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) SetLpToken(_lpToken common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetLpToken(&_LiquidityProviders.TransactOpts, _lpToken)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityProviders *LiquidityProvidersSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetTokenManager(&_LiquidityProviders.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetTokenManager(&_LiquidityProviders.TransactOpts, _tokenManager)
}

// SetWhiteListPeriodManager is a paid mutator transaction binding the contract method 0xe479b099.
//
// Solidity: function setWhiteListPeriodManager(address _whiteListPeriodManager) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) SetWhiteListPeriodManager(opts *bind.TransactOpts, _whiteListPeriodManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "setWhiteListPeriodManager", _whiteListPeriodManager)
}

// SetWhiteListPeriodManager is a paid mutator transaction binding the contract method 0xe479b099.
//
// Solidity: function setWhiteListPeriodManager(address _whiteListPeriodManager) returns()
func (_LiquidityProviders *LiquidityProvidersSession) SetWhiteListPeriodManager(_whiteListPeriodManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetWhiteListPeriodManager(&_LiquidityProviders.TransactOpts, _whiteListPeriodManager)
}

// SetWhiteListPeriodManager is a paid mutator transaction binding the contract method 0xe479b099.
//
// Solidity: function setWhiteListPeriodManager(address _whiteListPeriodManager) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) SetWhiteListPeriodManager(_whiteListPeriodManager common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.SetWhiteListPeriodManager(&_LiquidityProviders.TransactOpts, _whiteListPeriodManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityProviders *LiquidityProvidersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.TransferOwnership(&_LiquidityProviders.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityProviders.Contract.TransferOwnership(&_LiquidityProviders.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityProviders.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityProviders *LiquidityProvidersSession) Receive() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.Receive(&_LiquidityProviders.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityProviders *LiquidityProvidersTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidityProviders.Contract.Receive(&_LiquidityProviders.TransactOpts)
}

// LiquidityProvidersEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the LiquidityProviders contract.
type LiquidityProvidersEthReceivedIterator struct {
	Event *LiquidityProvidersEthReceived // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersEthReceived)
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
		it.Event = new(LiquidityProvidersEthReceived)
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
func (it *LiquidityProvidersEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersEthReceived represents a EthReceived event raised by the LiquidityProviders contract.
type LiquidityProvidersEthReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed sender, uint256 value)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterEthReceived(opts *bind.FilterOpts, sender []common.Address) (*LiquidityProvidersEthReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "EthReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersEthReceivedIterator{contract: _LiquidityProviders.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed sender, uint256 value)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersEthReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "EthReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersEthReceived)
				if err := _LiquidityProviders.contract.UnpackLog(event, "EthReceived", log); err != nil {
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
// Solidity: event EthReceived(address indexed sender, uint256 value)
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseEthReceived(log types.Log) (*LiquidityProvidersEthReceived, error) {
	event := new(LiquidityProvidersEthReceived)
	if err := _LiquidityProviders.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersFeeClaimedIterator is returned from FilterFeeClaimed and is used to iterate over the raw logs and unpacked data for FeeClaimed events raised by the LiquidityProviders contract.
type LiquidityProvidersFeeClaimedIterator struct {
	Event *LiquidityProvidersFeeClaimed // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersFeeClaimed)
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
		it.Event = new(LiquidityProvidersFeeClaimed)
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
func (it *LiquidityProvidersFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersFeeClaimed represents a FeeClaimed event raised by the LiquidityProviders contract.
type LiquidityProvidersFeeClaimed struct {
	TokenAddress common.Address
	Fee          *big.Int
	Lp           common.Address
	SharesBurnt  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeClaimed is a free log retrieval operation binding the contract event 0x8e7bf3ba9a828655d38a8746fe5f122e5b02ada0764a356944630cff7eff35e2.
//
// Solidity: event FeeClaimed(address indexed tokenAddress, uint256 indexed fee, address indexed lp, uint256 sharesBurnt)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterFeeClaimed(opts *bind.FilterOpts, tokenAddress []common.Address, fee []*big.Int, lp []common.Address) (*LiquidityProvidersFeeClaimedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "FeeClaimed", tokenAddressRule, feeRule, lpRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersFeeClaimedIterator{contract: _LiquidityProviders.contract, event: "FeeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeClaimed is a free log subscription operation binding the contract event 0x8e7bf3ba9a828655d38a8746fe5f122e5b02ada0764a356944630cff7eff35e2.
//
// Solidity: event FeeClaimed(address indexed tokenAddress, uint256 indexed fee, address indexed lp, uint256 sharesBurnt)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchFeeClaimed(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersFeeClaimed, tokenAddress []common.Address, fee []*big.Int, lp []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "FeeClaimed", tokenAddressRule, feeRule, lpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersFeeClaimed)
				if err := _LiquidityProviders.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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

// ParseFeeClaimed is a log parse operation binding the contract event 0x8e7bf3ba9a828655d38a8746fe5f122e5b02ada0764a356944630cff7eff35e2.
//
// Solidity: event FeeClaimed(address indexed tokenAddress, uint256 indexed fee, address indexed lp, uint256 sharesBurnt)
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseFeeClaimed(log types.Log) (*LiquidityProvidersFeeClaimed, error) {
	event := new(LiquidityProvidersFeeClaimed)
	if err := _LiquidityProviders.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the LiquidityProviders contract.
type LiquidityProvidersLiquidityAddedIterator struct {
	Event *LiquidityProvidersLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersLiquidityAdded)
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
		it.Event = new(LiquidityProvidersLiquidityAdded)
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
func (it *LiquidityProvidersLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersLiquidityAdded represents a LiquidityAdded event raised by the LiquidityProviders contract.
type LiquidityProvidersLiquidityAdded struct {
	TokenAddress common.Address
	Amount       *big.Int
	Lp           common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xa21288bdd948f634bcd5a8bfc9825db1b01914f370ef82149e123b7c8dc3b65b.
//
// Solidity: event LiquidityAdded(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, tokenAddress []common.Address, amount []*big.Int, lp []common.Address) (*LiquidityProvidersLiquidityAddedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "LiquidityAdded", tokenAddressRule, amountRule, lpRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersLiquidityAddedIterator{contract: _LiquidityProviders.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xa21288bdd948f634bcd5a8bfc9825db1b01914f370ef82149e123b7c8dc3b65b.
//
// Solidity: event LiquidityAdded(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersLiquidityAdded, tokenAddress []common.Address, amount []*big.Int, lp []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "LiquidityAdded", tokenAddressRule, amountRule, lpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersLiquidityAdded)
				if err := _LiquidityProviders.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xa21288bdd948f634bcd5a8bfc9825db1b01914f370ef82149e123b7c8dc3b65b.
//
// Solidity: event LiquidityAdded(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseLiquidityAdded(log types.Log) (*LiquidityProvidersLiquidityAdded, error) {
	event := new(LiquidityProvidersLiquidityAdded)
	if err := _LiquidityProviders.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the LiquidityProviders contract.
type LiquidityProvidersLiquidityRemovedIterator struct {
	Event *LiquidityProvidersLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersLiquidityRemoved)
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
		it.Event = new(LiquidityProvidersLiquidityRemoved)
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
func (it *LiquidityProvidersLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersLiquidityRemoved represents a LiquidityRemoved event raised by the LiquidityProviders contract.
type LiquidityProvidersLiquidityRemoved struct {
	TokenAddress common.Address
	Amount       *big.Int
	Lp           common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x70516e69d9b3069ff3184583d867f7a832772e850ba89b554ae06ff752474f9e.
//
// Solidity: event LiquidityRemoved(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, tokenAddress []common.Address, amount []*big.Int, lp []common.Address) (*LiquidityProvidersLiquidityRemovedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "LiquidityRemoved", tokenAddressRule, amountRule, lpRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersLiquidityRemovedIterator{contract: _LiquidityProviders.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x70516e69d9b3069ff3184583d867f7a832772e850ba89b554ae06ff752474f9e.
//
// Solidity: event LiquidityRemoved(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersLiquidityRemoved, tokenAddress []common.Address, amount []*big.Int, lp []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var lpRule []interface{}
	for _, lpItem := range lp {
		lpRule = append(lpRule, lpItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "LiquidityRemoved", tokenAddressRule, amountRule, lpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersLiquidityRemoved)
				if err := _LiquidityProviders.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0x70516e69d9b3069ff3184583d867f7a832772e850ba89b554ae06ff752474f9e.
//
// Solidity: event LiquidityRemoved(address indexed tokenAddress, uint256 indexed amount, address indexed lp)
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseLiquidityRemoved(log types.Log) (*LiquidityProvidersLiquidityRemoved, error) {
	event := new(LiquidityProvidersLiquidityRemoved)
	if err := _LiquidityProviders.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidityProviders contract.
type LiquidityProvidersOwnershipTransferredIterator struct {
	Event *LiquidityProvidersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersOwnershipTransferred)
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
		it.Event = new(LiquidityProvidersOwnershipTransferred)
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
func (it *LiquidityProvidersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidityProviders contract.
type LiquidityProvidersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityProvidersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersOwnershipTransferredIterator{contract: _LiquidityProviders.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersOwnershipTransferred)
				if err := _LiquidityProviders.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityProvidersOwnershipTransferred, error) {
	event := new(LiquidityProvidersOwnershipTransferred)
	if err := _LiquidityProviders.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LiquidityProviders contract.
type LiquidityProvidersPausedIterator struct {
	Event *LiquidityProvidersPaused // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersPaused)
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
		it.Event = new(LiquidityProvidersPaused)
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
func (it *LiquidityProvidersPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersPaused represents a Paused event raised by the LiquidityProviders contract.
type LiquidityProvidersPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterPaused(opts *bind.FilterOpts) (*LiquidityProvidersPausedIterator, error) {

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersPausedIterator{contract: _LiquidityProviders.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersPaused) (event.Subscription, error) {

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersPaused)
				if err := _LiquidityProviders.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LiquidityProviders *LiquidityProvidersFilterer) ParsePaused(log types.Log) (*LiquidityProvidersPaused, error) {
	event := new(LiquidityProvidersPaused)
	if err := _LiquidityProviders.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the LiquidityProviders contract.
type LiquidityProvidersPauserChangedIterator struct {
	Event *LiquidityProvidersPauserChanged // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersPauserChanged)
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
		it.Event = new(LiquidityProvidersPauserChanged)
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
func (it *LiquidityProvidersPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersPauserChanged represents a PauserChanged event raised by the LiquidityProviders contract.
type LiquidityProvidersPauserChanged struct {
	PreviousPauser common.Address
	NewPauser      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed previousPauser, address indexed newPauser)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterPauserChanged(opts *bind.FilterOpts, previousPauser []common.Address, newPauser []common.Address) (*LiquidityProvidersPauserChangedIterator, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "PauserChanged", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersPauserChangedIterator{contract: _LiquidityProviders.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed previousPauser, address indexed newPauser)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersPauserChanged, previousPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "PauserChanged", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersPauserChanged)
				if err := _LiquidityProviders.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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
func (_LiquidityProviders *LiquidityProvidersFilterer) ParsePauserChanged(log types.Log) (*LiquidityProvidersPauserChanged, error) {
	event := new(LiquidityProvidersPauserChanged)
	if err := _LiquidityProviders.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityProvidersUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LiquidityProviders contract.
type LiquidityProvidersUnpausedIterator struct {
	Event *LiquidityProvidersUnpaused // Event containing the contract specifics and raw log

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
func (it *LiquidityProvidersUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityProvidersUnpaused)
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
		it.Event = new(LiquidityProvidersUnpaused)
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
func (it *LiquidityProvidersUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityProvidersUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityProvidersUnpaused represents a Unpaused event raised by the LiquidityProviders contract.
type LiquidityProvidersUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LiquidityProviders *LiquidityProvidersFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LiquidityProvidersUnpausedIterator, error) {

	logs, sub, err := _LiquidityProviders.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LiquidityProvidersUnpausedIterator{contract: _LiquidityProviders.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LiquidityProviders *LiquidityProvidersFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LiquidityProvidersUnpaused) (event.Subscription, error) {

	logs, sub, err := _LiquidityProviders.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityProvidersUnpaused)
				if err := _LiquidityProviders.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LiquidityProviders *LiquidityProvidersFilterer) ParseUnpaused(log types.Log) (*LiquidityProvidersUnpaused, error) {
	event := new(LiquidityProvidersUnpaused)
	if err := _LiquidityProviders.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
