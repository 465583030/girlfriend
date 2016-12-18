package gf

import 	(
		"sync"
		"bytes"
		)

type HandlerFunction func (*Request) *ResponseStatus

type Registry map[string]func (*Request) *ResponseStatus

type Config struct {
	host string
	rootRegistry Registry
	handlerRegistry Registry
	activeHandlers map[*Handler]struct{}
	clientJS *bytes.Buffer
	isAppEngine bool
	countries map[string]*Country
	languages map[string]*Language
	sync.RWMutex
}

func (config *Config) getRootFunction(functionKey string) HandlerFunction {

	if config.rootRegistry == nil { return nil }

	config.RLock()
		function := config.rootRegistry[functionKey]
	config.RUnlock()

	return function
}

func (config *Config) getHandlerFunction(functionKey string) HandlerFunction {

	config.RLock()
		function := config.handlerRegistry[functionKey]
	config.RUnlock()

	return function
}
