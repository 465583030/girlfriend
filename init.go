package gf

import	(
		"bytes"
		)

var gfConfig *Config

func NewRouter(host string) *Node {

	root := &Node{
		routes:			map[string]*Node{},
		methods:		map[string]*Handler{},
		validations:	[]*ValidationConfig{},
	}

	gfConfig = &Config{
		host:				host,
		activeHandlers:		map[*Handler]struct{}{},
		countries:			Countries(),
		languages:			Languages(),
		isAppEngine:		true,
		clientJS:			bytes.NewBuffer(nil),
	}

	root.config = gfConfig

	return root
}
/*
// exposes the http router
func (node *Node) Router() *WildcardRouter {

	if node.config == nil { panic("INVALID CONFIG") }

	if node.config.router == nil { panic("INVALID ROUTER") }

	return node.config.router
}
*/

// registry config

// Allows the setting of the registry
func (node *Node) SetRootRegistry(reg Registry) {

	node.config.Lock()
	
		node.config.rootRegistry = reg

	node.config.Unlock()

}

func (node *Node) SetHandlerRegistry(reg Registry) {

	node.config.Lock()
	
		node.config.handlerRegistry = reg

	node.config.Unlock()

}