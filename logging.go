package gf

import 	(
		"reflect"
		"encoding/json"
		//
		"github.com/fatih/color"
		)

func (req *Request) Debug(msg string) { color.Blue(req.Path() + ": %v", msg) }

func (req *Request) NewError(msg string) { color.Red(req.Path() + ": %v", msg) }

func (req *Request) Error(msg error) { color.Red(req.Path() + ": %v", msg) }

func (req *Request) Reflect(e interface{}) {

	msg := "REFLECT VALUE IS NIL"
	if e != nil { msg = "REFLECT VALUE IS "+reflect.TypeOf(e).String() }

	req.NewError(msg)
}

func (req *Request) DebugJSON(i interface{}) {
	b, err := json.Marshal(i); if err != nil { req.Error(err); return }
	req.Debug(string(b))
}

func (req *Request) HttpError(msg string, statusCode int) {

	req.ctx.Error(msg, statusCode)
	req.NewError(msg)
}
