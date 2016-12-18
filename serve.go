package gf

import (
		"strings"
		"encoding/json"
		//
		"github.com/valyala/fasthttp"
		)

const	(
		ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
		)

// main handler
func (node *Node) mainHandler(ctx *fasthttp.RequestCtx) {

	req := node.NewRequestObject(ctx)

	fullPath := string(req.ctx.Path())

	switch fullPath {

		case "/_.js":

			req.ctx.Request.Header.Set("Content-Type", "application/javascript")
			req.ctx.Write(node.config.clientJS.Bytes())
			return

		case "/_.json":

			// render the handler documentation

			tree := map[string]*HandlerSpec{}

			node.config.RLock()

			for handler, _ := range node.config.activeHandlers {

				spec := handler.Spec()

				tree[spec.Endpoint] = spec

			}

			node.config.RUnlock()

			b, err := json.Marshal(tree); if err != nil { req.HttpError(err.Error(), 500); return }
			req.ctx.Request.Header.Set("Content-Type", "application/json")
			req.ctx.Write(b)
			return

		case "/robots.txt":

			req.ctx.Write([]byte(ROBOTS_TXT))
			return

		default:

			rootFunc := node.config.getRootFunction(fullPath)

			if rootFunc != nil {

				req.HandleStatus(rootFunc(req))
				return

			}

	}

	segments := strings.Split(fullPath, "/")[1:]

	next := req.Node

	for _, pathParam := range segments {

		if len(pathParam) == 0 { break }

		n, status := next.next(req, pathParam)

		if status != nil {

			req.HandleStatus(status)
			return

		}

		if n != nil {

			next = n
			continue

		}

		break
	}

	// resolve handler

	handler := next.handler(req)

	if handler == nil {

		req.HttpError("NO CONTROLLER FOUND AT "+next.Path, 500)
		return

	}

	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" { return }

	handler.Handle(req)

}


