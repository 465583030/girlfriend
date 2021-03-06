package gf

type Array []interface{}
type Object map[string]interface{}

type Headers map[string]string

type HandlerFunction func (RequestInterface) *ResponseStatus

type Registry map[string]HandlerFunction

type ModuleFunction func (RequestInterface, interface{}) *ResponseStatus

type ModuleRegistry map[string]ModuleFunction
