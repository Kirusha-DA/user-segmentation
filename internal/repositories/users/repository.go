package users

import (
	"github.com/Kirusha-DA/user-segmentation/internal/entities"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllClients() ([]entities.User, error) {
	users := make([]entities.User, 0)
	err := r.db.Select(&users, "SELECT id, name, last_login FROM users")
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return users, nil
}
