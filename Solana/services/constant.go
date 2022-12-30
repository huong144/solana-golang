package services

import "github.com/portto/solana-go-sdk/rpc"

const (
	rpc_endpoint   = rpc.TestnetRPCEndpoint
	topic          = "order_created"
	topicAccount   = "account_output"
	broker1Address = "localhost:9093"
	nftOutTopic    = "nft_metadata_output"
	nftInTopic     = "nft_metadata_input"
)
