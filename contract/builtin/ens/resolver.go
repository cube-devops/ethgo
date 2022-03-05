// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: 0c04761c33a894f2ef98ca34c1107d52c1e79c1913d2e632ceec744a37453c33
// Version: 0.1.0
package ens

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

// Resolver is a solidity contract
type Resolver struct {
	c *contract.Contract
}

// DeployResolver deploys a new Resolver contract
func DeployResolver(provider *jsonrpc.Client, from ethgo.Address, args ...interface{}) *contract.Txn {
	return contract.DeployContract(provider, from, abiResolver, binResolver, args...)
}

// NewResolver creates a new instance of the contract at a specific address
func NewResolver(addr ethgo.Address, provider *jsonrpc.Client) *Resolver {
	return &Resolver{c: contract.NewContract(addr, abiResolver, provider)}
}

// Contract returns the contract object
func (r *Resolver) Contract() *contract.Contract {
	return r.c
}

// calls

// ABI calls the ABI method in the solidity contract
func (r *Resolver) ABI(node [32]byte, contentTypes *big.Int, block ...ethgo.BlockNumber) (retval0 *big.Int, retval1 []byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("ABI", ethgo.EncodeBlock(block...), node, contentTypes)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["contentType"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	retval1, ok = out["data"].([]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 1")
		return
	}
	
	return
}

// Addr calls the addr method in the solidity contract
func (r *Resolver) Addr(node [32]byte, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("addr", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Content calls the content method in the solidity contract
func (r *Resolver) Content(node [32]byte, block ...ethgo.BlockNumber) (retval0 [32]byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("content", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Name calls the name method in the solidity contract
func (r *Resolver) Name(node [32]byte, block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("name", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Pubkey calls the pubkey method in the solidity contract
func (r *Resolver) Pubkey(node [32]byte, block ...ethgo.BlockNumber) (retval0 [32]byte, retval1 [32]byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("pubkey", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["x"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	retval1, ok = out["y"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 1")
		return
	}
	
	return
}

// SupportsInterface calls the supportsInterface method in the solidity contract
func (r *Resolver) SupportsInterface(interfaceID [4]byte, block ...ethgo.BlockNumber) (retval0 bool, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("supportsInterface", ethgo.EncodeBlock(block...), interfaceID)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(bool)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// SetABI sends a setABI transaction in the solidity contract
func (r *Resolver) SetABI(node [32]byte, contentType *big.Int, data []byte) *contract.Txn {
	return r.c.Txn("setABI", node, contentType, data)
}

// SetAddr sends a setAddr transaction in the solidity contract
func (r *Resolver) SetAddr(node [32]byte, addr ethgo.Address) *contract.Txn {
	return r.c.Txn("setAddr", node, addr)
}

// SetContent sends a setContent transaction in the solidity contract
func (r *Resolver) SetContent(node [32]byte, hash [32]byte) *contract.Txn {
	return r.c.Txn("setContent", node, hash)
}

// SetName sends a setName transaction in the solidity contract
func (r *Resolver) SetName(node [32]byte, name string) *contract.Txn {
	return r.c.Txn("setName", node, name)
}

// SetPubkey sends a setPubkey transaction in the solidity contract
func (r *Resolver) SetPubkey(node [32]byte, x [32]byte, y [32]byte) *contract.Txn {
	return r.c.Txn("setPubkey", node, x, y)
}

// events

func (r *Resolver) ABIChangedEventSig() ethgo.Hash {
	return r.c.ABI().Events["ABIChanged"].ID()
}

func (r *Resolver) AddrChangedEventSig() ethgo.Hash {
	return r.c.ABI().Events["AddrChanged"].ID()
}

func (r *Resolver) ContentChangedEventSig() ethgo.Hash {
	return r.c.ABI().Events["ContentChanged"].ID()
}

func (r *Resolver) NameChangedEventSig() ethgo.Hash {
	return r.c.ABI().Events["NameChanged"].ID()
}

func (r *Resolver) PubkeyChangedEventSig() ethgo.Hash {
	return r.c.ABI().Events["PubkeyChanged"].ID()
}
