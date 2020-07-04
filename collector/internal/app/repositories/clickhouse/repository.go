package clickhouse

import (
	"context"

	"github.com/gadavy/tracing"
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, template string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, template string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, template string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, template string, args ...interface{})
}

type Repository struct {
	db     *sqlx.DB
	logger Logger
}

func NewRepository(db *sqlx.DB, logger Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) Ping(ctx context.Context) error {
	defer tracing.ChildSpan(&ctx).Finish()

	return r.db.PingContext(ctx)
}
