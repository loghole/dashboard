package clickhouse

import (
	"context"
	"fmt"

	"github.com/loghole/tracing"

	"github.com/loghole/dashboard/internal/app/domain"
	"github.com/loghole/dashboard/internal/app/repositories/clickhouse/models"
)

func (r *Repository) ListEntry(ctx context.Context, input *domain.Query) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query, args, err := models.NewBuilder().Build(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("can't build query: %w", err)
	}

	var dest []*models.Entry

	if err := r.db.SelectContext(ctx, &dest, query, args...); err != nil {
		return nil, fmt.Errorf("can't select data: %w", err)
	}

	result := make([]*domain.Entry, 0, len(dest))

	for _, val := range dest {
		result = append(result, val.ToDomain())
	}

	return result, nil
}
