package gf

import 	(
		"io"
		"net/http"
		)

type RequestInterface interface {
	Path() string
	Method() string
	Body(string) interface{}
	Param(string) interface{}
	StrParam(string) string
	SetParam(string, interface{})
	SetHeader(string, string)
	ReadBody() *ResponseStatus
	ReadBodyArray() *ResponseStatus
	BodyArray() []interface{}
	BodyObject() map[string]interface{}
	Redirect(string, int) *ResponseStatus
	HttpError(string, int)
	Writer() io.Writer
	Write([]byte)
	Fail() *ResponseStatus
	Respond(args ...interface{}) *ResponseStatus
	// logging
	Debug(string)
	NewError(string) error
	Error(error)
	DebugJSON(interface{})
	Reflect(interface{})
	//
	Res() http.ResponseWriter
	R() *http.Request
}
