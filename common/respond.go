package gf

import (
		"reflect"
		"strconv"
		"encoding/json"
		//
		"github.com/golangdaddy/gaml"
		)

type ResponseStatus struct {
	Value interface{}
	Code int
	Message string
}

func Fail() *ResponseStatus {

	return Respond(500, "UNEXPECTED APPLICATION ERROR")
}

func Respond(args ...interface{}) *ResponseStatus {

	var ok bool
	s := &ResponseStatus{}

	switch len(args) {

		case 1:

			s.Value = args[0]
			s.Code = 200
			return s

		case 2:

			s.Code, ok = args[0].(int); if !ok {
				return &ResponseStatus{nil, 501, "UNEXPECTED RESPONSE PARAMETER 0 TYPE: " + reflect.TypeOf(args[0]).String()}
			}
			s.Message, ok = args[1].(string); if !ok {
				return &ResponseStatus{nil, 501, "UNEXPECTED RESPONSE PARAMETER 1 TYPE: " + reflect.TypeOf(args[1]).String()}
			}
			return s

		default:

			return &ResponseStatus{nil, 400, "INVALID STATUS ARGS LENGTH: "+strconv.Itoa(len(args))}

	}

	return nil
}

func HandleStatus(req RequestInterface, status *ResponseStatus) {

	req.DebugJSON(status)

	// return with no action if handler returns nil
	if status == nil { return }

	if status.Code == 200 {

		switch v := status.Value.(type) {

			case nil:

				return

			case *gaml.ELEMENT:

				b, err := v.Render(); if err != nil { req.Error(err); break }
				req.Write(b)
				return

			case []byte:

				req.Write(v)
				return

			default:

				req.SetHeader("Content-Type", "application/json")
				b, err := json.Marshal(status.Value); if err != nil { req.Error(err); break }
				req.Write(b)
				return

		}

		return

	}

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.Message

	req.NewError(statusMessage)
	req.HttpError(statusMessage, status.Code)
}
