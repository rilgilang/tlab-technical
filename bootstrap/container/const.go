package container

// Config and package
const (
	ContainerDefName  string = "container"
	ConfigDefName     string = "config"
	DBDefName         string = "db"
	UnitOfWorkDefName string = "unitOfWork"
	LoggerDefName     string = "logger"
	EchoDefName       string = "echo"
	ValidatorDefName  string = "validator"

	JWTDefName string = "client.jwt"

	WalletRepositoryDefName         string = "repo.wallet"
	UserRepositoryDefName           string = "repo.user"
	TransactionLogRepositoryDefName string = "repo.transaction_log"
)

// Middleware
const (
	AuthMiddlewareDefName string = "authMiddleware"
)

// Application
const (
	UserApplicationDefName   string = "application.user"
	WalletApplicationDefName string = "application.wallet"
)
