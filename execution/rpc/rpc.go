package rpc

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/smsunarto/cookie-clicker-rollup/execution/mempool"
	"github.com/smsunarto/cookie-clicker-rollup/execution/types"
	"io"
	"net/http"
)

type RPCServer struct {
	mp *mempool.Mempool
}

func NewRPCServer(mp *mempool.Mempool) *RPCServer {
	return &RPCServer{
		mp: mp,
	}
}

func (rpc *RPCServer) Start() {
	http.HandleFunc("/rpc/submit-tx", rpc.handleSubmitTx)
	log.Fatal().Err(http.ListenAndServe(":8080", nil))
}

func (rpc *RPCServer) handleSubmitTx(w http.ResponseWriter, r *http.Request) {
	var tx types.Transaction
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading JSON body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &tx)
	if err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	// Submit transaction to mempool
	txHash := rpc.mp.SubmitTX(tx)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(txHash))
}
