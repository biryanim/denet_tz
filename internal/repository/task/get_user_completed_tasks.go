package task

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/biryanim/denet_tz/internal/model"
)

func (r *repo) GetUserCompletedTasks(ctx context.Context, userId int64) ([]*model.Task, error) {
	query, args, err := r.qb.
		Select("t.id", "t.name", "t.description", "t.points_reward", "t.created_at").
		From("tasks t").
		Join("user_tasks ut on t.id = ut.task_id").
		Where(
			squirrel.Eq{"ut.user_id": userId},
		).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	tasks := make([]*model.Task, 0)
	for rows.Next() {
		t := &model.Task{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.PointsReward,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows: %w", err)
	}

	return tasks, nil
}
