package tests

import "github.com/robfig/revel"

type SignatureControllerTest struct {
	revel.TestSuite
}

func (t SignatureControllerTest) Before() { }
func (t SignatureControllerTest) After() { }

func (t SignatureControllerTest) TestIndex() { }
func (t SignatureControllerTest) TestCreate() { }
func (t SignatureControllerTest) TestShow() { }
func (t SignatureControllerTest) TestReport() { }
