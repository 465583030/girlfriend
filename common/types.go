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
	MSI(interface{}) (bool, Object)
	IA(interface{}) (bool, Array)
	String(interface{}) (bool, string)
	Float64(interface{}) (bool, float64)
	Bool(interface{}) (bool, bool)
	Int(interface{}) (bool, int)
	HttpError(string, int)
	Writer() io.Writer
}