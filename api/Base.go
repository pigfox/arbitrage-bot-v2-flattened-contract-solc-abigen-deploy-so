// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// BaseMetaData contains all meta data concerning the Base contract.
var BaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"Adding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delployedAt\",\"type\":\"address\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"q\",\"type\":\"uint256\"}],\"name\":\"SettingQ\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_q\",\"type\":\"uint256\"}],\"name\":\"setQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052600360809081526261626360e81b60a0526001906100229082610104565b5034801561002e575f80fd5b50604080513381523060208201527f587ece4cd19692c5be1a4184503d607d45542d2aca0698c0068f52e09ccb541c910160405180910390a16101be565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061009457607f821691505b6020821081036100b257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100ff57805f5260205f20601f840160051c810160208510156100dd5750805b601f840160051c820191505b818110156100fc575f81556001016100e9565b50505b505050565b81516001600160401b0381111561011d5761011d61006c565b6101318161012b8454610080565b846100b8565b6020601f821160018114610163575f831561014c5750848201515b5f19600385901b1c1916600184901b1784556100fc565b5f84815260208120601f198516915b828110156101925787850151825560209485019460019092019101610172565b50848210156101af57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6103a7806101cb5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063771602f71461004e5780637881c417146100745780638f2d812814610094578063e8e4197a146100a9575b5f80fd5b61006161005c366004610161565b6100b0565b6040519081526020015b60405180910390f35b610087610082366004610195565b6100fe565b60405161006b9190610248565b6100a76100a236600461027d565b61012a565b005b5f54610061565b60408051838152602081018390525f917f8cee9bf80efda447f9d9d047e2e937b140e1c653cb2af0f39d1f6c52a1ffd7dc910160405180910390a16100f58284610294565b90505b92915050565b60606001826040516020016101149291906102ca565b6040516020818303038152906040529050919050565b6040518181527f4af6c9c7470c3b6ab60fab4fef0d6053a5a05afd3df0b7be9243f972be4eb66b9060200160405180910390a15f55565b5f8060408385031215610172575f80fd5b50508035926020909101359150565b634e487b7160e01b5f52604160045260245ffd5b5f602082840312156101a5575f80fd5b813567ffffffffffffffff8111156101bb575f80fd5b8201601f810184136101cb575f80fd5b803567ffffffffffffffff8111156101e5576101e5610181565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561021457610214610181565b60405281815282820160200186101561022b575f80fd5b816020840160208301375f91810160200191909152949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f6020828403121561028d575f80fd5b5035919050565b808201808211156100f857634e487b7160e01b5f52601160045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f8084545f8160011c905060018216806102e557607f821691505b60208210810361030357634e487b7160e01b5f52602260045260245ffd5b808015610317576001811461032c5761035a565b60ff198416875282151583028701945061035a565b5f898152602090205f5b8481101561035257815489820152600190910190602001610336565b505082870194505b5050505061036881856102b3565b9594505050505056fea2646970667358221220a440e08fee15de0a971a851d8fb5e52dcfa2533f1e71d863f57a4843c049e6dd64736f6c634300081a0033",
}

// BaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMetaData.ABI instead.
var BaseABI = BaseMetaData.ABI

// BaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMetaData.Bin instead.
var BaseBin = BaseMetaData.Bin

// DeployBase deploys a new Ethereum contract, binding an instance of Base to it.
func DeployBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// Base is an auto generated Go binding around an Ethereum contract.
type Base struct {
	BaseCaller     // Read-only binding to the contract
	BaseTransactor // Write-only binding to the contract
	BaseFilterer   // Log filterer for contract events
}

// BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseSession struct {
	Contract     *Base             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCallerSession struct {
	Contract *BaseCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseTransactorSession struct {
	Contract     *BaseTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRaw struct {
	Contract *Base // Generic contract binding to access the raw methods on
}

// BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCallerRaw struct {
	Contract *BaseCaller // Generic read-only contract binding to access the raw methods on
}

// BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseTransactorRaw struct {
	Contract *BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase creates a new instance of Base, bound to a specific deployed contract.
func NewBase(address common.Address, backend bind.ContractBackend) (*Base, error) {
	contract, err := bindBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// NewBaseCaller creates a new read-only instance of Base, bound to a specific deployed contract.
func NewBaseCaller(address common.Address, caller bind.ContractCaller) (*BaseCaller, error) {
	contract, err := bindBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCaller{contract: contract}, nil
}

// NewBaseTransactor creates a new write-only instance of Base, bound to a specific deployed contract.
func NewBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseTransactor, error) {
	contract, err := bindBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTransactor{contract: contract}, nil
}

// NewBaseFilterer creates a new log filterer instance of Base, bound to a specific deployed contract.
func NewBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFilterer, error) {
	contract, err := bindBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFilterer{contract: contract}, nil
}

