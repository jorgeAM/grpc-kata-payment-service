package domain

import (
	"context"

	"github.com/jorgeAM/grpc-kata-payment-service/pkg/model"
)

//go:generate mockgen -source=./payment.go -destination=../mock/payment.go -package=mock -mock_names=Repository=MockPaymentRepository
type PaymentRepository interface {
	Save(ctx context.Context, user *Payment) error
	FindByID(ctx context.Context, id string) (*Payment, error)
}

type Payment struct {
	ID         model.ID
	CustomerID model.ID
	Status     PaymentStatus
	OrderId    model.ID
	TotalPrice float32
	Timestamps model.Timestamps
}

func NewPayment(customerID, orderID model.ID, totalPrice float32) *Payment {
	return &Payment{
		ID:         model.GenerateUUID(),
		CustomerID: customerID,
		Status:     Pending,
		OrderId:    orderID,
		TotalPrice: totalPrice,
		Timestamps: model.NewTimestamps(),
	}
}
