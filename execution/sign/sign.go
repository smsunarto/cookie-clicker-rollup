package sign

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
)

// SignRawTransaction signs a raw transaction using the private key
func SignRawTransaction(pk *ecdsa.PrivateKey, rawTx *types.RawTransaction) (*types.Transaction, error) {
	if rawTx == nil {
		return nil, errors.New("raw transaction is nil")
	}

	// Signed the encoded raw transaction using the private key
	buf, err := crypto.Sign(rawTx.Hash(), pk)
	if err != nil {
		return nil, err
	}

	tx := &types.Transaction{
		Nonce:     rawTx.Nonce,
		ChainID:   rawTx.ChainID,
		From:      rawTx.From,
		Signature: common.Bytes2Hex(buf),
	}

	return tx, nil
}

// VerifyTransactionSignature verifies the signature of a transaction
func VerifyTransactionSignature(tx *types.Transaction) bool {
	// Decode the raw transaction from the transaction
	rawTx := &types.RawTransaction{
		Nonce:   tx.Nonce,
		ChainID: tx.ChainID,
		From:    tx.From,
	}

	// Decode the signature from the transaction
	signerPubKey, err := crypto.SigToPub(rawTx.Hash(), common.Hex2Bytes(tx.Signature))
	if err != nil {
		return false
	}
	signerAddr := crypto.PubkeyToAddress(*signerPubKey).Hex()

	// Verify the signature using the public key
	if signerAddr == tx.From {
		return true
	}

	return false
}
