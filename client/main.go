package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/RonaldCrb/blockchain/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "create and append new block to the blockchain")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("[ERROR] missing subcommand")
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERROR] no connection => %v", err)
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("[ERROR] cant append block => %v", err)
	}
	log.Printf("[SUCCESS] block %v has been succesfully appended to the server blockchain", block.Hash)
}

func getBlockchain() {
	blockchain, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("[ERROR] cant get to the blockchain => %s", err)
	}

	for _, b := range blockchain.Blocks {
		log.Printf("[SUCCESS]\n Hash: %s\n Previous Block Hash: %s\n Data: %s\n ===================================================================================================\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
