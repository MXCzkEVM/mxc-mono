// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mxcl2

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
	_ = abi.ConvertType
)

// MxcL2EIP1559Config is an auto generated low-level Go binding around an user-defined struct.
type MxcL2EIP1559Config struct {
	Yscale             *big.Int
	Xscale             uint64
	GasIssuedPerSecond uint64
}

// MxcL2EIP1559Params is an auto generated low-level Go binding around an user-defined struct.
type MxcL2EIP1559Params struct {
	Basefee            uint64
	GasIssuedPerSecond uint64
	GasExcessMax       uint64
	GasTarget          uint64
	Ratio2x1x          uint64
}

// MxcL2MetaData contains all meta data concerning the MxcL2 contract.
var MxcL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"L2_BASEFEE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_1559_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_CHAIN_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_GOLDEN_TOUCH_K\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_SENDER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L2_PUBLIC_INPUT_HASH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_TOO_LATE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"M1559_OUT_OF_STOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"M1559_OUT_OF_STOCK\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"M1559_UNEXPECTED_CHANGE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"actual\",\"type\":\"uint64\"}],\"name\":\"M1559_UNEXPECTED_CHANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"number\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"basefee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gaslimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevrandao\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainid\",\"type\":\"uint32\"}],\"name\":\"Anchored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOLDEN_TOUCH_PRIVATEKEY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1SignalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"parentGasUsed\",\"type\":\"uint64\"}],\"name\":\"anchor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasExcess\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"timeSinceParent\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"parentGasUsed\",\"type\":\"uint64\"}],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEIP1559Config\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"yscale\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"xscale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasIssuedPerSecond\",\"type\":\"uint64\"}],\"internalType\":\"structMxcL2.EIP1559Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"basefee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasIssuedPerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasExcessMax\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ratio2x1x\",\"type\":\"uint64\"}],\"internalType\":\"structMxcL2.EIP1559Params\",\"name\":\"_param1559\",\"type\":\"tuple\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSyncedL1Height\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signAnchor\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MxcL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MxcL2MetaData.ABI instead.
var MxcL2ABI = MxcL2MetaData.ABI

// MxcL2 is an auto generated Go binding around an Ethereum contract.
type MxcL2 struct {
	MxcL2Caller     // Read-only binding to the contract
	MxcL2Transactor // Write-only binding to the contract
	MxcL2Filterer   // Log filterer for contract events
}

