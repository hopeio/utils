package errcode

import (
	"github.com/gin-gonic/gin/render"
	"github.com/hopeio/utils/log"
	stringsi "github.com/hopeio/utils/strings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type IErrRep interface {
	ErrRep() *ErrRep
}

type GrpcStatus interface {
	GrpcStatus() *status.Status
}

type ErrCode uint32

func (x ErrCode) String() string {
	if x == Success {
		return "OK"
	}
	value, ok := codeMap[x]
	if ok {
		return value
	}
	return strconv.Itoa(int(x))
}

func (x ErrCode) ErrRep() *ErrRep {
	return &ErrRep{Code: x, Message: x.String()}
}

// example 实现
func (x ErrCode) GrpcStatus() *status.Status {
	return status.New(codes.Code(x), x.String())
}

func (x ErrCode) Message(msg string) *ErrRep {
	return &ErrRep{Code: x, Message: msg}
}

func (x ErrCode) MarshalJSON() ([]byte, error) {
	return stringsi.ToBytes(`{"code":` + strconv.Itoa(int(x)) + `,"message":"` + x.String() + `"}`), nil
}

func (x ErrCode) Warp(err error) *ErrRep {
	return &ErrRep{Code: x, Message: err.Error()}
}

func (x ErrCode) Log(err error) *ErrRep {
	log.Error(err)
	return &ErrRep{Code: x, Message: x.String()}
}

func (x ErrCode) Error() string {
	return x.String()
}

func (x ErrCode) Response(w http.ResponseWriter) {
	render.WriteJSON(w, x)
}
