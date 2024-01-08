// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Role

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
)

// ServiceABI is the input ABI used to generate the binding from.
const ServiceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_user_addr\",\"type\":\"address\"}],\"name\":\"onlyDRRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user_addr\",\"type\":\"address\"}],\"name\":\"onlyPRRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user_addr\",\"type\":\"address\"}],\"name\":\"onlyRRRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ServiceBin is the compiled bytecode used for deploying new contracts.
var ServiceBin = "0x608060405234801561001057600080fd5b5061001961007a565b604051809103906000f080158015610035573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061008a565b60405161169380610c9383390190565b610bfa806100996000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806334fc80721461005c5780635ee471121461009f57806377973d2f146100e2575b600080fd5b34801561006857600080fd5b5061009d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610125565b005b3480156100ab57600080fd5b506100e0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610376565b005b3480156100ee57600080fd5b50610123600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105c7565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324e079136040805190810160405280600281526020017f44520000000000000000000000000000000000000000000000000000000000008152506101a184610818565b6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835285818151815260200191508051906020019080838360005b838110156102105780820151818401526020810190506101f5565b50505050905090810190601f16801561023d5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561027657808201518184015260208101905061025b565b50505050905090810190601f1680156102a35780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156102c457600080fd5b505af11580156102d8573d6000803e3d6000fd5b505050506040513d60208110156102ee57600080fd5b81019080805190602001909291905050501515610373576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420686173206472207065726d697373696f6e000000000000000000000081525060200191505060405180910390fd5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324e079136040805190810160405280600281526020017f50520000000000000000000000000000000000000000000000000000000000008152506103f284610818565b6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610461578082015181840152602081019050610446565b50505050905090810190601f16801561048e5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156104c75780820151818401526020810190506104ac565b50505050905090810190601f1680156104f45780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561051557600080fd5b505af1158015610529573d6000803e3d6000fd5b505050506040513d602081101561053f57600080fd5b810190808051906020019092919050505015156105c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420686173207072207065726d697373696f6e000000000000000000000081525060200191505060405180910390fd5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166324e079136040805190810160405280600281526020017f525200000000000000000000000000000000000000000000000000000000000081525061064384610818565b6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835285818151815260200191508051906020019080838360005b838110156106b2578082015181840152602081019050610697565b50505050905090810190601f1680156106df5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156107185780820151818401526020810190506106fd565b50505050905090810190601f1680156107455780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561076657600080fd5b505af115801561077a573d6000803e3d6000fd5b505050506040513d602081101561079057600080fd5b81019080805190602001909291905050501515610815576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420686173207272207065726d697373696f6e000000000000000000000081525060200191505060405180910390fd5b50565b6060600060608060008573ffffffffffffffffffffffffffffffffffffffff1660010293506040805190810160405280601081526020017f30313233343536373839616263646566000000000000000000000000000000008152509250602a6040519080825280601f01601f1916602001820160405280156108a95781602001602082028038833980820191505090505b5091507f30000000000000000000000000000000000000000000000000000000000000008260008151811015156108dc57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000082600181518110151561093c57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600090505b6014811015610bc25782600485600c840160208110151561098e57fe5b1a7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060020a90047f0100000000000000000000000000000000000000000000000000000000000000900460ff16815181101515610a0b57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000028260028302600201815181101515610a6a57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535082600f7f01000000000000000000000000000000000000000000000000000000000000000285600c8401602081101515610ad057fe5b1a7f010000000000000000000000000000000000000000000000000000000000000002167f0100000000000000000000000000000000000000000000000000000000000000900460ff16815181101515610b2657fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000028260028302600301815181101515610b8557fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610971565b819450505050509190505600a165627a7a7230582040ca75b28503600c95374bd9b4b670a816b7acea4265b4484a1bb010ea77db28002960806040523480156200001157600080fd5b506110016000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016200010191906200024a565b602060405180830381600087803b1580156200011c57600080fd5b505af115801562000131573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525062000157919081019062000174565b50620002f4565b60006200016c8251620002a3565b905092915050565b6000602082840312156200018757600080fd5b600062000197848285016200015e565b91505092915050565b6000620001ad8262000298565b808452620001c3816020860160208601620002ad565b620001ce81620002e3565b602085010191505092915050565b6000600482527f726f6c65000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000600982527f757365725f6164647200000000000000000000000000000000000000000000006020830152604082019050919050565b60006060820190508181036000830152620002668184620001a0565b905081810360208301526200027b81620001dc565b90508181036040830152620002908162000213565b905092915050565b600081519050919050565b6000819050919050565b60005b83811015620002cd578082015181840152602081019050620002b0565b83811115620002dd576000848401525b50505050565b6000601f19601f8301169050919050565b61138f80620003046000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806310158b891461006757806324e0791314610090578063a6343012146100cd578063dd517fae1461010a575b600080fd5b34801561007357600080fd5b5061008e60048036036100899190810190610f58565b610147565b005b34801561009c57600080fd5b506100b760048036036100b29190810190610f58565b610196565b6040516100c491906110da565b60405180910390f35b3480156100d957600080fd5b506100f460048036036100ef9190810190610f58565b6102ac565b60405161010191906110f5565b60405180910390f35b34801561011657600080fd5b50610131600480360361012c9190810190610f58565b6105c8565b60405161013e91906110f5565b60405180910390f35b6101518282610196565b1515610192576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610189906111c7565b60405180910390fd5b5050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016102449190611110565b602060405180830381600087803b15801561025e57600080fd5b505af1158015610272573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102969190810190610f06565b90506102a3818585610908565b91505092915050565b600080600080600092506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016103619190611110565b602060405180830381600087803b15801561037b57600080fd5b505af115801561038f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103b39190810190610f06565b91508173ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561041957600080fd5b505af115801561042d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104519190810190610e8b565b90508073ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104a89190611192565b600060405180830381600087803b1580156104c257600080fd5b505af11580156104d6573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166328bb211787836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610531929190611132565b602060405180830381600087803b15801561054b57600080fd5b505af115801561055f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105839190810190610f2f565b92507fee205d8383de0d8cc6823192e93e3086dd317c369858ed095ed2d719d25bfd23836040516105b491906110f5565b60405180910390a182935050505092915050565b600080600080600092506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161067d9190611110565b602060405180830381600087803b15801561069757600080fd5b505af11580156106ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506106cf9190810190610f06565b91506106da86610b81565b1580156106ef57506106ed828787610908565b155b156108c5578173ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561075857600080fd5b505af115801561076c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107909190810190610edd565b90508073ffffffffffffffffffffffffffffffffffffffff1663e942b516866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016107e79190611192565b600060405180830381600087803b15801561080157600080fd5b505af1158015610815573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166331afac3687836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610870929190611162565b602060405180830381600087803b15801561088a57600080fd5b505af115801561089e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108c29190810190610f2f565b92505b7f14a121ebed99a8737803a36f9a293fa68078956202e0a990b3d04694d93b71e4836040516108f491906110f5565b60405180910390a182935050505092915050565b6000808473ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561096f57600080fd5b505af1158015610983573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506109a79190810190610e8b565b90508073ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016109fe9190611192565b600060405180830381600087803b158015610a1857600080fd5b505af1158015610a2c573d6000803e3d6000fd5b5050505060008573ffffffffffffffffffffffffffffffffffffffff1663e8434e3986846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610a89929190611132565b602060405180830381600087803b158015610aa357600080fd5b505af1158015610ab7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610adb9190810190610eb4565b73ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610b3e57600080fd5b505af1158015610b52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b769190810190610f2f565b139150509392505050565b60006060600080849250600083511415610b9e5760019350610dc9565b600091505b8251821015610dc4578282815181101515610bba57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002905060207f010000000000000000000000000000000000000000000000000000000000000002817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614158015610cc3575060097f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b8015610d365750600a7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b8015610da95750600d7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b15610db75760009350610dc9565b8180600101925050610ba3565b600193505b505050919050565b6000610ddd825161128c565b905092915050565b6000610df1825161129e565b905092915050565b6000610e0582516112b0565b905092915050565b6000610e1982516112c2565b905092915050565b6000610e2d82516112d4565b905092915050565b600082601f8301121515610e4857600080fd5b8135610e5b610e5682611214565b6111e7565b91508082526020830160208301858383011115610e7757600080fd5b610e82838284611302565b50505092915050565b600060208284031215610e9d57600080fd5b6000610eab84828501610dd1565b91505092915050565b600060208284031215610ec657600080fd5b6000610ed484828501610de5565b91505092915050565b600060208284031215610eef57600080fd5b6000610efd84828501610df9565b91505092915050565b600060208284031215610f1857600080fd5b6000610f2684828501610e0d565b91505092915050565b600060208284031215610f4157600080fd5b6000610f4f84828501610e21565b91505092915050565b60008060408385031215610f6b57600080fd5b600083013567ffffffffffffffff811115610f8557600080fd5b610f9185828601610e35565b925050602083013567ffffffffffffffff811115610fae57600080fd5b610fba85828601610e35565b9150509250929050565b610fcd81611276565b82525050565b610fdc816112de565b82525050565b610feb816112f0565b82525050565b610ffa81611282565b82525050565b600061100b8261124b565b80845261101f816020860160208601611311565b61102881611344565b602085010191505092915050565b600061104182611240565b808452611055816020860160208601611311565b61105e81611344565b602085010191505092915050565b6000600982527f757365725f6164647200000000000000000000000000000000000000000000006020830152604082019050919050565b6000601382527f6e6f742068617665207065726d697373696f6e000000000000000000000000006020830152604082019050919050565b60006020820190506110ef6000830184610fc4565b92915050565b600060208201905061110a6000830184610ff1565b92915050565b6000602082019050818103600083015261112a8184611036565b905092915050565b6000604082019050818103600083015261114c8185611000565b905061115b6020830184610fd3565b9392505050565b6000604082019050818103600083015261117c8185611000565b905061118b6020830184610fe2565b9392505050565b600060408201905081810360008301526111ab8161106c565b905081810360208301526111bf8184611000565b905092915050565b600060208201905081810360008301526111e0816110a3565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171561120a57600080fd5b8060405250919050565b600067ffffffffffffffff82111561122b57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60008115159050919050565b6000819050919050565b600061129782611256565b9050919050565b60006112a982611256565b9050919050565b60006112bb82611256565b9050919050565b60006112cd82611256565b9050919050565b6000819050919050565b60006112e982611256565b9050919050565b60006112fb82611256565b9050919050565b82818337600083830152505050565b60005b8381101561132f578082015181840152602081019050611314565b8381111561133e576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a72305820c64e4a459dbd5a43f86bab8215410901e6b9169a4da9b48fa6ad2609c937754d6c6578706572696d656e74616cf50037"

