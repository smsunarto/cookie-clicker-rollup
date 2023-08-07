package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rs/zerolog/log"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
)

type StateDB interface {
	GetBlock(blockNumber uint64) Block
	SetBlock(blockNumber uint64, block Block)

	// State hashed part of things
	CalculateStateHash() string

	GetBlockHeight() uint64
	SetBlockHeight(blockHeight uint64)

	GetCookieClickerCount() uint64
	SetCookieClickerCount(count uint64)

	GetAccount(address common.Address) types.Account
	GetAggregateAccountHash() []byte
	SetAccount(address common.Address, nonce uint64, clickCount uint64)
}

type MockStateDB struct {
	CookieClickerCount uint64
	BlockHeight        uint64
	Blocks             map[uint64]Block
	Accounts           map[common.Address]types.Account
}

func NewMockStateDB() *StateDB {
	var sdb StateDB
	sdb = &MockStateDB{
		CookieClickerCount: 0,
		BlockHeight:        0,
		Blocks:             map[uint64]Block{},
		Accounts:           map[common.Address]types.Account{},
	}

	return &sdb
}

func (msd *MockStateDB) GetBlock(blockNumber uint64) Block {
	return msd.Blocks[blockNumber]
}

func (msd *MockStateDB) SetBlock(blockNumber uint64, block Block) {
	msd.Blocks[blockNumber] = block
}

func (msd *MockStateDB) CalculateStateHash() string {
	state := struct {
		CookieClickerCount uint64
		BlockHeight        uint64
		AccountHash        []byte
	}{
		CookieClickerCount: msd.CookieClickerCount,
		BlockHeight:        msd.BlockHeight,
		AccountHash:        msd.GetAggregateAccountHash(),
	}

	encState, err := rlp.EncodeToBytes(state)
	if err != nil {
		panic(err)
	}

	return common.Bytes2Hex(crypto.Keccak256(encState))
}

func (msd *MockStateDB) GetBlockHeight() uint64 {
	return msd.BlockHeight
}

func (msd *MockStateDB) SetBlockHeight(blockHeight uint64) {
	msd.BlockHeight = blockHeight
}

func (msd *MockStateDB) GetCookieClickerCount() uint64 {
	return msd.CookieClickerCount
}

func (msd *MockStateDB) SetCookieClickerCount(count uint64) {
	msd.CookieClickerCount = count
}

func (msd *MockStateDB) GetAccount(address common.Address) types.Account {
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

func (msd *MockStateDB) SetAccount(address common.Address, nonce uint64, clickCount uint64) {
	acctData := struct {
		Address    common.Address
		Nonce      uint64
		ClickCount uint64
	}{
		Address:    address,
		Nonce:      nonce,
		ClickCount: clickCount,
	}

	encAcctData, err := rlp.EncodeToBytes(acctData)
	if err != nil {
		panic(err)
	}

	log.Info().Msg("Setting account hash")
	msd.Accounts[address] = types.Account{
		AccountHash: common.Bytes2Hex(crypto.Keccak256(encAcctData)),
		Address:     address,
		Nonce:       nonce,
		ClickCount:  clickCount,
	}
}
