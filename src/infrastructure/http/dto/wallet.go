package dto

import "time"

type (
	Transfer struct {
		WalletID string `json:"wallet_id" validate:"required"`
		Amount   int64  `json:"amount" validate:"required,gt=0"`
	}

	TopUp struct {
		Amount int64 `json:"amount" validate:"required,gt=0"`
	}

	TopUpResponse struct {
		Amount  int64 `json:"amount"`
		Balance int64 `json:"balance"`
	}

	BalanceResponse struct {
		WalletId string `json:"wallet_id"`
		Balance  int64  `json:"balance"`
	}

	DetailedTransactionLog struct {
		ID           string    `json:"id,omitempty"`
		Sender       string    `json:"sender,omitempty"`
		Receiver     string    `json:"receiver,omitempty"`
		Amount       int64     `json:"amount,omitempty"`
		SenderName   string    `json:"sender_name"`
		ReceiverName string    `json:"received_name"`
		Status       string    `json:"status"`
		Reason       string    `json:"reason"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	}

	TransactionLog struct {
		ID        string    `json:"id,omitempty"`
		Sender    string    `json:"sender,omitempty"`
		Receiver  string    `json:"receiver,omitempty"`
		Amount    int64     `json:"amount,omitempty"`
		Status    string    `json:"status"`
		Reason    string    `json:"reason"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	TransferResponse struct {
		WalletID          string `json:"wallet_id"`
		TransferredAmount int64  `json:"transferred_amount"`
	}
)
