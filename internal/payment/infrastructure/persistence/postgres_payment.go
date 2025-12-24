package persistence

import (
	"time"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/domain"
	"github.com/jorgeAM/grpc-kata-payment-service/pkg/model"
)

type postgresPayment struct {
	ID         string     `db:"id"`
	CustomerID string     `db:"customer_id"`
	Status     string     `db:"status"`
	OrderID    string     `db:"order_id"`
	TotalPrice float32    `db:"total_price"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" goqu:"omitnil"`
}

func (dto postgresPayment) toDomain() (*domain.Payment, error) {
	id, err := model.NewID(dto.ID)
	if err != nil {
		return nil, err
	}

	customerID, err := model.NewID(dto.CustomerID)
	if err != nil {
		return nil, err
	}

	orderID, err := model.NewID(dto.OrderID)
	if err != nil {
		return nil, err
	}

	status, err := domain.NewOrderStatus(dto.Status)
	if err != nil {
		return nil, err
	}

	return &domain.Payment{
		ID:         id,
		CustomerID: customerID,
		Status:     status,
		OrderId:    orderID,
		TotalPrice: dto.TotalPrice,
		Timestamps: model.Timestamps{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
			DeletedAt: dto.DeletedAt,
		},
	}, nil
}

func fromDomain(entity *domain.Payment) (*postgresPayment, error) {
	return &postgresPayment{
		ID:         entity.ID.String(),
		CustomerID: entity.CustomerID.String(),
		Status:     entity.Status.String(),
		OrderID:    entity.OrderId.String(),
		TotalPrice: entity.TotalPrice,
		CreatedAt:  entity.Timestamps.CreatedAt,
		UpdatedAt:  entity.Timestamps.UpdatedAt,
		DeletedAt:  entity.Timestamps.DeletedAt,
	}, nil
}
