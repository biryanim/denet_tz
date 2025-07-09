package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/biryanim/denet_tz/internal/model"
)

func (r *repo) GetReferrers(ctx context.Context, userId int64) (*model.Referrers, error) {
	query, args, err := r.qb.
		Select("r.referred_id").
		From("users u").
		Join("referrals r ON u.id = r.referrer_id").
		Where(squirrel.Eq{
			"r.referrer_id": userId,
		}).ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows: %w", err)
	}

	return &model.Referrers{
		UserIds: ids,
	}, nil
}
