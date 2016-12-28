package gf

import 	(
		"sync"
		"html"
		//
		"github.com/microcosm-cc/bluemonday"
		)

type Config struct {
	Host string
	RootRegistry Registry
	HandlerRegistry Registry
	ModuleRegistry ModuleRegistry
	headers Headers
	lDelim, rDelim string
	activeHandlers map[*Handler]struct{}
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

// Sets the root registry to the specified map
func (config *Config) SetRootRegistry(reg Registry) {

	config.Lock()
	
		config.RootRegistry = reg

	config.Unlock()

}

// Sets the handler registry to the specified map
func (config *Config) SetHandlerRegistry(reg Registry) {

	config.Lock()
	
		config.HandlerRegistry = reg

	config.Unlock()

}

// Sets the module registry to the specified map
func (config *Config) SetModuleRegistry(reg ModuleRegistry) {

	config.Lock()
	
		config.ModuleRegistry = reg

	config.Unlock()

}

// Sets the http preflight headers to the specified map
func (config *Config) SetHeaders(h Headers) {

	config.Lock()
	
		config.headers = h

	config.Unlock()

}

// Returns the root function if present in the registry
func (config *Config) GetRootFunction(functionKey string) HandlerFunction {

	if config.RootRegistry == nil { return nil }

	config.RLock()
		function := config.RootRegistry[functionKey]
	config.RUnlock()

	return function
}

// Returns the handler function if present in the registry
func (config *Config) GetHandlerFunction(functionKey string) HandlerFunction {

	config.RLock()
		function := config.HandlerRegistry[functionKey]
	config.RUnlock()

	return function
}

// Returns the handler function if present in the registry
func (config *Config) GetModuleFunction(functionKey string) ModuleFunction {

	config.RLock()
		function := config.ModuleRegistry[functionKey]
	config.RUnlock()

	return function
}


