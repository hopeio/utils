package http

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/errors/errcode"
	"io"
	"net/http"
	"time"
)

type Body map[string]any

// ResData 主要用来接收返回，发送请使用ResAnyData
type ResData[T any] struct {
	Code errcode.ErrCode `json:"code"`
	Msg  string          `json:"msg,omitempty"`
	//验证码
	Data T `json:"data,omitempty"`
}

func (res *ResData[T]) Response(w http.ResponseWriter, httpcode int) {
	w.WriteHeader(httpcode)
	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
	jsonBytes, _ := json.Marshal(res)
	w.Write(jsonBytes)
}

func NewResData[T any](code errcode.ErrCode, msg string, data T) *ResData[T] {
	return &ResData[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

type ResAnyData = ResData[any]

func NewResAnyData(code errcode.ErrCode, msg string, data any) *ResAnyData {
	return &ResAnyData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func RespErrcode(w http.ResponseWriter, code errcode.ErrCode) {
	NewResData[any](code, code.Error(), nil).Response(w, http.StatusOK)
}

func RespError(w http.ResponseWriter, code errcode.ErrCode, msg string) {
	NewResData[any](code, msg, nil).Response(w, http.StatusOK)
}

func RespSuccess[T any](w http.ResponseWriter, msg string, data T) {
	NewResData(errcode.Success, msg, data).Response(w, http.StatusOK)
}

func RespSuccessMsg(w http.ResponseWriter, msg string) {
	NewResData[any](errcode.Success, msg, nil).Response(w, http.StatusOK)
}

func RespSuccessData(w http.ResponseWriter, data any) {
	NewResData[any](errcode.Success, errcode.Success.String(), data).Response(w, http.StatusOK)
}

func RespErrRep(w http.ResponseWriter, rep *errcode.ErrRep) {
	NewResData[any](rep.Code, rep.Msg, nil).Response(w, http.StatusOK)
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
		fmt.Fprintf(w, "Msg number %d<br>", ints[i])
		time.Sleep(500 * time.Millisecond) // simulate delay.
		if i == len(ints)-1 {
			return false //关闭并刷新
		}
		i++
		return true //继续写入数据
	})
}

var ResponseSysErr = []byte(`{"code":-1,"msg":"system error"}`)
var ResponseOk = []byte(`{"code":0}`)

type ReceiveData = ResData[json.RawMessage]

func NewReceiveData(code errcode.ErrCode, msg string, data any) *ReceiveData {
	jsonBytes, _ := json.Marshal(data)
	return &ReceiveData{
		Code: code,
		Msg:  msg,
		Data: jsonBytes,
	}
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
