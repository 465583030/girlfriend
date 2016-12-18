package girlfriend

import 	(
		"reflect"
		"net/http"
		"encoding/json"
		//
		"google.golang.org/appengine"
		"google.golang.org/appengine/log"
		)

func (req *Request) Debug(msg string) {

	ctx := appengine.NewContext(req.R)
	log.Debugf(ctx, req.Path() + ": %v", msg)

}

func (req *Request) NewError(msg string) {

	ctx := appengine.NewContext(req.R)
	log.Errorf(ctx, req.Path() + ": %v", msg)

}

func (req *Request) Error(msg error) {

	ctx := appengine.NewContext(req.R)
	log.Errorf(ctx, req.Path() + ": %v", msg.Error())

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

	http.Error(req.Res, msg, statusCode)
	req.NewError(msg)
}
