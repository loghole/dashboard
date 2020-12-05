package suggest

import (
	validation "github.com/gadavy/ozzo-validation/v4"
	"github.com/lissteron/simplerr"

	"github.com/loghole/dashboard/internal/app/codes"
	"github.com/loghole/dashboard/internal/app/domain"
)

const (
	maxSuggestValue = 50
)

func (m *ListReq) Validate() error {
	err := validation.ValidateStruct(m,
		validation.Field(&m.Type, m.typeRules()...),
		validation.Field(&m.Value, m.valueRules()...),
	)
	if err != nil {
		return simplerr.WithCode(err, codes.ValidationErr)
	}

	return nil
}

func (m *ListReq) typeRules() []validation.Rule {
	return []validation.Rule{
		validation.In(
			domain.SuggestLevel,
			domain.SuggestHost,
			domain.SuggestSource,
			domain.SuggestNamespace).ErrorCode(codes.ValidSuggestTypeIn.String()),
	}
}

func (m *ListReq) valueRules() []validation.Rule {
	return []validation.Rule{
		validation.Length(0, maxSuggestValue).ErrorCode(codes.ValidSuggestValueLength.String()),
	}
}
