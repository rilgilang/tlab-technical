package database

import (
	"fmt"
	"time"
	"tlab/bootstrap/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sarulabs/di/v2"
)

func LoadDatabase() *[]di.Def {
	return &[]di.Def{
		{
			Name:  "db",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				cfg := ctn.Get("config").(config.Config)

				// Generate DSN string from config
				var generateConnectionString = func() string {
					return fmt.Sprintf(
						"host=%s port=%v dbname=%s user=%s password=%s sslmode=%s application_name=%s",
						cfg.DBHost,
						cfg.DBPort,
						cfg.DBName,
						cfg.DBUser,
						cfg.DBPassword,
						cfg.DBSSLMode,
						cfg.DBApplicationName,
					)
				}
				db, err := sqlx.Connect("postgres", generateConnectionString())
				if err != nil {
					fmt.Println("error : ", err)
					return nil, err
				}
				db.SetMaxOpenConns(50)
				db.SetConnMaxLifetime(time.Minute * 15)
				db.SetMaxIdleConns(10)
				return db, err
			},
			Close: func(obj interface{}) error {
				return obj.(*sqlx.DB).Close()
			},
		},
	}
}
