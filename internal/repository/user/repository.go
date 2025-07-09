package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/biryanim/denet_tz/internal/client/db"
	"github.com/biryanim/denet_tz/internal/repository"
)

type repo struct {
	db db.Client
	qb squirrel.StatementBuilderType
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