// MxcL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MxcL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MxcL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MxcL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MxcL2Session struct {
	Contract     *MxcL2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MxcL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MxcL2CallerSession struct {
	Contract *MxcL2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MxcL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MxcL2TransactorSession struct {
	Contract     *MxcL2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MxcL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MxcL2Raw struct {
	Contract *MxcL2 // Generic contract binding to access the raw methods on
}

// MxcL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MxcL2CallerRaw struct {
	Contract *MxcL2Caller // Generic read-only contract binding to access the raw methods on
}

// MxcL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MxcL2TransactorRaw struct {
	Contract *MxcL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMxcL2 creates a new instance of MxcL2, bound to a specific deployed contract.
func NewMxcL2(address common.Address, backend bind.ContractBackend) (*MxcL2, error) {
	contract, err := bindMxcL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MxcL2{MxcL2Caller: MxcL2Caller{contract: contract}, MxcL2Transactor: MxcL2Transactor{contract: contract}, MxcL2Filterer: MxcL2Filterer{contract: contract}}, nil
}

// NewMxcL2Caller creates a new read-only instance of MxcL2, bound to a specific deployed contract.
func NewMxcL2Caller(address common.Address, caller bind.ContractCaller) (*MxcL2Caller, error) {
	contract, err := bindMxcL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MxcL2Caller{contract: contract}, nil
}

// NewMxcL2Transactor creates a new write-only instance of MxcL2, bound to a specific deployed contract.
func NewMxcL2Transactor(address common.Address, transactor bind.ContractTransactor) (*MxcL2Transactor, error) {
	contract, err := bindMxcL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MxcL2Transactor{contract: contract}, nil
}

// NewMxcL2Filterer creates a new log filterer instance of MxcL2, bound to a specific deployed contract.
func NewMxcL2Filterer(address common.Address, filterer bind.ContractFilterer) (*MxcL2Filterer, error) {
	contract, err := bindMxcL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MxcL2Filterer{contract: contract}, nil
}

// bindMxcL2 binds a generic wrapper to an already deployed contract.
func bindMxcL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MxcL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcL2 *MxcL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcL2.Contract.MxcL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcL2 *MxcL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL2.Contract.MxcL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcL2 *MxcL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcL2.Contract.MxcL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcL2 *MxcL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcL2 *MxcL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcL2 *MxcL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcL2.Contract.contract.Transact(opts, method, params...)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_MxcL2 *MxcL2Caller) GOLDENTOUCHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "GOLDEN_TOUCH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_MxcL2 *MxcL2Session) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _MxcL2.Contract.GOLDENTOUCHADDRESS(&_MxcL2.CallOpts)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_MxcL2 *MxcL2CallerSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _MxcL2.Contract.GOLDENTOUCHADDRESS(&_MxcL2.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_MxcL2 *MxcL2Caller) GOLDENTOUCHPRIVATEKEY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "GOLDEN_TOUCH_PRIVATEKEY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_MxcL2 *MxcL2Session) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _MxcL2.Contract.GOLDENTOUCHPRIVATEKEY(&_MxcL2.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_MxcL2 *MxcL2CallerSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _MxcL2.Contract.GOLDENTOUCHPRIVATEKEY(&_MxcL2.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL2 *MxcL2Caller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL2 *MxcL2Session) AddressManager() (common.Address, error) {
	return _MxcL2.Contract.AddressManager(&_MxcL2.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL2 *MxcL2CallerSession) AddressManager() (common.Address, error) {
	return _MxcL2.Contract.AddressManager(&_MxcL2.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_MxcL2 *MxcL2Caller) GasExcess(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "gasExcess")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_MxcL2 *MxcL2Session) GasExcess() (uint64, error) {
	return _MxcL2.Contract.GasExcess(&_MxcL2.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_MxcL2 *MxcL2CallerSession) GasExcess() (uint64, error) {
	return _MxcL2.Contract.GasExcess(&_MxcL2.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_MxcL2 *MxcL2Caller) GetBasefee(opts *bind.CallOpts, timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "getBasefee", timeSinceParent, gasLimit, parentGasUsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_MxcL2 *MxcL2Session) GetBasefee(timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	return _MxcL2.Contract.GetBasefee(&_MxcL2.CallOpts, timeSinceParent, gasLimit, parentGasUsed)
}

// GetBasefee is a free data retrieval call binding the contract method 0xe1848cb0.
//
// Solidity: function getBasefee(uint32 timeSinceParent, uint64 gasLimit, uint64 parentGasUsed) view returns(uint256 _basefee)
func (_MxcL2 *MxcL2CallerSession) GetBasefee(timeSinceParent uint32, gasLimit uint64, parentGasUsed uint64) (*big.Int, error) {
	return _MxcL2.Contract.GetBasefee(&_MxcL2.CallOpts, timeSinceParent, gasLimit, parentGasUsed)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Caller) GetBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "getBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Session) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetBlockHash(&_MxcL2.CallOpts, number)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2CallerSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetBlockHash(&_MxcL2.CallOpts, number)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Caller) GetCrossChainBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "getCrossChainBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Session) GetCrossChainBlockHash(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetCrossChainBlockHash(&_MxcL2.CallOpts, number)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2CallerSession) GetCrossChainBlockHash(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetCrossChainBlockHash(&_MxcL2.CallOpts, number)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Caller) GetCrossChainSignalRoot(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "getCrossChainSignalRoot", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2Session) GetCrossChainSignalRoot(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetCrossChainSignalRoot(&_MxcL2.CallOpts, number)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 number) view returns(bytes32)
func (_MxcL2 *MxcL2CallerSession) GetCrossChainSignalRoot(number *big.Int) ([32]byte, error) {
	return _MxcL2.Contract.GetCrossChainSignalRoot(&_MxcL2.CallOpts, number)
}

// GetEIP1559Config is a free data retrieval call binding the contract method 0x4e755573.
//
// Solidity: function getEIP1559Config() view returns((uint128,uint64,uint64))
func (_MxcL2 *MxcL2Caller) GetEIP1559Config(opts *bind.CallOpts) (MxcL2EIP1559Config, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "getEIP1559Config")

	if err != nil {
		return *new(MxcL2EIP1559Config), err
	}

	out0 := *abi.ConvertType(out[0], new(MxcL2EIP1559Config)).(*MxcL2EIP1559Config)

	return out0, err

}

// GetEIP1559Config is a free data retrieval call binding the contract method 0x4e755573.
//
// Solidity: function getEIP1559Config() view returns((uint128,uint64,uint64))
func (_MxcL2 *MxcL2Session) GetEIP1559Config() (MxcL2EIP1559Config, error) {
	return _MxcL2.Contract.GetEIP1559Config(&_MxcL2.CallOpts)
}

// GetEIP1559Config is a free data retrieval call binding the contract method 0x4e755573.
//
// Solidity: function getEIP1559Config() view returns((uint128,uint64,uint64))
func (_MxcL2 *MxcL2CallerSession) GetEIP1559Config() (MxcL2EIP1559Config, error) {
	return _MxcL2.Contract.GetEIP1559Config(&_MxcL2.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_MxcL2 *MxcL2Caller) LatestSyncedL1Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "latestSyncedL1Height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_MxcL2 *MxcL2Session) LatestSyncedL1Height() (uint64, error) {
	return _MxcL2.Contract.LatestSyncedL1Height(&_MxcL2.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_MxcL2 *MxcL2CallerSession) LatestSyncedL1Height() (uint64, error) {
	return _MxcL2.Contract.LatestSyncedL1Height(&_MxcL2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL2 *MxcL2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL2 *MxcL2Session) Owner() (common.Address, error) {
	return _MxcL2.Contract.Owner(&_MxcL2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL2 *MxcL2CallerSession) Owner() (common.Address, error) {
	return _MxcL2.Contract.Owner(&_MxcL2.CallOpts)
}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_MxcL2 *MxcL2Caller) ParentTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "parentTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_MxcL2 *MxcL2Session) ParentTimestamp() (uint64, error) {
	return _MxcL2.Contract.ParentTimestamp(&_MxcL2.CallOpts)
}

// ParentTimestamp is a free data retrieval call binding the contract method 0x539b8ade.
//
// Solidity: function parentTimestamp() view returns(uint64)
func (_MxcL2 *MxcL2CallerSession) ParentTimestamp() (uint64, error) {
	return _MxcL2.Contract.ParentTimestamp(&_MxcL2.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_MxcL2 *MxcL2Caller) PublicInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "publicInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_MxcL2 *MxcL2Session) PublicInputHash() ([32]byte, error) {
	return _MxcL2.Contract.PublicInputHash(&_MxcL2.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_MxcL2 *MxcL2CallerSession) PublicInputHash() ([32]byte, error) {
	return _MxcL2.Contract.PublicInputHash(&_MxcL2.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2Caller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2Session) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL2.Contract.Resolve(&_MxcL2.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2CallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL2.Contract.Resolve(&_MxcL2.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2Caller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2Session) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL2.Contract.Resolve0(&_MxcL2.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL2 *MxcL2CallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL2.Contract.Resolve0(&_MxcL2.CallOpts, name, allowZeroAddress)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MxcL2 *MxcL2Caller) SignAnchor(opts *bind.CallOpts, digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _MxcL2.contract.Call(opts, &out, "signAnchor", digest, k)

	outstruct := new(struct {
		V uint8
		R *big.Int
		S *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MxcL2 *MxcL2Session) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _MxcL2.Contract.SignAnchor(&_MxcL2.CallOpts, digest, k)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MxcL2 *MxcL2CallerSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _MxcL2.Contract.SignAnchor(&_MxcL2.CallOpts, digest, k)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_MxcL2 *MxcL2Transactor) Anchor(opts *bind.TransactOpts, l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _MxcL2.contract.Transact(opts, "anchor", l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_MxcL2 *MxcL2Session) Anchor(l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _MxcL2.Contract.Anchor(&_MxcL2.TransactOpts, l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0x3d384a4b.
//
// Solidity: function anchor(bytes32 l1Hash, bytes32 l1SignalRoot, uint64 l1Height, uint64 parentGasUsed) returns()
func (_MxcL2 *MxcL2TransactorSession) Anchor(l1Hash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint64) (*types.Transaction, error) {
	return _MxcL2.Contract.Anchor(&_MxcL2.TransactOpts, l1Hash, l1SignalRoot, l1Height, parentGasUsed)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_MxcL2 *MxcL2Transactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _param1559 MxcL2EIP1559Params) (*types.Transaction, error) {
	return _MxcL2.contract.Transact(opts, "init", _addressManager, _param1559)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_MxcL2 *MxcL2Session) Init(_addressManager common.Address, _param1559 MxcL2EIP1559Params) (*types.Transaction, error) {
	return _MxcL2.Contract.Init(&_MxcL2.TransactOpts, _addressManager, _param1559)
}

// Init is a paid mutator transaction binding the contract method 0x8f3ca30d.
//
// Solidity: function init(address _addressManager, (uint64,uint64,uint64,uint64,uint64) _param1559) returns()
func (_MxcL2 *MxcL2TransactorSession) Init(_addressManager common.Address, _param1559 MxcL2EIP1559Params) (*types.Transaction, error) {
	return _MxcL2.Contract.Init(&_MxcL2.TransactOpts, _addressManager, _param1559)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL2 *MxcL2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL2 *MxcL2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MxcL2.Contract.RenounceOwnership(&_MxcL2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL2 *MxcL2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MxcL2.Contract.RenounceOwnership(&_MxcL2.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL2 *MxcL2Transactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL2.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL2 *MxcL2Session) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL2.Contract.SetAddressManager(&_MxcL2.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL2 *MxcL2TransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL2.Contract.SetAddressManager(&_MxcL2.TransactOpts, newAddressManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL2 *MxcL2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MxcL2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL2 *MxcL2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcL2.Contract.TransferOwnership(&_MxcL2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL2 *MxcL2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcL2.Contract.TransferOwnership(&_MxcL2.TransactOpts, newOwner)
}

// MxcL2AddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the MxcL2 contract.
type MxcL2AddressManagerChangedIterator struct {
	Event *MxcL2AddressManagerChanged // Event containing the contract specifics and raw log

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
func (it *MxcL2AddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL2AddressManagerChanged)
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
		it.Event = new(MxcL2AddressManagerChanged)
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
func (it *MxcL2AddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL2AddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL2AddressManagerChanged represents a AddressManagerChanged event raised by the MxcL2 contract.
type MxcL2AddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcL2 *MxcL2Filterer) FilterAddressManagerChanged(opts *bind.FilterOpts) (*MxcL2AddressManagerChangedIterator, error) {

	logs, sub, err := _MxcL2.contract.FilterLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return &MxcL2AddressManagerChangedIterator{contract: _MxcL2.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcL2 *MxcL2Filterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *MxcL2AddressManagerChanged) (event.Subscription, error) {

	logs, sub, err := _MxcL2.contract.WatchLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL2AddressManagerChanged)
				if err := _MxcL2.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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

// ParseAddressManagerChanged is a log parse operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcL2 *MxcL2Filterer) ParseAddressManagerChanged(log types.Log) (*MxcL2AddressManagerChanged, error) {
	event := new(MxcL2AddressManagerChanged)
	if err := _MxcL2.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL2AnchoredIterator is returned from FilterAnchored and is used to iterate over the raw logs and unpacked data for Anchored events raised by the MxcL2 contract.
type MxcL2AnchoredIterator struct {
	Event *MxcL2Anchored // Event containing the contract specifics and raw log

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
func (it *MxcL2AnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL2Anchored)
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
		it.Event = new(MxcL2Anchored)
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
func (it *MxcL2AnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL2AnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL2Anchored represents a Anchored event raised by the MxcL2 contract.
type MxcL2Anchored struct {
	Number     uint64
	Basefee    uint64
	Gaslimit   uint64
	Timestamp  uint64
	ParentHash [32]byte
	Prevrandao *big.Int
	Coinbase   common.Address
	Chainid    uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAnchored is a free log retrieval operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_MxcL2 *MxcL2Filterer) FilterAnchored(opts *bind.FilterOpts) (*MxcL2AnchoredIterator, error) {

	logs, sub, err := _MxcL2.contract.FilterLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return &MxcL2AnchoredIterator{contract: _MxcL2.contract, event: "Anchored", logs: logs, sub: sub}, nil
}

// WatchAnchored is a free log subscription operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_MxcL2 *MxcL2Filterer) WatchAnchored(opts *bind.WatchOpts, sink chan<- *MxcL2Anchored) (event.Subscription, error) {

	logs, sub, err := _MxcL2.contract.WatchLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL2Anchored)
				if err := _MxcL2.contract.UnpackLog(event, "Anchored", log); err != nil {
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

// ParseAnchored is a log parse operation binding the contract event 0x4dcb01f99c4a2c27a16ab38d00ec92434f8231be81fa62e058f260d3c7156029.
//
// Solidity: event Anchored(uint64 number, uint64 basefee, uint64 gaslimit, uint64 timestamp, bytes32 parentHash, uint256 prevrandao, address coinbase, uint32 chainid)
func (_MxcL2 *MxcL2Filterer) ParseAnchored(log types.Log) (*MxcL2Anchored, error) {
	event := new(MxcL2Anchored)
	if err := _MxcL2.contract.UnpackLog(event, "Anchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL2CrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the MxcL2 contract.
type MxcL2CrossChainSyncedIterator struct {
	Event *MxcL2CrossChainSynced // Event containing the contract specifics and raw log

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
func (it *MxcL2CrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL2CrossChainSynced)
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
		it.Event = new(MxcL2CrossChainSynced)
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
func (it *MxcL2CrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL2CrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL2CrossChainSynced represents a CrossChainSynced event raised by the MxcL2 contract.
type MxcL2CrossChainSynced struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL2 *MxcL2Filterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*MxcL2CrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MxcL2.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &MxcL2CrossChainSyncedIterator{contract: _MxcL2.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL2 *MxcL2Filterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *MxcL2CrossChainSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MxcL2.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL2CrossChainSynced)
				if err := _MxcL2.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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

// ParseCrossChainSynced is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL2 *MxcL2Filterer) ParseCrossChainSynced(log types.Log) (*MxcL2CrossChainSynced, error) {
	event := new(MxcL2CrossChainSynced)
	if err := _MxcL2.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MxcL2 contract.
type MxcL2InitializedIterator struct {
	Event *MxcL2Initialized // Event containing the contract specifics and raw log

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
func (it *MxcL2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL2Initialized)
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
		it.Event = new(MxcL2Initialized)
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
func (it *MxcL2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL2Initialized represents a Initialized event raised by the MxcL2 contract.
type MxcL2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcL2 *MxcL2Filterer) FilterInitialized(opts *bind.FilterOpts) (*MxcL2InitializedIterator, error) {

	logs, sub, err := _MxcL2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MxcL2InitializedIterator{contract: _MxcL2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcL2 *MxcL2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MxcL2Initialized) (event.Subscription, error) {

	logs, sub, err := _MxcL2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL2Initialized)
				if err := _MxcL2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcL2 *MxcL2Filterer) ParseInitialized(log types.Log) (*MxcL2Initialized, error) {
	event := new(MxcL2Initialized)
	if err := _MxcL2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MxcL2 contract.
type MxcL2OwnershipTransferredIterator struct {
	Event *MxcL2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MxcL2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL2OwnershipTransferred)
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
		it.Event = new(MxcL2OwnershipTransferred)
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
func (it *MxcL2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL2OwnershipTransferred represents a OwnershipTransferred event raised by the MxcL2 contract.
type MxcL2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcL2 *MxcL2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MxcL2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcL2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MxcL2OwnershipTransferredIterator{contract: _MxcL2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcL2 *MxcL2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MxcL2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcL2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL2OwnershipTransferred)
				if err := _MxcL2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MxcL2 *MxcL2Filterer) ParseOwnershipTransferred(log types.Log) (*MxcL2OwnershipTransferred, error) {
	event := new(MxcL2OwnershipTransferred)
	if err := _MxcL2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}