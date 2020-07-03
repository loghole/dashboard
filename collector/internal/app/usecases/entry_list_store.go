package usecases

import (
	"context"

	"github.com/gadavy/tracing"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/loghole/collector/internal/app/codes"
	"github.com/lissteron/loghole/collector/internal/app/domain"
	"github.com/lissteron/loghole/collector/internal/app/usecases/interfaces"
)

type StoreEntryList struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewStoreEntryList(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *StoreEntryList {
	return &StoreEntryList{
		storage: storage,
		logger:  logger,
	}
}

func (s *StoreEntryList) Do(ctx context.Context, data []byte) (err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	list, err := s.parseEntryList(ctx, data[:])
	if err != nil {
		s.logger.Errorf(ctx, "parse entry list failed: %v", err)

		return simplerr.WrapWithCode(err, codes.UnmarshalError, "parse json failed")
	}

	err = s.storage.StoreEntryList(ctx, list)
	if err != nil {
		s.logger.Errorf(ctx, "store entry list failed: %v", err)

		return simplerr.WrapWithCode(err, codes.DatabaseError, "store failed")
	}

	return nil
}

func (s *StoreEntryList) parseEntryList(ctx context.Context, data []byte) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	list := domain.EntryList{}

	if err := list.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	return list, nil
}
