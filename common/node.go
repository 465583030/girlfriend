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
	validation *ValidationConfig
	validations []*ValidationConfig
	sync.RWMutex
}

func (node *Node) new(path string) *Node {

	n := &Node{
		Config:			node.Config,
		routes:			map[string]*Node{},
		methods:		map[string]*Handler{},
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

// traversal

// finds next node according to supplied URL path segment
func (node *Node) Next(req RequestInterface, pathSegment string) (*Node, *ResponseStatus) {

	// check for child routes

	next := node.routes[pathSegment]

	if next != nil { return next, nil }

	// check for path param

	next = node.param

	if next == nil { return nil, nil }

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

func (node *Node) Template(templatePath string) *Node {

	t, err := temp.New("").Delims(node.Config.lDelim, node.Config.rDelim).ParseFiles(templatePath); if err != nil { panic(err) }

	h := &Handler{
		handlerType:			"file",
		template:				t,
		templatePath:			templatePath,
	}

	node.addHandler("GET", h)

	return node
}

// methods

func (node *Node) addHandler(method string, h *Handler) {

	h.method = method
	h.Config = node.Config
	h.node = node

	node.Lock()
		node.methods[method] = h
	node.Unlock()

	node.Config.Lock()
		node.Config.activeHandlers[h] = struct{}{}
	node.Config.Unlock()

	h.GenerateClientJS()
}

// Allows GET requests to the node's handler
func (node *Node) GET(functionKey string, responseSchema ...interface{}) *Node {

	h := &Handler{
		functionKey:			functionKey,
	}

	if len(responseSchema) > 0 {
		h.responseSchema = responseSchema[0]
	}

	node.addHandler("GET", h)

	return node
}

// Allows POST requests to the node's handler
func (node *Node) POST(functionKey string, schemes ...interface{}) *Node {

	h := &Handler{
		functionKey:			functionKey,
	}

	if len(schemes) > 0 {
		h.payloadSchema = schemes[0]
	}

	if len(schemes) > 1 {
		h.responseSchema = schemes[1]
	}

	node.addHandler("POST", h)

	return node
}
