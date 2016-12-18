package girlfriend

import 	(
		"io"
		"strconv"
		"net/http"
		"io/ioutil"
		"encoding/json"
		//
		"github.com/golangdaddy/gaml"
		//
		"github.com/golangdaddy/girlfriend/common"
		)

type Request struct {
	config *gf.Config
	Node *gf.Node
	method string
	Res http.ResponseWriter
	R *http.Request
	Params map[string]interface{}
	Object map[string]interface{}
	Array []interface{}
}

func NewRequestObject(node *gf.Node, res http.ResponseWriter, r *http.Request) *Request {

	return &Request{
		config:			node.Config,
		Node:			node,
		Res:			res,
		R: 				r,
		Params:			map[string]interface{}{},
		Object:			gf.Object{},
		Array:			gf.Array{},
	}
}

func (req *Request) Path() string {
	
	return req.Node.Path
}

func (req *Request) Method() string {
	
	return req.method
}

func (req *Request) Writer() io.Writer {

	return req.Res
}

func (req *Request) Write(b []byte) {

	req.Res.Write(b)
}

func (req *Request) SetParam(k string, v interface{}) {
	
	req.Params[k] = v
}

func (req *Request) SetHeader(k, v string) {

	req.Res.Header().Set(k, v)
}

func (req *Request) ReadBodyArray() *gf.ResponseStatus {

	body := req.R.Body

	b, err := ioutil.ReadAll(body)

	if body != nil { body.Close() }

	if err != nil { return gf.Respond(400, err.Error()) }

	err = json.Unmarshal(b, &req.Array); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) ReadBody() *gf.ResponseStatus {

	body := req.R.Body

	b, err := ioutil.ReadAll(body)

	if body != nil { body.Close() }

	if err != nil { return gf.Respond(400, err.Error()) }

	err = json.Unmarshal(b, &req.Object); if err != nil { return gf.Respond(400, err.Error()) }

	return nil
}

func (req *Request) Redirect(path string, code int) *gf.ResponseStatus {

	http.Redirect(req.Res, req.R, path, code)

	return nil
}

func (req *Request) HandleStatus(status *gf.ResponseStatus) {

	// return with no action if handler returns nil
	if status == nil { return }

	if status.Code == 200 {

		switch v := status.Value.(type) {

			case nil:

				return

			case *gaml.ELEMENT:

				b, err := v.Render(); if err != nil { req.Error(err); break }
				req.Res.Write(b)
				return

			case []byte:

				req.Res.Write(v)
				return

			default:

				req.Res.Header().Set("Content-Type", "application/json")
				b, err := json.Marshal(status.Value); if err != nil { req.Error(err); break }
				req.Res.Write(b)
				return

		}

		return

	}

	statusMessage := "HTTP ERROR " + strconv.Itoa(status.Code) + ": " + status.Message

	req.NewError(statusMessage)
	http.Error(req.Res, statusMessage, status.Code)
}