// DeployService deploys a new contract, binding an instance of Service to it.
func DeployService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Service, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Service{ServiceCaller: ServiceCaller{contract: contract}, ServiceTransactor: ServiceTransactor{contract: contract}, ServiceFilterer: ServiceFilterer{contract: contract}}, nil
}

func AsyncDeployService(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ServiceBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Service is an auto generated Go binding around a Solidity contract.
type Service struct {
	ServiceCaller     // Read-only binding to the contract
	ServiceTransactor // Write-only binding to the contract
	ServiceFilterer   // Log filterer for contract events
}

// ServiceCaller is an auto generated read-only Go binding around a Solidity contract.
type ServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceTransactor is an auto generated write-only Go binding around a Solidity contract.
type ServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ServiceSession struct {
	Contract     *Service          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServiceCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ServiceCallerSession struct {
	Contract *ServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ServiceTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ServiceTransactorSession struct {
	Contract     *ServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ServiceRaw is an auto generated low-level Go binding around a Solidity contract.
type ServiceRaw struct {
	Contract *Service // Generic contract binding to access the raw methods on
}

// ServiceCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ServiceCallerRaw struct {
	Contract *ServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ServiceTransactorRaw struct {
	Contract *ServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewService creates a new instance of Service, bound to a specific deployed contract.
func NewService(address common.Address, backend bind.ContractBackend) (*Service, error) {
	contract, err := bindService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Service{ServiceCaller: ServiceCaller{contract: contract}, ServiceTransactor: ServiceTransactor{contract: contract}, ServiceFilterer: ServiceFilterer{contract: contract}}, nil
}

// NewServiceCaller creates a new read-only instance of Service, bound to a specific deployed contract.
func NewServiceCaller(address common.Address, caller bind.ContractCaller) (*ServiceCaller, error) {
	contract, err := bindService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceCaller{contract: contract}, nil
}

// NewServiceTransactor creates a new write-only instance of Service, bound to a specific deployed contract.
func NewServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceTransactor, error) {
	contract, err := bindService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceTransactor{contract: contract}, nil
}

// NewServiceFilterer creates a new log filterer instance of Service, bound to a specific deployed contract.
func NewServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceFilterer, error) {
	contract, err := bindService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceFilterer{contract: contract}, nil
}

// bindService binds a generic wrapper to an already deployed contract.
func bindService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Service.Contract.ServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Service.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.contract.Transact(opts, method, params...)
}

// OnlyDRRole is a free data retrieval call binding the contract method 0x34fc8072.
//
// Solidity: function onlyDRRole(address _user_addr) constant returns()
func (_Service *ServiceCaller) OnlyDRRole(opts *bind.CallOpts, _user_addr common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _Service.contract.Call(opts, out, "onlyDRRole", _user_addr)
	return err
}

// OnlyDRRole is a free data retrieval call binding the contract method 0x34fc8072.
//
// Solidity: function onlyDRRole(address _user_addr) constant returns()
func (_Service *ServiceSession) OnlyDRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyDRRole(&_Service.CallOpts, _user_addr)
}

// OnlyDRRole is a free data retrieval call binding the contract method 0x34fc8072.
//
// Solidity: function onlyDRRole(address _user_addr) constant returns()
func (_Service *ServiceCallerSession) OnlyDRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyDRRole(&_Service.CallOpts, _user_addr)
}

