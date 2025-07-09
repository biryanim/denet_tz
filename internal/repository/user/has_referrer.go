package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperror "github.com/biryanim/denet_tz/internal/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func (r *repo) HasReferrer(ctx context.Context, userId int64) (bool, error) {
	query, args, err := r.qb.
		Select("1").
		From("referrals").
		Where(squirrel.Eq{
			"referred_id": userId,
		}).ToSql()
	query = "SELECT EXISTS(" + query + ")"
	if err != nil {
		return false, fmt.Errorf("failed to build select query: %w", err)
	}

	var exists bool
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return false, apperror.ErrUserAlreadyHasReferrer
		}
		return false, fmt.Errorf("failed to create query: %w", err)
	}
	return exists, nil
}
