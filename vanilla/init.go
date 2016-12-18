package girlfriend

import	(
		"github.com/valyala/fasthttp"
		//
		"github.com/golangdaddy/girlfriend/common"
		)

func NewRouter(host string) (*gf.Node, func (*fasthttp.RequestCtx)) {

	root := gf.Root()

	root.Config.Host = host

	f := func (ctx *fasthttp.RequestCtx) {

		node := gf.Root()

		req := NewRequestObject(node, ctx)

		node.MainHandler(req, string(ctx.Path()))

	}

	return root, f
}
