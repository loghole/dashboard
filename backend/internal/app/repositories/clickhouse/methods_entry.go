package clickhouse

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/gadavy/tracing"

	"github.com/loghole/dashboard/internal/app/domain"
	"github.com/loghole/dashboard/internal/app/repositories/clickhouse/models"
)

func (r *Repository) ListEntry(ctx context.Context, input *domain.Query) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query, args, err := buildListEntryQuery(ctx, input)
	if err != nil {
		return nil, err
	}

	`
		SELECT time, nsec, namespace, source, host, level, trace_id, message, remote_ip, params, build_commit, config_hash
		FROM internal_logs_buffer WHERE time>='2020-09-01 00:00:00'
			AND (
				row_id IN(
					SELECT row_id FROM internal_logs_buffer ARRAY JOIN params_float.keys as keys, params_float.values as values 
					WHERE time>='2020-09-01 00:00:00' AND (keys='json_key' AND values=3) AND (keys='json_key' AND values=4) group by row_id
				) OR row_id IN(
					SELECT row_id FROM internal_logs_buffer ARRAY JOIN params_string.keys as keys, params_string.values as values 
					WHERE time>='2020-09-01 00:00:00' AND keys='json_key' AND values='3' group by row_id
				)
			) LIMIT 10
`

	log.Println(query)

	var dest []*models.Entry

	if err := r.db.SelectContext(ctx, &dest, query, args...); err != nil {
		return nil, err
	}

	result := make([]*domain.Entry, 0, len(dest))

	for _, val := range dest {
		result = append(result, val.ToDomain())
	}

	return result, nil
}

type Builder struct {
	mainBuilder   squirrel.SelectBuilder
	floatBuilder  squirrel.SelectBuilder
	stringBuilder squirrel.SelectBuilder
}

func buildListEntryQuery(
	ctx context.Context,
	input *domain.Query,
) (query string, args []interface{}, err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	builder := squirrel.Select("time", "nsec", "namespace", "source", "host", "level",
		"trace_id", "message", "remote_ip", "params", "build_commit", "config_hash").From("internal_logs_buffer")

	for _, param := range input.Params {
		if param.IsTypeJSON() {
			//	builder = builder.Where(models.JSONParamFromDomain(param))
			continue
		}

		builder = builder.Where(models.ColumnParamFromDomain(param))
	}

	for _, param := range input.Params {

	}

	return builder.OrderBy("nsec DESC").
		Suffix(fmt.Sprintf("LIMIT %d, %d", input.Offset, input.Limit)).
		PlaceholderFormat(squirrel.Question).
		ToSql()
}
