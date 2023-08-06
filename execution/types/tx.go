package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type TransactionHash string

type RawTransaction struct {
	Nonce   uint64
	ChainID string
	From    string
}

func (rawTx *RawTransaction) Hash() []byte {
	encRawTx, err := rlp.EncodeToBytes(rawTx)
	if err != nil {
		panic(err)
	}
	return crypto.Keccak256(encRawTx)
}

func (rawTx *RawTransaction) HashHex() TransactionHash {
	return TransactionHash(common.Bytes2Hex(rawTx.Hash()))
}

type Transaction struct {
	Nonce     uint64
	ChainID   string
	From      string
	Signature string
}

func (tx *Transaction) Hash() []byte {
	encTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		panic(err)
	}
	return crypto.Keccak256(encTx)
}

func (tx *Transaction) HashHex() TransactionHash {
	return TransactionHash(common.Bytes2Hex(tx.Hash()))
}

func (tx *Transaction) NewTxReceipt(error string) TransactionReceipt {
	return TransactionReceipt{
		Tx:     *tx,
		TxHash: tx.HashHex(),
		Error:  error,
	}
}

type TransactionReceipt struct {
	Tx     Transaction
	TxHash TransactionHash // Hash(Transaction)
	Error  string
}
