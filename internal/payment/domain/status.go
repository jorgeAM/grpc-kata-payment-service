package domain

import (
	"strings"

	"github.com/jorgeAM/grpc-kata-payment-service/pkg/errors"
)

var (
	ErrInvalidPaymentStatus = errors.Define("payment.invalid_status")
)

type PaymentStatus string

const (
	Pending PaymentStatus = "PENDING"
)

var allowedPaymentStatus = map[string]PaymentStatus{
	Pending.String(): Pending,
}

func NewOrderStatus(t string) (PaymentStatus, error) {
	if status, ok := allowedPaymentStatus[strings.ToUpper(t)]; ok {
		return status, nil
	}

	return "", errors.New(
		ErrInvalidPaymentStatus,
		"invalid payment status",
		errors.WithMetadata("payment_status", t),
	)
}

func (o PaymentStatus) String() string {
	return string(o)
}
