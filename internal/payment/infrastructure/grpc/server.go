package grpc

import (
	"context"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/application/command"
	paymentpb "github.com/jorgeAM/grpc-kata-proto/gen/go/payment/v1"
)

var _ paymentpb.PaymentServer = (*PaymentGRPCServer)(nil)

type PaymentGRPCServer struct {
	createPaymentApp *command.CreatePayment
	*paymentpb.UnimplementedPaymentServer
}

func NewPaymentGRPCServer(createPaymentApp *command.CreatePayment) *PaymentGRPCServer {
	return &PaymentGRPCServer{
		createPaymentApp:           createPaymentApp,
		UnimplementedPaymentServer: &paymentpb.UnimplementedPaymentServer{},
	}
}

func (p *PaymentGRPCServer) Create(ctx context.Context, request *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	panic("unimplemented")
}
