package wallet

import (
	"context"
	"github.com/google/uuid"
	"time"
	errDomain "tlab/src/domain/error"
)

const (
	TransferStatusSuccess string = "success"
	TransferStatusFail    string = "fail"
)

type (
	TrxID    string
	WalletId string
	UserId   string
	Wallet   struct {
		Id        string    `json:"id"`
		UserID    string    `json:"merchant_id"`
		Amount    int64     `json:"gross_amount"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Transfer struct {
		WalletID string `json:"wallet_id"`
		Amount   int64  `json:"amount" validate:"required,gt=0"`
	}
)

func (w *Wallet) CreateNewWallet(ctx context.Context, repositories WalletRepository) error {
	return repositories.CreateWallet(ctx, w)
}

func (w *WalletId) UpdateBalance(ctx context.Context, repositories WalletRepository, amount int64) error {
	return repositories.UpdateBalance(ctx, amount, string(*w))
}

func (t *Transfer) Transfer(ctx context.Context, walletRepo WalletRepository, trxLogRepo TransactionLogRepository) error {
	var (
		now    = time.Now()
		userId = UserId(ctx.Value("user_id").(string))
	)

	trxLogId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	receiverWalletId := WalletId(t.WalletID)

	receiverWallet, err := receiverWalletId.GetWallet(ctx, walletRepo)

	if err != nil {
		return err
	}

	currentWallet, err := userId.GetWallet(ctx, walletRepo)

	if err != nil {
		return err
	}

	currentWalletId := WalletId(currentWallet.Id)

	if currentWallet.Amount < t.Amount {
		log := &TransactionLog{
			ID:        trxLogId.String(),
			Sender:    currentWallet.Id,
			Receiver:  receiverWallet.Id,
			Amount:    t.Amount,
			Status:    TransferStatusFail,
			Reason:    errDomain.ErrInsufficientAmount.Error(),
			CreatedAt: now,
			UpdatedAt: now,
		}

		// if we want the log data not be lost better push it on queue process
		log.SaveLog(context.Background(), trxLogRepo)

		return errDomain.ErrInsufficientAmount
	}

	if err = receiverWalletId.UpdateBalance(ctx, walletRepo, receiverWallet.Amount+t.Amount); err != nil {
		log := &TransactionLog{
			ID:        trxLogId.String(),
			Sender:    currentWallet.Id,
			Receiver:  receiverWallet.Id,
			Amount:    t.Amount,
			Status:    TransferStatusFail,
			Reason:    errDomain.ErrInsufficientAmount.Error(),
			CreatedAt: now,
			UpdatedAt: now,
		}

		// if we want the log data not be lost better push it on queue process
		log.SaveLog(context.Background(), trxLogRepo)

		return errDomain.ErrInsufficientAmount
	}

	if err = currentWalletId.UpdateBalance(ctx, walletRepo, currentWallet.Amount-t.Amount); err != nil {
		log := &TransactionLog{
			ID:        trxLogId.String(),
			Sender:    currentWallet.Id,
			Receiver:  receiverWallet.Id,
			Amount:    t.Amount,
			Status:    TransferStatusFail,
			Reason:    errDomain.ErrTransferError.Error(),
			CreatedAt: now,
			UpdatedAt: now,
		}

		// if we want the log data not be lost better push it on queue process
		log.SaveLog(context.Background(), trxLogRepo)
		return err
	}

	log := &TransactionLog{
		ID:        trxLogId.String(),
		Sender:    currentWallet.Id,
		Receiver:  receiverWallet.Id,
		Amount:    t.Amount,
		Status:    TransferStatusSuccess,
		Reason:    "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// if we want the log data not be lost better push it on queue process
	log.SaveLog(ctx, trxLogRepo)

	return nil
}

func (wid *WalletId) GetWallet(ctx context.Context, repositories WalletRepository) (*Wallet, error) {
	wallet, err := repositories.GetWallet(ctx, &Wallet{Id: string(*wid)})

	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errDomain.ErrWalletNotFound
	}

	return wallet, nil
}

func (uid *UserId) GetWallet(ctx context.Context, repositories WalletRepository) (*Wallet, error) {
	wallet, err := repositories.GetWallet(ctx, &Wallet{UserID: string(*uid)})

	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errDomain.ErrWalletNotFound
	}

	return wallet, nil
}
