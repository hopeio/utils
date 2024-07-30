package errcode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type IErrRep interface {
	ErrRep() *ErrRep
}

type GRPCStatus interface {
	GRPCStatus() *status.Status
}

type ErrCode uint32

func (x ErrCode) String() string {
	value, ok := codeMap[x]
	if ok {
		return value
	}
	return strconv.Itoa(int(x))
}

func (x ErrCode) ErrRep() *ErrRep {
	return &ErrRep{Code: x, Msg: x.String()}
}

// example 实现
func (x ErrCode) GRPCStatus() *status.Status {
	return status.New(codes.Code(x), x.String())
}

func (x ErrCode) Msg(msg string) *ErrRep {
	return &ErrRep{Code: x, Msg: msg}
}

func (x ErrCode) Wrap(err error) *ErrRep {
	return &ErrRep{Code: x, Msg: err.Error()}
}

func (x ErrCode) Error() string {
	return x.String()
}
