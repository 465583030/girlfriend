package girlfriend

import	(
		"regexp"
		"net/http"
		//
		"github.com/golangdaddy/girlfriend/common"
		)

type WildcardRouter struct {
	handler http.Handler
}

func (router *WildcardRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {}

func (router *WildcardRouter) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {}

func (router *WildcardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) { router.handler.ServeHTTP(w, r) }

// create a new router for a GF app
func NewRouter(host string) (*gf.Node, *WildcardRouter) {

	root := gf.Root()

	root.Config.Host = host

	f := func (res http.ResponseWriter, r *http.Request) {

		node := gf.Root()

		req := NewRequestObject(node, res, r)

		node.MainHandler(req, r.URL.Path)

	}

	return root, &WildcardRouter{http.HandlerFunc(f)}
}
