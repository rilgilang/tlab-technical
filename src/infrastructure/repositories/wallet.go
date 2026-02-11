package repositories

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"time"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/wallet"
)

type WalletRepository struct {
	DB  *sqlx.DB
	log logger.Logger
}

func NewWalletRepository(db *sqlx.DB, log logger.Logger) *WalletRepository {
	return &WalletRepository{
		DB:  db,
		log: log,
	}
}

func (r *WalletRepository) GetWallet(ctx context.Context, payload *wallet.Wallet) (*wallet.Wallet, error) {
	filter := sq.And{}

	if payload.Id != "" {
		filter = sq.And{
			sq.Eq{"id": payload.Id},
		}
	} else {
		filter = sq.And{
			sq.Eq{"user_id": payload.UserID},
		}
	}

	q := sq.Select(
		"id",
		"user_id",
		"amount",
		"created_at",
		"updated_at",
	).
		From("wallet w").
		Where(filter)

	q = q.PlaceholderFormat(sq.Dollar)

	sqlString, i, err := q.ToSql()
	if err != nil {
		r.log.Error("error_get_wallet", err)
		return nil, err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sqlString)
	if err != nil {
		r.log.Error("error_get_wallet", err)
		return nil, err
	}

	defer stmt.Close()

	wallet := wallet.Wallet{}
	err = stmt.QueryRowContext(ctx, i...).Scan(
		&wallet.Id,
		&wallet.UserID,
		&wallet.Amount,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.log.Error("error_get_wallet", err)
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) CreateWallet(ctx context.Context, payload *wallet.Wallet) error {

	q := sq.Insert("wallet").
		Columns(
			"id",
			"user_id",
			"amount",
			"created_at",
			"updated_at").
		Values(
			&payload.Id,
			&payload.UserID,
			&payload.Amount,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)

	q = q.PlaceholderFormat(sq.Dollar)

	sql, i, err := q.ToSql()
	if err != nil {
		r.log.Error("error_create_wallet", err)
		return err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sql)
	if err != nil {
		r.log.Error("error_create_wallet", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, i...)
	if err != nil {
		r.log.Error("error_create_wallet", err)
		return err
	}

	return nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, balance int64, id string) error {

	q := sq.Update("wallet").
		Set("amount", balance).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": id})

	q = q.PlaceholderFormat(sq.Dollar)

	sql, i, err := q.ToSql()
	if err != nil {
		r.log.Error("error_update_wallet", err)
		return err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sql)
	if err != nil {
		r.log.Error("error_update_wallet", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, i...)
	if err != nil {
		r.log.Error("error_update_wallet", err)
		return err
	}

	return nil
}
