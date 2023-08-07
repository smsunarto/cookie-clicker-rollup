package state

import (
	"github.com/smsunarto/cookie-clicker-rollup/execution/mempool"
	"github.com/smsunarto/cookie-clicker-rollup/execution/sign"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
	"time"
)

type Block struct {
	Timestamp  int64
	Number     uint64
	StateHash  string
	Txs        []types.Transaction
	TxReceipts []types.TransactionReceipt
}

type StateProcessor struct {
	VerifySignature bool
}

func NewStateProcessor(verifySignature bool) *StateProcessor {
	return &StateProcessor{
		VerifySignature: verifySignature,
	}
}

// ProcessState processes the state and returns the new state and latest block
func (sp *StateProcessor) ProcessState(sdb StateDB, mp *mempool.Mempool) (Block, error) {
	var (
		txs        []types.Transaction
		txReceipts []types.TransactionReceipt
	)

	// For each transaction in mempool
	for {
		if mp.IsEmpty() {
			break
		}
		tx := mp.PopTX()

		// Apply transaction
		txReceipt := sp.applyTx(sdb, tx)

		// Record transaction and receipt
		txs = append(txs, tx)
		txReceipts = append(txReceipts, txReceipt)
	}

	sdb.SetBlockHeight(sdb.GetBlockHeight() + 1)
	block := Block{
		Timestamp:  time.Now().Unix(),
		Number:     sdb.GetBlockHeight(),
		StateHash:  sdb.CalculateStateHash(),
		Txs:        txs,
		TxReceipts: txReceipts,
	}

	// Store block in state db
	sdb.SetBlock(block.Number, block)

	return block, nil
}

func (sp *StateProcessor) applyTx(sdb StateDB, tx types.Transaction) types.TransactionReceipt {
	// Get account info from state db
	acct := sdb.GetAccount(tx.From)

	// If tx.Nonce is smaller or equal to the account latest nonce, ignore
	if sp.VerifySignature && tx.Nonce <= acct.Nonce {
		return tx.NewTxReceipt("invalid nonce")
	}

	// Verify signature
	if sp.VerifySignature && !sign.VerifyTransactionSignature(&tx) {
		return tx.NewTxReceipt("invalid signature")
	}

	// Update account info
	sdb.SetAccount(tx.From, tx.Nonce, acct.ClickCount+1)

	// Update cookie clicker count
	sdb.SetCookieClickerCount(sdb.GetCookieClickerCount() + 1)

	return tx.NewTxReceipt("")
}
