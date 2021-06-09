package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/repository"
)

type ExampleRepository struct {
	conn *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) repository.ExampleRepositoryI {
	return &ExampleRepository{conn: db}
}
