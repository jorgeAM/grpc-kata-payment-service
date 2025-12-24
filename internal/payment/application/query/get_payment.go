package query

import (
	"context"
	"errors"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/domain"
)

type GetPayment struct {
	paymentRepository domain.PaymentRepository
}

func NewGetPayment(paymentRepository domain.PaymentRepository) *GetPayment {
	return &GetPayment{
		paymentRepository: paymentRepository,
	}
}

func (g *GetPayment) Exec(ctx context.Context, paymentID string) (*domain.Payment, error) {
	if paymentID == "" {
		return nil, errors.New("payment id cannot be empty")
	}

	return g.paymentRepository.FindByID(ctx, paymentID)
}