// bindBase binds a generic wrapper to an already deployed contract.
func bindBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.contract.Transact(opts, method, params...)
}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseCaller) Concat(opts *bind.CallOpts, s string) (string, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "concat", s)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseSession) Concat(s string) (string, error) {
	return _Base.Contract.Concat(&_Base.CallOpts, s)
}

// Concat is a free data retrieval call binding the contract method 0x7881c417.
//
// Solidity: function concat(string s) view returns(string)
func (_Base *BaseCallerSession) Concat(s string) (string, error) {
	return _Base.Contract.Concat(&_Base.CallOpts, s)
}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseCaller) GetQ(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Base.contract.Call(opts, &out, "getQ")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseSession) GetQ() (*big.Int, error) {
	return _Base.Contract.GetQ(&_Base.CallOpts)
}

// GetQ is a free data retrieval call binding the contract method 0xe8e4197a.
//
// Solidity: function getQ() view returns(uint256)
func (_Base *BaseCallerSession) GetQ() (*big.Int, error) {
	return _Base.Contract.GetQ(&_Base.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) returns(uint256)
func (_Base *BaseTransactor) Add(opts *bind.TransactOpts, _x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "add", _x, _y)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) returns(uint256)
func (_Base *BaseSession) Add(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Base.Contract.Add(&_Base.TransactOpts, _x, _y)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 _x, uint256 _y) returns(uint256)
func (_Base *BaseTransactorSession) Add(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Base.Contract.Add(&_Base.TransactOpts, _x, _y)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseTransactor) SetQ(opts *bind.TransactOpts, _q *big.Int) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "setQ", _q)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseSession) SetQ(_q *big.Int) (*types.Transaction, error) {
	return _Base.Contract.SetQ(&_Base.TransactOpts, _q)
}

// SetQ is a paid mutator transaction binding the contract method 0x8f2d8128.
//
// Solidity: function setQ(uint256 _q) returns()
func (_Base *BaseTransactorSession) SetQ(_q *big.Int) (*types.Transaction, error) {
	return _Base.Contract.SetQ(&_Base.TransactOpts, _q)
}

