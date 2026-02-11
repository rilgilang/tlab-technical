package repositories

import (
	"context"
	"tlab/src/domain/sharedkernel/unitofwork"

	"github.com/jmoiron/sqlx"
)

type UnitOfWork struct {
	db *sqlx.DB
}

func NewUnitOfWork(db *sqlx.DB) *UnitOfWork {
	return &UnitOfWork{
		db,
	}
}

func (u *UnitOfWork) Execute(ctx context.Context, fun func(ctx context.Context) (result *unitofwork.Result, err error)) (result *unitofwork.Result, err error) {
	tx, err := u.db.Beginx()
	if err != nil {
		return result, err
	}

	ctx = context.WithValue(ctx, unitofwork.TransactionContextKey, tx)

	result, err = fun(ctx)
	if err != nil {
		tx.Rollback()
		return result, err
	}

	tx.Commit()
	return result, err
}

func GenerateStatement(ctx context.Context, db *sqlx.DB, sql string) (*sqlx.Stmt, error) {
	stmt, err := db.Preparex(sql)
	// In case opration is in transaction context
	tx, ok := ctx.Value(unitofwork.TransactionContextKey).(*sqlx.Tx)
	if ok {
		stmt, err = tx.Preparex(sql)
	}

	return stmt, err
}
