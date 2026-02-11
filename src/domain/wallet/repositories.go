package wallet

import (
	"context"
)

type WalletRepository interface {
	GetWallet(ctx context.Context, payload *Wallet) (*Wallet, error)
}

type TransactionLogRepository interface {
	CreateTransactionLog(ctx context.Context, merchant TransactionLog) error
	UpdateCallbackTransactionLog(ctx context.Context, callbackPayload []byte, trxId TrxID) error
}
