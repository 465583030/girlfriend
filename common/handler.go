package gf

import	(
		"html"
		"strings"
		"reflect"
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

			status := req.ReadBodyArray(); if status != nil { req.HandleStatus(status); return }

		case map[string]interface{}:

			status := req.ReadBody(); if status != nil { req.HandleStatus(status); return }

			for key, value := range v {

				switch v := value.(type) {

					case string:

						ok, s := req.String(key); if !ok { break }

						req.SetParam("_" + key, html.UnescapeString(s))

						continue

					case float64:

						ok, f := req.Float64(key); if !ok { break }

						req.SetParam("_" + key, f)

						continue

					case bool:

						ok, b := req.Bool(key); if !ok { break }

						req.SetParam("_" + key, b)

						continue

					case []interface{}:

						ok, a := req.IA(key); if !ok { break }

						req.SetParam("_" + key, a)

						continue

					case map[string]interface{}:

						ok, m := req.MSI(key); if !ok { break }

						req.SetParam("_" + key, m)

						continue

					default:

						req.HttpError("INVALID POST DATATYPE: " + key + " - " + reflect.TypeOf(v).String(), 400)
						return

				}

			}

		default:

	}

	// serve output

	switch handler.handlerType {

		case "":

			function := handler.Config.getHandlerFunction(handler.functionKey)

			if function == nil { panic("FAILED TO GET FUNCTION WITH KEY: "+handler.functionKey) }

			req.HandleStatus(function(req))

		case "file":

			err := handler.template.Execute(req.Writer(), nil); if err != nil { panic(err) }

		case "folder":

			// serve a template added with the .Template(...) method
			err := handler.template.ExecuteTemplate(req.Writer(), handler.templatePath, nil); if err != nil { panic(err) }

	}

}