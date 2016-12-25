package gf

import	(
		"fmt"
		"sync"
		"strings"
		temp "html/template"
		)

type Node struct {
	Config *Config
	Path string
	param *Node
	routes map[string]*Node
	methods map[string]*Handler
	modules []*Module
	validation *ValidationConfig
	validations []*ValidationConfig
	sync.RWMutex
}

func (node *Node) new(path string) *Node {

	n := &Node{
		Config:			node.Config,
		routes:			map[string]*Node{},
		methods:		map[string]*Handler{},
		modules:		[]*Module{},
		// inherited properties
		Path: 			node.Path + "/" + path,
		validations:	node.validations,
	}

	return n
}

// Adds a new path-node to the tree
func (node *Node) Add(path string) *Node {

	path = strings.TrimSpace(strings.Replace(path, "/", "", -1))

	n := node.new(path)

	node.Lock()
		node.routes[path] = n
	node.Unlock()

	return n
}

// Adds a new param-node
func (node *Node) Param(config *ValidationConfig, keys ...string) *Node {

	if len(keys) == 0 { panic("NO KEYS SUPPLIED FOR NEW PARAMETER") }

	n := node.new(":" + keys[0])

	config.keys = keys

	n.Lock()
		n.validation = config
		n.validations = append(n.validations, config)
	n.Unlock()

	node.param = n

	return n
}

// Adds a new path-node to the tree
func (node *Node) Mod(functionKey string, keys ...string) *Node {

	if node.Config.ModuleRegistry == nil { panic("Config has no ModuleRegistry setting!") }

	module := &Module{
		config:					node.Config,
		functionKey:			functionKey,
		paramKeys:				keys,
	}

	node.Lock()
		node.modules = append(node.modules, module)
	node.Unlock()

	return node
}

// traversal

// finds next node according to supplied URL path segment
func (node *Node) Next(req RequestInterface, pathSegment string) (*Node, *ResponseStatus) {

	// check for child routes

	next := node.routes[pathSegment]

	if next != nil { return next, nil }

	// check for path param

	next = node.param

	if next == nil { return nil, nil }

	// check for file glob

	handler := node.methods["GET"]

	if handler != nil && handler.isFolder {

		return node, nil

	}

	if next.validation != nil {

		ok, value := next.validation.pathFunction(req, pathSegment); if !ok {

			return nil, &ResponseStatus{nil, 400, fmt.Sprintf("UNEXPECTED VALUE  %v, %v", pathSegment, next.validation.Expecting())}

		}

		// write route params into request object

		for _, key := range next.validation.keys { req.SetParam(key, value) }

	}

	return next, nil
}

func (node *Node) handler(req RequestInterface) *Handler {

	node.RLock()

		handler := node.methods[req.Method()]

	node.RUnlock()

	return handler
}

// templates

func (node *Node) newTemplate() *temp.Template {

	return temp.New("").Delims(node.Config.lDelim, node.Config.rDelim)

}

func (node *Node) Template(templatePath string) *Node {

	t, err := node.newTemplate().ParseFiles(templatePath); if err != nil { panic(err) }

	h := &Handler{
		isFile:					true,
		template:				t,
		templatePath:			templatePath,
		templateType:			"text/html",
	}

	node.addHandler("GET", h)

	return node
}

func (node *Node) TemplateFolder(folderPath string) *Node {

	h := &Handler{
		isFolder:				true,
		templatePath:			folderPath,
		templateType:			"text/html",
	}

	node.addHandler("GET", h)

	return node
}

func (node *Node) File(templatePath, contentType string) *Node {

	t, err := node.newTemplate().ParseFiles(templatePath); if err != nil { panic(err) }

	h := &Handler{
		isFile:					true,
		template:				t,
		templatePath:			templatePath,
		templateType:			contentType,
	}

	node.addHandler("GET", h)

	return node
}

// methods

// Allows GET requests to the node's handler
func (node *Node) GET(functionKeys ...string) *Handler {

	h := &Handler{}
	
	if len(functionKeys) > 0 {
		h.functionKey = functionKeys[0]
	}

	node.addHandler("GET", h)

	return h
}

// Allows POST requests to the node's handler
func (node *Node) POST(functionKeys ...string) *Handler {

	h := &Handler{}
	
	if len(functionKeys) > 0 {
		h.functionKey = functionKeys[0]
	}

	node.addHandler("POST", h)

	return h
}

// Allows PUT requests to the node's handler
func (node *Node) PUT(functionKeys ...string) *Handler {

	h := &Handler{}
	
	if len(functionKeys) > 0 {
		h.functionKey = functionKeys[0]
	}

	node.addHandler("PUT", h)

	return h
}

// Allows POST requests to the node's handler
func (node *Node) DELETE(functionKeys ...string) *Handler {

	h := &Handler{}
	
	if len(functionKeys) > 0 {
		h.functionKey = functionKeys[0]
	}

	node.addHandler("DELETE", h)

	return h
}