package types

import "github.com/ethereum/go-ethereum/common"

type Account struct {
	// AccountHash is stored in StateDB so we don't have to recompute it
	// everytime we calculate state hash
	AccountHash string // Hash(Nonce, ClickCount) -> AccountHash
	Address     common.Address
	Nonce       uint64
	ClickCount  uint64
}
