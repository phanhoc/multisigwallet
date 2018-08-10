// // Code generated - DO NOT EDIT.
// // This file is a generated binding and any manual changes will be lost.
//
package contracts

//
// import (
// 	"math/big"
// 	"strings"
//
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// )
//
// // MultisigWalletABI is the input ABI used to generate the binding from.
// const MultisigWalletABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signature1\",\"type\":\"bytes\"},{\"name\":\"signature2\",\"type\":\"bytes\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owners_\",\"type\":\"address[]\"},{\"name\":\"threshold_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"
//
// // MultisigWalletBin is the compiled bytecode used for deploying new contracts.
// const MultisigWalletBin = `0x608060405234801561001057600080fd5b506040516105cc3803806105cc83398101604052805160208201519101805190919060009060021061004157600080fd5b600a835111158015610054575082518211155b80156100605750600082115b151561006b57600080fd5b5060005b82518110156100c657600160026000858481518110151561008c57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905560010161006f565b82516100d99060039060208601906100e4565b505060015550610170565b828054828255906000526020600020908101928215610139579160200282015b828111156101395782518254600160a060020a031916600160a060020a03909116178255602090920191600190910190610104565b50610145929150610149565b5090565b61016d91905b80821115610145578054600160a060020a031916815560010161014f565b90565b61044d8061017f6000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461006357806342cde4e814610097578063aab327e8146100be578063affed0e0146101a9575b005b34801561006f57600080fd5b5061007b6004356101be565b60408051600160a060020a039092168252519081900360200190f35b3480156100a357600080fd5b506100ac6101e6565b60408051918252519081900360200190f35b3480156100ca57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261006194369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020888301358a018035601f8101839004830284018301909452838352979a8935600160a060020a03169a8a8301359a919990985060609091019650919450908101925081908401838280828437509497506101ec9650505050505050565b3480156101b557600080fd5b506100ac610335565b60038054829081106101cc57fe5b600091825260209091200154600160a060020a0316905081565b60015481565b6000808484846000546040518085600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140184815260200183805190602001908083835b602083106102515780518252601f199092019160209182019101610232565b51815160209384036101000a60001901801990921691161790529201938452506040519283900301909120955061029193508892508591508a905061033b565b905061029e85838861033b565b905084600160a060020a0316848460405180828051906020019080838360005b838110156102d65781810151838201526020016102be565b50505050905090810190601f1680156103035780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af192505050151561032357600080fd5b50506000805460010190555050505050565b60005481565b600080610348848461037a565b600160a060020a03811660009081526002602052604090205490915060ff16151561037257600080fd5b949350505050565b6000806000808451604114151561039057600080fd5b50505060208201516040830151604184015160ff16601b8110156103b257601b015b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561040c573d6000803e3d6000fd5b5050604051601f1901519796505050505050505600a165627a7a7230582049b8caf97e3a520aee4f4344d88ea161fb6798a8baaeb3d594a924b02d2317fd0029`
//
// // DeployMultisigWallet deploys a new Ethereum contract, binding an instance of MultisigWallet to it.
// func DeployMultisigWallet(auth *bind.TransactOpts, backend bind.ContractBackend, owners_ []common.Address, threshold_ *big.Int) (common.Address, *types.Transaction, *MultisigWallet, error) {
// 	parsed, err := abi.JSON(strings.NewReader(MultisigWalletABI))
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultisigWalletBin), backend, owners_, threshold_)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	return address, tx, &MultisigWallet{MultisigWalletCaller: MultisigWalletCaller{contract: contract}, MultisigWalletTransactor: MultisigWalletTransactor{contract: contract}, MultisigWalletFilterer: MultisigWalletFilterer{contract: contract}}, nil
// }
//
// // MultisigWallet is an auto generated Go binding around an Ethereum contract.
// type MultisigWallet struct {
// 	MultisigWalletCaller     // Read-only binding to the contract
// 	MultisigWalletTransactor // Write-only binding to the contract
// 	MultisigWalletFilterer   // Log filterer for contract events
// }
//
// // MultisigWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
// type MultisigWalletCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }
//
// // MultisigWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type MultisigWalletTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }
//
// // MultisigWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type MultisigWalletFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }
//
// // MultisigWalletSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type MultisigWalletSession struct {
// 	Contract     *MultisigWallet   // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }
//
// // MultisigWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type MultisigWalletCallerSession struct {
// 	Contract *MultisigWalletCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts         // Call options to use throughout this session
// }
//
// // MultisigWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type MultisigWalletTransactorSession struct {
// 	Contract     *MultisigWalletTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
// }
//
// // MultisigWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
// type MultisigWalletRaw struct {
// 	Contract *MultisigWallet // Generic contract binding to access the raw methods on
// }
//
// // MultisigWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type MultisigWalletCallerRaw struct {
// 	Contract *MultisigWalletCaller // Generic read-only contract binding to access the raw methods on
// }
//
// // MultisigWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type MultisigWalletTransactorRaw struct {
// 	Contract *MultisigWalletTransactor // Generic write-only contract binding to access the raw methods on
// }
//
// // NewMultisigWallet creates a new instance of MultisigWallet, bound to a specific deployed contract.
// func NewMultisigWallet(address common.Address, backend bind.ContractBackend) (*MultisigWallet, error) {
// 	contract, err := bindMultisigWallet(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &MultisigWallet{MultisigWalletCaller: MultisigWalletCaller{contract: contract}, MultisigWalletTransactor: MultisigWalletTransactor{contract: contract}, MultisigWalletFilterer: MultisigWalletFilterer{contract: contract}}, nil
// }
//
// // NewMultisigWalletCaller creates a new read-only instance of MultisigWallet, bound to a specific deployed contract.
// func NewMultisigWalletCaller(address common.Address, caller bind.ContractCaller) (*MultisigWalletCaller, error) {
// 	contract, err := bindMultisigWallet(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &MultisigWalletCaller{contract: contract}, nil
// }
//
// // NewMultisigWalletTransactor creates a new write-only instance of MultisigWallet, bound to a specific deployed contract.
// func NewMultisigWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigWalletTransactor, error) {
// 	contract, err := bindMultisigWallet(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &MultisigWalletTransactor{contract: contract}, nil
// }
//
// // NewMultisigWalletFilterer creates a new log filterer instance of MultisigWallet, bound to a specific deployed contract.
// func NewMultisigWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigWalletFilterer, error) {
// 	contract, err := bindMultisigWallet(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &MultisigWalletFilterer{contract: contract}, nil
// }
//
// // bindMultisigWallet binds a generic wrapper to an already deployed contract.
// func bindMultisigWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(MultisigWalletABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }
//
// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_MultisigWallet *MultisigWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
// 	return _MultisigWallet.Contract.MultisigWalletCaller.contract.Call(opts, result, method, params...)
// }
//
// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_MultisigWallet *MultisigWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.MultisigWalletTransactor.contract.Transfer(opts)
// }
//
// // Transact invokes the (paid) contract method with params as input values.
// func (_MultisigWallet *MultisigWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.MultisigWalletTransactor.contract.Transact(opts, method, params...)
// }
//
// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_MultisigWallet *MultisigWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
// 	return _MultisigWallet.Contract.contract.Call(opts, result, method, params...)
// }
//
// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_MultisigWallet *MultisigWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.contract.Transfer(opts)
// }
//
// // Transact invokes the (paid) contract method with params as input values.
// func (_MultisigWallet *MultisigWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.contract.Transact(opts, method, params...)
// }
//
// // Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
// //
// // Solidity: function nonce() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
// 	var (
// 		ret0 = new(*big.Int)
// 	)
// 	out := ret0
// 	err := _MultisigWallet.contract.Call(opts, out, "nonce")
// 	return *ret0, err
// }
//
// // Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
// //
// // Solidity: function nonce() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletSession) Nonce() (*big.Int, error) {
// 	return _MultisigWallet.Contract.Nonce(&_MultisigWallet.CallOpts)
// }
//
// // Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
// //
// // Solidity: function nonce() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletCallerSession) Nonce() (*big.Int, error) {
// 	return _MultisigWallet.Contract.Nonce(&_MultisigWallet.CallOpts)
// }
//
// // Owners is a free data retrieval call binding the contract method 0x025e7c27.
// //
// // Solidity: function owners( uint256) constant returns(address)
// func (_MultisigWallet *MultisigWalletCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
// 	var (
// 		ret0 = new(common.Address)
// 	)
// 	out := ret0
// 	err := _MultisigWallet.contract.Call(opts, out, "owners", arg0)
// 	return *ret0, err
// }
//
// // Owners is a free data retrieval call binding the contract method 0x025e7c27.
// //
// // Solidity: function owners( uint256) constant returns(address)
// func (_MultisigWallet *MultisigWalletSession) Owners(arg0 *big.Int) (common.Address, error) {
// 	return _MultisigWallet.Contract.Owners(&_MultisigWallet.CallOpts, arg0)
// }
//
// // Owners is a free data retrieval call binding the contract method 0x025e7c27.
// //
// // Solidity: function owners( uint256) constant returns(address)
// func (_MultisigWallet *MultisigWalletCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
// 	return _MultisigWallet.Contract.Owners(&_MultisigWallet.CallOpts, arg0)
// }
//
// // Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
// //
// // Solidity: function threshold() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
// 	var (
// 		ret0 = new(*big.Int)
// 	)
// 	out := ret0
// 	err := _MultisigWallet.contract.Call(opts, out, "threshold")
// 	return *ret0, err
// }
//
// // Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
// //
// // Solidity: function threshold() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletSession) Threshold() (*big.Int, error) {
// 	return _MultisigWallet.Contract.Threshold(&_MultisigWallet.CallOpts)
// }
//
// // Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
// //
// // Solidity: function threshold() constant returns(uint256)
// func (_MultisigWallet *MultisigWalletCallerSession) Threshold() (*big.Int, error) {
// 	return _MultisigWallet.Contract.Threshold(&_MultisigWallet.CallOpts)
// }
//
// // Execute is a paid mutator transaction binding the contract method 0xaab327e8.
// //
// // Solidity: function execute(signature1 bytes, signature2 bytes, destination address, value uint256, data bytes) returns()
// func (_MultisigWallet *MultisigWalletTransactor) Execute(opts *bind.TransactOpts, signature1 []byte, signature2 []byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
// 	return _MultisigWallet.contract.Transact(opts, "execute", signature1, signature2, destination, value, data)
// }
//
// // Execute is a paid mutator transaction binding the contract method 0xaab327e8.
// //
// // Solidity: function execute(signature1 bytes, signature2 bytes, destination address, value uint256, data bytes) returns()
// func (_MultisigWallet *MultisigWalletSession) Execute(signature1 []byte, signature2 []byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.Execute(&_MultisigWallet.TransactOpts, signature1, signature2, destination, value, data)
// }
//
// // Execute is a paid mutator transaction binding the contract method 0xaab327e8.
// //
// // Solidity: function execute(signature1 bytes, signature2 bytes, destination address, value uint256, data bytes) returns()
// func (_MultisigWallet *MultisigWalletTransactorSession) Execute(signature1 []byte, signature2 []byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
// 	return _MultisigWallet.Contract.Execute(&_MultisigWallet.TransactOpts, signature1, signature2, destination, value, data)
// }
