package handlers

import (
	"github.com/jmoiron/sqlx"
)

type handler struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) handler {
	return handler{db}
}
