// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: 2f4ca06adf3ede4e375fbb39f8806c16a90078c56d3b7703325b97d8d194729b
// Version: 0.1.0
package erc20

import (
	"fmt"
	"math/big"

	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
)

var (
	_ = big.NewInt
)

// ERC20 is a solidity contract
type ERC20 struct {
	c *contract.Contract
}

// NewERC20 creates a new instance of the contract at a specific address
func NewERC20(addr ethgo.Address, provider *jsonrpc.Client) *ERC20 {
	return &ERC20{c: contract.NewContract(addr, abiERC20, provider)}
}

// Contract returns the contract object
func (e *ERC20) Contract() *contract.Contract {
	return e.c
}

// calls

// Allowance calls the allowance method in the solidity contract
func (e *ERC20) Allowance(owner ethgo.Address, spender ethgo.Address, block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("allowance", ethgo.EncodeBlock(block...), owner, spender)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// BalanceOf calls the balanceOf method in the solidity contract
func (e *ERC20) BalanceOf(owner ethgo.Address, block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("balanceOf", ethgo.EncodeBlock(block...), owner)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["balance"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Decimals calls the decimals method in the solidity contract
func (e *ERC20) Decimals(block ...ethgo.BlockNumber) (retval0 uint8, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("decimals", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(uint8)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Name calls the name method in the solidity contract
func (e *ERC20) Name(block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("name", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Symbol calls the symbol method in the solidity contract
func (e *ERC20) Symbol(block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("symbol", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// TotalSupply calls the totalSupply method in the solidity contract
func (e *ERC20) TotalSupply(block ...ethgo.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = e.c.Call("totalSupply", ethgo.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// Approve sends a approve transaction in the solidity contract
func (e *ERC20) Approve(spender ethgo.Address, value *big.Int) *contract.Txn {
	return e.c.Txn("approve", spender, value)
}

// Transfer sends a transfer transaction in the solidity contract
func (e *ERC20) Transfer(to ethgo.Address, value *big.Int) *contract.Txn {
	return e.c.Txn("transfer", to, value)
}

// TransferFrom sends a transferFrom transaction in the solidity contract
func (e *ERC20) TransferFrom(from ethgo.Address, to ethgo.Address, value *big.Int) *contract.Txn {
	return e.c.Txn("transferFrom", from, to, value)
}

// events

func (e *ERC20) ApprovalEventSig() ethgo.Hash {
	return e.c.ABI().Events["Approval"].ID()
}

func (e *ERC20) TransferEventSig() ethgo.Hash {
	return e.c.ABI().Events["Transfer"].ID()
}
