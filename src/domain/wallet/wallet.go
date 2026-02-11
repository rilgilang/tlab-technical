package wallet

import (
	"context"
	"time"
)

const (
	TransactionStatusPending string = "success"
	TransactionStatusFail    string = "fail"
)

type (
	TrxID  string
	Wallet struct {
		Id        string    `json:"id"`
		UserID    string    `json:"merchant_id"`
		Amount    float64   `json:"gross_amount"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (t *Wallet) CreateNewWallet(ctx context.Context, trxRepo WalletRepository) error {
	//return trxRepo.C(ctx, *t)
	return nil
}
