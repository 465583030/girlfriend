package gf

import (
		"strings"
		"encoding/json"
		)

const	(
		ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
		)

// main handler
func (node *Node) MainHandler(req RequestInterface, fullPath string) {

	switch fullPath {

		case "/_.js":

			req.SetHeader("Content-Type", "application/javascript")
			req.Write(node.Config.clientJS.Bytes())
			return

		case "/_.json":

			// render the handler documentation

			tree := map[string]*HandlerSpec{}

			node.Config.RLock()

			for handler, _ := range node.Config.activeHandlers {

				spec := handler.Spec()

				tree[spec.Endpoint] = spec

			}

			node.Config.RUnlock()

			b, err := json.Marshal(tree); if err != nil { req.HttpError(err.Error(), 500); return }
			req.SetHeader("Content-Type", "application/json")
			req.Write(b)
			return

		case "/robots.txt":

			req.Write([]byte(ROBOTS_TXT))
			return

		default:

			rootFunc := node.Config.GetRootFunction(fullPath)

			if rootFunc != nil {

				HandleStatus(req, rootFunc(req))
				return

			}

	}

	segments := strings.Split(fullPath, "/")[1:]

	next := node

	for _, pathParam := range segments {

		if len(pathParam) == 0 { break }

		n, status := next.Next(req, pathParam)

		if status != nil {

			HandleStatus(req, status)
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

	req.SetHeader("Access-Control-Allow-Origin", "*")
	if req.Method() == "OPTIONS" { return }

	handler.Handle(req)

}