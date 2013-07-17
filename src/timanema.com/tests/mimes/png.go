package tests

import (
  "github.com/robfig/revel"
  "timanema.com/app/mimes"
  "github.com/tanema/revel_mock"
  "net/http/httptest"
  "encoding/base64"
  "strings"
  "bytes"
)

type PngMimeTest struct {
	revel.TestSuite
}

func (t PngMimeTest) Before() { }
func (t PngMimeTest) After() { }

func (t PngMimeTest) TestApply() {
  test_string := "data blah, this is a test string that will be made into a png"
  req := revel_mock.BuildEmptyRequest()
  resp := httptest.NewRecorder()
  mimes.Png(test_string).Apply(req, revel.NewResponse(resp))
  data := strings.Split(test_string, ",")
  b64s := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[1]))
  buf := new(bytes.Buffer)
  buf.ReadFrom(b64s)
  t.Assert(resp.Body.String() == buf.String())
}
