package main

import (
	"github.com/rs/zerolog/log"
	"github.com/smsunarto/cookie-clicker-rollup/execution/mempool"
	"github.com/smsunarto/cookie-clicker-rollup/execution/rpc"
	"github.com/smsunarto/cookie-clicker-rollup/execution/state"
	"time"
)

func main() {
	mp := mempool.NewMempool()

	// RPC
	rpcServer := rpc.NewRPCServer(mp)
	go rpcServer.Start()

	// State processor
	sp := state.NewStateProcessor(true)
	sdb := *state.NewMockStateDB()

	for {
		// Skip block if mempool is empty
		if mp.IsEmpty() {
			time.Sleep(1 * time.Second)
			continue
		}

		NewBlock(sp, sdb, mp)
		time.Sleep(1 * time.Second)
	}
}

func NewBlock(sp *state.StateProcessor, sdb state.StateDB, mp *mempool.Mempool) {
	block, err := sp.ProcessState(sdb, mp)
	if err != nil {
		log.Error().Err(err).Msg("Error processing state")
	} else {
		//log.Info().Msgf("Processed block %d", block.Number)
		log.Debug().Msgf("Block: %+v", block)
	}
}

func SubmitBlock() {
	// TODO: implement
	// Submits block to DA
}
