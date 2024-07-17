package http

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/protobuf/errcode"
	"io"
	"net/http"
	"time"
)

type Body map[string]interface{}

type ResData[T any] struct {
	Code    errcode.ErrCode `json:"code"`
	Message string          `json:"message,omitempty"`
	//验证码
	Details T `json:"details,omitempty"`
}

func (res *ResData[T]) Response(w http.ResponseWriter, httpcode int) {
	w.WriteHeader(httpcode)
	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

type ResAnyData = ResData[any]

func NewResData[T any](code errcode.ErrCode, msg string, data T) *ResData[T] {
	return &ResData[T]{
		Code:    code,
		Message: msg,
		Details: data,
	}
}

func RespErrcode(w http.ResponseWriter, code errcode.ErrCode) {
	NewResData[any](code, code.Error(), nil).Response(w, http.StatusOK)
}

func RespErr(w http.ResponseWriter, err error) {
	NewResData[any](errcode.Unknown, err.Error(), nil).Response(w, http.StatusOK)
}

func RespErrMsg(w http.ResponseWriter, msg string) {
	NewResData[any](errcode.Success, msg, nil).Response(w, http.StatusOK)
}

func RespErrRep(w http.ResponseWriter, rep *errcode.ErrRep) {
	NewResData[any](rep.Code, rep.Message, nil).Response(w, http.StatusOK)
}

func Response[T any](w http.ResponseWriter, code errcode.ErrCode, msg string, data T) {
	NewResData(code, msg, data).Response(w, http.StatusOK)
}

func StreamWriter(w http.ResponseWriter, writer func(w io.Writer) bool) {
	notifyClosed := w.(http.CloseNotifier).CloseNotify()
	for {
		select {
		// response writer forced to close, exit.
		case <-notifyClosed:
			return
		default:
			shouldContinue := writer(w)
			w.(http.Flusher).Flush()
			if !shouldContinue {
				return
			}
		}
	}
}

func Stream(w http.ResponseWriter) {
	w.Header().Set(HeaderXAccelBuffering, "no") //nginx的锅必须加
	w.Header().Set(HeaderTransferEncoding, "chunked")
	i := 0
	ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
	StreamWriter(w, func(w io.Writer) bool {
		fmt.Fprintf(w, "Message number %d<br>", ints[i])
		time.Sleep(500 * time.Millisecond) // simulate delay.
		if i == len(ints)-1 {
			return false //关闭并刷新
		}
		i++
		return true //继续写入数据
	})
}

var ResponseSysErr = []byte(`{"code":10000,"message":"系统错误"}`)
var ResponseOk = []byte(`{"code":0,"message":"OK"}`)

type ReceiveData struct {
	Code    errcode.ErrCode `json:"code"`
	Message string          `json:"message,omitempty"`
	//验证码
	Details json.RawMessage `json:"details,omitempty"`
}

func NewReceivesData(code errcode.ErrCode, msg string, data any) *ReceiveData {
	jsonBytes, _ := json.Marshal(data)
	return &ReceiveData{
		Code:    code,
		Message: msg,
		Details: jsonBytes,
	}
}

func (r *ReceiveData) Response(w http.ResponseWriter, httpcode int) {
	w.WriteHeader(httpcode)
	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
	jsonBytes, _ := json.Marshal(r)
	w.Write(jsonBytes)
}

func (r *ReceiveData) UnmarshalData(v any) error {
	return json.Unmarshal(r.Details, v)
}

type IHttpResponse interface {
	Header() http.Header
	Body() []byte
	StatusCode() int
}

type HttpResponse struct {
	Header     map[string]string `json:"header"`
	Body       []byte            `json:"body"`
	StatusCode int               `json:"status"`
}
