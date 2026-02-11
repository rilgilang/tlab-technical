package container

import (
	"tlab/src/domain/sharedkernel/logger"
	repositories2 "tlab/src/infrastructure/repositories"

	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di/v2"
)

func NewRepository() *[]di.Def {
	return &[]di.Def{
		{
			Name: WalletRepositoryDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories2.NewWalletRepository(
					ctn.Get(DBDefName).(*sqlx.DB),
					ctn.Get(LoggerDefName).(logger.Logger),
				), nil
			},
		},

		{
			Name: UserRepositoryDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories2.NewUserRepository(
					ctn.Get(DBDefName).(*sqlx.DB),
					ctn.Get(LoggerDefName).(logger.Logger),
				), nil
			},
		},

		{
			Name: TransactionLogRepositoryDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories2.NewTransactionLogRepository(
					ctn.Get(DBDefName).(*sqlx.DB),
					ctn.Get(LoggerDefName).(logger.Logger),
				), nil
			},
		},
	}
}
