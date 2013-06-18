package mimes

import (
  "github.com/robfig/revel"
  "net/http"
)

type Png string

func (r Png) Apply(req *revel.Request, resp *revel.Response) {
  resp.WriteHeader(http.StatusOK, "image/png")
  resp.Out.Write([]byte(r))
}
