/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/errors/errcode"
	"io"
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
	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
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

func RespErrcode(w http.ResponseWriter, code errcode.ErrCode) {
	NewResData[any](code, code.Error(), nil).Response(w, http.StatusOK)
}

func RespError(w http.ResponseWriter, code errcode.ErrCode, msg string) {
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

func ResponseStreamWrite(w http.ResponseWriter, writer func(w io.Writer) bool) {
	w.Header().Set(HeaderXAccelBuffering, "no") //nginx的锅必须加
	w.Header().Set(HeaderTransferEncoding, "chunked")
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
	RespHeader() map[string]string
	RespBody() io.ReadCloser
}

func ResponseWrite(w http.ResponseWriter, httpres IHttpResponse) (int, error) {
	w.WriteHeader(httpres.StatusCode())
	for k, v := range httpres.RespHeader() {
		w.Header().Set(k, v)
	}
	body := httpres.RespBody()
	i, err := io.Copy(w, body)
	if err != nil {
		return int(i), err
	}
	err = body.Close()
	if err != nil {
		return int(i), err
	}
	return int(i), err
}

type HttpResponseRawBody struct {
	Status int               `json:"status,omitempty"`
	Header map[string]string `json:"header,omitempty"`
	Body   []byte            `json:"body,omitempty"`
}

func (res *HttpResponseRawBody) RespHeader() map[string]string {
	return res.Header
}

func (res *HttpResponseRawBody) RespBody() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(res.Body))
}

func (res *HttpResponseRawBody) StatusCode() int {
	return res.Status
}

func (res *HttpResponseRawBody) Response(w http.ResponseWriter) (int, error) {
	w.WriteHeader(res.Status)
	for k, v := range res.Header {
		w.Header().Set(k, v)
	}
	return w.Write(res.Body)
}

type HttpResponse struct {
	Status int               `json:"status,omitempty"`
	Header map[string]string `json:"header,omitempty"`
	Body   io.ReadCloser     `json:"body,omitempty"`
}

func (res *HttpResponse) RespHeader() map[string]string {
	return res.Header
}

func (res *HttpResponse) RespBody() io.ReadCloser {
	return res.Body
}

func (res *HttpResponse) StatusCode() int {
	return res.Status
}

func (res *HttpResponse) Flush() error {
	return nil
}

func (res *HttpResponse) Response(w http.ResponseWriter) (int, error) {
	w.WriteHeader(res.Status)
	for k, v := range res.Header {
		w.Header().Set(k, v)
	}
	i, err := io.Copy(w, res.Body)
	if err != nil {
		return int(i), err
	}
	err = res.Body.Close()
	if err != nil {
		return int(i), err
	}
	return int(i), err
}

type ResError struct {
	Code errcode.ErrCode `json:"code"`
	Msg  string          `json:"msg,omitempty"`
}

func (res *ResError) Response(w http.ResponseWriter, statusCode int) (int, error) {
	w.WriteHeader(statusCode)
	w.Header().Set(HeaderContentType, ContentTypeJsonUtf8)
	jsonBytes, _ := json.Marshal(res)
	return w.Write(jsonBytes)
}

type ResponseFile struct {
	Name string        `json:"name"`
	Body io.ReadCloser `json:"body,omitempty"`
}

func (res *ResponseFile) RespHeader() map[string]string {
	return map[string]string{HeaderContentType: ContentTypeOctetStream, HeaderContentDisposition: fmt.Sprintf(AttachmentTmpl, res.Name)}
}

func (res *ResponseFile) RespBody() io.ReadCloser {
	return res.Body
}

func (res *ResponseFile) StatusCode() int {
	return http.StatusOK
}

type StreamWriter interface {
	Write(io.Writer) (n int, err error)
	Flush() error
}

type HttpResponseWriteTo struct {
	Status int               `json:"status,omitempty"`
	Header map[string]string `json:"header,omitempty"`
	Body   io.WriterTo       `json:"body,omitempty"`
}

func (res *HttpResponseWriteTo) RespHeader() map[string]string {
	return res.Header
}

func (res *HttpResponseWriteTo) StatusCode() int {
	return res.Status
}

func (res *HttpResponseWriteTo) ResBody() io.WriterTo {
	return res.Body
}

func (res *HttpResponseWriteTo) Response(w http.ResponseWriter) (int, error) {
	w.WriteHeader(res.Status)
	body := res.ResBody()
	i, err := body.WriteTo(w)
	return int(i), err
}
