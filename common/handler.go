package gf

import	(
		"strings"
		temp "html/template"
		)

type Handler struct {
	Config *Config
	node *Node
	method string
	functionKey string
	handlerType string
	template *temp.Template
	templateFolder *temp.Template
	templatePath string
	responseSchema interface{}
	payloadSchema interface{}
}

func (handler *Handler) Name() string {

	var name string

	parts := strings.Split(handler.node.Path, "/")

	if len(parts) == 1 { return "" }

	for _, part := range parts[1:] {

		if string(part[0]) == ":" { continue }

		name += strings.Title(part)

	}

	return name
}

func (handler *Handler) ApiUrl() string {

	var name string

	parts := strings.Split(handler.node.Path, "/")

	if len(parts) == 1 { return "" }

	for _, part := range parts[1:] {

		if string(part[0]) == ":" {

			part = "'+" + part[1:] + "+'"

		}

		name += "/" + part

	}

	return "'" + name + "'"
}

func (handler *Handler) Handle(req RequestInterface) {

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

	// serve output

	switch handler.handlerType {

		case "":

			function := handler.Config.GetHandlerFunction(handler.functionKey)

			if function == nil { panic("FAILED TO GET FUNCTION WITH KEY: "+handler.functionKey) }

			HandleStatus(req, function(req))

		case "file":

			err := handler.template.Execute(req.Writer(), nil); if err != nil { panic(err) }

		case "folder":

			// serve a template added with the .Template(...) method
			err := handler.template.ExecuteTemplate(req.Writer(), handler.templatePath, nil); if err != nil { panic(err) }

	}

}