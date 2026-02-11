package repositories

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/wallet"
)

type TransactionLogRepository struct {
	DB  *sqlx.DB
	log logger.Logger
}

func NewTransactionLogRepository(db *sqlx.DB, log logger.Logger) *TransactionLogRepository {
	return &TransactionLogRepository{
		DB:  db,
		log: log,
	}
}

func (r *TransactionLogRepository) CreateTransactionLog(ctx context.Context, payload wallet.TransactionLog) error {

	q := sq.Insert("trx_log").
		Columns(
			"id",
			"sender",
			"receiver",
			"amount",
			"status",
			"reason",
			"created_at",
			"updated_at").
		Values(
			&payload.ID,
			&payload.Sender,
			&payload.Receiver,
			&payload.Amount,
			&payload.Status,
			&payload.Reason,
			&payload.CreatedAt,
			&payload.UpdatedAt,
		)

	q = q.PlaceholderFormat(sq.Dollar)

	sql, i, err := q.ToSql()
	if err != nil {
		r.log.Error("error_create_trx_log", err)
		return err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sql)
	if err != nil {
		r.log.Error("error_create_trx_log", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, i...)
	if err != nil {
		r.log.Error("error_create_trx_log", err)
		return err
	}

	return nil
}

func (r *TransactionLogRepository) GetTransactionHistory(
	ctx context.Context,
	walletID string,
) ([]wallet.DetailedTransactionLog, error) {

	// Convert string walletID to UUID format
	q := sq.Select(
		"t.id",
		"t.sender",
		"su.name AS sender_name",
		"t.receiver",
		"ru.name AS receiver_name",
		"t.amount",
		"t.status",
		"t.reason",
		"t.created_at",
		"t.updated_at",
	).
		From("trx_log t").
		Join("wallet sw ON sw.id = t.sender").
		Join(`"user" su ON su.id = sw.user_id`).
		Join("wallet rw ON rw.id = t.receiver").
		Join(`"user" ru ON ru.id = rw.user_id`).
		Where(
			// Use ::uuid cast to explicitly convert the string to UUID
			sq.Expr("(t.sender = ?::uuid OR t.receiver = ?::uuid)", walletID, walletID),
		).
		OrderBy("t.created_at DESC")

	q = q.PlaceholderFormat(sq.Dollar)

	sqlString, args, err := q.ToSql()
	if err != nil {
		r.log.Error("error_get_trx_history", err)
		return nil, err
	}

	stmt, err := GenerateStatement(ctx, r.DB, sqlString)
	if err != nil {
		r.log.Error("error_get_trx_history", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		r.log.Error("error_get_trx_history_query", err)
		return nil, err
	}
	defer rows.Close()

	var logs []wallet.DetailedTransactionLog

	for rows.Next() {
		var log wallet.DetailedTransactionLog

		err := rows.Scan(
			&log.ID,
			&log.Sender,
			&log.SenderName,
			&log.Receiver,
			&log.ReceiverName,
			&log.Amount,
			&log.Status,
			&log.Reason,
			&log.CreatedAt,
			&log.UpdatedAt,
		)
		if err != nil {
			r.log.Error("error_get_trx_history_scan", err)
			return nil, err
		}

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		r.log.Error("error_get_trx_history_rows", err)
		return nil, err
	}

	return logs, nil
}
