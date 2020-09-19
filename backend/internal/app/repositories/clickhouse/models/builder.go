package models

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/loghole/tracing"

	"github.com/loghole/dashboard/internal/app/domain"
)

type fieldType string

const (
	typeString fieldType = "string"
	typeFloat  fieldType = "float"
)

type Builder struct {
	mainBuilder   squirrel.SelectBuilder
	floatBuilder  squirrel.SelectBuilder
	stringBuilder squirrel.SelectBuilder
}

func NewBuilder() *Builder {
	return &Builder{
		mainBuilder: squirrel.Select("time", "nsec", "namespace", "source", "host", "level", "trace_id",
			"message", "remote_ip", "params", "build_commit", "config_hash").From("internal_logs_buffer"),
		floatBuilder: squirrel.Select("row_id").
			From("internal_logs_buffer ARRAY JOIN params_float.keys as keys, params_float.values as values"),
		stringBuilder: squirrel.Select("row_id").
			From("internal_logs_buffer ARRAY JOIN params_string.keys as keys, params_string.values as values"),
	}
}

// nolint:funlen // builder
func (b *Builder) Build(ctx context.Context, input *domain.Query) (query string, args []interface{}, err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	for _, param := range input.Params {
		if param.IsTypeJSON() {
			continue
		}

		b.mainBuilder = b.mainBuilder.Where(ColumnParamFromDomain(param))
		b.floatBuilder = b.floatBuilder.Where(ColumnParamFromDomain(param))
		b.stringBuilder = b.stringBuilder.Where(ColumnParamFromDomain(param))
	}

	for _, param := range input.Params {
		if !param.IsTypeJSON() {
			continue
		}

		var (
			flt = b.floatBuilder.Where(JSONParamFromDomain(param, typeFloat))
			str = b.stringBuilder.Where(JSONParamFromDomain(param, typeString))
		)

		switch param.Operator {
		case domain.OperatorEq:
			flt = flt.Prefix("row_id IN(").Suffix("GROUP BY row_id)")
			str = str.Prefix("row_id IN(").Suffix("GROUP BY row_id)")

			if hasFloat(param) {
				query, args, err = squirrel.Or{flt, str}.ToSql()
			} else {
				query, args, err = str.ToSql()
			}
		case domain.OperatorNotEq:
			flt = flt.Prefix("row_id NOT IN(").Suffix("GROUP BY row_id)")
			str = str.Prefix("row_id NOT IN(").Suffix("GROUP BY row_id)")

			if hasFloat(param) {
				query, args, err = squirrel.Or{flt, str}.ToSql()
			} else {
				query, args, err = str.ToSql()
			}
		case domain.OperatorLike:
			str = str.Prefix("row_id IN(").Suffix("GROUP BY row_id)")

			query, args, err = str.ToSql()
		case domain.OperatorNotLike:
			str = str.Prefix("row_id NOT IN(").Suffix("GROUP BY row_id)")

			query, args, err = str.ToSql()
		case domain.OperatorLt, domain.OperatorLte, domain.OperatorGt, domain.OperatorGte:
			flt = flt.Prefix("row_id IN(").Suffix("GROUP BY row_id)")
			str = str.Prefix("row_id IN(").Suffix("GROUP BY row_id)")

			if hasFloat(param) {
				query, args, err = flt.ToSql()
			} else {
				query, args, err = str.ToSql()
			}
		}

		if err != nil {
			return "", nil, err
		}

		b.mainBuilder = b.mainBuilder.Where(query, args...)
	}

	return b.mainBuilder.OrderBy("nsec DESC").
		Suffix(fmt.Sprintf("LIMIT %d, %d", input.Offset, input.Limit)).
		PlaceholderFormat(squirrel.Question).
		ToSql()
}

func hasFloat(param *domain.QueryParam) bool {
	if _, err := strconv.ParseFloat(param.Value.Item, 64); err == nil {
		return true
	}

	for _, value := range param.Value.List {
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			return true
		}
	}

	return false
}
