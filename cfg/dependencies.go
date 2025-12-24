package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	paymentDomain "github.com/jorgeAM/grpc-kata-payment-service/internal/payment/domain"
	paymentPersistence "github.com/jorgeAM/grpc-kata-payment-service/internal/payment/infrastructure/persistence"
	_ "github.com/lib/pq"
)

type Dependencies struct {
	PaymentRepository paymentDomain.PaymentRepository
}

func BuildDependencies(cfg *Config) (*Dependencies, error) {
	postgresClient, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return nil, err
	}

	postgresPaymentRepository := paymentPersistence.NewPostgresPaymentRepository(postgresClient)

	return &Dependencies{
		PaymentRepository: postgresPaymentRepository,
	}, nil
}
