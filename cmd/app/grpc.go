package main

import (
	"fmt"
	"net"

	config "github.com/jorgeAM/grpc-kata-payment-service/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg *config.Config, deps *config.Dependencies) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	orderGrpcServer := ordergrpc.NewOrderGrpcServer(
		command.NewCreateOrder(deps.OrderRepository),
	)

	grpcServer := grpc.NewServer(opts...)
	orderpb.RegisterOrderServiceServer(grpcServer, orderGrpcServer)

	if cfg.AppEnv == "local" {
		reflection.Register(grpcServer)
	}

	return grpcServer.Serve(lis)
}