// BaseAddingIterator is returned from FilterAdding and is used to iterate over the raw logs and unpacked data for Adding events raised by the Base contract.
type BaseAddingIterator struct {
	Event *BaseAdding // Event containing the contract specifics and raw log

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
func (it *BaseAddingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseAdding)
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
		it.Event = new(BaseAdding)
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
func (it *BaseAddingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseAddingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseAdding represents a Adding event raised by the Base contract.
type BaseAdding struct {
	X   *big.Int
	Y   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdding is a free log retrieval operation binding the contract event 0x8cee9bf80efda447f9d9d047e2e937b140e1c653cb2af0f39d1f6c52a1ffd7dc.
//
// Solidity: event Adding(uint256 x, uint256 y)
func (_Base *BaseFilterer) FilterAdding(opts *bind.FilterOpts) (*BaseAddingIterator, error) {

	logs, sub, err := _Base.contract.FilterLogs(opts, "Adding")
	if err != nil {
		return nil, err
	}
	return &BaseAddingIterator{contract: _Base.contract, event: "Adding", logs: logs, sub: sub}, nil
}

// WatchAdding is a free log subscription operation binding the contract event 0x8cee9bf80efda447f9d9d047e2e937b140e1c653cb2af0f39d1f6c52a1ffd7dc.
//
// Solidity: event Adding(uint256 x, uint256 y)
func (_Base *BaseFilterer) WatchAdding(opts *bind.WatchOpts, sink chan<- *BaseAdding) (event.Subscription, error) {

	logs, sub, err := _Base.contract.WatchLogs(opts, "Adding")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseAdding)
				if err := _Base.contract.UnpackLog(event, "Adding", log); err != nil {
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

// ParseAdding is a log parse operation binding the contract event 0x8cee9bf80efda447f9d9d047e2e937b140e1c653cb2af0f39d1f6c52a1ffd7dc.
//
// Solidity: event Adding(uint256 x, uint256 y)
func (_Base *BaseFilterer) ParseAdding(log types.Log) (*BaseAdding, error) {
	event := new(BaseAdding)
	if err := _Base.contract.UnpackLog(event, "Adding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Base contract.
type BaseCreatedIterator struct {
	Event *BaseCreated // Event containing the contract specifics and raw log

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
func (it *BaseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseCreated)
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
		it.Event = new(BaseCreated)
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
func (it *BaseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseCreated represents a Created event raised by the Base contract.
type BaseCreated struct {
	Owner       common.Address
	DelployedAt common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x587ece4cd19692c5be1a4184503d607d45542d2aca0698c0068f52e09ccb541c.
//
// Solidity: event Created(address owner, address delployedAt)
func (_Base *BaseFilterer) FilterCreated(opts *bind.FilterOpts) (*BaseCreatedIterator, error) {

	logs, sub, err := _Base.contract.FilterLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return &BaseCreatedIterator{contract: _Base.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x587ece4cd19692c5be1a4184503d607d45542d2aca0698c0068f52e09ccb541c.
//
// Solidity: event Created(address owner, address delployedAt)
func (_Base *BaseFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *BaseCreated) (event.Subscription, error) {

	logs, sub, err := _Base.contract.WatchLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseCreated)
				if err := _Base.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x587ece4cd19692c5be1a4184503d607d45542d2aca0698c0068f52e09ccb541c.
//
// Solidity: event Created(address owner, address delployedAt)
func (_Base *BaseFilterer) ParseCreated(log types.Log) (*BaseCreated, error) {
	event := new(BaseCreated)
	if err := _Base.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseSettingQIterator is returned from FilterSettingQ and is used to iterate over the raw logs and unpacked data for SettingQ events raised by the Base contract.
type BaseSettingQIterator struct {
	Event *BaseSettingQ // Event containing the contract specifics and raw log

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
func (it *BaseSettingQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseSettingQ)
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
		it.Event = new(BaseSettingQ)
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
func (it *BaseSettingQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseSettingQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseSettingQ represents a SettingQ event raised by the Base contract.
type BaseSettingQ struct {
	Q   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettingQ is a free log retrieval operation binding the contract event 0x4af6c9c7470c3b6ab60fab4fef0d6053a5a05afd3df0b7be9243f972be4eb66b.
//
// Solidity: event SettingQ(uint256 q)
func (_Base *BaseFilterer) FilterSettingQ(opts *bind.FilterOpts) (*BaseSettingQIterator, error) {

	logs, sub, err := _Base.contract.FilterLogs(opts, "SettingQ")
	if err != nil {
		return nil, err
	}
	return &BaseSettingQIterator{contract: _Base.contract, event: "SettingQ", logs: logs, sub: sub}, nil
}

// WatchSettingQ is a free log subscription operation binding the contract event 0x4af6c9c7470c3b6ab60fab4fef0d6053a5a05afd3df0b7be9243f972be4eb66b.
//
// Solidity: event SettingQ(uint256 q)
func (_Base *BaseFilterer) WatchSettingQ(opts *bind.WatchOpts, sink chan<- *BaseSettingQ) (event.Subscription, error) {

	logs, sub, err := _Base.contract.WatchLogs(opts, "SettingQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseSettingQ)
				if err := _Base.contract.UnpackLog(event, "SettingQ", log); err != nil {
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

// ParseSettingQ is a log parse operation binding the contract event 0x4af6c9c7470c3b6ab60fab4fef0d6053a5a05afd3df0b7be9243f972be4eb66b.
//
// Solidity: event SettingQ(uint256 q)
func (_Base *BaseFilterer) ParseSettingQ(log types.Log) (*BaseSettingQ, error) {
	event := new(BaseSettingQ)
	if err := _Base.contract.UnpackLog(event, "SettingQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
