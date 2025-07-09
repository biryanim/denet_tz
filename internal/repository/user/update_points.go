package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) UpdatePoints(ctx context.Context, id int64, points int64) (*model.User, error) {
	query, args, err := r.qb.
		Update("users").
		Set("points", squirrel.Expr("points + ?", points)).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING id, username, email, points, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	var user model.User
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Points,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound
		}

		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}
