package usecases

import (
	"context"

	"github.com/lissteron/simplerr"
	"github.com/loghole/tracing"

	"github.com/loghole/dashboard/internal/app/codes"
	"github.com/loghole/dashboard/internal/app/domain"
)

type ListEntryIn struct {
	*domain.Query
}

type ListEntry struct {
	storage Storage
	logger  Logger
}

func NewListEntry(
	storage Storage,
	logger Logger,
) *ListEntry {
	return &ListEntry{
		storage: storage,
		logger:  logger,
	}
}

func (l *ListEntry) Do(ctx context.Context, input *ListEntryIn) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	result, err := l.storage.ListEntry(ctx, input.Query)
	if err != nil {
		l.logger.Errorf(ctx, "get entry list failed: %v", err)

		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "get entry list failed")
	}

	return result, nil
}
