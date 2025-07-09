package user

import (
	"context"
	"fmt"

	apperror "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func (r *repo) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	query, args, err := r.qb.
		Insert("users").
		Columns("username", "email", "password").
		Values(user.Info.Username, user.Info.Email, user.Password).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build insert query: %w", err)
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, apperror.ErrUserAlreadyExists
		}
		return 0, fmt.Errorf("failed to create query: %w", err)
	}

	return id, nil
}
