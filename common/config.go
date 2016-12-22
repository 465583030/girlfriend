package gf

import 	(
		"sync"
		"html"
		"bytes"
		//
		"github.com/microcosm-cc/bluemonday"
		)

type Config struct {
	Host string
	RootRegistry Registry
	HandlerRegistry Registry
	lDelim, rDelim string
	activeHandlers map[*Handler]struct{}
	clientJS *bytes.Buffer
	countries map[string]*Country
	languages map[string]*Language
	sanitizer *bluemonday.Policy
	sync.RWMutex
}

func (config *Config) Sanitize(s string) string {

	return config.sanitizer.Sanitize(html.UnescapeString(s))
}

func (config *Config) SetDelims(l, r string) {

	config.Lock()
		config.lDelim = l
		config.rDelim = r
	config.Unlock()
}

// Allows the setting of the registry
func (config *Config) SetRootRegistry(reg Registry) {

	config.Lock()
	
		config.RootRegistry = reg

	config.Unlock()

}

func (config *Config) SetHandlerRegistry(reg Registry) {

	config.Lock()
	
		config.HandlerRegistry = reg

	config.Unlock()

}

func (config *Config) GetRootFunction(functionKey string) HandlerFunction {

	if config.RootRegistry == nil { return nil }

	config.RLock()
		function := config.RootRegistry[functionKey]
	config.RUnlock()

	return function
}

func (config *Config) GetHandlerFunction(functionKey string) HandlerFunction {

	config.RLock()
		function := config.HandlerRegistry[functionKey]
	config.RUnlock()

	return function
}


