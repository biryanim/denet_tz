package user

import (
	"context"
	"fmt"
	apperror "github.com/biryanim/denet_tz/internal/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func (r *repo) CompleteTask(ctx context.Context, userId int64, taskId int64) error {
	query, args, err := r.qb.
		Insert("user_tasks").
		Columns("user_id", "task_id").
		Values(userId, taskId).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return apperror.ErrUserAndTaskAlreadyExists
		}
		return fmt.Errorf("failed to create query: %w", err)
	}

	return nil
}
