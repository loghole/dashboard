// Code generated by tron. You must modify it.

package suggest

import (
	context "context"

	"github.com/loghole/dashboard/internal/app/usecases"
	suggestV1 "github.com/loghole/dashboard/pkg/suggest/v1"
)

func (i *Implementation) List(
	ctx context.Context,
	req *suggestV1.ListReq,
) (resp *suggestV1.ListResp, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	result, err := i.listSuggest.Do(ctx, &usecases.ListSuggestIn{Type: req.Type, Value: req.Value})
	if err != nil {
		return nil, err
	}

	resp = &suggestV1.ListResp{
		Data: result,
	}

	return resp, nil
}