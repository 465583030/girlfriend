package gf

import 	(
		"bytes"
		//
		"github.com/microcosm-cc/bluemonday"
		)

var globalNode *Node

func init() {

	globalNode = &Node{
		routes:			map[string]*Node{},
		methods:		map[string]*Handler{},
		validations:	[]*ValidationConfig{},
	}

	globalNode.Config = &Config{
		activeHandlers:		map[*Handler]struct{}{},
		countries:			Countries(),
		languages:			Languages(),
		clientJS:			bytes.NewBuffer(nil),
		sanitizer:			bluemonday.StrictPolicy(),
	}

}

func Root() *Node {

	return globalNode
}
