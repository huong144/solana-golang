package services

import (
	"context"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"log"
)

func GetTransactionReceipt() *client.GetTransactionResponse {
	c := client.NewClient(rpc.TestnetRPCEndpoint)
	var (
		txHash = string("4b9DWhuwWd4MP393MegFS2q6xNtSxmFim6rUDutDTF8EDLDaXjQkGz69yZgbV4p61N2oQeK1U4Af6qhd9nH8NLBV")
	)

	res, error := c.GetTransaction(context.TODO(), txHash)
	if error != nil {
		log.Fatalln("failed to get tx Solana, err: %v", error)
	}
	return res
}
