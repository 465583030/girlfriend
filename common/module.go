package gf

import 	(
		)

type Module struct {
	config *Config
	functionKey string
	function ModuleFunction
	paramKeys []string
}

func (mod *Module) Run(req RequestInterface) *ResponseStatus {

	if mod.function == nil {

		mod.function = mod.config.GetModuleFunction(mod.functionKey)
		
		if mod.function == nil { return req.Fail() }

	}

	return mod.function(req, mod.paramKeys...)
} 

