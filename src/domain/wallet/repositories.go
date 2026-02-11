package wallet

import (
	"context"
)

type WalletRepository interface {
	GetWallet(ctx context.Context, payload *Wallet) (*Wallet, error)
	CreateWallet(ctx context.Context, payload *Wallet) error
	UpdateBalance(ctx context.Context, balance int64, id string) error
}

type TransactionLogRepository interface {
	CreateTransactionLog(ctx context.Context, payload TransactionLog) error
	GetTransactionHistory(ctx context.Context, walletID string) ([]DetailedTransactionLog, error)
}
