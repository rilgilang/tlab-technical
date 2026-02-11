package container

import (
	"tlab/src/application"
	"tlab/src/domain/sharedkernel/jwt"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/sharedkernel/unitofwork"
	"tlab/src/infrastructure/repositories"

	"github.com/sarulabs/di/v2"
)

func NewApplication() *[]di.Def {
	return &[]di.Def{
		{
			Name: WalletApplicationDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				walletRepo := ctn.Get(WalletRepositoryDefName).(*repositories.WalletRepository)
				transactionLogRepo := ctn.Get(TransactionLogRepositoryDefName).(*repositories.TransactionLogRepository)
				log := ctn.Get(LoggerDefName).(logger.Logger)
				uow := ctn.Get(UnitOfWorkDefName).(unitofwork.UnitOfWork)
				return application.NewWallet(uow, log, walletRepo, transactionLogRepo), nil
			},
		},

		{
			Name: UserApplicationDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				userRepo := ctn.Get(UserRepositoryDefName).(*repositories.UserRepository)
				walletRepo := ctn.Get(WalletRepositoryDefName).(*repositories.WalletRepository)
				log := ctn.Get(LoggerDefName).(logger.Logger)
				uow := ctn.Get(UnitOfWorkDefName).(unitofwork.UnitOfWork)
				jwtMiddleware := ctn.Get(JWTDefName).(jwt.JWT)
				return application.NewUser(uow, jwtMiddleware, log, userRepo, walletRepo), nil
			},
		},
	}
}