// OnlyPRRole is a free data retrieval call binding the contract method 0x5ee47112.
//
// Solidity: function onlyPRRole(address _user_addr) constant returns()
func (_Service *ServiceCaller) OnlyPRRole(opts *bind.CallOpts, _user_addr common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _Service.contract.Call(opts, out, "onlyPRRole", _user_addr)
	return err
}

// OnlyPRRole is a free data retrieval call binding the contract method 0x5ee47112.
//
// Solidity: function onlyPRRole(address _user_addr) constant returns()
func (_Service *ServiceSession) OnlyPRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyPRRole(&_Service.CallOpts, _user_addr)
}

// OnlyPRRole is a free data retrieval call binding the contract method 0x5ee47112.
//
// Solidity: function onlyPRRole(address _user_addr) constant returns()
func (_Service *ServiceCallerSession) OnlyPRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyPRRole(&_Service.CallOpts, _user_addr)
}

// OnlyRRRole is a free data retrieval call binding the contract method 0x77973d2f.
//
// Solidity: function onlyRRRole(address _user_addr) constant returns()
func (_Service *ServiceCaller) OnlyRRRole(opts *bind.CallOpts, _user_addr common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _Service.contract.Call(opts, out, "onlyRRRole", _user_addr)
	return err
}

// OnlyRRRole is a free data retrieval call binding the contract method 0x77973d2f.
//
// Solidity: function onlyRRRole(address _user_addr) constant returns()
func (_Service *ServiceSession) OnlyRRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyRRRole(&_Service.CallOpts, _user_addr)
}

// OnlyRRRole is a free data retrieval call binding the contract method 0x77973d2f.
//
// Solidity: function onlyRRRole(address _user_addr) constant returns()
func (_Service *ServiceCallerSession) OnlyRRRole(_user_addr common.Address) error {
	return _Service.Contract.OnlyRRRole(&_Service.CallOpts, _user_addr)
}
