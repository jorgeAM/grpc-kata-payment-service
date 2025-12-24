package command

import (
	"context"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/domain"
	"github.com/jorgeAM/grpc-kata-payment-service/pkg/model"
)

type CreatePaymentCommand struct {
	CustomerID string  `json:"customer_id"`
	OrderId    string  `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
}

type CreatePayment struct {
	paymentRepository domain.PaymentRepository
}

func NewCreatePayment(paymentRepository domain.PaymentRepository) *CreatePayment {
	return &CreatePayment{
		paymentRepository: paymentRepository,
	}
}

func (c *CreatePayment) Exec(ctx context.Context, cmd *CreatePaymentCommand) (string, error) {
	orderID, err := model.NewID(cmd.OrderId)
	if err != nil {
		return "", err
	}

	customerID, err := model.NewID(cmd.CustomerID)
	if err != nil {
		return "", err
	}

	payment := domain.NewPayment(customerID, orderID, cmd.TotalPrice)

	if err := c.paymentRepository.Save(ctx, payment); err != nil {
		return "", err
	}

	return payment.ID.String(), nil
}
