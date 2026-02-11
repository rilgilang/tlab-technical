package repositories

import (
	"context"
	"database/sql"
	"tlab/src/domain/sharedkernel/logger"
	"tlab/src/domain/user"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB  *sqlx.DB
	log logger.Logger
}

func NewUserRepository(db *sqlx.DB, log logger.Logger) *UserRepository {
	return &UserRepository{
		DB:  db,
		log: log,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, payload user.User) error {

	query := `
		INSERT INTO "user" (
		                          id, 
		                          name, 
		                          email,
		                          password, 
		                          created_at, 
		                          updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	vals := []interface{}{
		payload.ID,
		payload.Name,
		payload.Email,
		payload.Password,
		payload.CreatedAt,
		payload.UpdatedAt,
	}

	stmt, err := GenerateStatement(ctx, r.DB, query)
	if err != nil {
		r.log.Error("error_create_user", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, vals...)
	if err != nil {
		r.log.Error("error_create_user", err)
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {

	query := `
		SELECT
		                          id, 
		                          name, 
		                          email,
		                          password, 
		                          created_at, 
		                          updated_at
		FROM "user"
		WHERE email = $1
		`

	user := &user.User{}

	err := r.DB.QueryRowxContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.log.Error("error_get_user_by_email", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*user.User, error) {

	query := `
		SELECT
		                          id, 
		                          name, 
		                          email,
		                          password, 
		                          created_at, 
		                          updated_at
		FROM "user"
		WHERE id = $1
		`

	user := &user.User{}

	err := r.DB.QueryRowxContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.log.Error("error_get_user_by_email", err)
		return nil, err
	}
	return user, nil
}
