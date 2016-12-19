package girlfriend

import (
  "testing"
  //"net/http"
  //"net/http/httptest"
  //
  //"github.com/golangdaddy/girlfriend/common"
)

func BenchmarkMainHandler1000000(b *testing.B) {
  /* Doesn't work
  root, _ := NewRouter("peach")
  config := root.Config

  node := root.Add("path")

  config.SetHandlerRegistry(gf.Registry{
    "handler": func (req gf.RequestInterface) *gf.ResponseStatus {
      return gf.Respond("")
    },
  })
  node.GET("handler")

  w := httptest.NewRecorder()

  req, err := http.NewRequest("GET", "/path", nil)
  if err != nil { b.Error(err) }

  for i := 0; i < b.N; i++ {
    root.MainHandler(NewRequestObject(root, w, req), "/path")
  }
  */
}
