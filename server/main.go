package main

import (
	"context"
	"log"
	"net"

	"github.com/RonaldCrb/blockchain/proto"
	"github.com/RonaldCrb/blockchain/server/blockchain"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("unable to listen on port 8080: %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})

	log.Print("Blockchain gRPC server online! 🤑")
	srv.Serve(listener)
}

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s Server) AddBlock(ctx context.Context, req *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AppendBlock(req.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s Server) GetBlockchain(ctx context.Context, req *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			Hash:          b.Hash,
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
