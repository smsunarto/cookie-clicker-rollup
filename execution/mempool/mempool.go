package mempool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
)

type Mempool struct {
	transactions []types.Transaction
}

func NewMempool() *Mempool {
	return &Mempool{
		transactions: []types.Transaction{},
	}
}

func (mp *Mempool) SubmitTX(tx types.Transaction) types.TransactionHash {
	mp.transactions = append(mp.transactions, tx)
	return types.TransactionHash(common.Bytes2Hex(tx.Hash()))
}

func (mp *Mempool) PopTX() types.Transaction {
	transaction := mp.transactions[0]
	mp.transactions = mp.transactions[1:]
	return transaction
}

func (mp *Mempool) IsEmpty() bool {
	return len(mp.transactions) == 0
}
