package repository

import (
	"database/sql"

	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	"github.com/nuttchai/go-rest/internal/utils/context"
)

type TUserRepository struct {
	DB *sql.DB
}

var (
	UserRepository irepository.IUserRepository
)

func (m *TUserRepository) RetrieveOne(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	// NOTE: cannot directly use 'user' as table name because it is a reserved keyword
	query := "select * from public.user where id = $1"
	row := m.DB.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.Username,
	)

	return &user, err
}