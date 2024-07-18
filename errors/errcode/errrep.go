package errcode

import (
	"github.com/gin-gonic/gin/render"
	stringsi "github.com/hopeio/utils/strings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type ErrRep struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message,omitempty"`
}

func NewErrRep(code ErrCode, msg string) *ErrRep {
	return &ErrRep{
		Code:    code,
		Message: msg,
	}
}

func (x *ErrRep) Error() string {
	return x.Message
}

func (x *ErrRep) GrpcStatus() *status.Status {
	return status.New(codes.Code(x.Code), x.Message)
}

func (x *ErrRep) MarshalJSON() ([]byte, error) {
	return stringsi.ToBytes(`{"code":` + strconv.Itoa(int(x.Code)) + `,"message":"` + x.Message + `"}`), nil
}

func (x *ErrRep) Response(w http.ResponseWriter) {
	render.WriteJSON(w, x)
}

func (x *ErrRep) AppendErr(err error) *ErrRep {
	x.Message += " " + err.Error()
	return x
}

func (x *ErrRep) Warp(err error) *WarpErrRep {
	return &WarpErrRep{*x, err}
}
