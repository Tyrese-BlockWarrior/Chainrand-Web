// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cc

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

// CcMetaData contains all meta data concerning the Cc contract.
var CcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"RandomnessFufilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SeedKeyRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"VerificationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Verified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_seedKeyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_codeHash\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_codeURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_projectURI\",\"type\":\"string\"}],\"name\":\"fakeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_seedKeyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_codeHash\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_codeURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_projectURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"revokeVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_imageURI\",\"type\":\"string\"}],\"name\":\"setImageURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"}],\"name\":\"setMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_projectURI\",\"type\":\"string\"}],\"name\":\"setProjectURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_seedKey\",\"type\":\"string\"}],\"name\":\"setSeedKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"tokenData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"codeURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seedKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"projectURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seedKeyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"codeHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CcABI is the input ABI used to generate the binding from.
// Deprecated: Use CcMetaData.ABI instead.
var CcABI = CcMetaData.ABI

// Cc is an auto generated Go binding around an Ethereum contract.
type Cc struct {
	CcCaller     // Read-only binding to the contract
	CcTransactor // Write-only binding to the contract
	CcFilterer   // Log filterer for contract events
}

// CcCaller is an auto generated read-only Go binding around an Ethereum contract.
type CcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CcSession struct {
	Contract     *Cc               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CcCallerSession struct {
	Contract *CcCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CcTransactorSession struct {
	Contract     *CcTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CcRaw is an auto generated low-level Go binding around an Ethereum contract.
type CcRaw struct {
	Contract *Cc // Generic contract binding to access the raw methods on
}

// CcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CcCallerRaw struct {
	Contract *CcCaller // Generic read-only contract binding to access the raw methods on
}

// CcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CcTransactorRaw struct {
	Contract *CcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCc creates a new instance of Cc, bound to a specific deployed contract.
func NewCc(address common.Address, backend bind.ContractBackend) (*Cc, error) {
	contract, err := bindCc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cc{CcCaller: CcCaller{contract: contract}, CcTransactor: CcTransactor{contract: contract}, CcFilterer: CcFilterer{contract: contract}}, nil
}

// NewCcCaller creates a new read-only instance of Cc, bound to a specific deployed contract.
func NewCcCaller(address common.Address, caller bind.ContractCaller) (*CcCaller, error) {
	contract, err := bindCc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CcCaller{contract: contract}, nil
}

// NewCcTransactor creates a new write-only instance of Cc, bound to a specific deployed contract.
func NewCcTransactor(address common.Address, transactor bind.ContractTransactor) (*CcTransactor, error) {
	contract, err := bindCc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CcTransactor{contract: contract}, nil
}

// NewCcFilterer creates a new log filterer instance of Cc, bound to a specific deployed contract.
func NewCcFilterer(address common.Address, filterer bind.ContractFilterer) (*CcFilterer, error) {
	contract, err := bindCc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CcFilterer{contract: contract}, nil
}

// bindCc binds a generic wrapper to an already deployed contract.
func bindCc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cc *CcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cc.Contract.CcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cc *CcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cc.Contract.CcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cc *CcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cc.Contract.CcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cc *CcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cc *CcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cc *CcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cc.Contract.contract.Transact(opts, method, params...)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Cc *CcCaller) MAXTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "MAX_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Cc *CcSession) MAXTOKENS() (*big.Int, error) {
	return _Cc.Contract.MAXTOKENS(&_Cc.CallOpts)
}

