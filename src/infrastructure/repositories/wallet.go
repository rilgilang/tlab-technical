package repositories

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"tlab/src/domain/wallet"
)

type WalletRepository struct {
	DB *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) GetWallet(ctx context.Context, payload *wallet.Wallet) (*wallet.Wallet, error) {
	filter := sq.Or{
		sq.Eq{"id": payload.Id},
		sq.Eq{"user_id": payload.UserID},
	}

	q := sq.Select(
		"id",
		"merchant_id",
		"gross_amount",
		"margin",
		"status",
		"payment_gateway_fee",
		"transaction_type",
		"net_revenue",
		"created_at",
		"updated_at",
	).From("transaction t").Where(filter)

	q = q.PlaceholderFormat(sq.Dollar)

	sql, i, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	wallet := wallet.Wallet{}
	err = stmt.QueryRowContext(ctx, i...).Scan()
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
