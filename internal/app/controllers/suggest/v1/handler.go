// Code generated by tron. You must modify it.

package suggest

import (
	"context"

	transport "github.com/loghole/tron/transport"

	"github.com/loghole/dashboard/internal/app/usecases"
	suggestV1 "github.com/loghole/dashboard/pkg/suggest/v1"
)

type ListSuggest interface {
	Do(ctx context.Context, input *usecases.ListSuggestIn) (result []string, err error)
}

type Implementation struct {
	suggestV1.UnimplementedSuggestServer
	listSuggest ListSuggest
}

func NewImplementation(listSuggest ListSuggest) *Implementation {
	return &Implementation{listSuggest: listSuggest}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Implementation) GetDescription() transport.ServiceDesc {
	return suggestV1.NewSuggestServiceDesc(i)
}
