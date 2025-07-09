package user

import (
	"context"
	"fmt"
	"github.com/biryanim/denet_tz/internal/model"
)

func (r *repo) AddReferrer(ctx context.Context, referrer *model.Referrals) error {
	query, args, err := r.qb.
		Insert("referrals").
		Columns("referrer_id", "referred_id").
		Values(referrer.ReferrerUserId, referrer.ReferredUserId).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to add referrer: %w", err)
	}

	return nil
}
