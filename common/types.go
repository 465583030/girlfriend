package gf

import 	(
		"io"
		)

type Array []interface{}
type Object map[string]interface{}

type HandlerFunction func (RequestInterface) *ResponseStatus

type Registry map[string]func (RequestInterface) *ResponseStatus

type RequestInterface interface {
	Path() string
	Method() string
	SetParam(string, interface{})
	ReadBody() *ResponseStatus
	ReadBodyArray() *ResponseStatus
	Redirect() *ResponseStatus
	HandleStatus(*ResponseStatus)
	MSI(string) (bool, Object)
	IA(string) (bool, Array)
	String(string) (bool, string)
	Float64(string) (bool, float64)
	Bool(string) (bool, bool)
	Int(string) (bool, int)
	HttpError(string, int)
	Writer() io.Writer
}