package girlfriend

import 	(
		"io"
		"net/http"
		"io/ioutil"
		"encoding/json"
		//
		"github.com/golangdaddy/girlfriend/common"
		)

type Request struct {
	config *gf.Config
	Node *gf.Node
	method string
	res http.ResponseWriter
	r *http.Request
	Params map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
}

func NewRequestObject(node *gf.Node, res http.ResponseWriter, r *http.Request) *Request {

	return &Request{
		config:			node.Config,
		Node:			node,
		res:			res,
		r: 				r,
		Params:			map[string]interface{}{},
		Object:			gf.Object{},
		Array:			gf.Array{},
	}
}

func (req *Request) Res() http.ResponseWriter {

	return req.res
}

func (req *Request) R() *http.Request {

	return req.r
}

func (req *Request) BodyArray() []interface{} {

	return req.Array
}

func (req *Request) BodyObject() map[string]interface{} {

	return req.Object
}

func (req *Request) Path() string {

	return req.Node.Path
}

func (req *Request) Method() string {

	return req.method
}

func (req *Request) Writer() io.Writer {

	return req.res
}

func (req *Request) Write(b []byte) {

	req.res.Write(b)
}

func (req *Request) Body(k string) interface{} {

	return req.Object[k]
}

func (req *Request) Param(k string) interface{} {

	return req.Params[k]
}

func (req *Request) SetParam(k string, v interface{}) {

	req.Params[k] = v
}

func (req *Request) SetHeader(k, v string) {

	req.res.Header().Set(k, v)
}

func (req *Request) ReadBodyArray() *gf.ResponseStatus {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)

	if body != nil { body.Close() }

	if err != nil { return gf.Respond(400, err.Error()) }

	err = json.Unmarshal(b, &req.Array); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBody() *gf.ResponseStatus {

	body := req.r.Body

	b, err := ioutil.ReadAll(body)

	if body != nil { body.Close() }

	if err != nil { return gf.Respond(400, err.Error()) }

	err = json.Unmarshal(b, &req.Object); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) Fail() *gf.ResponseStatus {

	return gf.Fail()
}

func (req *Request) Respond(args ...interface{}) *gf.ResponseStatus {

	return gf.Respond(args)
}

func (req *Request) Redirect(path string, code int) *gf.ResponseStatus {

	http.Redirect(req.res, req.r, path, code)

	return nil
}

func (req *Request) HttpError(msg string, code int) {

	http.Error(req.res, msg, code)
	req.NewError(msg)
}
