package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAccountInfo(addr string) {
	params := [2]interface{}{addr, map[string]string{"encoding": "jsonParsed"}}
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "getAccountInfo",
		"params":  params,
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(rpc_endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	emitAccountDataKafka(body, addr)

	fmt.Println(body)
}

func emitAccountDataKafka(data []byte, signature string) {
	ctx := context.Background()

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topicAccount,
	})

	//reqBodyBytes := new(bytes.Buffer)
	//json.NewEncoder(reqBodyBytes).Encode(data)
	//item := reqBodyBytes.Bytes()
	//fmt.Println(item)

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(signature),
		// create an arbitrary message payload for the value
		Value: data,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func GetMetaData(add string) {

	// NFT in solana is a normal mint but only mint 1.
	// If you want to get its metadata, you need to know where it stored.
	// and you can use `tokenmeta.GetTokenMetaPubkey` to get the metadata account key
	// here I take a random Degenerate Ape Academy as an example
	mint := common.PublicKeyFromString(add)
	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}

	// new a client
	c := client.NewClient(rpc_endpoint)

	// get data which stored in metadataAccount
	accountInfo, err := c.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		log.Fatalf("failed to get accountInfo, err: %v", err)
	}

	// parse it
	metadata, err := tokenmeta.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		log.Fatalf("failed to parse metaAccount, err: %v", err)
	}
	//fmt.Println(metadata)
	//spew.Dump(metadata)
	emitMetadata(metadata, add)
}

func emitMetadata(data tokenmeta.Metadata, address string) {
	spew.Dump(data)
	//UpdateAuthority := data.UpdateAuthority.String()
	//Mint := data.Mint.String()

	ctx := context.Background()

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   nftOutTopic,
	})

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	item := reqBodyBytes.Bytes()
	//fmt.Println(item)

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(address),
		// create an arbitrary message payload for the value
		Value: item,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
