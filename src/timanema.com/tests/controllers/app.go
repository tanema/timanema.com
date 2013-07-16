package tests

import (
  "github.com/robfig/revel"
  "timanema.com/app/controllers"
  "timanema.com/tests/helpers"
  "timanema.com/app/models"
  "labix.org/v2/mgo/bson"
  "reflect"
)

type AppControllerTest struct {
	revel.TestSuite
}

func (t AppControllerTest) Before() { }
func (t AppControllerTest) After() { }

func (t AppControllerTest) TestIndexFunctional() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t AppControllerTest) TestIndexResult() { 
  result, ok := (controllers.App{helpers.MockController("App","Index")}.Index()).(*revel.RenderTemplateResult)
  t.Assert(ok) //succeeded rendering

  signatures := []models.Signature{}
  models.Signatures().All(&signatures, bson.M{"order": "-_id", "limit": 5})
  count, _ := models.Signatures().Count(nil)

  t.Assert(result.RenderArgs["count"] == count)
  t.Assert(reflect.ValueOf(result.RenderArgs["signatures"]).Len() == len(signatures))
}
