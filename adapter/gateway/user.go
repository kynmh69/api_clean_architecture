package gateway

import (
	"ca/v2/entity"
	"ca/v2/usecase/port"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type UserRepository struct {
	conn *sql.DB
}

// GetUserByID implements port.UserRepository
func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "select * from user where id = ?", userID)
	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found. user id = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("internal server error. adapter/gateway/getuserbyiD")
	}
	return &user, nil
}

func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
