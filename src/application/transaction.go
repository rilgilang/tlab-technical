package application

import (
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/sharedkernel/unitofwork"
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

//func (a *Transaction) Charge(ctx context.Context, payload dto.ChargeTransaction) (*dto.ChargeResponse, error) {
//	var result = &dto.ChargeResponse{}
//	uowResult, err := a.uow.Execute(ctx, func(ctx context.Context) (result *unitofwork.Result, err error) {
//		//
//		trxId, err := uuid.NewUUID()
//		if err != nil {
//			return nil, err
//		}
//
//		now := time.Now()
//
//		trx := &transaction.Transaction{
//			Id:                  trxId.String(),
//			MerchantId:          payload.MerchantId,
//			GrossAmount:         payload.Amount,
//			ClientTransactionID: payload.ClientTransactionID,
//			Status:              transaction.TransactionStatusPending,
//			TransactionType:     transaction.TransactionTypePayment,
//			Margin:              1000,
//			PaymentGatewayFee:   1000,
//			NetRevenue:          1000,
//			CreatedAt:           now,
//			UpdatedAt:           now,
//		}
//
//		if err := trx.SaveNewTransaction(ctx, a.transactionRepo); err != nil {
//			a.logger.Error("create_transaction_error: ", err)
//			return nil, err
//		}
//
//		payload.ExternalId = trxId.String()
//
//		chargeTransaction := transaction.ChargeTransaction{
//			ExternalId:     payload.ExternalId,
//			MerchantId:     payload.MerchantId,
//			Amount:         payload.Amount,
//			Description:    payload.Description,
//			PaymentMethod:  payload.PaymentMethod,
//			VirtualAccount: transaction.ChargeTransactionVA(payload.VirtualAccount),
//		}
//
//		chargePayloadBytes, chargeResponseBytes, err := chargeTransaction.ChargeToPaymentGateway(ctx, a.xendit)
//		if err != nil {
//			a.logger.Error("charge_transaction_error: ", err)
//			return nil, err
//		}
//
//		trxLogId, err := uuid.NewUUID()
//		if err != nil {
//			return nil, err
//		}
//
//		requestPayloadBytes, err := json.Marshal(trx)
//		if err != nil {
//			return nil, err
//		}
//
//		if err = a.transactionLogRepo.CreateTransactionLog(ctx, transaction.TransactionLog{
//			ID:                            trxLogId.String(),
//			TrxID:                         trxId.String(),
//			PaymentGatewayPayload:         chargePayloadBytes,
//			PaymentGatewayResponse:        chargeResponseBytes,
//			InfoLevel:                     transaction.TransactionLogLevelInfo,
//			RequestPayload:                requestPayloadBytes,
//			PaymentGatewayCallbackPayload: nil,
//			CreatedAt:                     now,
//			UpdatedAt:                     now,
//		}); err != nil {
//			a.logger.Error("create_transaction_log_error: ", err)
//			return nil, err
//		}
//
//		chargeResponse := &dto.ChargeResponse{}
//
//		err = json.Unmarshal(chargeResponseBytes, chargeResponse)
//		if err != nil {
//			return nil, err
//		}
//
//		return &unitofwork.Result{
//			Body: chargeResponse,
//		}, err
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	// Handle unit of work result
//	if uowResult != nil && uowResult.Body != nil {
//		if t, ok := uowResult.Body.(*dto.ChargeResponse); ok {
//			result = t
//		}
//	}
//
//	return result, nil
//}
