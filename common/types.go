package gf

type Array []interface{}
type Object map[string]interface{}

type HandlerFunction func (RequestInterface) *ResponseStatus

type Registry map[string]HandlerFunction

type ModuleFunction func (RequestInterface, ...string) *ResponseStatus

type ModuleRegistry map[string]ModuleFunction