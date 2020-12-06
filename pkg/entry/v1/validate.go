package entry

import (
	validation "github.com/gadavy/ozzo-validation/v4"
	"github.com/lissteron/simplerr"

	"github.com/loghole/dashboard/internal/app/codes"
	"github.com/loghole/dashboard/internal/app/domain"
)

const (
	maxLimit = 1000
)

func (m *ListReq) Validate() error {
	err := validation.ValidateStruct(m,
		validation.Field(&m.Params),
		validation.Field(&m.Limit, m.limitRules()...),
		validation.Field(&m.Offset, m.offsetRules()...),
	)
	if err != nil {
		return simplerr.WithCode(err, codes.ValidationErr)
	}

	return nil
}

func (m *ListReq) limitRules() []validation.Rule {
	return []validation.Rule{
		validation.Min(int64(0)).ErrorCode(codes.ValidMinLimit.String()),
		validation.Max(int64(maxLimit)).ErrorCode(codes.ValidMaxLimit.String()),
	}
}

func (m *ListReq) offsetRules() []validation.Rule {
	return []validation.Rule{
		validation.Min(int64(0)).ErrorCode(codes.ValidMinOffset.String()),
	}
}

func (m *QueryParam) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Type, m.typeRules()...),
		validation.Field(&m.Key, m.keyRules()...),
		validation.Field(&m.Value, m.valueRules()...),
		validation.Field(&m.Operator, m.operatorRules()...),
	)
}

func (m *QueryParam) typeRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsTypeRequired.String()),
		validation.In(domain.TypeKey, domain.TypeColumn).ErrorCode(codes.ValidQueryParamsTypeIn.String()),
	}
}

func (m *QueryParam) keyRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsKeyRequired.String()),
	}
}

func (m *QueryParam) valueRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsValueRequired.String()),
	}
}

func (m *QueryParam) operatorRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsOperatorRequired.String()),
		validation.In(
			domain.OperatorLt,
			domain.OperatorLte,
			domain.OperatorGt,
			domain.OperatorGte,
			domain.OperatorEq,
			domain.OperatorNotEq,
			domain.OperatorLike,
			domain.OperatorNotLike).ErrorCode(codes.ValidQueryParamsOperatorIn.String()),
	}
}

func (m *ParamValue) IsEmpty() bool {
	return len(m.List) == 0 && m.Item == ""
}

func (m *ParamValue) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Item, m.itemRules()...),
		validation.Field(&m.List, m.listRules()...),
	)
}

func (m *ParamValue) itemRules() []validation.Rule {
	return []validation.Rule{
		validation.When(m.List != nil, validation.NotIn("").ErrorCode(codes.ValidQueryParamsValueItemEmpty.String())),
	}
}

func (m *ParamValue) listRules() []validation.Rule {
	return []validation.Rule{
		validation.When(m.Item != "", validation.Length(0, 0).ErrorCode(codes.ValidQueryParamsValueListEmpty.String())),
	}
}
