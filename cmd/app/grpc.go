package main

import (
	"fmt"
	"net"

	config "github.com/jorgeAM/grpc-kata-payment-service/cfg"
	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/application/command"
	paymentgrpc "github.com/jorgeAM/grpc-kata-payment-service/internal/payment/infrastructure/grpc"
	paymentpb "github.com/jorgeAM/grpc-kata-proto/gen/go/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg *config.Config, deps *config.Dependencies) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	paymentGRPCServer := paymentgrpc.NewPaymentGRPCServer(
		command.NewCreatePayment(deps.PaymentRepository),
	)

	grpcServer := grpc.NewServer(opts...)
	paymentpb.RegisterPaymentServer(grpcServer, paymentGRPCServer)

	if cfg.AppEnv == "local" {
		reflection.Register(grpcServer)
	}

	return grpcServer.Serve(lis)
}
