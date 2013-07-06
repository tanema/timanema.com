package tests

import (
  "github.com/robfig/revel"
  "timanema.com/app/controllers"
)

type AppControllerTest struct {
	revel.TestSuite
}

func (t AppControllerTest) Before() { }
func (t AppControllerTest) After() { }

func (t AppControllerTest) TestIndex() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
  controllers.App{}.Index()
}
