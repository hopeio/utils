/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

import (
	"encoding/json"
	"github.com/hopeio/utils/errors/errcode"
	"github.com/hopeio/utils/net/http/consts"
	"io"
	"iter"
	"net/http"
)

type Body map[string]any

// ResData 主要用来接收返回，发送请使用ResAnyData
type ResData[T any] struct {
	Code errcode.ErrCode `json:"code"`
	Msg  string          `json:"msg,omitempty"`
	//验证码
	Data T `json:"data,omitempty"`
}

func (res *ResData[T]) Response(w http.ResponseWriter, statusCode int) (int, error) {
	w.WriteHeader(statusCode)
	w.Header().Set(consts.HeaderContentType, "application/json; charset=utf-8")
	jsonBytes, _ := json.Marshal(res)
	return w.Write(jsonBytes)
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

func NewSuccessResData(data any) *ResAnyData {
	return &ResAnyData{
		Data: data,
	}
}

func NewErrorResData(code errcode.ErrCode, msg string) *ResAnyData {
	return &ResAnyData{
		Code: code,
		Msg:  msg,
	}
}

func RespErrCode(w http.ResponseWriter, code errcode.ErrCode) {
	NewResData[any](code, code.Error(), nil).Response(w, http.StatusOK)
}

func RespErrCodeMsg(w http.ResponseWriter, code errcode.ErrCode, msg string) {
	NewResData[any](code, msg, nil).Response(w, http.StatusOK)
}

func RespSuccess[T any](w http.ResponseWriter, msg string, data T) (int, error) {
	return NewResData(errcode.Success, msg, data).Response(w, http.StatusOK)
}

func RespSuccessMsg(w http.ResponseWriter, msg string) (int, error) {
	return NewResData[any](errcode.Success, msg, nil).Response(w, http.StatusOK)
}

func RespSuccessData(w http.ResponseWriter, data any) (int, error) {
	return NewResData[any](errcode.Success, errcode.Success.String(), data).Response(w, http.StatusOK)
}

func RespErrRep(w http.ResponseWriter, rep *errcode.ErrRep) (int, error) {
	return NewResData[any](rep.Code, rep.Msg, nil).Response(w, http.StatusOK)
}

func Response[T any](w http.ResponseWriter, code errcode.ErrCode, msg string, data T) (int, error) {
	return NewResData(code, msg, data).Response(w, http.StatusOK)
}

func ResponseStreamWrite(w http.ResponseWriter, dataSource iter.Seq[[]byte]) {
	w.Header().Set(consts.HeaderXAccelBuffering, "no") //nginx的锅必须加
	w.Header().Set(consts.HeaderTransferEncoding, "chunked")
	notifyClosed := w.(http.CloseNotifier).CloseNotify()
	for data := range dataSource {
		select {
		// response writer forced to close, exit.
		case <-notifyClosed:
			return
		default:
			w.Write(data)
			w.(http.Flusher).Flush()
		}
	}
}

var ResponseSysErr = json.RawMessage(`{"code":-1,"msg":"system error"}`)
var ResponseOk = json.RawMessage(`{"code":0}`)

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
	StatusCode() int
	Header() Header
	io.WriterTo
	io.Closer
}

func ResponseWrite(w http.ResponseWriter, httpres IHttpResponse) (int, error) {
	w.WriteHeader(httpres.StatusCode())
	header := w.Header()
	httpres.Header().Range(header.Set)
	i, err := httpres.WriteTo(w)
	if err != nil {
		return int(i), err
	}
	err = httpres.Close()
	if err != nil {
		return int(i), err
	}
	return int(i), err
}

type HttpResponseRawBody struct {
	Status  int       `json:"status,omitempty"`
	Headers MapHeader `json:"header,omitempty"`
	Body    []byte    `json:"body,omitempty"`
}

func (res *HttpResponseRawBody) Header() Header {
	return res.Headers
}

func (res *HttpResponseRawBody) WriteTo(writer io.Writer) (int64, error) {
	i, err := writer.Write(res.Body)
	return int64(i), err
}

func (res *HttpResponseRawBody) Close() error {
	return nil
}

func (res *HttpResponseRawBody) StatusCode() int {
	return res.Status
}

func (res *HttpResponseRawBody) Response(w http.ResponseWriter) (int, error) {
	w.WriteHeader(res.Status)
	header := w.Header()
	for k, v := range res.Headers {
		header.Set(k, v)
	}
	return w.Write(res.Body)
}

type HttpResponse struct {
	Status  int            `json:"status,omitempty"`
	Headers MapHeader      `json:"header,omitempty"`
	Body    WriterToCloser `json:"body,omitempty"`
}

type WriterToCloser interface {
	io.WriterTo
	io.Closer
}

func (res *HttpResponse) Header() Header {
	return res.Headers
}

func (res *HttpResponse) WriteTo(writer io.Writer) (int64, error) {
	return res.Body.WriteTo(writer)
}

func (res *HttpResponse) Close() error {
	return res.Body.Close()
}

func (res *HttpResponse) StatusCode() int {
	return res.Status
}

func (res *HttpResponse) Response(w http.ResponseWriter) (int, error) {
	w.WriteHeader(res.Status)
	for k, v := range res.Headers {
		w.Header().Set(k, v)
	}
	i, err := res.Body.WriteTo(w)
	if err != nil {
		return int(i), err
	}
	err = res.Body.Close()
	if err != nil {
		return int(i), err
	}
	return int(i), err
}

type ErrRep errcode.ErrRep

func ErrRepFrom(err error) *ErrRep {
	if errrep, ok := err.(*errcode.ErrRep); ok {
		return (*ErrRep)(errrep)
	}
	if errcode, ok := err.(errcode.ErrCode); ok {
		return &ErrRep{Code: errcode, Msg: errcode.Error()}
	}
	return &ErrRep{Code: errcode.Unknown, Msg: err.Error()}
}

func (res *ErrRep) Response(w http.ResponseWriter, statusCode int) (int, error) {
	w.WriteHeader(statusCode)
	w.Header().Set(consts.HeaderContentType, consts.ContentTypeJsonUtf8)
	jsonBytes, _ := json.Marshal(res)
	return w.Write(jsonBytes)
}

func ResponseError(w http.ResponseWriter, err error) {
	ErrRepFrom(err).Response(w, http.StatusOK)
}

type IHttpResponseTo interface {
	Response(w http.ResponseWriter) (int, error)
}

type HttpResponseStream struct {
	Status  int              `json:"status,omitempty"`
	Headers MapHeader        `json:"header,omitempty"`
	Body    iter.Seq[[]byte] `json:"body,omitempty"`
}

func (res *HttpResponseStream) Header() Header {
	res.Headers[consts.HeaderTransferEncoding] = "chunked"
	return res.Headers
}

func (res *HttpResponseStream) WriteTo(writer io.Writer) (int64, error) {
	notifyClosed := writer.(http.CloseNotifier).CloseNotify()
	var n int64
	for data := range res.Body {
		select {
		// response writer forced to close, exit.
		case <-notifyClosed:
			return n, nil
		default:
			write, err := writer.Write(data)
			if err != nil {
				return 0, err
			}
			n += int64(write)
			writer.(http.Flusher).Flush()
		}
	}
	return n, nil
}

func (res *HttpResponseStream) Close() error {
	return nil
}

func (res *HttpResponseStream) StatusCode() int {
	return res.Status
}

type RawBody []byte

func (res RawBody) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(res)
	return int64(n), err
}

func (res RawBody) Closer() error {
	return nil
}
