package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
)

type StateDB interface {
	GetBlock(blockNumber int) Block
	SetBlock(blockNumber int, block Block)

	// State hashed part of things
	CalculateStateHash() string

	GetBlockHeight() int
	SetBlockHeight(blockHeight int)

	GetCookieClickerCount() int
	SetCookieClickerCount(count int)

	GetAccount(address string) types.Account
	GetAggregateAccountHash() string
	SetAccount(address string, nonce int, clickCount int)
}

type MockStateDB struct {
	CookieClickerCount int
	BlockHeight        int
	Blocks             map[int]Block
	Accounts           map[string]types.Account
}

func NewMockStateDB() *MockStateDB {
	return &MockStateDB{
		CookieClickerCount: 0,
		Blocks:             map[int]Block{},
		Accounts:           map[string]types.Account{},
	}
}

func (msd *MockStateDB) GetBlock(blockNumber int) Block {
	return msd.Blocks[blockNumber]
}

func (msd *MockStateDB) SetBlock(blockNumber int, block Block) {
	msd.Blocks[blockNumber] = block
}

func (msd *MockStateDB) CalculateStateHash() string {
	state := struct {
		cookieClickerCount int
		blockHeight        int
		accountHash        []byte
	}{
		cookieClickerCount: msd.CookieClickerCount,
		blockHeight:        msd.BlockHeight,
		accountHash:        msd.GetAggregateAccountHash(),
	}

	encState, err := rlp.EncodeToBytes(state)
	if err != nil {
		panic(err)
	}

	return common.Bytes2Hex(crypto.Keccak256(encState))
}

func (msd *MockStateDB) GetBlockHeight() int {
	return msd.BlockHeight
}

func (msd *MockStateDB) SetBlockHeight(blockHeight int) {
	msd.BlockHeight = blockHeight
}

func (msd *MockStateDB) GetCookieClickerCount() int {
	return msd.CookieClickerCount
}

func (msd *MockStateDB) SetCookieClickerCount(count int) {
	msd.CookieClickerCount = count
}

func (msd *MockStateDB) GetAccount(address string) types.Account {
	return msd.Accounts[address]
}

func (msd *MockStateDB) GetAggregateAccountHash() []byte {
	// Iterate through all accounts and get their hashes
	// Then hash all the hashes together
	hash := crypto.NewKeccakState()
	for _, acct := range msd.Accounts {
		_, err := hash.Write(common.Hex2Bytes(acct.AccountHash))
		if err != nil {
			panic(err)
		}
	}
	return hash.Sum(nil)
}

func (msd *MockStateDB) SetAccount(address string, nonce int, clickCount int) {
	acctData := struct {
		nonce      int
		clickCount int
	}{
		nonce:      nonce,
		clickCount: clickCount,
	}

	encAcctData, err := rlp.EncodeToBytes(acctData)
	if err != nil {
		panic(err)
	}

	msd.Accounts[address] = types.Account{
		AccountHash: common.Bytes2Hex(crypto.Keccak256(encAcctData)),
		Nonce:       nonce,
		ClickCount:  clickCount,
	}
}
