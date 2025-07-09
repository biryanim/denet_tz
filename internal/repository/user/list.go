package user

import (
	"context"
	"fmt"
	"github.com/biryanim/denet_tz/internal/model"
)

func (r *repo) List(ctx context.Context, limit int64) ([]*model.User, error) {
	query, args, err := r.qb.
		Select("id", "username", "email", "password", "points", "created_at", "updated_at").
		From("users").
		OrderBy("points DESC").
		Limit(uint64(limit)).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var users []*model.User
	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		u := new(model.User)
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Password,
			&u.Points,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate users: %w", err)
	}

	return users, nil
}
