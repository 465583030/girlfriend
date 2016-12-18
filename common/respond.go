package gf

import (
		"reflect"
		"strconv"
		)

type ResponseStatus struct {
	Value interface{}
	Code int
	Message string
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

func Fail() *ResponseStatus {

	return Respond(500, "UNEXPECTED APPLICATION ERROR")
}
