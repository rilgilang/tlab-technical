package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	"tlab/src/domain/wallet"
)

type TransactionLogRepository struct {
	DB *sqlx.DB
}

func NewTransactionLogRepository(db *sqlx.DB) *TransactionLogRepository {
	return &TransactionLogRepository{
		DB: db,
	}
}

func (r *TransactionLogRepository) CreateTransactionLog(ctx context.Context, payload wallet.TransactionLog) error {

	query := `
		INSERT INTO trx_log (
		                          id, 
		                          trx_id, 
		                          payment_gateway_payload,
		                          payment_gateway_response, 
		                          info_level,
		                     	  payload, 
		                          payment_gateway_callback_payload, 
		                          created_at, 
		                          updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`

	vals := []interface{}{
		payload.ID,
		payload.TrxID,
		payload.PaymentGatewayPayload,
		payload.PaymentGatewayResponse,
		payload.InfoLevel,
		payload.RequestPayload,
		"{}",
		payload.CreatedAt,
		payload.UpdatedAt,
	}

	stmt, err := GenerateStatement(ctx, r.DB, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, vals...)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionLogRepository) UpdateCallbackTransactionLog(ctx context.Context, callbackPayload []byte, trxId wallet.TrxID) error {

	query := `
		UPDATE  trx_log SET
			payment_gateway_callback_payload = $1, 
			updated_at = $2
		WHERE trx_id = $3
		`

	vals := []interface{}{
		callbackPayload,
		time.Now(),
		trxId,
	}

	stmt, err := GenerateStatement(ctx, r.DB, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, vals...); err != nil {
		return nil
	}

	return nil
}
