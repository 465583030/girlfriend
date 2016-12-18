package gf

import	(
		"strings"
		"html/template"
		)


func (handler *Handler) GenerateClientJS() error {

	domain := strings.Title(handler.config.host)
	handlerName := handler.Name()

	// skip handler name which breaks the js
	if len(handlerName) == 0 { return nil }
	if string(handlerName[0]) == "_" { return nil }

	script := []string{
		"\n // " + handler.functionKey,
		"\n this." + strings.ToLower(handler.method) + domain + handlerName + " = function (",
	}

	payload := "null"

	args := []string{}
	for _, vc := range handler.node.validations { args = append(args, vc.Keys()) }

	if len(args) > 0 {
		script = append(script, strings.Join(args, ", "))
		script = append(script, ", ")
	}

	if handler.method == "POST" {

		script = append(script, "payload")
		payload = "payload"
		script = append(script, ", ")

	}

	str := "success, fail) { " + strings.ToLower(handler.method) + "('" + handler.config.host + "', " + handler.ApiUrl() + ", success, " + payload + ", fail); }; \n"

	script = append(script, str)

	t, err := template.New("").Parse(strings.Join(script, "")); if err != nil { return err }
	err = t.Execute(handler.config.clientJS, handler); if err != nil { return err }

	return nil
}