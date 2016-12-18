package girlfriend

import (
  "testing"
  //
  "github.com/valyala/fasthttp"
  "github.com/golangdaddy/girlfriend/common"
)

func BenchmarkMainHandler1000000(b *testing.B) {
  root, _ := NewRouter("peach")
  config := root.Config

  node := root.Add("path")

  config.SetHandlerRegistry(gf.Registry{
    "handler": func (req gf.RequestInterface) *gf.ResponseStatus {
      return gf.Respond("")
    },
  })
  node.GET("handler")

  ctx := &fasthttp.RequestCtx{}

  for i := 0; i < b.N; i++ {
    root.MainHandler(NewRequestObject(root, ctx), "/path")
  }
}
