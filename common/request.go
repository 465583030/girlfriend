package gf

import 	(
		"io"
		)

type RequestInterface interface {
	Path() string
	Method() string
	BodyParam(string) interface{}
	GetParam(string) interface{}
	SetParam(string, interface{})
	SetHeader(string, string)
	ReadBody() *ResponseStatus
	ReadBodyArray() *ResponseStatus
	Redirect(string, int) *ResponseStatus
	HandleStatus(*ResponseStatus)
	MSI(string) (bool, Object)
	IA(string) (bool, Array)
	String(string) (bool, string)
	Float64(string) (bool, float64)
	Bool(string) (bool, bool)
	Int(string) (bool, int)
	HttpError(string, int)
	Writer() io.Writer
	Write([]byte)
	Fail() *ResponseStatus
	Respond(args ...interface{}) *ResponseStatus
}