package main

import (
	"github.com/smsunarto/cookie-clicker-rollup/execution/mempool"
	"github.com/smsunarto/cookie-clicker-rollup/execution/rpc"
)

func main() {
	mp := mempool.NewMempool()
	rpcServer := rpc.NewRPCServer(mp)
	rpcServer.Start()
}

func NewBlock() {
	// TODO: implement
	// Get transactions from mempool and handle them
}

func SubmitBlock() {
	// TODO: implement
	// Submits block to DA
}
