package grpc

import (
	"context"
	"fmt"

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
	cmd := command.CreatePaymentCommand{
		CustomerID: request.UserId,
		OrderId:    request.OrderId,
		TotalPrice: request.TotalPrice,
	}

	paymentID, err := p.createPaymentApp.Exec(ctx, &cmd)
	if err != nil {
		return nil, err
	}

	return &paymentpb.CreatePaymentResponse{
		PaymentId: paymentID,
		BillId:    fmt.Sprintf("bill-%s", paymentID),
	}, nil
}
