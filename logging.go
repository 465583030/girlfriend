package gf

import 	(
		"reflect"
		"net/http"
		"encoding/json"
		//
		"github.com/fatih/color"
		//
		"google.golang.org/appengine"
		"google.golang.org/appengine/log"
		)

func (req *Request) Debug(msg string) {

	if req.config.isAppEngine {

		ctx := appengine.NewContext(req.ctx.Request)
		log.Debugf(ctx, req.Path() + ": %v", msg)

	} else {

		color.Blue(req.Path() + ": %v", msg)

	}
}

func (req *Request) NewError(msg string) {

	if req.config.isAppEngine {

		ctx := appengine.NewContext(req.ctx.Request)
		log.Errorf(ctx, req.Path() + ": %v", msg)

	} else {

		color.Red(req.Path() + ": %v", msg)

	}
}

func (req *Request) Error(msg error) {

	if req.config.isAppEngine {

		ctx := appengine.NewContext(req.ctx.Request)
		log.Errorf(ctx, req.Path() + ": %v", msg.Error())

	} else {

		color.Red(req.Path() + ": %v", msg)

	}
}

func (req *Request) Reflect(e interface{}) {
	msg := "REFLECT VALUE IS NIL"
	if e != nil {
		msg = "REFLECT VALUE IS "+reflect.TypeOf(e).String()
	}
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
