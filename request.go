package gf

import 	(
		"strconv"
		"encoding/json"
		//
		"github.com/golangdaddy/gaml"
		"github.com/valyala/fasthttp"
		"github.com/microcosm-cc/bluemonday"
		)

type Request struct {
	ctx *fasthttp.RequestCtx
	config *Config
	Node *Node
	Method string
	Params map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
	BlueMonday *bluemonday.Policy
}

func (node *Node) NewRequestObject(ctx *fasthttp.RequestCtx) *Request {

	return &Request{
		ctx:			ctx,
		config:			node.config,
		Node:			node,
		Method:			string(ctx.Method()),
		Params:			map[string]interface{}{},
		Object:			Object{},
		Array:			Array{},
		BlueMonday:		bluemonday.StrictPolicy(),
	}
}

func (req *Request) Path() string {
	
	return req.Node.Path
}

func (req *Request) ReadBody(dst interface{}) *ResponseStatus {

	err := json.Unmarshal(req.ctx.PostBody(), dst); if err != nil { return Respond(400, err.Error()) }

	return nil
}

func (req *Request) Redirect(path string, code int) *ResponseStatus {

	req.ctx.Redirect(path, code)

	return nil
}

func (req *Request) HandleStatus(status *ResponseStatus) {

	// return with no action if handler returns nil
	if status == nil { return }

	if status.Code == 200 {

		switch v := status.Value.(type) {

			case nil:

				return

			case *gaml.ELEMENT:

				b, err := v.Render(); if err != nil { req.Error(err); break }
				req.ctx.Write(b)
				return

			case []byte:

				req.ctx.Write(v)
				return

			default:

				req.ctx.Response.Header.Set("Content-Type", "application/json")
				b, err := json.Marshal(status.Value); if err != nil { req.Error(err); break }
				req.ctx.Write(b)
				return

		}

		return

	}

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.Message

	req.NewError(statusMessage)
	req.ctx.Error(statusMessage, status.Code)
}