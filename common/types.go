package gf

type Array []interface{}
type Object map[string]interface{}

type HandlerFunction func (RequestInterface) *ResponseStatus

type Registry map[string]func (RequestInterface) *ResponseStatus