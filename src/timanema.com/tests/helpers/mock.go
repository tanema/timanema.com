package helpers

import (
  "github.com/robfig/revel"
  "net/http/httptest"
  "net/http"
)

func MockController(controller, action string) *revel.Controller {
  c :=  &revel.Controller{
    Request:  buildEmptyRequest(),
    Response: revel.NewResponse(httptest.NewRecorder()),
    Params:   new(revel.Params),
    Args:     map[string]interface{}{},
    Flash:    revel.Flash{Data: map[string]string{}, Out: map[string]string{}},
    Validation: &revel.Validation{},
    RenderArgs: map[string]interface{}{
      "RunMode": revel.RunMode,
      "DevMode": revel.DevMode,
    },
  }

  c.SetAction(controller, action)

  return c
}

func buildEmptyRequest() *revel.Request {
  httpRequest, _ := http.NewRequest("GET", "/", nil)
  request := revel.NewRequest(httpRequest)
  return request
}
