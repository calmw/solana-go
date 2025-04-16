package rpc

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetBlockHeight() (uint64, error) {
	endpoint := "https://api.devnet.solana.com"
	cli := rpc.New(endpoint)
	return cli.GetBlockHeight(context.Background(), rpc.CommitmentConfirmed)
}
