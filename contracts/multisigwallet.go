// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MultisigWalletABI is the input ABI used to generate the binding from.
const MultisigWalletABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownersArr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owners_\",\"type\":\"address[]\"},{\"name\":\"threshold_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MultisigWalletBin is the compiled bytecode used for deploying new contracts.
const MultisigWalletBin = `0x608060405234801561001057600080fd5b5060405161062c38038061062c833981016040528051602082015191018051909190600090600a10801590610046575082518211155b8015610053575060008210155b151561005e57600080fd5b5060005b82518110156100b957600160026000858481518110151561007f57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610062565b82516100cc9060039060208601906100d7565b505060015550610163565b82805482825590600052602060002090810192821561012c579160200282015b8281111561012c5782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906100f7565b5061013892915061013c565b5090565b61016091905b80821115610138578054600160a060020a0319168155600101610142565b90565b6104ba806101726000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166342cde4e88114610063578063aa5df9e21461008a578063affed0e0146100cb578063f12d394f146100e0575b005b34801561006f57600080fd5b50610078610208565b60408051918252519081900360200190f35b34801561009657600080fd5b506100a260043561020e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100d757600080fd5b50610078610243565b3480156100ec57600080fd5b506040805160206004803580820135838102808601850190965280855261006195369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a99890198929750908201955093508392508501908490808284375050604080516020888301358a018035601f8101839004830284018301909452838352979a893573ffffffffffffffffffffffffffffffffffffffff169a8a8301359a919990985060609091019650919450908101925081908401838280828437509497506102499650505050505050565b60015481565b600380548290811061021c57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005481565b600080600080600154895114151561026057600080fd5b87518951148015610272575089518951145b151561027d57600080fd5b600080546040517f190000000000000000000000000000000000000000000000000000000000000080825260018201849052306c01000000000000000000000000818102600285015273ffffffffffffffffffffffffffffffffffffffff8d16026016840152602a83018b9052895191949390928c928c928c929091604a82019060208501908083835b602083106103265780518252601f199092019160209182019101610307565b51815160209384036101000a600019018019909216911617905292019384525060405192839003019091209a5060009950505050505050505b60015483101561045a576001848b8581518110151561037a57fe5b906020019060200201518b8681518110151561039257fe5b906020019060200201518b878151811015156103aa57fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af115801561040b573d6000803e3d6000fd5b505060408051601f19015173ffffffffffffffffffffffffffffffffffffffff811660009081526002602052919091205490935060ff161515905061044f57600080fd5b60019092019161035f565b60009050600080865160208801898b5af1905080151561047957600080fd5b505060008054600101905550505050505050505600a165627a7a723058200d81fa39989c7aaa56bb19bfd15dd5ddeb10f806072119fb1093cbb273fde0a60029`

// DeployMultisigWallet deploys a new Ethereum contract, binding an instance of MultisigWallet to it.
func DeployMultisigWallet(auth *bind.TransactOpts, backend bind.ContractBackend, owners_ []common.Address, threshold_ *big.Int) (common.Address, *types.Transaction, *MultisigWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultisigWalletBin), backend, owners_, threshold_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultisigWallet{MultisigWalletCaller: MultisigWalletCaller{contract: contract}, MultisigWalletTransactor: MultisigWalletTransactor{contract: contract}, MultisigWalletFilterer: MultisigWalletFilterer{contract: contract}}, nil
}

// MultisigWallet is an auto generated Go binding around an Ethereum contract.
type MultisigWallet struct {
	MultisigWalletCaller     // Read-only binding to the contract
	MultisigWalletTransactor // Write-only binding to the contract
	MultisigWalletFilterer   // Log filterer for contract events
}

// MultisigWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigWalletSession struct {
	Contract     *MultisigWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigWalletCallerSession struct {
	Contract *MultisigWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultisigWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigWalletTransactorSession struct {
	Contract     *MultisigWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultisigWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigWalletRaw struct {
	Contract *MultisigWallet // Generic contract binding to access the raw methods on
}

// MultisigWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigWalletCallerRaw struct {
	Contract *MultisigWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigWalletTransactorRaw struct {
	Contract *MultisigWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisigWallet creates a new instance of MultisigWallet, bound to a specific deployed contract.
func NewMultisigWallet(address common.Address, backend bind.ContractBackend) (*MultisigWallet, error) {
	contract, err := bindMultisigWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultisigWallet{MultisigWalletCaller: MultisigWalletCaller{contract: contract}, MultisigWalletTransactor: MultisigWalletTransactor{contract: contract}, MultisigWalletFilterer: MultisigWalletFilterer{contract: contract}}, nil
}

// NewMultisigWalletCaller creates a new read-only instance of MultisigWallet, bound to a specific deployed contract.
func NewMultisigWalletCaller(address common.Address, caller bind.ContractCaller) (*MultisigWalletCaller, error) {
	contract, err := bindMultisigWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigWalletCaller{contract: contract}, nil
}

// NewMultisigWalletTransactor creates a new write-only instance of MultisigWallet, bound to a specific deployed contract.
func NewMultisigWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigWalletTransactor, error) {
	contract, err := bindMultisigWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigWalletTransactor{contract: contract}, nil
}

// NewMultisigWalletFilterer creates a new log filterer instance of MultisigWallet, bound to a specific deployed contract.
func NewMultisigWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigWalletFilterer, error) {
	contract, err := bindMultisigWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigWalletFilterer{contract: contract}, nil
}

// bindMultisigWallet binds a generic wrapper to an already deployed contract.
func bindMultisigWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultisigWallet *MultisigWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultisigWallet.Contract.MultisigWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultisigWallet *MultisigWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultisigWallet.Contract.MultisigWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultisigWallet *MultisigWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultisigWallet.Contract.MultisigWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultisigWallet *MultisigWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultisigWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultisigWallet *MultisigWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultisigWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultisigWallet *MultisigWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultisigWallet.Contract.contract.Transact(opts, method, params...)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MultisigWallet *MultisigWalletCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultisigWallet.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MultisigWallet *MultisigWalletSession) Nonce() (*big.Int, error) {
	return _MultisigWallet.Contract.Nonce(&_MultisigWallet.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MultisigWallet *MultisigWalletCallerSession) Nonce() (*big.Int, error) {
	return _MultisigWallet.Contract.Nonce(&_MultisigWallet.CallOpts)
}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr( uint256) constant returns(address)
func (_MultisigWallet *MultisigWalletCaller) OwnersArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultisigWallet.contract.Call(opts, out, "ownersArr", arg0)
	return *ret0, err
}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr( uint256) constant returns(address)
func (_MultisigWallet *MultisigWalletSession) OwnersArr(arg0 *big.Int) (common.Address, error) {
	return _MultisigWallet.Contract.OwnersArr(&_MultisigWallet.CallOpts, arg0)
}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr( uint256) constant returns(address)
func (_MultisigWallet *MultisigWalletCallerSession) OwnersArr(arg0 *big.Int) (common.Address, error) {
	return _MultisigWallet.Contract.OwnersArr(&_MultisigWallet.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_MultisigWallet *MultisigWalletCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultisigWallet.contract.Call(opts, out, "threshold")
	return *ret0, err
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_MultisigWallet *MultisigWalletSession) Threshold() (*big.Int, error) {
	return _MultisigWallet.Contract.Threshold(&_MultisigWallet.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() constant returns(uint256)
func (_MultisigWallet *MultisigWalletCallerSession) Threshold() (*big.Int, error) {
	return _MultisigWallet.Contract.Threshold(&_MultisigWallet.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xf12d394f.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], destination address, value uint256, data bytes) returns()
func (_MultisigWallet *MultisigWalletTransactor) Execute(opts *bind.TransactOpts, sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultisigWallet.contract.Transact(opts, "execute", sigV, sigR, sigS, destination, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xf12d394f.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], destination address, value uint256, data bytes) returns()
func (_MultisigWallet *MultisigWalletSession) Execute(sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultisigWallet.Contract.Execute(&_MultisigWallet.TransactOpts, sigV, sigR, sigS, destination, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xf12d394f.
//
// Solidity: function execute(sigV uint8[], sigR bytes32[], sigS bytes32[], destination address, value uint256, data bytes) returns()
func (_MultisigWallet *MultisigWalletTransactorSession) Execute(sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultisigWallet.Contract.Execute(&_MultisigWallet.TransactOpts, sigV, sigR, sigS, destination, value, data)
}
