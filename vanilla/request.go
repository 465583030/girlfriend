package girlfriend

import 	(
		"io"
		"net/http"
		"encoding/json"
		//
		"github.com/valyala/fasthttp"
		//
		"github.com/golangdaddy/girlfriend/common"
		)

type Request struct {
	ctx *fasthttp.RequestCtx
	config *gf.Config
	Node *gf.Node
	method string
	Params map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
}

func NewRequestObject(node *gf.Node, ctx *fasthttp.RequestCtx) *Request {

	return &Request{
		ctx:			ctx,
		config:			node.Config,
		Node:			node,
		method:			string(ctx.Method()),
		Params:			map[string]interface{}{},
		Object:			gf.Object{},
		Array:			gf.Array{},
	}
}

func (req *Request) Res() http.ResponseWriter {

	x := new(http.ResponseWriter)

	return *x
}

func (req *Request) R() *http.Request {

	return &http.Request{}
}

func (req *Request) BodyArray() []interface{} {

	return req.Array
}

func (req *Request) Path() string {

	return req.Node.Path
}

func (req *Request) Method() string {

	return req.method
}

func (req *Request) Writer() io.Writer {
	return req.ctx.Response.BodyWriter()
}

func (req *Request) Write(b []byte) {

	req.ctx.Write(b)
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

	req.ctx.Request.Header.Set(k, v)
}

func (req *Request) ReadBodyArray() *gf.ResponseStatus {

	err := json.Unmarshal(req.ctx.PostBody(), &req.Array); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBody() *gf.ResponseStatus {

	err := json.Unmarshal(req.ctx.PostBody(), &req.Object); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) Fail() *gf.ResponseStatus {

	return gf.Fail()
}

func (req *Request) Respond(args ...interface{}) *gf.ResponseStatus {

	return gf.Respond(args)
}

func (req *Request) Redirect(path string, code int) *gf.ResponseStatus {

	req.ctx.Redirect(path, code)

	return nil
}

func (req *Request) HttpError(msg string, code int) {

	req.ctx.Error(msg, code)
	req.NewError(msg)
}

