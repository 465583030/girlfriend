package gf

import 	(
		"io"
		)

type RequestInterface interface {
	Path() string
	Method() string
	Body(string) interface{}
	Param(string) interface{}
	SetParam(string, interface{})
	SetHeader(string, string)
	ReadBody() *ResponseStatus
	ReadBodyArray() *ResponseStatus
	Redirect(string, int) *ResponseStatus
	HandleStatus(*ResponseStatus)
	HttpError(string, int)
	Writer() io.Writer
	Write([]byte)
	Fail() *ResponseStatus
	Respond(args ...interface{}) *ResponseStatus
}