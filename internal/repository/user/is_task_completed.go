package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperror "github.com/biryanim/denet_tz/internal/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func (r *repo) IsTaskCompleted(ctx context.Context, userId int64, taskId int64) (bool, error) {
	query, args, err := r.qb.
		Select("1").
		From("user_tasks").
		Where(squirrel.Eq{
			"user_id": userId,
			"task_id": taskId,
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
			return false, apperror.ErrTaskAlreadyCompleted
		}
		return false, fmt.Errorf("failed to create query: %w", err)
	}
	return exists, nil
}
