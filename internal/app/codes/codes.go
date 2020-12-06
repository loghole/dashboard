package codes

import (
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"
)

const (
	system     = 1000
	logic      = 2000
	validation = 200_000
)

// Validation error codes
type System int

func (s System) Int() int {
	return int(s) + system
}

func (s System) GRPC() int {
	return int(codes.Internal)
}

func (s System) HTTP() int {
	return http.StatusInternalServerError
}

// System error codes
const (
	DatabaseError System = iota + 1
)

// Logic error codes
type Logic int

func (s Logic) Int() int {
	return int(s) + logic
}

func (s Logic) GRPC() int {
	return int(codes.InvalidArgument)
}

func (s Logic) HTTP() int {
	return http.StatusBadRequest
}

const (
	InvalidSuggestType Logic = iota + 1
)

type Validation int

func (s Validation) Int() int {
	return int(s) + validation
}

func (s Validation) GRPC() int {
	return int(codes.InvalidArgument)
}

func (s Validation) HTTP() int {
	return http.StatusBadRequest
}

func (s Validation) String() string {
	return strconv.Itoa(int(s))
}

// Validation error codes
const (
	ValidationErr Validation = iota + 1
	ValidMinLimit
	ValidMaxLimit
	ValidMinOffset
	ValidQueryParamsTypeRequired
	ValidQueryParamsTypeIn
	ValidQueryParamsKeyRequired
	ValidQueryParamsValueRequired
	ValidQueryParamsValueItemEmpty
	ValidQueryParamsValueListEmpty
	ValidQueryParamsOperatorRequired
	ValidQueryParamsOperatorIn
	ValidSuggestTypeIn
	ValidSuggestValueLength
)
