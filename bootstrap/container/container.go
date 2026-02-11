package container

import (
	"tlab/bootstrap/config"
	"tlab/bootstrap/database"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/infrastructure/pkg/jwt"
	repositories2 "tlab/src/infrastructure/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
)

func NewContainer() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return di.Container{}, err
	}

	defs := []di.Def{
		{
			Name: ConfigDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.LoadConfig()
			},
		},
		{
			Name: LoggerDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return logger.NewLogger(), nil
			},
		},
		{
			Name: UnitOfWorkDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				uow := repositories2.NewUnitOfWork(ctn.Get(DBDefName).(*sqlx.DB))
				return uow, nil
			},
		},
		{
			Name: ValidatorDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return validator.New(), nil
			},
		},
		{
			Name: EchoDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				e := echo.New()
				validate := ctn.Get(ValidatorDefName).(*validator.Validate)
				e.Validator = &CustomValidator{validator: validate}
				return e, nil
			},
		},

		{
			Name: JWTDefName,
			Build: func(ctn di.Container) (interface{}, error) {
				return jwt.NewJWT(ctn.Get(ConfigDefName).(config.Config)), nil
			},
		},
	}

	if err := builder.Add(defs...); err != nil {
		return di.Container{}, err
	}

	if err := builder.Add(*NewApplication()...); err != nil {
		return di.Container{}, err
	}

	if err := builder.Add(*NewRepository()...); err != nil {
		return di.Container{}, err
	}

	if err := builder.Add(*database.LoadDatabase()...); err != nil {
		return di.Container{}, err
	}

	return builder.Build(), nil
}

// CustomValidator adalah custom validator untuk Echo
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