// MAXTOKENS is a free data retrieval call binding the contract method 0xf47c84c5.
//
// Solidity: function MAX_TOKENS() view returns(uint256)
func (_Cc *CcCallerSession) MAXTOKENS() (*big.Int, error) {
	return _Cc.Contract.MAXTOKENS(&_Cc.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cc *CcCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cc *CcSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cc.Contract.BalanceOf(&_Cc.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cc *CcCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cc.Contract.BalanceOf(&_Cc.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cc *CcCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cc *CcSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cc.Contract.GetApproved(&_Cc.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Cc *CcCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Cc.Contract.GetApproved(&_Cc.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cc *CcCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cc *CcSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cc.Contract.IsApprovedForAll(&_Cc.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Cc *CcCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Cc.Contract.IsApprovedForAll(&_Cc.CallOpts, owner, operator)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Cc *CcCaller) MintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "mintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Cc *CcSession) MintFee() (*big.Int, error) {
	return _Cc.Contract.MintFee(&_Cc.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Cc *CcCallerSession) MintFee() (*big.Int, error) {
	return _Cc.Contract.MintFee(&_Cc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cc *CcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cc *CcSession) Name() (string, error) {
	return _Cc.Contract.Name(&_Cc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cc *CcCallerSession) Name() (string, error) {
	return _Cc.Contract.Name(&_Cc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cc *CcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cc *CcSession) Owner() (common.Address, error) {
	return _Cc.Contract.Owner(&_Cc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cc *CcCallerSession) Owner() (common.Address, error) {
	return _Cc.Contract.Owner(&_Cc.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cc *CcCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cc *CcSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cc.Contract.OwnerOf(&_Cc.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Cc *CcCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Cc.Contract.OwnerOf(&_Cc.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cc *CcCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cc *CcSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cc.Contract.SupportsInterface(&_Cc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cc *CcCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cc.Contract.SupportsInterface(&_Cc.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cc *CcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cc *CcSession) Symbol() (string, error) {
	return _Cc.Contract.Symbol(&_Cc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cc *CcCallerSession) Symbol() (string, error) {
	return _Cc.Contract.Symbol(&_Cc.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cc *CcCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cc *CcSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cc.Contract.TokenByIndex(&_Cc.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Cc *CcCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Cc.Contract.TokenByIndex(&_Cc.CallOpts, index)
}

// TokenData is a free data retrieval call binding the contract method 0xb2b6c9a6.
//
// Solidity: function tokenData(uint256[] _tokenIds) view returns(string[])
func (_Cc *CcCaller) TokenData(opts *bind.CallOpts, _tokenIds []*big.Int) ([]string, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokenData", _tokenIds)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TokenData is a free data retrieval call binding the contract method 0xb2b6c9a6.
//
// Solidity: function tokenData(uint256[] _tokenIds) view returns(string[])
func (_Cc *CcSession) TokenData(_tokenIds []*big.Int) ([]string, error) {
	return _Cc.Contract.TokenData(&_Cc.CallOpts, _tokenIds)
}

// TokenData is a free data retrieval call binding the contract method 0xb2b6c9a6.
//
// Solidity: function tokenData(uint256[] _tokenIds) view returns(string[])
func (_Cc *CcCallerSession) TokenData(_tokenIds []*big.Int) ([]string, error) {
	return _Cc.Contract.TokenData(&_Cc.CallOpts, _tokenIds)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cc *CcCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cc *CcSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cc.Contract.TokenOfOwnerByIndex(&_Cc.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Cc *CcCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Cc.Contract.TokenOfOwnerByIndex(&_Cc.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Cc *CcCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Cc *CcSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Cc.Contract.TokenURI(&_Cc.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Cc *CcCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Cc.Contract.TokenURI(&_Cc.CallOpts, _tokenId)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(string name, string codeURI, string seedKey, string imageURI, string projectURI, address minter, uint256 seedKeyHash, uint256 codeHash, uint256 randomness, uint256 paid, bool verified)
func (_Cc *CcCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	CodeURI     string
	SeedKey     string
	ImageURI    string
	ProjectURI  string
	Minter      common.Address
	SeedKeyHash *big.Int
	CodeHash    *big.Int
	Randomness  *big.Int
	Paid        *big.Int
	Verified    bool
}, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokens", arg0)

	outstruct := new(struct {
		Name        string
		CodeURI     string
		SeedKey     string
		ImageURI    string
		ProjectURI  string
		Minter      common.Address
		SeedKeyHash *big.Int
		CodeHash    *big.Int
		Randomness  *big.Int
		Paid        *big.Int
		Verified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CodeURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SeedKey = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ImageURI = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ProjectURI = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Minter = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.SeedKeyHash = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CodeHash = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Randomness = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Paid = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Verified = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(string name, string codeURI, string seedKey, string imageURI, string projectURI, address minter, uint256 seedKeyHash, uint256 codeHash, uint256 randomness, uint256 paid, bool verified)
func (_Cc *CcSession) Tokens(arg0 *big.Int) (struct {
	Name        string
	CodeURI     string
	SeedKey     string
	ImageURI    string
	ProjectURI  string
	Minter      common.Address
	SeedKeyHash *big.Int
	CodeHash    *big.Int
	Randomness  *big.Int
	Paid        *big.Int
	Verified    bool
}, error) {
	return _Cc.Contract.Tokens(&_Cc.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(string name, string codeURI, string seedKey, string imageURI, string projectURI, address minter, uint256 seedKeyHash, uint256 codeHash, uint256 randomness, uint256 paid, bool verified)
func (_Cc *CcCallerSession) Tokens(arg0 *big.Int) (struct {
	Name        string
	CodeURI     string
	SeedKey     string
	ImageURI    string
	ProjectURI  string
	Minter      common.Address
	SeedKeyHash *big.Int
	CodeHash    *big.Int
	Randomness  *big.Int
	Paid        *big.Int
	Verified    bool
}, error) {
	return _Cc.Contract.Tokens(&_Cc.CallOpts, arg0)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Cc *CcCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Cc *CcSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Cc.Contract.TokensOfOwner(&_Cc.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[])
func (_Cc *CcCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Cc.Contract.TokensOfOwner(&_Cc.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cc *CcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cc *CcSession) TotalSupply() (*big.Int, error) {
	return _Cc.Contract.TotalSupply(&_Cc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cc *CcCallerSession) TotalSupply() (*big.Int, error) {
	return _Cc.Contract.TotalSupply(&_Cc.CallOpts)
}

// VrfFee is a free data retrieval call binding the contract method 0x1017507d.
//
// Solidity: function vrfFee() view returns(uint256)
func (_Cc *CcCaller) VrfFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cc.contract.Call(opts, &out, "vrfFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VrfFee is a free data retrieval call binding the contract method 0x1017507d.
//
// Solidity: function vrfFee() view returns(uint256)
func (_Cc *CcSession) VrfFee() (*big.Int, error) {
	return _Cc.Contract.VrfFee(&_Cc.CallOpts)
}

// VrfFee is a free data retrieval call binding the contract method 0x1017507d.
//
// Solidity: function vrfFee() view returns(uint256)
func (_Cc *CcCallerSession) VrfFee() (*big.Int, error) {
	return _Cc.Contract.VrfFee(&_Cc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cc *CcTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cc *CcSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Approve(&_Cc.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Cc *CcTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Approve(&_Cc.TransactOpts, to, tokenId)
}

// FakeMint is a paid mutator transaction binding the contract method 0xe263396f.
//
// Solidity: function fakeMint(uint256 _count, string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) returns()
func (_Cc *CcTransactor) FakeMint(opts *bind.TransactOpts, _count *big.Int, _name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "fakeMint", _count, _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// FakeMint is a paid mutator transaction binding the contract method 0xe263396f.
//
// Solidity: function fakeMint(uint256 _count, string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) returns()
func (_Cc *CcSession) FakeMint(_count *big.Int, _name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.FakeMint(&_Cc.TransactOpts, _count, _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// FakeMint is a paid mutator transaction binding the contract method 0xe263396f.
//
// Solidity: function fakeMint(uint256 _count, string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) returns()
func (_Cc *CcTransactorSession) FakeMint(_count *big.Int, _name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.FakeMint(&_Cc.TransactOpts, _count, _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// Mint is a paid mutator transaction binding the contract method 0x7cd9de40.
//
// Solidity: function mint(string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) payable returns()
func (_Cc *CcTransactor) Mint(opts *bind.TransactOpts, _name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "mint", _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// Mint is a paid mutator transaction binding the contract method 0x7cd9de40.
//
// Solidity: function mint(string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) payable returns()
func (_Cc *CcSession) Mint(_name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.Mint(&_Cc.TransactOpts, _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// Mint is a paid mutator transaction binding the contract method 0x7cd9de40.
//
// Solidity: function mint(string _name, uint256 _seedKeyHash, uint256 _codeHash, string _codeURI, string _imageURI, string _projectURI) payable returns()
func (_Cc *CcTransactorSession) Mint(_name string, _seedKeyHash *big.Int, _codeHash *big.Int, _codeURI string, _imageURI string, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.Mint(&_Cc.TransactOpts, _name, _seedKeyHash, _codeHash, _codeURI, _imageURI, _projectURI)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _tokenId) payable returns()
func (_Cc *CcTransactor) Pay(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "pay", _tokenId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _tokenId) payable returns()
func (_Cc *CcSession) Pay(_tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Pay(&_Cc.TransactOpts, _tokenId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 _tokenId) payable returns()
func (_Cc *CcTransactorSession) Pay(_tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Pay(&_Cc.TransactOpts, _tokenId)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Cc *CcTransactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Cc *CcSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.RawFulfillRandomness(&_Cc.TransactOpts, requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Cc *CcTransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.RawFulfillRandomness(&_Cc.TransactOpts, requestId, randomness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cc *CcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cc *CcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cc.Contract.RenounceOwnership(&_Cc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cc *CcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cc.Contract.RenounceOwnership(&_Cc.TransactOpts)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0xb09adc35.
//
// Solidity: function revokeVerification(uint256[] _tokenIds) returns()
func (_Cc *CcTransactor) RevokeVerification(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "revokeVerification", _tokenIds)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0xb09adc35.
//
// Solidity: function revokeVerification(uint256[] _tokenIds) returns()
func (_Cc *CcSession) RevokeVerification(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.Contract.RevokeVerification(&_Cc.TransactOpts, _tokenIds)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0xb09adc35.
//
// Solidity: function revokeVerification(uint256[] _tokenIds) returns()
func (_Cc *CcTransactorSession) RevokeVerification(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.Contract.RevokeVerification(&_Cc.TransactOpts, _tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.SafeTransferFrom(&_Cc.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.SafeTransferFrom(&_Cc.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cc *CcTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cc *CcSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cc.Contract.SafeTransferFrom0(&_Cc.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Cc *CcTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cc.Contract.SafeTransferFrom0(&_Cc.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cc *CcTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cc *CcSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cc.Contract.SetApprovalForAll(&_Cc.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Cc *CcTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Cc.Contract.SetApprovalForAll(&_Cc.TransactOpts, operator, approved)
}

// SetImageURI is a paid mutator transaction binding the contract method 0x029624e0.
//
// Solidity: function setImageURI(uint256 _tokenId, string _imageURI) returns()
func (_Cc *CcTransactor) SetImageURI(opts *bind.TransactOpts, _tokenId *big.Int, _imageURI string) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "setImageURI", _tokenId, _imageURI)
}

// SetImageURI is a paid mutator transaction binding the contract method 0x029624e0.
//
// Solidity: function setImageURI(uint256 _tokenId, string _imageURI) returns()
func (_Cc *CcSession) SetImageURI(_tokenId *big.Int, _imageURI string) (*types.Transaction, error) {
	return _Cc.Contract.SetImageURI(&_Cc.TransactOpts, _tokenId, _imageURI)
}

// SetImageURI is a paid mutator transaction binding the contract method 0x029624e0.
//
// Solidity: function setImageURI(uint256 _tokenId, string _imageURI) returns()
func (_Cc *CcTransactorSession) SetImageURI(_tokenId *big.Int, _imageURI string) (*types.Transaction, error) {
	return _Cc.Contract.SetImageURI(&_Cc.TransactOpts, _tokenId, _imageURI)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 _mintFee) returns()
func (_Cc *CcTransactor) SetMintFee(opts *bind.TransactOpts, _mintFee *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "setMintFee", _mintFee)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 _mintFee) returns()
func (_Cc *CcSession) SetMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.SetMintFee(&_Cc.TransactOpts, _mintFee)
}

// SetMintFee is a paid mutator transaction binding the contract method 0xeddd0d9c.
//
// Solidity: function setMintFee(uint256 _mintFee) returns()
func (_Cc *CcTransactorSession) SetMintFee(_mintFee *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.SetMintFee(&_Cc.TransactOpts, _mintFee)
}

// SetProjectURI is a paid mutator transaction binding the contract method 0xb9782a19.
//
// Solidity: function setProjectURI(uint256 _tokenId, string _projectURI) returns()
func (_Cc *CcTransactor) SetProjectURI(opts *bind.TransactOpts, _tokenId *big.Int, _projectURI string) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "setProjectURI", _tokenId, _projectURI)
}

// SetProjectURI is a paid mutator transaction binding the contract method 0xb9782a19.
//
// Solidity: function setProjectURI(uint256 _tokenId, string _projectURI) returns()
func (_Cc *CcSession) SetProjectURI(_tokenId *big.Int, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.SetProjectURI(&_Cc.TransactOpts, _tokenId, _projectURI)
}

// SetProjectURI is a paid mutator transaction binding the contract method 0xb9782a19.
//
// Solidity: function setProjectURI(uint256 _tokenId, string _projectURI) returns()
func (_Cc *CcTransactorSession) SetProjectURI(_tokenId *big.Int, _projectURI string) (*types.Transaction, error) {
	return _Cc.Contract.SetProjectURI(&_Cc.TransactOpts, _tokenId, _projectURI)
}

// SetSeedKey is a paid mutator transaction binding the contract method 0x0cf12f0b.
//
// Solidity: function setSeedKey(uint256 _tokenId, string _seedKey) returns()
func (_Cc *CcTransactor) SetSeedKey(opts *bind.TransactOpts, _tokenId *big.Int, _seedKey string) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "setSeedKey", _tokenId, _seedKey)
}

// SetSeedKey is a paid mutator transaction binding the contract method 0x0cf12f0b.
//
// Solidity: function setSeedKey(uint256 _tokenId, string _seedKey) returns()
func (_Cc *CcSession) SetSeedKey(_tokenId *big.Int, _seedKey string) (*types.Transaction, error) {
	return _Cc.Contract.SetSeedKey(&_Cc.TransactOpts, _tokenId, _seedKey)
}

// SetSeedKey is a paid mutator transaction binding the contract method 0x0cf12f0b.
//
// Solidity: function setSeedKey(uint256 _tokenId, string _seedKey) returns()
func (_Cc *CcTransactorSession) SetSeedKey(_tokenId *big.Int, _seedKey string) (*types.Transaction, error) {
	return _Cc.Contract.SetSeedKey(&_Cc.TransactOpts, _tokenId, _seedKey)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.TransferFrom(&_Cc.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Cc *CcTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cc.Contract.TransferFrom(&_Cc.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cc *CcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cc *CcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cc.Contract.TransferOwnership(&_Cc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cc *CcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cc.Contract.TransferOwnership(&_Cc.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0x0ca5fc25.
//
// Solidity: function verify(uint256[] _tokenIds) returns()
func (_Cc *CcTransactor) Verify(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "verify", _tokenIds)
}

// Verify is a paid mutator transaction binding the contract method 0x0ca5fc25.
//
// Solidity: function verify(uint256[] _tokenIds) returns()
func (_Cc *CcSession) Verify(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Verify(&_Cc.TransactOpts, _tokenIds)
}

// Verify is a paid mutator transaction binding the contract method 0x0ca5fc25.
//
// Solidity: function verify(uint256[] _tokenIds) returns()
func (_Cc *CcTransactorSession) Verify(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _Cc.Contract.Verify(&_Cc.TransactOpts, _tokenIds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Cc *CcTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cc.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Cc *CcSession) Withdraw() (*types.Transaction, error) {
	return _Cc.Contract.Withdraw(&_Cc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Cc *CcTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Cc.Contract.Withdraw(&_Cc.TransactOpts)
}

// CcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cc contract.
type CcApprovalIterator struct {
	Event *CcApproval // Event containing the contract specifics and raw log

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
func (it *CcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcApproval)
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
		it.Event = new(CcApproval)
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
func (it *CcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcApproval represents a Approval event raised by the Cc contract.
type CcApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cc *CcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cc.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CcApprovalIterator{contract: _Cc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cc *CcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CcApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cc.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcApproval)
				if err := _Cc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Cc *CcFilterer) ParseApproval(log types.Log) (*CcApproval, error) {
	event := new(CcApproval)
	if err := _Cc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Cc contract.
type CcApprovalForAllIterator struct {
	Event *CcApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CcApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcApprovalForAll)
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
		it.Event = new(CcApprovalForAll)
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
func (it *CcApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcApprovalForAll represents a ApprovalForAll event raised by the Cc contract.
type CcApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cc *CcFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CcApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cc.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CcApprovalForAllIterator{contract: _Cc.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cc *CcFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CcApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Cc.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcApprovalForAll)
				if err := _Cc.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Cc *CcFilterer) ParseApprovalForAll(log types.Log) (*CcApprovalForAll, error) {
	event := new(CcApprovalForAll)
	if err := _Cc.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cc contract.
type CcOwnershipTransferredIterator struct {
	Event *CcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcOwnershipTransferred)
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
		it.Event = new(CcOwnershipTransferred)
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
func (it *CcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcOwnershipTransferred represents a OwnershipTransferred event raised by the Cc contract.
type CcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cc *CcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CcOwnershipTransferredIterator{contract: _Cc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cc *CcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcOwnershipTransferred)
				if err := _Cc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cc *CcFilterer) ParseOwnershipTransferred(log types.Log) (*CcOwnershipTransferred, error) {
	event := new(CcOwnershipTransferred)
	if err := _Cc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the Cc contract.
type CcPaidIterator struct {
	Event *CcPaid // Event containing the contract specifics and raw log

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
func (it *CcPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcPaid)
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
		it.Event = new(CcPaid)
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
func (it *CcPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcPaid represents a Paid event raised by the Cc contract.
type CcPaid struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x581d416ae9dff30c9305c2b35cb09ed5991897ab97804db29ccf92678e953160.
//
// Solidity: event Paid(uint256 _tokenId)
func (_Cc *CcFilterer) FilterPaid(opts *bind.FilterOpts) (*CcPaidIterator, error) {

	logs, sub, err := _Cc.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &CcPaidIterator{contract: _Cc.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x581d416ae9dff30c9305c2b35cb09ed5991897ab97804db29ccf92678e953160.
//
// Solidity: event Paid(uint256 _tokenId)
func (_Cc *CcFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *CcPaid) (event.Subscription, error) {

	logs, sub, err := _Cc.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcPaid)
				if err := _Cc.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x581d416ae9dff30c9305c2b35cb09ed5991897ab97804db29ccf92678e953160.
//
// Solidity: event Paid(uint256 _tokenId)
func (_Cc *CcFilterer) ParsePaid(log types.Log) (*CcPaid, error) {
	event := new(CcPaid)
	if err := _Cc.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcRandomnessFufilledIterator is returned from FilterRandomnessFufilled and is used to iterate over the raw logs and unpacked data for RandomnessFufilled events raised by the Cc contract.
type CcRandomnessFufilledIterator struct {
	Event *CcRandomnessFufilled // Event containing the contract specifics and raw log

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
func (it *CcRandomnessFufilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcRandomnessFufilled)
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
		it.Event = new(CcRandomnessFufilled)
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
func (it *CcRandomnessFufilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcRandomnessFufilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcRandomnessFufilled represents a RandomnessFufilled event raised by the Cc contract.
type CcRandomnessFufilled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRandomnessFufilled is a free log retrieval operation binding the contract event 0x5d4f36b7dd38017cd2516313ad530ff4db44718c48d388657f2e0a2282765064.
//
// Solidity: event RandomnessFufilled(uint256 _tokenId)
func (_Cc *CcFilterer) FilterRandomnessFufilled(opts *bind.FilterOpts) (*CcRandomnessFufilledIterator, error) {

	logs, sub, err := _Cc.contract.FilterLogs(opts, "RandomnessFufilled")
	if err != nil {
		return nil, err
	}
	return &CcRandomnessFufilledIterator{contract: _Cc.contract, event: "RandomnessFufilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessFufilled is a free log subscription operation binding the contract event 0x5d4f36b7dd38017cd2516313ad530ff4db44718c48d388657f2e0a2282765064.
//
// Solidity: event RandomnessFufilled(uint256 _tokenId)
func (_Cc *CcFilterer) WatchRandomnessFufilled(opts *bind.WatchOpts, sink chan<- *CcRandomnessFufilled) (event.Subscription, error) {

	logs, sub, err := _Cc.contract.WatchLogs(opts, "RandomnessFufilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcRandomnessFufilled)
				if err := _Cc.contract.UnpackLog(event, "RandomnessFufilled", log); err != nil {
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

// ParseRandomnessFufilled is a log parse operation binding the contract event 0x5d4f36b7dd38017cd2516313ad530ff4db44718c48d388657f2e0a2282765064.
//
// Solidity: event RandomnessFufilled(uint256 _tokenId)
func (_Cc *CcFilterer) ParseRandomnessFufilled(log types.Log) (*CcRandomnessFufilled, error) {
	event := new(CcRandomnessFufilled)
	if err := _Cc.contract.UnpackLog(event, "RandomnessFufilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcSeedKeyRevealedIterator is returned from FilterSeedKeyRevealed and is used to iterate over the raw logs and unpacked data for SeedKeyRevealed events raised by the Cc contract.
type CcSeedKeyRevealedIterator struct {
	Event *CcSeedKeyRevealed // Event containing the contract specifics and raw log

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
func (it *CcSeedKeyRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcSeedKeyRevealed)
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
		it.Event = new(CcSeedKeyRevealed)
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
func (it *CcSeedKeyRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcSeedKeyRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcSeedKeyRevealed represents a SeedKeyRevealed event raised by the Cc contract.
type CcSeedKeyRevealed struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSeedKeyRevealed is a free log retrieval operation binding the contract event 0xf6413cc198458f655d2e4aa681fb661791b49f70af0ba7cbacc4e2ac652ff8f4.
//
// Solidity: event SeedKeyRevealed(uint256 _tokenId)
func (_Cc *CcFilterer) FilterSeedKeyRevealed(opts *bind.FilterOpts) (*CcSeedKeyRevealedIterator, error) {

	logs, sub, err := _Cc.contract.FilterLogs(opts, "SeedKeyRevealed")
	if err != nil {
		return nil, err
	}
	return &CcSeedKeyRevealedIterator{contract: _Cc.contract, event: "SeedKeyRevealed", logs: logs, sub: sub}, nil
}

// WatchSeedKeyRevealed is a free log subscription operation binding the contract event 0xf6413cc198458f655d2e4aa681fb661791b49f70af0ba7cbacc4e2ac652ff8f4.
//
// Solidity: event SeedKeyRevealed(uint256 _tokenId)
func (_Cc *CcFilterer) WatchSeedKeyRevealed(opts *bind.WatchOpts, sink chan<- *CcSeedKeyRevealed) (event.Subscription, error) {

	logs, sub, err := _Cc.contract.WatchLogs(opts, "SeedKeyRevealed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcSeedKeyRevealed)
				if err := _Cc.contract.UnpackLog(event, "SeedKeyRevealed", log); err != nil {
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

// ParseSeedKeyRevealed is a log parse operation binding the contract event 0xf6413cc198458f655d2e4aa681fb661791b49f70af0ba7cbacc4e2ac652ff8f4.
//
// Solidity: event SeedKeyRevealed(uint256 _tokenId)
func (_Cc *CcFilterer) ParseSeedKeyRevealed(log types.Log) (*CcSeedKeyRevealed, error) {
	event := new(CcSeedKeyRevealed)
	if err := _Cc.contract.UnpackLog(event, "SeedKeyRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cc contract.
type CcTransferIterator struct {
	Event *CcTransfer // Event containing the contract specifics and raw log

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
func (it *CcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcTransfer)
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
		it.Event = new(CcTransfer)
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
func (it *CcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcTransfer represents a Transfer event raised by the Cc contract.
type CcTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cc *CcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cc.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CcTransferIterator{contract: _Cc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cc *CcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CcTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Cc.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcTransfer)
				if err := _Cc.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Cc *CcFilterer) ParseTransfer(log types.Log) (*CcTransfer, error) {
	event := new(CcTransfer)
	if err := _Cc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcVerificationRevokedIterator is returned from FilterVerificationRevoked and is used to iterate over the raw logs and unpacked data for VerificationRevoked events raised by the Cc contract.
type CcVerificationRevokedIterator struct {
	Event *CcVerificationRevoked // Event containing the contract specifics and raw log

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
func (it *CcVerificationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcVerificationRevoked)
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
		it.Event = new(CcVerificationRevoked)
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
func (it *CcVerificationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcVerificationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcVerificationRevoked represents a VerificationRevoked event raised by the Cc contract.
type CcVerificationRevoked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerificationRevoked is a free log retrieval operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 _tokenId)
func (_Cc *CcFilterer) FilterVerificationRevoked(opts *bind.FilterOpts) (*CcVerificationRevokedIterator, error) {

	logs, sub, err := _Cc.contract.FilterLogs(opts, "VerificationRevoked")
	if err != nil {
		return nil, err
	}
	return &CcVerificationRevokedIterator{contract: _Cc.contract, event: "VerificationRevoked", logs: logs, sub: sub}, nil
}

// WatchVerificationRevoked is a free log subscription operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 _tokenId)
func (_Cc *CcFilterer) WatchVerificationRevoked(opts *bind.WatchOpts, sink chan<- *CcVerificationRevoked) (event.Subscription, error) {

	logs, sub, err := _Cc.contract.WatchLogs(opts, "VerificationRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcVerificationRevoked)
				if err := _Cc.contract.UnpackLog(event, "VerificationRevoked", log); err != nil {
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

// ParseVerificationRevoked is a log parse operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 _tokenId)
func (_Cc *CcFilterer) ParseVerificationRevoked(log types.Log) (*CcVerificationRevoked, error) {
	event := new(CcVerificationRevoked)
	if err := _Cc.contract.UnpackLog(event, "VerificationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CcVerifiedIterator is returned from FilterVerified and is used to iterate over the raw logs and unpacked data for Verified events raised by the Cc contract.
type CcVerifiedIterator struct {
	Event *CcVerified // Event containing the contract specifics and raw log

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
func (it *CcVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CcVerified)
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
		it.Event = new(CcVerified)
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
func (it *CcVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CcVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CcVerified represents a Verified event raised by the Cc contract.
type CcVerified struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerified is a free log retrieval operation binding the contract event 0xf786e7f77ede00a02a5464f8f0555798f42ba99a4a920ef2778db8d75e4656f7.
//
// Solidity: event Verified(uint256 _tokenId)
func (_Cc *CcFilterer) FilterVerified(opts *bind.FilterOpts) (*CcVerifiedIterator, error) {

	logs, sub, err := _Cc.contract.FilterLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return &CcVerifiedIterator{contract: _Cc.contract, event: "Verified", logs: logs, sub: sub}, nil
}

// WatchVerified is a free log subscription operation binding the contract event 0xf786e7f77ede00a02a5464f8f0555798f42ba99a4a920ef2778db8d75e4656f7.
//
// Solidity: event Verified(uint256 _tokenId)
func (_Cc *CcFilterer) WatchVerified(opts *bind.WatchOpts, sink chan<- *CcVerified) (event.Subscription, error) {

	logs, sub, err := _Cc.contract.WatchLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CcVerified)
				if err := _Cc.contract.UnpackLog(event, "Verified", log); err != nil {
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

// ParseVerified is a log parse operation binding the contract event 0xf786e7f77ede00a02a5464f8f0555798f42ba99a4a920ef2778db8d75e4656f7.
//
// Solidity: event Verified(uint256 _tokenId)
func (_Cc *CcFilterer) ParseVerified(log types.Log) (*CcVerified, error) {
	event := new(CcVerified)
	if err := _Cc.contract.UnpackLog(event, "Verified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
