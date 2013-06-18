package mimes

import (
  "github.com/robfig/revel"
  "net/http"
  "io"
  "encoding/base64"
  "strings"
)

type Png string

func (r Png) Apply(req *revel.Request, resp *revel.Response) {
  resp.WriteHeader(http.StatusOK, "image/png")
  data := strings.Split(string(r), ",")
  if len(data) > 1 {
    png := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[1]))
    io.Copy(resp.Out, png)
  }
}
