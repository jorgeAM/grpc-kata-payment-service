package grpc

import (
	"context"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/application/command"
	paymentpb "github.com/jorgeAM/grpc-kata-proto/gen/go/payment/v1"
)

var _ paymentpb.PaymentServiceServer = (*PaymentGRPCServer)(nil)

type PaymentGRPCServer struct {
	createPaymentApp *command.CreatePayment
	*paymentpb.UnimplementedPaymentServiceServer
}

func NewPaymentGRPCServer(createPaymentApp *command.CreatePayment) *PaymentGRPCServer {
	return &PaymentGRPCServer{
		createPaymentApp:                  createPaymentApp,
		UnimplementedPaymentServiceServer: &paymentpb.UnimplementedPaymentServiceServer{},
	}
}

func (p *PaymentGRPCServer) Create(ctx context.Context, request *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	panic("unimplemented")
}
