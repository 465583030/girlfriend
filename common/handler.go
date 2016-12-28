package gf

import	(
		"sync"
		"bytes"
		"strings"
		temp "html/template"
		)

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

type Handler struct {
	Config *Config
	node *Node
	method string
	functionKey string
	function HandlerFunction
	isFile bool
	isFolder bool
	template *temp.Template
	templatePath string
	templateType string
	responseSchema interface{}
	payloadSchema interface{}
	clientJS *bytes.Buffer
	sync.RWMutex
}

func (handler *Handler) Name() string {

	var name string

	parts := strings.Split(handler.node.Path(), "/")

	if len(parts) == 1 { return "" }

	for _, part := range parts[1:] {

		if string(part[0]) == ":" { continue }

		name += strings.Title(part)

	}

	return name
}

func (handler *Handler) ApiUrl() string {

	var name string

	parts := strings.Split(handler.node.Path(), "/")

	if len(parts) == 1 { return "" }

	for _, part := range parts[1:] {

		if string(part[0]) == ":" {

			part = "'+" + part[1:] + "+'"

		}

		name += "/" + part

	}

	return "'" + name + "'"
}

// Applys model which describes request payload
func (handler *Handler) Payload(schema ...interface{}) *Handler {

	handler.payloadSchema = schema[0]

	return handler
}

// Applys model which describes response schema
func (handler *Handler) Response(schema ...interface{}) *Handler {

	handler.responseSchema = schema[0]

	return handler
}

// Validates any payload present in the request body, according to the payloadSchema
func (handler *Handler) ReadPayload(req RequestInterface) {

	// handle payload

	switch v := handler.payloadSchema.(type) {

		case nil:

			// do nothing

		case []interface{}:

			status := req.ReadBodyArray(); if status != nil { HandleStatus(req, status); return }

		case map[string]*ValidationConfig:

			status := req.ReadBody(); if status != nil { HandleStatus(req, status); return }
			
			for key, validation := range v {

				value := req.Body(key)

				ok, x := validation.bodyFunction(req, value); if !ok { break }

				req.SetParam("_" + key, x)

			}

		default:

	}

}

// Executes the modules and any hander-function, template or folder
func (handler *Handler) Handle(req RequestInterface, pathSegment string) {

	// execute all module functions

	for _, module := range handler.node.modules {

		status := module.Run(req); if status != nil { HandleStatus(req, status); return }

	}

	if !handler.isFile && !handler.isFolder {

		function := handler.Config.GetHandlerFunction(handler.functionKey)

		if function == nil { panic("FAILED TO GET FUNCTION WITH KEY: "+handler.functionKey) }

		HandleStatus(req, function(req))

		return
	}

	if handler.isFolder {

		req.SetHeader("Content-Type", handler.templateType)

		req.ServeFile(handler.templatePath + "/" + pathSegment)

		return

	}

	if handler.isFile {

		path := strings.Split(handler.templatePath, "/")

		name := path[len(path)-1]

		req.SetHeader("Content-Type", handler.templateType)

		// serve a template added with the .Template(...) method
		err := handler.template.ExecuteTemplate(req.Writer(), name, nil); if err != nil { panic(err) }

		return

	}

}