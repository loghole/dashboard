package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/loghole/dashboard/internal/app/domain"
	"github.com/loghole/dashboard/internal/app/repositories/clickhouse/tools"
)

const (
	anyParams  = "(keys=? AND values=?)"
	likeParams = "(keys=? AND values LIKE ?)"
	ltgtParams = "(keys=? AND values%s?)"
)

type JSONParam struct {
	domain.QueryParam
	Type fieldType
}

func JSONParamFromDomain(param *domain.QueryParam, t fieldType) *JSONParam {
	return &JSONParam{QueryParam: *param, Type: t}
}

// nolint:golint,stylecheck,gocritic
func (p *JSONParam) ToSql() (query string, args []interface{}, err error) {
	switch {
	case p.IsIn(), p.IsNotIn():
		return p.getIn()
	case p.IsLike():
		return p.prepareParamLike()
	case p.IsLtGt():
		return p.prepareParamLtGt()
	default:
		return p.getDefault()
	}
}

func (p *JSONParam) getIn() (query string, args []interface{}, err error) {
	builder := make([]string, 0, len(p.Value.List))

	if p.Type == typeString {
		for _, value := range p.GetValueList() {
			builder = append(builder, anyParams)
			args = append(args, p.Key, value)
		}
	} else {
		for _, value := range p.GetValueList() {
			if val, err := strconv.ParseFloat(value, 64); err == nil {
				builder = append(builder, anyParams)
				args = append(args, p.Key, val)
			}
		}
	}

	return strings.Join([]string{"(", strings.Join(builder, " OR "), ")"}, ""), args, nil
}

func (p *JSONParam) prepareParamLike() (query string, args []interface{}, err error) {
	if p.IsList() {
		return p.prepareParamListLike()
	}

	switch {
	case p.IsOperator(domain.OperatorLike), p.IsOperator(domain.OperatorNotLike):
		return p.getLike()
	default:
		panic(ErrNotImplemented)
	}
}

func (p *JSONParam) prepareParamListLike() (query string, args []interface{}, err error) {
	var (
		queries = make([]string, 0, len(p.GetValueList()))
		a       = make([]interface{}, 0)
		q       string
	)

	for _, value := range p.GetValueList() {
		switch {
		case p.IsOperator(domain.OperatorLike), p.IsOperator(domain.OperatorNotLike):
			q, a, err = p.getLikeWithValue(value)
		default:
			panic(ErrNotImplemented)
		}

		if err != nil {
			return "", nil, err
		}

		queries = append(queries, q)
		args = append(args, a...)
	}

	return strings.Join(queries, " OR "), args, nil
}

func (p *JSONParam) getLike() (query string, args []interface{}, err error) {
	args = []interface{}{p.Key, tools.CreateLike(p.GetValueItem())}

	return likeParams, args, nil
}

// nolint:unparam // need
func (p *JSONParam) getLikeWithValue(val string) (query string, args []interface{}, err error) {
	args = []interface{}{p.Key, tools.CreateLike(val)}

	return likeParams, args, nil
}

func (p *JSONParam) prepareParamLtGt() (query string, args []interface{}, err error) {
	if p.IsList() {
		return "", nil, ErrArrayNotAccepted
	}

	if value, ok := valueToFloat(p.GetValueItem()); ok {
		return p.getLtGtFloat(value)
	}

	return p.getLtGtString()
}

func (p *JSONParam) getLtGtString() (query string, args []interface{}, err error) {
	return fmt.Sprintf(ltgtParams, getOperator(p.Operator)), []interface{}{p.Key, p.Value.Item}, nil
}

func (p *JSONParam) getLtGtFloat(val float64) (query string, args []interface{}, err error) {
	return fmt.Sprintf(ltgtParams, getOperator(p.Operator)), []interface{}{p.Key, val}, nil
}

func (p *JSONParam) getDefault() (query string, args []interface{}, err error) {
	return fmt.Sprintf(ltgtParams, getOperator(p.Operator)), []interface{}{p.Key, p.Value.Item}, nil
}
