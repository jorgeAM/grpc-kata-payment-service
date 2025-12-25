package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/jorgeAM/grpc-kata-payment-service/internal/payment/domain"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var _ domain.PaymentRepository = (*PostgresPaymentRepository)(nil)

type PostgresPaymentRepository struct {
	db     *sqlx.DB
	schema string
	table  string
}

func NewPostgresPaymentRepository(db *sqlx.DB) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{
		db:     db,
		schema: "payment_schema",
		table:  "payments",
	}
}

func (p *PostgresPaymentRepository) FindByID(ctx context.Context, id string) (*domain.Payment, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.
		From(fmt.Sprintf("%s.%s", p.schema, p.table)).
		Where(goqu.Ex{
			"id": id,
		})

	sqlQuery, args, err := ds.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	var dto postgresPayment
	if err := p.db.GetContext(
		ctx,
		&dto,
		sqlQuery,
		args...,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment not found")
		}

		return nil, err
	}

	return dto.toDomain()

}

func (p *PostgresPaymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	dto, err := fromDomain(payment)
	if err != nil {
		return err
	}

	ds := goqu.
		Insert(fmt.Sprintf("%s.%s", p.schema, p.table)).
		Rows(dto).
		OnConflict(goqu.DoUpdate("id", goqu.Record{
			"customer_id": dto.CustomerID,
			"status":      dto.Status,
			"order_id":    dto.OrderID,
			"total_price": dto.TotalPrice,
			"updated_at":  dto.UpdatedAt,
		}))

	sql, _, err := ds.ToSQL()
	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}

	return nil
}
