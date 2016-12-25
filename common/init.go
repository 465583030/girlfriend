package gf

import 	(
		//
		"github.com/microcosm-cc/bluemonday"
		)

var globalNode *Node

func init() {

	globalNode = &Node{
		routes:			map[string]*Node{},
		methods:		map[string]*Handler{},
		modules:		[]*Module{},
		validations:	[]*ValidationConfig{},
	}

	globalNode.Config = &Config{
		activeHandlers:		map[*Handler]struct{}{},
		countries:			Countries(),
		languages:			Languages(),
		sanitizer:			bluemonday.StrictPolicy(),
		lDelim:				"{{",
		rDelim:				"}}",
	}

}

func Root() *Node {

	return globalNode
}
