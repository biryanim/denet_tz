package task

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) GetByID(ctx context.Context, id int64) (*model.Task, error) {
	query, args, err := r.qb.
		Select("id", "name", "description", "points_reward", "created_at").
		From("tasks").
		Where(squirrel.Eq{
			"id": id,
		}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var task model.Task
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&task.ID, &task.Name, &task.Description, &task.PointsReward, &task.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrTaskNotFound
		}

		return nil, fmt.Errorf("failed to query task: %w", err)
	}

	return &task, nil
}
