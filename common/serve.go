package gf

import (
		"strings"
		)

const	(
		ROBOTS_TXT = "User-agent: *\nDisallow: /api/"
		)

// main handler
func (node *Node) MainHandler(req RequestInterface, fullPath string) {

	req.SetHeader("Access-Control-Allow-Origin", "*")
	if req.Method() == "OPTIONS" { return }

	switch fullPath {

		case "/_.js":

			req.SetHeader("Content-Type", "application/javascript")
			
			for handler, _ := range node.Config.activeHandlers {

				req.Write(handler.clientJS.Bytes())

			}
			
			return

		case "/_.json":

			// render the handler documentation

			tree := []*HandlerSpec{}

			node.Config.RLock()

			for handler, _ := range node.Config.activeHandlers {

				tree = append(tree, handler.Spec())

			}

			node.Config.RUnlock()

			HandleStatus(req, req.Respond(tree))
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

	var lastSegment string

	for _, pathParam := range segments {

		if len(pathParam) == 0 { break }

		lastSegment = pathParam

		n, status := next.Next(req, pathParam)

		if n == next {

			// handler is a folder

			break
		}

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

		req.HttpError("NO CONTROLLER FOUND AT "+next.Path(), 500)
		return

	}

	handler.ReadPayload(req)

	handler.Handle(req, lastSegment)

}