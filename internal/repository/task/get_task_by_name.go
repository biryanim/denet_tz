package task

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperror "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) GetTaskByName(ctx context.Context, name string) (*model.Task, error) {
	query, args, err := r.qb.
		Select("id", "name", "description", "points_reward", "created_at").
		From("tasks").
		Where(squirrel.Eq{
			"name": name,
		}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build insert query: %w", err)
	}

	var task model.Task
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.PointsReward,
		&task.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrTaskNotFound
		}
		return nil, fmt.Errorf("failed to fetch task by name: %w", err)
	}

	return &task, nil
}
