package application

import (
	"context"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/sharedkernel/unitofwork"
	"tlab/src/domain/wallet"
	"tlab/src/infrastructure/http/dto"
	"tlab/src/infrastructure/repositories"
)

type Wallet struct {
	uow                unitofwork.UnitOfWork
	logger             logger.Logger
	walletRepo         *repositories.WalletRepository
	transactionLogRepo *repositories.TransactionLogRepository
}

func NewWallet(
	uow unitofwork.UnitOfWork,
	logger logger.Logger,
	walletRepo *repositories.WalletRepository,
	transactionLogRepo *repositories.TransactionLogRepository,
) *Wallet {
	return &Wallet{
		uow:                uow,
		logger:             logger,
		walletRepo:         walletRepo,
		transactionLogRepo: transactionLogRepo,
	}
}

func (a *Wallet) Transfer(ctx context.Context, payload dto.Transfer) (*dto.TransferResponse, error) {
	var (
		result = &dto.TransferResponse{}
	)
	uowResult, err := a.uow.Execute(ctx, func(ctx context.Context) (result *unitofwork.Result, err error) {
		trx := &wallet.Transfer{
			WalletID: payload.WalletID,
			Amount:   payload.Amount,
		}

		err = trx.Transfer(ctx, a.walletRepo, a.transactionLogRepo)

		if err != nil {
			return nil, err
		}

		return &unitofwork.Result{
			Body: &dto.TransferResponse{
				WalletID:          payload.WalletID,
				TransferredAmount: payload.Amount,
			},
		}, err
	})

	if err != nil {
		return nil, err
	}

	// Handle unit of work result
	if uowResult != nil && uowResult.Body != nil {
		if t, ok := uowResult.Body.(*dto.TransferResponse); ok {
			result = t
		}
	}

	return result, nil
}

func (a *Wallet) TopUpBalance(ctx context.Context, payload dto.TopUp) (*dto.TopUpResponse, error) {
	var (
		userId = wallet.UserId(ctx.Value("user_id").(string))
	)

	currentWallet, err := userId.GetWallet(ctx, a.walletRepo)

	if err != nil {
		return nil, err
	}

	currentWalletId := wallet.WalletId(currentWallet.Id)

	newAmount := currentWallet.Amount + payload.Amount

	err = currentWalletId.UpdateBalance(ctx, a.walletRepo, newAmount)
	if err != nil {
		return nil, err
	}

	return &dto.TopUpResponse{
		Amount:  payload.Amount,
		Balance: newAmount,
	}, nil
}
func (a *Wallet) GetTransactionHistory(ctx context.Context) ([]dto.DetailedTransactionLog, error) {
	var (
		userId = wallet.UserId(ctx.Value("user_id").(string))
	)

	currentWallet, err := userId.GetWallet(ctx, a.walletRepo)

	if err != nil {
		return nil, err
	}

	walletId := wallet.WalletId(currentWallet.Id)

	list, err := walletId.TransactionList(ctx, a.transactionLogRepo)
	if err != nil {
		return nil, err
	}

	response := []dto.DetailedTransactionLog{}

	for _, log := range list {
		response = append(response, dto.DetailedTransactionLog{
			ID:           log.ID,
			Sender:       log.Sender,
			Receiver:     log.Receiver,
			Amount:       log.Amount,
			SenderName:   log.SenderName,
			ReceiverName: log.ReceiverName,
			Status:       log.Status,
			Reason:       log.Reason,
			CreatedAt:    log.CreatedAt,
			UpdatedAt:    log.UpdatedAt,
		})
	}

	return response, nil
}

func (a *Wallet) GetBalance(ctx context.Context) (*dto.BalanceResponse, error) {
	var (
		userId = wallet.UserId(ctx.Value("user_id").(string))
	)

	currentWallet, err := userId.GetWallet(ctx, a.walletRepo)

	if err != nil {
		return nil, err
	}

	return &dto.BalanceResponse{
		WalletId: currentWallet.Id,
		Balance:  currentWallet.Amount,
	}, nil
}
